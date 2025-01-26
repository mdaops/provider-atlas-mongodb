package advancedcluster

import "github.com/crossplane/upjet/pkg/config"

// Configure "advanced_cluster" resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_advanced_cluster", func(r *config.Resource) {
		r.ShortGroup = "advancedcluster"
		r.References["project_id"] = config.Reference{
			Type:              "github.com/mdaops/provider-atlas-mongodb/apis/project/v1alpha.Project",
			RefFieldName:      "ProjectRef",
			SelectorFieldName: "ProjectSelector",
		}
	})
}
