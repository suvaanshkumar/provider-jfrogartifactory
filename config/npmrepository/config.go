package npmrepository

import "github.com/crossplane/upjet/pkg/config"

// Configure the "artifactory_npm_repository" resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_npm_repository", func(r *config.Resource) {
		r.ShortGroup = "repository"
	})
}
