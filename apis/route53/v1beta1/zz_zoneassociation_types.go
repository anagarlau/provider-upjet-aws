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

type ZoneAssociationInitParameters struct {

	// The VPC to associate with the private hosted zone.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The VPC's region. Defaults to the region of the AWS provider.
	VPCRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`

	// The private hosted zone to associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.Zone
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("zone_id",true)
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type ZoneAssociationObservation struct {

	// The calculated unique identifier for the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The account ID of the account that created the hosted zone.
	OwningAccount *string `json:"owningAccount,omitempty" tf:"owning_account,omitempty"`

	// The VPC to associate with the private hosted zone.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The VPC's region. Defaults to the region of the AWS provider.
	VPCRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`

	// The private hosted zone to associate.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ZoneAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The VPC to associate with the private hosted zone.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The VPC's region. Defaults to the region of the AWS provider.
	// +kubebuilder:validation:Optional
	VPCRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`

	// The private hosted zone to associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.Zone
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("zone_id",true)
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// ZoneAssociationSpec defines the desired state of ZoneAssociation
type ZoneAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneAssociationParameters `json:"forProvider"`
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
	InitProvider ZoneAssociationInitParameters `json:"initProvider,omitempty"`
}

// ZoneAssociationStatus defines the observed state of ZoneAssociation.
type ZoneAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ZoneAssociation is the Schema for the ZoneAssociations API. Manages a Route53 Hosted Zone VPC association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ZoneAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneAssociationSpec   `json:"spec"`
	Status            ZoneAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneAssociationList contains a list of ZoneAssociations
type ZoneAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneAssociation `json:"items"`
}

// Repository type metadata.
var (
	ZoneAssociation_Kind             = "ZoneAssociation"
	ZoneAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ZoneAssociation_Kind}.String()
	ZoneAssociation_KindAPIVersion   = ZoneAssociation_Kind + "." + CRDGroupVersion.String()
	ZoneAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ZoneAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ZoneAssociation{}, &ZoneAssociationList{})
}
