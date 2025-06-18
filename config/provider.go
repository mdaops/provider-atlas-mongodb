/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	advancedcluster "github.com/mdaops/provider-atlas-mongodb/config/advanced-cluster"
	alertconfiguration "github.com/mdaops/provider-atlas-mongodb/config/alert-configuration"
	"github.com/mdaops/provider-atlas-mongodb/config/project"
)

const (
	resourcePrefix = "atlas-mongodb"
	modulePath     = "github.com/mdaops/provider-atlas-mongodb"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("mongoatlas.io"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(),
		ujconfig.WithSkipList([]string{
			"mongodbatlas_project_ip_access_list",
			"mongodbatlas_push_based_log_export",
			"mongodbatlas_stream_connection",
			"mongodbatlas_stream_instance",
			"mongodbatlas_stream_processor",
			"mongodbatlas_search_deployment",
		}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		project.Configure,
		advancedcluster.Configure,
		alertconfiguration.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
