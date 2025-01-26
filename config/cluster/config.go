package cluster

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_cluster", func(r *config.Resource) {
		r.ShortGroup = "cluster"
		r.ExternalName = config.NameAsIdentifier
		// Upcoming deprecation - use the mongodbatlas_project_ip_addresses data source instead
		delete(r.TerraformResource.Schema, "ip_addresses")
	})
}
