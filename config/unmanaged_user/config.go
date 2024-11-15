package unmanaged_user

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_user", func(r *config.Resource) {
		r.ShortGroup = "unmanageduser"
	})
}
