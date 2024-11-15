package artifactory_local_generic_repository

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_generic_repository", func(r *config.Resource) {
		r.ShortGroup = "repository"
	})
}
