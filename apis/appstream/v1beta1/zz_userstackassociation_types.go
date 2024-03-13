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

type UserStackAssociationInitParameters struct {

	// Whether a welcome email is sent to a user after the user is created in the user pool.
	SendEmailNotification *bool `json:"sendEmailNotification,omitempty" tf:"send_email_notification,omitempty"`
}

type UserStackAssociationObservation struct {

	// Authentication type for the user.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Unique ID of the appstream User Stack association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether a welcome email is sent to a user after the user is created in the user pool.
	SendEmailNotification *bool `json:"sendEmailNotification,omitempty" tf:"send_email_notification,omitempty"`

	// Name of the stack that is associated with the user.
	StackName *string `json:"stackName,omitempty" tf:"stack_name,omitempty"`

	// Email address of the user who is associated with the stack.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UserStackAssociationParameters struct {

	// Authentication type for the user.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appstream/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("authentication_type",false)
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Reference to a User in appstream to populate authenticationType.
	// +kubebuilder:validation:Optional
	AuthenticationTypeRef *v1.Reference `json:"authenticationTypeRef,omitempty" tf:"-"`

	// Selector for a User in appstream to populate authenticationType.
	// +kubebuilder:validation:Optional
	AuthenticationTypeSelector *v1.Selector `json:"authenticationTypeSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether a welcome email is sent to a user after the user is created in the user pool.
	// +kubebuilder:validation:Optional
	SendEmailNotification *bool `json:"sendEmailNotification,omitempty" tf:"send_email_notification,omitempty"`

	// Name of the stack that is associated with the user.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appstream/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	StackName *string `json:"stackName,omitempty" tf:"stack_name,omitempty"`

	// Reference to a Stack in appstream to populate stackName.
	// +kubebuilder:validation:Optional
	StackNameRef *v1.Reference `json:"stackNameRef,omitempty" tf:"-"`

	// Selector for a Stack in appstream to populate stackName.
	// +kubebuilder:validation:Optional
	StackNameSelector *v1.Selector `json:"stackNameSelector,omitempty" tf:"-"`

	// Email address of the user who is associated with the stack.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appstream/v1beta1.User
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Reference to a User in appstream to populate userName.
	// +kubebuilder:validation:Optional
	UserNameRef *v1.Reference `json:"userNameRef,omitempty" tf:"-"`

	// Selector for a User in appstream to populate userName.
	// +kubebuilder:validation:Optional
	UserNameSelector *v1.Selector `json:"userNameSelector,omitempty" tf:"-"`
}

// UserStackAssociationSpec defines the desired state of UserStackAssociation
type UserStackAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserStackAssociationParameters `json:"forProvider"`
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
	InitProvider UserStackAssociationInitParameters `json:"initProvider,omitempty"`
}

// UserStackAssociationStatus defines the observed state of UserStackAssociation.
type UserStackAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserStackAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserStackAssociation is the Schema for the UserStackAssociations API. Manages an AppStream User Stack association.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserStackAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserStackAssociationSpec   `json:"spec"`
	Status            UserStackAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserStackAssociationList contains a list of UserStackAssociations
type UserStackAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserStackAssociation `json:"items"`
}

// Repository type metadata.
var (
	UserStackAssociation_Kind             = "UserStackAssociation"
	UserStackAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserStackAssociation_Kind}.String()
	UserStackAssociation_KindAPIVersion   = UserStackAssociation_Kind + "." + CRDGroupVersion.String()
	UserStackAssociation_GroupVersionKind = CRDGroupVersion.WithKind(UserStackAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&UserStackAssociation{}, &UserStackAssociationList{})
}
