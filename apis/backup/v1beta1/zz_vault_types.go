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

type VaultInitParameters struct {

	// A boolean that indicates that all recovery points stored in the vault are deleted so that the vault can be destroyed without error.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The server-side encryption key that is used to protect your backups.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VaultObservation struct {

	// The ARN of the vault.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A boolean that indicates that all recovery points stored in the vault are deleted so that the vault can be destroyed without error.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The name of the vault.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The server-side encryption key that is used to protect your backups.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints *float64 `json:"recoveryPoints,omitempty" tf:"recovery_points,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VaultParameters struct {

	// A boolean that indicates that all recovery points stored in the vault are deleted so that the vault can be destroyed without error.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The server-side encryption key that is used to protect your backups.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VaultSpec defines the desired state of Vault
type VaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultParameters `json:"forProvider"`
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
	InitProvider VaultInitParameters `json:"initProvider,omitempty"`
}

// VaultStatus defines the observed state of Vault.
type VaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Vault is the Schema for the Vaults API. Provides an AWS Backup vault resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Vault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultSpec   `json:"spec"`
	Status            VaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultList contains a list of Vaults
type VaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vault `json:"items"`
}

// Repository type metadata.
var (
	Vault_Kind             = "Vault"
	Vault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vault_Kind}.String()
	Vault_KindAPIVersion   = Vault_Kind + "." + CRDGroupVersion.String()
	Vault_GroupVersionKind = CRDGroupVersion.WithKind(Vault_Kind)
)

func init() {
	SchemeBuilder.Register(&Vault{}, &VaultList{})
}
