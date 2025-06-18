// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	listapikey "github.com/mdaops/provider-atlas-mongodb/internal/controller/access/listapikey"
	cluster "github.com/mdaops/provider-atlas-mongodb/internal/controller/advancedcluster/cluster"
	configuration "github.com/mdaops/provider-atlas-mongodb/internal/controller/alert/configuration"
	key "github.com/mdaops/provider-atlas-mongodb/internal/controller/api/key"
	compliancepolicy "github.com/mdaops/provider-atlas-mongodb/internal/controller/backup/compliancepolicy"
	backupschedule "github.com/mdaops/provider-atlas-mongodb/internal/controller/cloud/backupschedule"
	backupsnapshot "github.com/mdaops/provider-atlas-mongodb/internal/controller/cloud/backupsnapshot"
	backupsnapshotexportbucket "github.com/mdaops/provider-atlas-mongodb/internal/controller/cloud/backupsnapshotexportbucket"
	backupsnapshotexportjob "github.com/mdaops/provider-atlas-mongodb/internal/controller/cloud/backupsnapshotexportjob"
	backupsnapshotrestorejob "github.com/mdaops/provider-atlas-mongodb/internal/controller/cloud/backupsnapshotrestorejob"
	provideraccessauthorization "github.com/mdaops/provider-atlas-mongodb/internal/controller/cloud/provideraccessauthorization"
	provideraccesssetup "github.com/mdaops/provider-atlas-mongodb/internal/controller/cloud/provideraccesssetup"
	outagesimulation "github.com/mdaops/provider-atlas-mongodb/internal/controller/cluster/outagesimulation"
	dbrole "github.com/mdaops/provider-atlas-mongodb/internal/controller/custom/dbrole"
	dnsconfigurationclusteraws "github.com/mdaops/provider-atlas-mongodb/internal/controller/custom/dnsconfigurationclusteraws"
	lakepipeline "github.com/mdaops/provider-atlas-mongodb/internal/controller/data/lakepipeline"
	user "github.com/mdaops/provider-atlas-mongodb/internal/controller/database/user"
	atrest "github.com/mdaops/provider-atlas-mongodb/internal/controller/encryption/atrest"
	atrestprivateendpoint "github.com/mdaops/provider-atlas-mongodb/internal/controller/encryption/atrestprivateendpoint"
	trigger "github.com/mdaops/provider-atlas-mongodb/internal/controller/event/trigger"
	databaseinstance "github.com/mdaops/provider-atlas-mongodb/internal/controller/federated/databaseinstance"
	querylimit "github.com/mdaops/provider-atlas-mongodb/internal/controller/federated/querylimit"
	settingsidentityprovider "github.com/mdaops/provider-atlas-mongodb/internal/controller/federated/settingsidentityprovider"
	settingsorgconfig "github.com/mdaops/provider-atlas-mongodb/internal/controller/federated/settingsorgconfig"
	settingsorgrolemapping "github.com/mdaops/provider-atlas-mongodb/internal/controller/federated/settingsorgrolemapping"
	clusterconfig "github.com/mdaops/provider-atlas-mongodb/internal/controller/global/clusterconfig"
	configurationldap "github.com/mdaops/provider-atlas-mongodb/internal/controller/ldap/configuration"
	verify "github.com/mdaops/provider-atlas-mongodb/internal/controller/ldap/verify"
	window "github.com/mdaops/provider-atlas-mongodb/internal/controller/maintenance/window"
	employeeaccessgrant "github.com/mdaops/provider-atlas-mongodb/internal/controller/mongodb/employeeaccessgrant"
	auditing "github.com/mdaops/provider-atlas-mongodb/internal/controller/mongodbatlas/auditing"
	clustermongodbatlas "github.com/mdaops/provider-atlas-mongodb/internal/controller/mongodbatlas/cluster"
	organization "github.com/mdaops/provider-atlas-mongodb/internal/controller/mongodbatlas/organization"
	team "github.com/mdaops/provider-atlas-mongodb/internal/controller/mongodbatlas/team"
	teams "github.com/mdaops/provider-atlas-mongodb/internal/controller/mongodbatlas/teams"
	container "github.com/mdaops/provider-atlas-mongodb/internal/controller/network/container"
	peering "github.com/mdaops/provider-atlas-mongodb/internal/controller/network/peering"
	archive "github.com/mdaops/provider-atlas-mongodb/internal/controller/online/archive"
	invitation "github.com/mdaops/provider-atlas-mongodb/internal/controller/org/invitation"
	endpointregionalmode "github.com/mdaops/provider-atlas-mongodb/internal/controller/private/endpointregionalmode"
	endpoint "github.com/mdaops/provider-atlas-mongodb/internal/controller/privatelink/endpoint"
	endpointserverless "github.com/mdaops/provider-atlas-mongodb/internal/controller/privatelink/endpointserverless"
	endpointservice "github.com/mdaops/provider-atlas-mongodb/internal/controller/privatelink/endpointservice"
	endpointservicedatafederationonlinearchive "github.com/mdaops/provider-atlas-mongodb/internal/controller/privatelink/endpointservicedatafederationonlinearchive"
	endpointserviceserverless "github.com/mdaops/provider-atlas-mongodb/internal/controller/privatelink/endpointserviceserverless"
	apikey "github.com/mdaops/provider-atlas-mongodb/internal/controller/project/apikey"
	invitationproject "github.com/mdaops/provider-atlas-mongodb/internal/controller/project/invitation"
	project "github.com/mdaops/provider-atlas-mongodb/internal/controller/project/project"
	providerconfig "github.com/mdaops/provider-atlas-mongodb/internal/controller/providerconfig"
	index "github.com/mdaops/provider-atlas-mongodb/internal/controller/search/index"
	instance "github.com/mdaops/provider-atlas-mongodb/internal/controller/serverless/instance"
	privatelinkendpoint "github.com/mdaops/provider-atlas-mongodb/internal/controller/stream/privatelinkendpoint"
	partyintegration "github.com/mdaops/provider-atlas-mongodb/internal/controller/third/partyintegration"
	authenticationdatabaseuser "github.com/mdaops/provider-atlas-mongodb/internal/controller/x509/authenticationdatabaseuser"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		listapikey.Setup,
		cluster.Setup,
		configuration.Setup,
		key.Setup,
		compliancepolicy.Setup,
		backupschedule.Setup,
		backupsnapshot.Setup,
		backupsnapshotexportbucket.Setup,
		backupsnapshotexportjob.Setup,
		backupsnapshotrestorejob.Setup,
		provideraccessauthorization.Setup,
		provideraccesssetup.Setup,
		outagesimulation.Setup,
		dbrole.Setup,
		dnsconfigurationclusteraws.Setup,
		lakepipeline.Setup,
		user.Setup,
		atrest.Setup,
		atrestprivateendpoint.Setup,
		trigger.Setup,
		databaseinstance.Setup,
		querylimit.Setup,
		settingsidentityprovider.Setup,
		settingsorgconfig.Setup,
		settingsorgrolemapping.Setup,
		clusterconfig.Setup,
		configurationldap.Setup,
		verify.Setup,
		window.Setup,
		employeeaccessgrant.Setup,
		auditing.Setup,
		clustermongodbatlas.Setup,
		organization.Setup,
		team.Setup,
		teams.Setup,
		container.Setup,
		peering.Setup,
		archive.Setup,
		invitation.Setup,
		endpointregionalmode.Setup,
		endpoint.Setup,
		endpointserverless.Setup,
		endpointservice.Setup,
		endpointservicedatafederationonlinearchive.Setup,
		endpointserviceserverless.Setup,
		apikey.Setup,
		invitationproject.Setup,
		project.Setup,
		providerconfig.Setup,
		index.Setup,
		instance.Setup,
		privatelinkendpoint.Setup,
		partyintegration.Setup,
		authenticationdatabaseuser.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
