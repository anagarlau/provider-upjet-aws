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

type TagOptionInitParameters struct {

	// Whether tag option is active. Default is true.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Tag option key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Tag option value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagOptionObservation struct {

	// Whether tag option is active. Default is true.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Identifier (e.g., tag-pjtvagohlyo3m).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Tag option key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Tag option value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagOptionParameters struct {

	// Whether tag option is active. Default is true.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Tag option key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Tag option value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// TagOptionSpec defines the desired state of TagOption
type TagOptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagOptionParameters `json:"forProvider"`
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
	InitProvider TagOptionInitParameters `json:"initProvider,omitempty"`
}

// TagOptionStatus defines the observed state of TagOption.
type TagOptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagOptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TagOption is the Schema for the TagOptions API. Manages a Service Catalog Tag Option
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TagOption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   TagOptionSpec   `json:"spec"`
	Status TagOptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagOptionList contains a list of TagOptions
type TagOptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagOption `json:"items"`
}

// Repository type metadata.
var (
	TagOption_Kind             = "TagOption"
	TagOption_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TagOption_Kind}.String()
	TagOption_KindAPIVersion   = TagOption_Kind + "." + CRDGroupVersion.String()
	TagOption_GroupVersionKind = CRDGroupVersion.WithKind(TagOption_Kind)
)

func init() {
	SchemeBuilder.Register(&TagOption{}, &TagOptionList{})
}
