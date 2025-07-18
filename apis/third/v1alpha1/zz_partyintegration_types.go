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

type PartyIntegrationInitParameters struct {

	// Your API Key.
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	ChannelName *string `json:"channelName,omitempty" tf:"channel_name,omitempty"`

	// Whether your cluster has Prometheus enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Your Microsoft Teams incoming webhook URL.
	MicrosoftTeamsWebhookURLSecretRef *v1.SecretKeySelector `json:"microsoftTeamsWebhookUrlSecretRef,omitempty" tf:"-"`

	// Your Prometheus password.
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The unique ID for the project to get all Third-Party service integrations
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// PagerDuty region that indicates the API Uniform Resource Locator (URL) to use, either "US" or "EU". PagerDuty will use "US" by default.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An optional field for your Routing Key.
	RoutingKeySecretRef *v1.SecretKeySelector `json:"routingKeySecretRef,omitempty" tf:"-"`

	// An optional field for your webhook secret.
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// Indicates which service discovery method is used, either file or http.
	ServiceDiscoverySecretRef *v1.SecretKeySelector `json:"serviceDiscoverySecretRef,omitempty" tf:"-"`

	// Your Service Key.
	ServiceKeySecretRef *v1.SecretKeySelector `json:"serviceKeySecretRef,omitempty" tf:"-"`

	TeamName *string `json:"teamName,omitempty" tf:"team_name,omitempty"`

	// Third-Party Integration Settings type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Your webhook URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Your Prometheus username.
	UserNameSecretRef *v1.SecretKeySelector `json:"userNameSecretRef,omitempty" tf:"-"`
}

type PartyIntegrationObservation struct {
	ChannelName *string `json:"channelName,omitempty" tf:"channel_name,omitempty"`

	// Whether your cluster has Prometheus enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Unique identifier of the integration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique ID for the project to get all Third-Party service integrations
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// PagerDuty region that indicates the API Uniform Resource Locator (URL) to use, either "US" or "EU". PagerDuty will use "US" by default.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	TeamName *string `json:"teamName,omitempty" tf:"team_name,omitempty"`

	// Third-Party Integration Settings type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Your webhook URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PartyIntegrationParameters struct {

	// Your API Key.
	// +kubebuilder:validation:Optional
	APIKeySecretRef *v1.SecretKeySelector `json:"apiKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ChannelName *string `json:"channelName,omitempty" tf:"channel_name,omitempty"`

	// Whether your cluster has Prometheus enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Your Microsoft Teams incoming webhook URL.
	// +kubebuilder:validation:Optional
	MicrosoftTeamsWebhookURLSecretRef *v1.SecretKeySelector `json:"microsoftTeamsWebhookUrlSecretRef,omitempty" tf:"-"`

	// Your Prometheus password.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The unique ID for the project to get all Third-Party service integrations
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// PagerDuty region that indicates the API Uniform Resource Locator (URL) to use, either "US" or "EU". PagerDuty will use "US" by default.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An optional field for your Routing Key.
	// +kubebuilder:validation:Optional
	RoutingKeySecretRef *v1.SecretKeySelector `json:"routingKeySecretRef,omitempty" tf:"-"`

	// An optional field for your webhook secret.
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// Indicates which service discovery method is used, either file or http.
	// +kubebuilder:validation:Optional
	ServiceDiscoverySecretRef *v1.SecretKeySelector `json:"serviceDiscoverySecretRef,omitempty" tf:"-"`

	// Your Service Key.
	// +kubebuilder:validation:Optional
	ServiceKeySecretRef *v1.SecretKeySelector `json:"serviceKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TeamName *string `json:"teamName,omitempty" tf:"team_name,omitempty"`

	// Third-Party Integration Settings type
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Your webhook URL.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Your Prometheus username.
	// +kubebuilder:validation:Optional
	UserNameSecretRef *v1.SecretKeySelector `json:"userNameSecretRef,omitempty" tf:"-"`
}

// PartyIntegrationSpec defines the desired state of PartyIntegration
type PartyIntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PartyIntegrationParameters `json:"forProvider"`
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
	InitProvider PartyIntegrationInitParameters `json:"initProvider,omitempty"`
}

// PartyIntegrationStatus defines the observed state of PartyIntegration.
type PartyIntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PartyIntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PartyIntegration is the Schema for the PartyIntegrations API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,atlas-mongodb}
type PartyIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   PartyIntegrationSpec   `json:"spec"`
	Status PartyIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PartyIntegrationList contains a list of PartyIntegrations
type PartyIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PartyIntegration `json:"items"`
}

// Repository type metadata.
var (
	PartyIntegration_Kind             = "PartyIntegration"
	PartyIntegration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PartyIntegration_Kind}.String()
	PartyIntegration_KindAPIVersion   = PartyIntegration_Kind + "." + CRDGroupVersion.String()
	PartyIntegration_GroupVersionKind = CRDGroupVersion.WithKind(PartyIntegration_Kind)
)

func init() {
	SchemeBuilder.Register(&PartyIntegration{}, &PartyIntegrationList{})
}
