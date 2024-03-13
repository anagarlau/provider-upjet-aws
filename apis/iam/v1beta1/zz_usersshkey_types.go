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

type UserSSHKeyInitParameters struct {

	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use SSH. To retrieve the public key in PEM format, use PEM.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the IAM user to associate the SSH public key with.
	// +crossplane:generate:reference:type=User
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

type UserSSHKeyObservation struct {

	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use SSH. To retrieve the public key in PEM format, use PEM.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The MD5 message digest of the SSH public key.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The unique identifier for the SSH public key.
	SSHPublicKeyID *string `json:"sshPublicKeyId,omitempty" tf:"ssh_public_key_id,omitempty"`

	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the IAM user to associate the SSH public key with.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserSSHKeyParameters struct {

	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use SSH. To retrieve the public key in PEM format, use PEM.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is active.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the IAM user to associate the SSH public key with.
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

// UserSSHKeySpec defines the desired state of UserSSHKey
type UserSSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserSSHKeyParameters `json:"forProvider"`
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
	InitProvider UserSSHKeyInitParameters `json:"initProvider,omitempty"`
}

// UserSSHKeyStatus defines the observed state of UserSSHKey.
type UserSSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserSSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserSSHKey is the Schema for the UserSSHKeys API. Uploads an SSH public key and associates it with the specified IAM user.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserSSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.encoding) || (has(self.initProvider) && has(self.initProvider.encoding))",message="spec.forProvider.encoding is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicKey) || (has(self.initProvider) && has(self.initProvider.publicKey))",message="spec.forProvider.publicKey is a required parameter"
	Spec   UserSSHKeySpec   `json:"spec"`
	Status UserSSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserSSHKeyList contains a list of UserSSHKeys
type UserSSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserSSHKey `json:"items"`
}

// Repository type metadata.
var (
	UserSSHKey_Kind             = "UserSSHKey"
	UserSSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserSSHKey_Kind}.String()
	UserSSHKey_KindAPIVersion   = UserSSHKey_Kind + "." + CRDGroupVersion.String()
	UserSSHKey_GroupVersionKind = CRDGroupVersion.WithKind(UserSSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&UserSSHKey{}, &UserSSHKeyList{})
}
