// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SettingsOrgConfigInitParameters struct {

	// The collection of unique ids representing the identity providers that can be used for data access in this organization.
	DataAccessIdentityProviderIds []*string `json:"dataAccessIdentityProviderIds,omitempty" tf:"data_access_identity_provider_ids,omitempty"`

	// List that contains the approved domains from which organization users can log in.
	DomainAllowList []*string `json:"domainAllowList,omitempty" tf:"domain_allow_list,omitempty"`

	// Flag that indicates whether domain restriction is enabled for the connected organization.
	DomainRestrictionEnabled *bool `json:"domainRestrictionEnabled,omitempty" tf:"domain_restriction_enabled,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// Legacy 20-hexadecimal digit string that identifies the SAML access identity provider that this connected org config is associated with. Removing the attribute or providing the value "" will detach/remove the SAML identity provider. This id can be found in two ways:
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// List that contains the default roles granted to users who authenticate through the IdP in a connected organization.
	PostAuthRoleGrants []*string `json:"postAuthRoleGrants,omitempty" tf:"post_auth_role_grants,omitempty"`
}

type SettingsOrgConfigObservation struct {

	// The collection of unique ids representing the identity providers that can be used for data access in this organization.
	DataAccessIdentityProviderIds []*string `json:"dataAccessIdentityProviderIds,omitempty" tf:"data_access_identity_provider_ids,omitempty"`

	// List that contains the approved domains from which organization users can log in.
	DomainAllowList []*string `json:"domainAllowList,omitempty" tf:"domain_allow_list,omitempty"`

	// Flag that indicates whether domain restriction is enabled for the connected organization.
	DomainRestrictionEnabled *bool `json:"domainRestrictionEnabled,omitempty" tf:"domain_restriction_enabled,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Legacy 20-hexadecimal digit string that identifies the SAML access identity provider that this connected org config is associated with. Removing the attribute or providing the value "" will detach/remove the SAML identity provider. This id can be found in two ways:
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// List that contains the default roles granted to users who authenticate through the IdP in a connected organization.
	PostAuthRoleGrants []*string `json:"postAuthRoleGrants,omitempty" tf:"post_auth_role_grants,omitempty"`

	// List that contains the users who have an email address that doesn't match any domain on the allowed list. See below
	UserConflicts []UserConflictsObservation `json:"userConflicts,omitempty" tf:"user_conflicts,omitempty"`
}

type SettingsOrgConfigParameters struct {

	// The collection of unique ids representing the identity providers that can be used for data access in this organization.
	// +kubebuilder:validation:Optional
	DataAccessIdentityProviderIds []*string `json:"dataAccessIdentityProviderIds,omitempty" tf:"data_access_identity_provider_ids,omitempty"`

	// List that contains the approved domains from which organization users can log in.
	// +kubebuilder:validation:Optional
	DomainAllowList []*string `json:"domainAllowList,omitempty" tf:"domain_allow_list,omitempty"`

	// Flag that indicates whether domain restriction is enabled for the connected organization.
	// +kubebuilder:validation:Optional
	DomainRestrictionEnabled *bool `json:"domainRestrictionEnabled,omitempty" tf:"domain_restriction_enabled,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	// +kubebuilder:validation:Optional
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// Legacy 20-hexadecimal digit string that identifies the SAML access identity provider that this connected org config is associated with. Removing the attribute or providing the value "" will detach/remove the SAML identity provider. This id can be found in two ways:
	// +kubebuilder:validation:Optional
	IdentityProviderID *string `json:"identityProviderId,omitempty" tf:"identity_provider_id,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// List that contains the default roles granted to users who authenticate through the IdP in a connected organization.
	// +kubebuilder:validation:Optional
	PostAuthRoleGrants []*string `json:"postAuthRoleGrants,omitempty" tf:"post_auth_role_grants,omitempty"`
}

type UserConflictsInitParameters struct {
}

type UserConflictsObservation struct {

	// Email address of the the user that conflicts with selected domains.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsID *string `json:"federationSettingsId,omitempty" tf:"federation_settings_id,omitempty"`

	// First name of the the user that conflicts with selected domains.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// Last name of the the user that conflicts with selected domains.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// Name of the Atlas user that conflicts with selected domains.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserConflictsParameters struct {
}

// SettingsOrgConfigSpec defines the desired state of SettingsOrgConfig
type SettingsOrgConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingsOrgConfigParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SettingsOrgConfigInitParameters `json:"initProvider,omitempty"`
}

// SettingsOrgConfigStatus defines the observed state of SettingsOrgConfig.
type SettingsOrgConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingsOrgConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SettingsOrgConfig is the Schema for the SettingsOrgConfigs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,atlas-mongodb}
type SettingsOrgConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domainRestrictionEnabled) || (has(self.initProvider) && has(self.initProvider.domainRestrictionEnabled))",message="spec.forProvider.domainRestrictionEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.federationSettingsId) || (has(self.initProvider) && has(self.initProvider.federationSettingsId))",message="spec.forProvider.federationSettingsId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.orgId) || (has(self.initProvider) && has(self.initProvider.orgId))",message="spec.forProvider.orgId is a required parameter"
	Spec   SettingsOrgConfigSpec   `json:"spec"`
	Status SettingsOrgConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingsOrgConfigList contains a list of SettingsOrgConfigs
type SettingsOrgConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SettingsOrgConfig `json:"items"`
}

// Repository type metadata.
var (
	SettingsOrgConfig_Kind             = "SettingsOrgConfig"
	SettingsOrgConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SettingsOrgConfig_Kind}.String()
	SettingsOrgConfig_KindAPIVersion   = SettingsOrgConfig_Kind + "." + CRDGroupVersion.String()
	SettingsOrgConfig_GroupVersionKind = CRDGroupVersion.WithKind(SettingsOrgConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&SettingsOrgConfig{}, &SettingsOrgConfigList{})
}
