package e2e

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/magefile/mage/sh"
)

const (
	ARTIFACTORY_NAMESPACE    = "jfrog"
	ARTIFACTORY_HELM_RELEASE = "artifactory"
)

func KindClusterExists(name string) (bool, error) {
	output, err := sh.Output("kind", "get", "clusters")

	if err != nil {
		return false, err
	}

	// split the output into an array of lines
	clusters := strings.Split(output, "\n")

	// check if the cluster name is in the list of clusters
	for _, cluster := range clusters {
		if cluster == name {
			return true, nil
		}
	}

	return false, nil
}

func CreateKindCluster(name string) error {
	err := sh.RunV("kind", "create", "cluster", "--name", name)

	if err != nil {
		return err
	}

	return nil
}

func EnsureKindCluster(name string) error {
	exists, err := KindClusterExists(name)

	if err != nil {
		return err
	}

	if exists {
		fmt.Printf("Kind cluster %s already exists.\n", name)
		return nil
	}

	return CreateKindCluster(name)
}

func InstallArtifactory() error {
	const POD_CREATION_TIMEOUT = 30 * time.Second
	const POD_READY_TIMEOUT = 5 * time.Minute
	start := time.Now()

	fmt.Println("Installing Artifactory...")
	err := sh.Run(
		"helm",
		"install",
		ARTIFACTORY_HELM_RELEASE,
		"--create-namespace",
		"-n",
		ARTIFACTORY_NAMESPACE,
		"--set",
		"xray.enabled=false",
		"--set",
		os.ExpandEnv("artifactory.artifactory.license.licenseKey=${ARTIFACTORY_LICENSE_KEY}"),
		"--set",
		"artifactory.access.accessConfig.token.allow-basic-auth=true",
		"jfrog/jfrog-platform",
	)

	if err != nil {
		return err
	}

	fmt.Println("Waiting for Artifactory to be ready...")
	for {
		outsb := strings.Builder{}
		errsb := strings.Builder{}

		_, err := sh.Exec(
			nil,
			&outsb,
			&errsb,
			"kubectl",
			"wait",
			"--for=condition=Ready",
			"-n",
			ARTIFACTORY_NAMESPACE,
			"pod/"+ARTIFACTORY_HELM_RELEASE+"-0",
		)

		if err != nil {
			if strings.Contains(errsb.String(), "pods \""+ARTIFACTORY_HELM_RELEASE+"-0\" not found") {
				if time.Since(start) < POD_CREATION_TIMEOUT {
					continue
				}

				return fmt.Errorf("artifactory pod not found after %v; will not retry", POD_CREATION_TIMEOUT)
			}

			if strings.Contains(errsb.String(), "timed out waiting for the condition on pods/"+ARTIFACTORY_HELM_RELEASE+"-0") {
				if time.Since(start) < POD_READY_TIMEOUT {
					continue
				}

				return fmt.Errorf("artifactory pod not ready after %v; will not retry", POD_READY_TIMEOUT)
			}

			fmt.Printf("Unhandled error: %s\n", err.Error())
			fmt.Printf("Standard output: %s\n", outsb.String())
			fmt.Printf("Error output: %s\n", errsb.String())

			return err
		}

		fmt.Println("Artifactory is ready.")
		return nil
	}
}

func ArtifactoryExists() (bool, error) {
	output, err := sh.Output("helm", "list", "-n", ARTIFACTORY_NAMESPACE)

	if err != nil {
		return false, err
	}

	// split the output into an array of lines
	releaseInfos := strings.Split(output, "\n")

	// check if the release name is in the list of releases
	for _, releaseInfo := range releaseInfos {
		if strings.Contains(releaseInfo, ARTIFACTORY_HELM_RELEASE) {
			return true, nil
		}
	}

	return false, nil
}

func EnsureArtifactory() error {
	exists, err := ArtifactoryExists()

	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Artifactory is already installed.")
		return nil
	}

	return InstallArtifactory()
}

func UpdateCredentials() error {
	outsb := strings.Builder{}
	errsb := strings.Builder{}
	_, err := sh.Exec(nil, &outsb, &errsb, "kubectl", "delete", "job", "create-credentials")

	if err != nil {
		if !strings.Contains(errsb.String(), "jobs.batch \"create-credentials\" not found") {
			fmt.Printf("Unhandled error: %s\n", err.Error())
			fmt.Printf("Standard output: %s\n", outsb.String())
			fmt.Printf("Error output: %s\n", errsb.String())

			return err
		}
	}

	outsb.Reset()
	errsb.Reset()
	_, err = sh.Exec(nil, &outsb, &errsb, "kubectl", "apply", "-f", "e2e/create-credentials.yaml")

	if err != nil {
		fmt.Printf("Unhandled error: %s\n", err.Error())
		fmt.Printf("Standard output: %s\n", outsb.String())
		fmt.Printf("Error output: %s\n", errsb.String())

		return err
	}

	return nil
}
