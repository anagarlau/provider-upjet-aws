// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActivationInitParameters struct {

	// The description of the resource that you want to register.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// UTC timestamp in RFC3339 format by which this activation request should expire. The default value is 24 hours from resource creation time.
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The IAM Role to attach to the managed instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	IAMRole *string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`

	// Reference to a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleRef *v1.Reference `json:"iamRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// The default name of the registered managed instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit *float64 `json:"registrationLimit,omitempty" tf:"registration_limit,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ActivationObservation struct {

	// The code the system generates when it processes the activation.
	ActivationCode *string `json:"activationCode,omitempty" tf:"activation_code,omitempty"`

	// The description of the resource that you want to register.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// UTC timestamp in RFC3339 format by which this activation request should expire. The default value is 24 hours from resource creation time.
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// If the current activation has expired.
	Expired *bool `json:"expired,omitempty" tf:"expired,omitempty"`

	// The IAM Role to attach to the managed instance.
	IAMRole *string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`

	// The activation ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The default name of the registered managed instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of managed instances that are currently registered using this activation.
	RegistrationCount *float64 `json:"registrationCount,omitempty" tf:"registration_count,omitempty"`

	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit *float64 `json:"registrationLimit,omitempty" tf:"registration_limit,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ActivationParameters struct {

	// The description of the resource that you want to register.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// UTC timestamp in RFC3339 format by which this activation request should expire. The default value is 24 hours from resource creation time.
	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The IAM Role to attach to the managed instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IAMRole *string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`

	// Reference to a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleRef *v1.Reference `json:"iamRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// The default name of the registered managed instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The maximum number of managed instances you want to register. The default value is 1 instance.
	// +kubebuilder:validation:Optional
	RegistrationLimit *float64 `json:"registrationLimit,omitempty" tf:"registration_limit,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ActivationSpec defines the desired state of Activation
type ActivationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActivationParameters `json:"forProvider"`
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
	InitProvider ActivationInitParameters `json:"initProvider,omitempty"`
}

// ActivationStatus defines the observed state of Activation.
type ActivationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActivationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Activation is the Schema for the Activations API. Registers an on-premises server or virtual machine with Amazon EC2 so that it can be managed using Run Command.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Activation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActivationSpec   `json:"spec"`
	Status            ActivationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActivationList contains a list of Activations
type ActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Activation `json:"items"`
}

// Repository type metadata.
var (
	Activation_Kind             = "Activation"
	Activation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Activation_Kind}.String()
	Activation_KindAPIVersion   = Activation_Kind + "." + CRDGroupVersion.String()
	Activation_GroupVersionKind = CRDGroupVersion.WithKind(Activation_Kind)
)

func init() {
	SchemeBuilder.Register(&Activation{}, &ActivationList{})
}
