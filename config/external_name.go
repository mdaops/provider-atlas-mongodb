package config

import "github.com/crossplane/upjet/pkg/config"

var ExternalNameConfigs = map[string]config.ExternalName{
	"mongodbatlas_project":             config.IdentifierFromProvider,
	"mongodbatlas_advanced_cluster":    config.IdentifierFromProvider,
	"mongodbatlas_project_api_key":     config.IdentifierFromProvider,
	"mongodbatlas_alert_configuration": config.IdentifierFromProvider,
}

func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		l[i] = name + "$"
		i++
	}
	return l
}
