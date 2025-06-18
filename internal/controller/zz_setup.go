// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/mdaops/provider-atlas-mongodb/internal/controller/advanced/cluster"
	configuration "github.com/mdaops/provider-atlas-mongodb/internal/controller/alert/configuration"
	apikey "github.com/mdaops/provider-atlas-mongodb/internal/controller/project/apikey"
	project "github.com/mdaops/provider-atlas-mongodb/internal/controller/project/project"
	providerconfig "github.com/mdaops/provider-atlas-mongodb/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		configuration.Setup,
		apikey.Setup,
		project.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
