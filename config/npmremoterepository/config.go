package npmremoterepository

import "github.com/crossplane/upjet/pkg/config"

// Configure the "artifactory_remote_npm_repository" resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_remote_npm_repository", func(r *config.Resource) {
		r.ShortGroup = "repository"
	})
}
