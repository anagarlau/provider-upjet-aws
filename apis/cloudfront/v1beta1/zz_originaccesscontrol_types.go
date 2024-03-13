// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OriginAccessControlInitParameters struct {

	// The description of the Origin Access Control.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name that identifies the Origin Access Control.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of origin that this Origin Access Control is for. Valid values are s3, and mediastore.
	OriginAccessControlOriginType *string `json:"originAccessControlOriginType,omitempty" tf:"origin_access_control_origin_type,omitempty"`

	// Specifies which requests CloudFront signs. Specify always for the most common use case. Allowed values: always, never, and no-override.
	SigningBehavior *string `json:"signingBehavior,omitempty" tf:"signing_behavior,omitempty"`

	// Determines how CloudFront signs (authenticates) requests. The only valid value is sigv4.
	SigningProtocol *string `json:"signingProtocol,omitempty" tf:"signing_protocol,omitempty"`
}

type OriginAccessControlObservation struct {

	// The description of the Origin Access Control.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The current version of this Origin Access Control.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The unique identifier of this Origin Access Control.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A name that identifies the Origin Access Control.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of origin that this Origin Access Control is for. Valid values are s3, and mediastore.
	OriginAccessControlOriginType *string `json:"originAccessControlOriginType,omitempty" tf:"origin_access_control_origin_type,omitempty"`

	// Specifies which requests CloudFront signs. Specify always for the most common use case. Allowed values: always, never, and no-override.
	SigningBehavior *string `json:"signingBehavior,omitempty" tf:"signing_behavior,omitempty"`

	// Determines how CloudFront signs (authenticates) requests. The only valid value is sigv4.
	SigningProtocol *string `json:"signingProtocol,omitempty" tf:"signing_protocol,omitempty"`
}

type OriginAccessControlParameters struct {

	// The description of the Origin Access Control.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name that identifies the Origin Access Control.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of origin that this Origin Access Control is for. Valid values are s3, and mediastore.
	// +kubebuilder:validation:Optional
	OriginAccessControlOriginType *string `json:"originAccessControlOriginType,omitempty" tf:"origin_access_control_origin_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies which requests CloudFront signs. Specify always for the most common use case. Allowed values: always, never, and no-override.
	// +kubebuilder:validation:Optional
	SigningBehavior *string `json:"signingBehavior,omitempty" tf:"signing_behavior,omitempty"`

	// Determines how CloudFront signs (authenticates) requests. The only valid value is sigv4.
	// +kubebuilder:validation:Optional
	SigningProtocol *string `json:"signingProtocol,omitempty" tf:"signing_protocol,omitempty"`
}

// OriginAccessControlSpec defines the desired state of OriginAccessControl
type OriginAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginAccessControlParameters `json:"forProvider"`
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
	InitProvider OriginAccessControlInitParameters `json:"initProvider,omitempty"`
}

// OriginAccessControlStatus defines the observed state of OriginAccessControl.
type OriginAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OriginAccessControl is the Schema for the OriginAccessControls API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OriginAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originAccessControlOriginType) || (has(self.initProvider) && has(self.initProvider.originAccessControlOriginType))",message="spec.forProvider.originAccessControlOriginType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.signingBehavior) || (has(self.initProvider) && has(self.initProvider.signingBehavior))",message="spec.forProvider.signingBehavior is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.signingProtocol) || (has(self.initProvider) && has(self.initProvider.signingProtocol))",message="spec.forProvider.signingProtocol is a required parameter"
	Spec   OriginAccessControlSpec   `json:"spec"`
	Status OriginAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessControlList contains a list of OriginAccessControls
type OriginAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginAccessControl `json:"items"`
}

// Repository type metadata.
var (
	OriginAccessControl_Kind             = "OriginAccessControl"
	OriginAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginAccessControl_Kind}.String()
	OriginAccessControl_KindAPIVersion   = OriginAccessControl_Kind + "." + CRDGroupVersion.String()
	OriginAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(OriginAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginAccessControl{}, &OriginAccessControlList{})
}
