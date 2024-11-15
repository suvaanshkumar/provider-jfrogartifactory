/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/myorg/provider-artfactory/apis/v1beta1"
)

const (
	// Got this information from https://registry.terraform.io/providers/jfrog/artifactory/latest/docs#:~:text=provider%20%22artifactory%22%20%7B%0A%20%20url%20%20%20%20%20%20%20%20%20%20%20%3D%20%22%24%7Bvar.artifactory_url%7D/artifactory%22%0A%20%20access_token%20%20%3D%20%22%24%7Bvar.artifactory_access_token%7D%22%0A%7D
	KeyURL         = "url"
	KeyAccessToken = "access_token"
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal artfactory credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration

// Understanding: this function continuously monitors the providerconfig resource to fetch the credentials from the providerconfig and creates the terraform provider
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		if creds["url"] == "" || creds["access_token"] == "" {
			return ps, errors.New("missing required Artifactory credentials: url or access_token")
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{
			KeyURL:         creds["url"],
			KeyAccessToken: creds["access_token"],
		}
		return ps, nil
	}
}
