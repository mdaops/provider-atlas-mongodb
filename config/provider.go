/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

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
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		project.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
