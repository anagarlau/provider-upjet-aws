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

type BucketObjectLockConfigurationInitParameters struct {

	// Name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to Enabled. Valid values: Enabled.
	ObjectLockEnabled *string `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// Configuration block for specifying the Object Lock rule for the specified object. See below.
	Rule []BucketObjectLockConfigurationRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketObjectLockConfigurationObservation struct {

	// Name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to Enabled. Valid values: Enabled.
	ObjectLockEnabled *string `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// Configuration block for specifying the Object Lock rule for the specified object. See below.
	Rule []BucketObjectLockConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketObjectLockConfigurationParameters struct {

	// Name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to Enabled. Valid values: Enabled.
	// +kubebuilder:validation:Optional
	ObjectLockEnabled *string `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for specifying the Object Lock rule for the specified object. See below.
	// +kubebuilder:validation:Optional
	Rule []BucketObjectLockConfigurationRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	// The token is generated in the back-end when versioning is enabled on a bucket. For more details on versioning, see the aws_s3_bucket_versioning resource.
	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

type BucketObjectLockConfigurationRuleInitParameters struct {

	// Configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket. See below.
	DefaultRetention []RuleDefaultRetentionInitParameters `json:"defaultRetention,omitempty" tf:"default_retention,omitempty"`
}

type BucketObjectLockConfigurationRuleObservation struct {

	// Configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket. See below.
	DefaultRetention []RuleDefaultRetentionObservation `json:"defaultRetention,omitempty" tf:"default_retention,omitempty"`
}

type BucketObjectLockConfigurationRuleParameters struct {

	// Configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket. See below.
	// +kubebuilder:validation:Optional
	DefaultRetention []RuleDefaultRetentionParameters `json:"defaultRetention" tf:"default_retention,omitempty"`
}

type RuleDefaultRetentionInitParameters struct {

	// Number of days that you want to specify for the default retention period.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values: COMPLIANCE, GOVERNANCE.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Number of years that you want to specify for the default retention period.
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type RuleDefaultRetentionObservation struct {

	// Number of days that you want to specify for the default retention period.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values: COMPLIANCE, GOVERNANCE.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Number of years that you want to specify for the default retention period.
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type RuleDefaultRetentionParameters struct {

	// Number of days that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values: COMPLIANCE, GOVERNANCE.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Number of years that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

// BucketObjectLockConfigurationSpec defines the desired state of BucketObjectLockConfiguration
type BucketObjectLockConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketObjectLockConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketObjectLockConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketObjectLockConfigurationStatus defines the observed state of BucketObjectLockConfiguration.
type BucketObjectLockConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObjectLockConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketObjectLockConfiguration is the Schema for the BucketObjectLockConfigurations API. Provides an S3 bucket Object Lock configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketObjectLockConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketObjectLockConfigurationSpec   `json:"spec"`
	Status            BucketObjectLockConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObjectLockConfigurationList contains a list of BucketObjectLockConfigurations
type BucketObjectLockConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketObjectLockConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketObjectLockConfiguration_Kind             = "BucketObjectLockConfiguration"
	BucketObjectLockConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketObjectLockConfiguration_Kind}.String()
	BucketObjectLockConfiguration_KindAPIVersion   = BucketObjectLockConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketObjectLockConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketObjectLockConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketObjectLockConfiguration{}, &BucketObjectLockConfigurationList{})
}
