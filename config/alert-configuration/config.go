package alertconfiguration

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_alert_configuration", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
	})
}
