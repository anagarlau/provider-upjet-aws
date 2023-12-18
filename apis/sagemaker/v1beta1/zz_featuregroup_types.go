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

type DataCatalogConfigInitParameters struct {

	// The name of the Glue table catalog.
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// The name of the Glue table database.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The name of the Glue table.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type DataCatalogConfigObservation struct {

	// The name of the Glue table catalog.
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// The name of the Glue table database.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The name of the Glue table.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type DataCatalogConfigParameters struct {

	// The name of the Glue table catalog.
	// +kubebuilder:validation:Optional
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// The name of the Glue table database.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The name of the Glue table.
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type FeatureDefinitionInitParameters struct {

	// The name of a feature. feature_name cannot be any of the following: is_deleted, write_time, api_invocation_time.
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name,omitempty"`

	// The value type of a feature. Valid values are Integral, Fractional, or String.
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`
}

type FeatureDefinitionObservation struct {

	// The name of a feature. feature_name cannot be any of the following: is_deleted, write_time, api_invocation_time.
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name,omitempty"`

	// The value type of a feature. Valid values are Integral, Fractional, or String.
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`
}

type FeatureDefinitionParameters struct {

	// The name of a feature. feature_name cannot be any of the following: is_deleted, write_time, api_invocation_time.
	// +kubebuilder:validation:Optional
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name,omitempty"`

	// The value type of a feature. Valid values are Integral, Fractional, or String.
	// +kubebuilder:validation:Optional
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type,omitempty"`
}

type FeatureGroupInitParameters struct {

	// A free-form description of a Feature Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the feature that stores the EventTime of a Record in a Feature Group.
	EventTimeFeatureName *string `json:"eventTimeFeatureName,omitempty" tf:"event_time_feature_name,omitempty"`

	// A list of Feature names and types. See Feature Definition Below.
	FeatureDefinition []FeatureDefinitionInitParameters `json:"featureDefinition,omitempty" tf:"feature_definition,omitempty"`

	// The Offline Feature Store Configuration. See Offline Store Config Below.
	OfflineStoreConfig []OfflineStoreConfigInitParameters `json:"offlineStoreConfig,omitempty" tf:"offline_store_config,omitempty"`

	// The Online Feature Store Configuration. See Online Store Config Below.
	OnlineStoreConfig []OnlineStoreConfigInitParameters `json:"onlineStoreConfig,omitempty" tf:"online_store_config,omitempty"`

	// The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
	RecordIdentifierFeatureName *string `json:"recordIdentifierFeatureName,omitempty" tf:"record_identifier_feature_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FeatureGroupObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this feature_group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A free-form description of a Feature Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the feature that stores the EventTime of a Record in a Feature Group.
	EventTimeFeatureName *string `json:"eventTimeFeatureName,omitempty" tf:"event_time_feature_name,omitempty"`

	// A list of Feature names and types. See Feature Definition Below.
	FeatureDefinition []FeatureDefinitionObservation `json:"featureDefinition,omitempty" tf:"feature_definition,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Offline Feature Store Configuration. See Offline Store Config Below.
	OfflineStoreConfig []OfflineStoreConfigObservation `json:"offlineStoreConfig,omitempty" tf:"offline_store_config,omitempty"`

	// The Online Feature Store Configuration. See Online Store Config Below.
	OnlineStoreConfig []OnlineStoreConfigObservation `json:"onlineStoreConfig,omitempty" tf:"online_store_config,omitempty"`

	// The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
	RecordIdentifierFeatureName *string `json:"recordIdentifierFeatureName,omitempty" tf:"record_identifier_feature_name,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an offline_store_config is provided.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FeatureGroupParameters struct {

	// A free-form description of a Feature Group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the feature that stores the EventTime of a Record in a Feature Group.
	// +kubebuilder:validation:Optional
	EventTimeFeatureName *string `json:"eventTimeFeatureName,omitempty" tf:"event_time_feature_name,omitempty"`

	// A list of Feature names and types. See Feature Definition Below.
	// +kubebuilder:validation:Optional
	FeatureDefinition []FeatureDefinitionParameters `json:"featureDefinition,omitempty" tf:"feature_definition,omitempty"`

	// The Offline Feature Store Configuration. See Offline Store Config Below.
	// +kubebuilder:validation:Optional
	OfflineStoreConfig []OfflineStoreConfigParameters `json:"offlineStoreConfig,omitempty" tf:"offline_store_config,omitempty"`

	// The Online Feature Store Configuration. See Online Store Config Below.
	// +kubebuilder:validation:Optional
	OnlineStoreConfig []OnlineStoreConfigParameters `json:"onlineStoreConfig,omitempty" tf:"online_store_config,omitempty"`

	// The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
	// +kubebuilder:validation:Optional
	RecordIdentifierFeatureName *string `json:"recordIdentifierFeatureName,omitempty" tf:"record_identifier_feature_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an offline_store_config is provided.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OfflineStoreConfigInitParameters struct {

	// The meta data of the Glue table that is autogenerated when an OfflineStore is created. See Data Catalog Config Below.
	DataCatalogConfig []DataCatalogConfigInitParameters `json:"dataCatalogConfig,omitempty" tf:"data_catalog_config,omitempty"`

	// Set to true to turn Online Store On.
	DisableGlueTableCreation *bool `json:"disableGlueTableCreation,omitempty" tf:"disable_glue_table_creation,omitempty"`

	// The Amazon Simple Storage (Amazon S3) location of OfflineStore. See S3 Storage Config Below.
	S3StorageConfig []S3StorageConfigInitParameters `json:"s3StorageConfig,omitempty" tf:"s3_storage_config,omitempty"`

	// Format for the offline store table. Supported formats are Glue (Default) and Apache Iceberg (https://iceberg.apache.org/).
	TableFormat *string `json:"tableFormat,omitempty" tf:"table_format,omitempty"`
}

type OfflineStoreConfigObservation struct {

	// The meta data of the Glue table that is autogenerated when an OfflineStore is created. See Data Catalog Config Below.
	DataCatalogConfig []DataCatalogConfigObservation `json:"dataCatalogConfig,omitempty" tf:"data_catalog_config,omitempty"`

	// Set to true to turn Online Store On.
	DisableGlueTableCreation *bool `json:"disableGlueTableCreation,omitempty" tf:"disable_glue_table_creation,omitempty"`

	// The Amazon Simple Storage (Amazon S3) location of OfflineStore. See S3 Storage Config Below.
	S3StorageConfig []S3StorageConfigObservation `json:"s3StorageConfig,omitempty" tf:"s3_storage_config,omitempty"`

	// Format for the offline store table. Supported formats are Glue (Default) and Apache Iceberg (https://iceberg.apache.org/).
	TableFormat *string `json:"tableFormat,omitempty" tf:"table_format,omitempty"`
}

type OfflineStoreConfigParameters struct {

	// The meta data of the Glue table that is autogenerated when an OfflineStore is created. See Data Catalog Config Below.
	// +kubebuilder:validation:Optional
	DataCatalogConfig []DataCatalogConfigParameters `json:"dataCatalogConfig,omitempty" tf:"data_catalog_config,omitempty"`

	// Set to true to turn Online Store On.
	// +kubebuilder:validation:Optional
	DisableGlueTableCreation *bool `json:"disableGlueTableCreation,omitempty" tf:"disable_glue_table_creation,omitempty"`

	// The Amazon Simple Storage (Amazon S3) location of OfflineStore. See S3 Storage Config Below.
	// +kubebuilder:validation:Optional
	S3StorageConfig []S3StorageConfigParameters `json:"s3StorageConfig" tf:"s3_storage_config,omitempty"`

	// Format for the offline store table. Supported formats are Glue (Default) and Apache Iceberg (https://iceberg.apache.org/).
	// +kubebuilder:validation:Optional
	TableFormat *string `json:"tableFormat,omitempty" tf:"table_format,omitempty"`
}

type OnlineStoreConfigInitParameters struct {

	// Set to true to disable the automatic creation of an AWS Glue table when configuring an OfflineStore.
	EnableOnlineStore *bool `json:"enableOnlineStore,omitempty" tf:"enable_online_store,omitempty"`

	// Security config for at-rest encryption of your OnlineStore. See Security Config Below.
	SecurityConfig []SecurityConfigInitParameters `json:"securityConfig,omitempty" tf:"security_config,omitempty"`
}

type OnlineStoreConfigObservation struct {

	// Set to true to disable the automatic creation of an AWS Glue table when configuring an OfflineStore.
	EnableOnlineStore *bool `json:"enableOnlineStore,omitempty" tf:"enable_online_store,omitempty"`

	// Security config for at-rest encryption of your OnlineStore. See Security Config Below.
	SecurityConfig []SecurityConfigObservation `json:"securityConfig,omitempty" tf:"security_config,omitempty"`
}

type OnlineStoreConfigParameters struct {

	// Set to true to disable the automatic creation of an AWS Glue table when configuring an OfflineStore.
	// +kubebuilder:validation:Optional
	EnableOnlineStore *bool `json:"enableOnlineStore,omitempty" tf:"enable_online_store,omitempty"`

	// Security config for at-rest encryption of your OnlineStore. See Security Config Below.
	// +kubebuilder:validation:Optional
	SecurityConfig []SecurityConfigParameters `json:"securityConfig,omitempty" tf:"security_config,omitempty"`
}

type S3StorageConfigInitParameters struct {

	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The S3 URI, or location in Amazon S3, of OfflineStore.
	S3URI *string `json:"s3Uri,omitempty" tf:"s3_uri,omitempty"`
}

type S3StorageConfigObservation struct {

	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The S3 URI, or location in Amazon S3, of OfflineStore.
	S3URI *string `json:"s3Uri,omitempty" tf:"s3_uri,omitempty"`
}

type S3StorageConfigParameters struct {

	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The S3 URI, or location in Amazon S3, of OfflineStore.
	// +kubebuilder:validation:Optional
	S3URI *string `json:"s3Uri" tf:"s3_uri,omitempty"`
}

type SecurityConfigInitParameters struct {

	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type SecurityConfigObservation struct {

	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type SecurityConfigParameters struct {

	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the OfflineStore S3 location.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

// FeatureGroupSpec defines the desired state of FeatureGroup
type FeatureGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FeatureGroupParameters `json:"forProvider"`
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
	InitProvider FeatureGroupInitParameters `json:"initProvider,omitempty"`
}

// FeatureGroupStatus defines the observed state of FeatureGroup.
type FeatureGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FeatureGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FeatureGroup is the Schema for the FeatureGroups API. Provides a SageMaker Feature Group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FeatureGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventTimeFeatureName) || (has(self.initProvider) && has(self.initProvider.eventTimeFeatureName))",message="spec.forProvider.eventTimeFeatureName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.featureDefinition) || (has(self.initProvider) && has(self.initProvider.featureDefinition))",message="spec.forProvider.featureDefinition is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.recordIdentifierFeatureName) || (has(self.initProvider) && has(self.initProvider.recordIdentifierFeatureName))",message="spec.forProvider.recordIdentifierFeatureName is a required parameter"
	Spec   FeatureGroupSpec   `json:"spec"`
	Status FeatureGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FeatureGroupList contains a list of FeatureGroups
type FeatureGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FeatureGroup `json:"items"`
}

// Repository type metadata.
var (
	FeatureGroup_Kind             = "FeatureGroup"
	FeatureGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FeatureGroup_Kind}.String()
	FeatureGroup_KindAPIVersion   = FeatureGroup_Kind + "." + CRDGroupVersion.String()
	FeatureGroup_GroupVersionKind = CRDGroupVersion.WithKind(FeatureGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&FeatureGroup{}, &FeatureGroupList{})
}
