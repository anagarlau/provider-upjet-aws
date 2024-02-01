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

type StaticIPInitParameters struct {

	// The name for the allocated static IP
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StaticIPObservation struct {

	// The ARN of the Lightsail static IP
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The allocated static IP address
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name for the allocated static IP
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The support code.
	SupportCode *string `json:"supportCode,omitempty" tf:"support_code,omitempty"`
}

type StaticIPParameters struct {

	// The name for the allocated static IP
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// StaticIPSpec defines the desired state of StaticIP
type StaticIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StaticIPParameters `json:"forProvider"`
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
	InitProvider StaticIPInitParameters `json:"initProvider,omitempty"`
}

// StaticIPStatus defines the observed state of StaticIP.
type StaticIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StaticIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StaticIP is the Schema for the StaticIPs API. Provides an Lightsail Static IP
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StaticIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   StaticIPSpec   `json:"spec"`
	Status StaticIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StaticIPList contains a list of StaticIPs
type StaticIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticIP `json:"items"`
}

// Repository type metadata.
var (
	StaticIP_Kind             = "StaticIP"
	StaticIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StaticIP_Kind}.String()
	StaticIP_KindAPIVersion   = StaticIP_Kind + "." + CRDGroupVersion.String()
	StaticIP_GroupVersionKind = CRDGroupVersion.WithKind(StaticIP_Kind)
)

func init() {
	SchemeBuilder.Register(&StaticIP{}, &StaticIPList{})
}
