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

type ListenerPolicyInitParameters struct {

	// The load balancer to attach the policy to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Reference to a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameRef *v1.Reference `json:"loadBalancerNameRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameSelector *v1.Selector `json:"loadBalancerNameSelector,omitempty" tf:"-"`

	// The load balancer listener port to apply the policy to.
	LoadBalancerPort *float64 `json:"loadBalancerPort,omitempty" tf:"load_balancer_port,omitempty"`

	// List of Policy Names to apply to the backend server.
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger an update.
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type ListenerPolicyObservation struct {

	// The ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The load balancer to attach the policy to.
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// The load balancer listener port to apply the policy to.
	LoadBalancerPort *float64 `json:"loadBalancerPort,omitempty" tf:"load_balancer_port,omitempty"`

	// List of Policy Names to apply to the backend server.
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger an update.
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type ListenerPolicyParameters struct {

	// The load balancer to attach the policy to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +kubebuilder:validation:Optional
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Reference to a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameRef *v1.Reference `json:"loadBalancerNameRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameSelector *v1.Selector `json:"loadBalancerNameSelector,omitempty" tf:"-"`

	// The load balancer listener port to apply the policy to.
	// +kubebuilder:validation:Optional
	LoadBalancerPort *float64 `json:"loadBalancerPort,omitempty" tf:"load_balancer_port,omitempty"`

	// List of Policy Names to apply to the backend server.
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Map of arbitrary keys and values that, when changed, will trigger an update.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

// ListenerPolicySpec defines the desired state of ListenerPolicy
type ListenerPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListenerPolicyParameters `json:"forProvider"`
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
	InitProvider ListenerPolicyInitParameters `json:"initProvider,omitempty"`
}

// ListenerPolicyStatus defines the observed state of ListenerPolicy.
type ListenerPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListenerPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ListenerPolicy is the Schema for the ListenerPolicys API. Attaches a load balancer policy to an ELB Listener.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ListenerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loadBalancerPort) || (has(self.initProvider) && has(self.initProvider.loadBalancerPort))",message="spec.forProvider.loadBalancerPort is a required parameter"
	Spec   ListenerPolicySpec   `json:"spec"`
	Status ListenerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListenerPolicyList contains a list of ListenerPolicys
type ListenerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ListenerPolicy `json:"items"`
}

// Repository type metadata.
var (
	ListenerPolicy_Kind             = "ListenerPolicy"
	ListenerPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ListenerPolicy_Kind}.String()
	ListenerPolicy_KindAPIVersion   = ListenerPolicy_Kind + "." + CRDGroupVersion.String()
	ListenerPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ListenerPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ListenerPolicy{}, &ListenerPolicyList{})
}
