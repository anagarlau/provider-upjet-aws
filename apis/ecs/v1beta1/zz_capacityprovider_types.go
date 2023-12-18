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

type AutoScalingGroupProviderInitParameters struct {

	// - Configuration block defining the parameters of the auto scaling. Detailed below.
	ManagedScaling []ManagedScalingInitParameters `json:"managedScaling,omitempty" tf:"managed_scaling,omitempty"`

	// - Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are ENABLED and DISABLED.
	ManagedTerminationProtection *string `json:"managedTerminationProtection,omitempty" tf:"managed_termination_protection,omitempty"`
}

type AutoScalingGroupProviderObservation struct {

	// - ARN of the associated auto scaling group.
	AutoScalingGroupArn *string `json:"autoScalingGroupArn,omitempty" tf:"auto_scaling_group_arn,omitempty"`

	// - Configuration block defining the parameters of the auto scaling. Detailed below.
	ManagedScaling []ManagedScalingObservation `json:"managedScaling,omitempty" tf:"managed_scaling,omitempty"`

	// - Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are ENABLED and DISABLED.
	ManagedTerminationProtection *string `json:"managedTerminationProtection,omitempty" tf:"managed_termination_protection,omitempty"`
}

type AutoScalingGroupProviderParameters struct {

	// - ARN of the associated auto scaling group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	AutoScalingGroupArn *string `json:"autoScalingGroupArn,omitempty" tf:"auto_scaling_group_arn,omitempty"`

	// Reference to a AutoscalingGroup in autoscaling to populate autoScalingGroupArn.
	// +kubebuilder:validation:Optional
	AutoScalingGroupArnRef *v1.Reference `json:"autoScalingGroupArnRef,omitempty" tf:"-"`

	// Selector for a AutoscalingGroup in autoscaling to populate autoScalingGroupArn.
	// +kubebuilder:validation:Optional
	AutoScalingGroupArnSelector *v1.Selector `json:"autoScalingGroupArnSelector,omitempty" tf:"-"`

	// - Configuration block defining the parameters of the auto scaling. Detailed below.
	// +kubebuilder:validation:Optional
	ManagedScaling []ManagedScalingParameters `json:"managedScaling,omitempty" tf:"managed_scaling,omitempty"`

	// - Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are ENABLED and DISABLED.
	// +kubebuilder:validation:Optional
	ManagedTerminationProtection *string `json:"managedTerminationProtection,omitempty" tf:"managed_termination_protection,omitempty"`
}

type CapacityProviderInitParameters struct {

	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider []AutoScalingGroupProviderInitParameters `json:"autoScalingGroupProvider,omitempty" tf:"auto_scaling_group_provider,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CapacityProviderObservation struct {

	// ARN that identifies the capacity provider.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider []AutoScalingGroupProviderObservation `json:"autoScalingGroupProvider,omitempty" tf:"auto_scaling_group_provider,omitempty"`

	// ARN that identifies the capacity provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CapacityProviderParameters struct {

	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	// +kubebuilder:validation:Optional
	AutoScalingGroupProvider []AutoScalingGroupProviderParameters `json:"autoScalingGroupProvider,omitempty" tf:"auto_scaling_group_provider,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ManagedScalingInitParameters struct {

	// Period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group. If this parameter is omitted, the default value of 300 seconds is used.
	InstanceWarmupPeriod *float64 `json:"instanceWarmupPeriod,omitempty" tf:"instance_warmup_period,omitempty"`

	// Maximum step adjustment size. A number between 1 and 10,000.
	MaximumScalingStepSize *float64 `json:"maximumScalingStepSize,omitempty" tf:"maximum_scaling_step_size,omitempty"`

	// Minimum step adjustment size. A number between 1 and 10,000.
	MinimumScalingStepSize *float64 `json:"minimumScalingStepSize,omitempty" tf:"minimum_scaling_step_size,omitempty"`

	// Whether auto scaling is managed by ECS. Valid values are ENABLED and DISABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Target utilization for the capacity provider. A number between 1 and 100.
	TargetCapacity *float64 `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`
}

type ManagedScalingObservation struct {

	// Period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group. If this parameter is omitted, the default value of 300 seconds is used.
	InstanceWarmupPeriod *float64 `json:"instanceWarmupPeriod,omitempty" tf:"instance_warmup_period,omitempty"`

	// Maximum step adjustment size. A number between 1 and 10,000.
	MaximumScalingStepSize *float64 `json:"maximumScalingStepSize,omitempty" tf:"maximum_scaling_step_size,omitempty"`

	// Minimum step adjustment size. A number between 1 and 10,000.
	MinimumScalingStepSize *float64 `json:"minimumScalingStepSize,omitempty" tf:"minimum_scaling_step_size,omitempty"`

	// Whether auto scaling is managed by ECS. Valid values are ENABLED and DISABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Target utilization for the capacity provider. A number between 1 and 100.
	TargetCapacity *float64 `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`
}

type ManagedScalingParameters struct {

	// Period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group. If this parameter is omitted, the default value of 300 seconds is used.
	// +kubebuilder:validation:Optional
	InstanceWarmupPeriod *float64 `json:"instanceWarmupPeriod,omitempty" tf:"instance_warmup_period,omitempty"`

	// Maximum step adjustment size. A number between 1 and 10,000.
	// +kubebuilder:validation:Optional
	MaximumScalingStepSize *float64 `json:"maximumScalingStepSize,omitempty" tf:"maximum_scaling_step_size,omitempty"`

	// Minimum step adjustment size. A number between 1 and 10,000.
	// +kubebuilder:validation:Optional
	MinimumScalingStepSize *float64 `json:"minimumScalingStepSize,omitempty" tf:"minimum_scaling_step_size,omitempty"`

	// Whether auto scaling is managed by ECS. Valid values are ENABLED and DISABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Target utilization for the capacity provider. A number between 1 and 100.
	// +kubebuilder:validation:Optional
	TargetCapacity *float64 `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`
}

// CapacityProviderSpec defines the desired state of CapacityProvider
type CapacityProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CapacityProviderParameters `json:"forProvider"`
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
	InitProvider CapacityProviderInitParameters `json:"initProvider,omitempty"`
}

// CapacityProviderStatus defines the observed state of CapacityProvider.
type CapacityProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CapacityProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CapacityProvider is the Schema for the CapacityProviders API. Provides an ECS cluster capacity provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CapacityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoScalingGroupProvider) || (has(self.initProvider) && has(self.initProvider.autoScalingGroupProvider))",message="spec.forProvider.autoScalingGroupProvider is a required parameter"
	Spec   CapacityProviderSpec   `json:"spec"`
	Status CapacityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CapacityProviderList contains a list of CapacityProviders
type CapacityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CapacityProvider `json:"items"`
}

// Repository type metadata.
var (
	CapacityProvider_Kind             = "CapacityProvider"
	CapacityProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CapacityProvider_Kind}.String()
	CapacityProvider_KindAPIVersion   = CapacityProvider_Kind + "." + CRDGroupVersion.String()
	CapacityProvider_GroupVersionKind = CRDGroupVersion.WithKind(CapacityProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&CapacityProvider{}, &CapacityProviderList{})
}
