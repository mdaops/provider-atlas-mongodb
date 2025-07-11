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

type AuthenticationDatabaseUserInitParameters struct {

	// PEM string containing one or more customer CAs for database user authentication.
	CustomerX509CasSecretRef *v1.SecretKeySelector `json:"customerX509CasSecretRef,omitempty" tf:"-"`

	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	MonthsUntilExpiration *float64 `json:"monthsUntilExpiration,omitempty" tf:"months_until_expiration,omitempty"`

	// Identifier for the Atlas project associated with the X.509 configuration.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Username of the database user to create a certificate for.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthenticationDatabaseUserObservation struct {

	// Array of objects where each details one unexpired database user certificate.
	Certificates []CertificatesObservation `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Serial number of this certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	MonthsUntilExpiration *float64 `json:"monthsUntilExpiration,omitempty" tf:"months_until_expiration,omitempty"`

	// Identifier for the Atlas project associated with the X.509 configuration.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Username of the database user to create a certificate for.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthenticationDatabaseUserParameters struct {

	// PEM string containing one or more customer CAs for database user authentication.
	// +kubebuilder:validation:Optional
	CustomerX509CasSecretRef *v1.SecretKeySelector `json:"customerX509CasSecretRef,omitempty" tf:"-"`

	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	// +kubebuilder:validation:Optional
	MonthsUntilExpiration *float64 `json:"monthsUntilExpiration,omitempty" tf:"months_until_expiration,omitempty"`

	// Identifier for the Atlas project associated with the X.509 configuration.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Username of the database user to create a certificate for.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CertificatesInitParameters struct {
}

type CertificatesObservation struct {

	// Timestamp in ISO 8601 date and time format in UTC when Atlas created this X.509 certificate.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Unique identifier of the Atlas project to which this certificate belongs.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Serial number of this certificate.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Timestamp in ISO 8601 date and time format in UTC when this certificate expires.
	NotAfter *string `json:"notAfter,omitempty" tf:"not_after,omitempty"`

	// Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type CertificatesParameters struct {
}

// AuthenticationDatabaseUserSpec defines the desired state of AuthenticationDatabaseUser
type AuthenticationDatabaseUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthenticationDatabaseUserParameters `json:"forProvider"`
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
	InitProvider AuthenticationDatabaseUserInitParameters `json:"initProvider,omitempty"`
}

// AuthenticationDatabaseUserStatus defines the observed state of AuthenticationDatabaseUser.
type AuthenticationDatabaseUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthenticationDatabaseUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthenticationDatabaseUser is the Schema for the AuthenticationDatabaseUsers API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,atlas-mongodb}
type AuthenticationDatabaseUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	Spec   AuthenticationDatabaseUserSpec   `json:"spec"`
	Status AuthenticationDatabaseUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthenticationDatabaseUserList contains a list of AuthenticationDatabaseUsers
type AuthenticationDatabaseUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthenticationDatabaseUser `json:"items"`
}

// Repository type metadata.
var (
	AuthenticationDatabaseUser_Kind             = "AuthenticationDatabaseUser"
	AuthenticationDatabaseUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthenticationDatabaseUser_Kind}.String()
	AuthenticationDatabaseUser_KindAPIVersion   = AuthenticationDatabaseUser_Kind + "." + CRDGroupVersion.String()
	AuthenticationDatabaseUser_GroupVersionKind = CRDGroupVersion.WithKind(AuthenticationDatabaseUser_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthenticationDatabaseUser{}, &AuthenticationDatabaseUserList{})
}
