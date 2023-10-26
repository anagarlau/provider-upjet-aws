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

type PrincipalPortfolioAssociationInitParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Principal type. Setting this argument empty (e.g., principal_type = "") will result in an error. Valid value is IAM. Default is IAM.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`
}

type PrincipalPortfolioAssociationObservation struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Identifier of the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Portfolio identifier.
	PortfolioID *string `json:"portfolioId,omitempty" tf:"portfolio_id,omitempty"`

	// Principal ARN.
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`

	// Principal type. Setting this argument empty (e.g., principal_type = "") will result in an error. Valid value is IAM. Default is IAM.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`
}

type PrincipalPortfolioAssociationParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Portfolio identifier.
	// +crossplane:generate:reference:type=Portfolio
	// +kubebuilder:validation:Optional
	PortfolioID *string `json:"portfolioId,omitempty" tf:"portfolio_id,omitempty"`

	// Reference to a Portfolio to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDRef *v1.Reference `json:"portfolioIdRef,omitempty" tf:"-"`

	// Selector for a Portfolio to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDSelector *v1.Selector `json:"portfolioIdSelector,omitempty" tf:"-"`

	// Principal ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`

	// Reference to a User in iam to populate principalArn.
	// +kubebuilder:validation:Optional
	PrincipalArnRef *v1.Reference `json:"principalArnRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate principalArn.
	// +kubebuilder:validation:Optional
	PrincipalArnSelector *v1.Selector `json:"principalArnSelector,omitempty" tf:"-"`

	// Principal type. Setting this argument empty (e.g., principal_type = "") will result in an error. Valid value is IAM. Default is IAM.
	// +kubebuilder:validation:Optional
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PrincipalPortfolioAssociationSpec defines the desired state of PrincipalPortfolioAssociation
type PrincipalPortfolioAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrincipalPortfolioAssociationParameters `json:"forProvider"`
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
	InitProvider PrincipalPortfolioAssociationInitParameters `json:"initProvider,omitempty"`
}

// PrincipalPortfolioAssociationStatus defines the observed state of PrincipalPortfolioAssociation.
type PrincipalPortfolioAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrincipalPortfolioAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrincipalPortfolioAssociation is the Schema for the PrincipalPortfolioAssociations API. Manages a Service Catalog Principal Portfolio Association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PrincipalPortfolioAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.acceptLanguage) || (has(self.initProvider) && has(self.initProvider.acceptLanguage))",message="spec.forProvider.acceptLanguage is a required parameter"
	Spec   PrincipalPortfolioAssociationSpec   `json:"spec"`
	Status PrincipalPortfolioAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrincipalPortfolioAssociationList contains a list of PrincipalPortfolioAssociations
type PrincipalPortfolioAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrincipalPortfolioAssociation `json:"items"`
}

// Repository type metadata.
var (
	PrincipalPortfolioAssociation_Kind             = "PrincipalPortfolioAssociation"
	PrincipalPortfolioAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrincipalPortfolioAssociation_Kind}.String()
	PrincipalPortfolioAssociation_KindAPIVersion   = PrincipalPortfolioAssociation_Kind + "." + CRDGroupVersion.String()
	PrincipalPortfolioAssociation_GroupVersionKind = CRDGroupVersion.WithKind(PrincipalPortfolioAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&PrincipalPortfolioAssociation{}, &PrincipalPortfolioAssociationList{})
}
