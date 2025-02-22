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

type CacheNodesInitParameters struct {
}

type CacheNodesObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use preferred_availability_zones instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the created ElastiCache Cluster.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// create the resource.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type CacheNodesParameters struct {
}

type ClusterInitParameters struct {

	// Whether any database modifications are applied immediately, or during the next maintenance window. Default is false. See Amazon ElastiCache Documentation for more information..
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
	// Only supported for engine type "redis" and if the engine version is 6 or higher.
	// Defaults to true.
	AutoMinorVersionUpgrade *string `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use preferred_availability_zones instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are single-az or cross-az, default is single-az. If you want to choose cross-az, num_cache_nodes must be greater than 1.
	AzMode *string `json:"azMode,omitempty" tf:"az_mode,omitempty"`

	// –  Name of the cache engine to be used for this cache cluster. Valid values are memcached and redis.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// –  Version number of the cache engine to be used.
	// If not set, defaults to the latest version.
	// See Describe Cache Engine Versions in the AWS Documentation for supported versions.
	// When engine is redis and the version is 7 or higher, the major and minor version should be set, e.g., 7.2.
	// When the version is 6, the major and minor version can be set, e.g., 6.2,
	// or the minor version can be unspecified which will use the latest version at creation time, e.g., 6.x.
	// Otherwise, specify the full version desired, e.g., 5.0.6.
	// The actual engine version used is returned in the attribute engine_version_actual, see Attribute Reference below. Cannot be provided with replication_group_id.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Name of your final cluster snapshot. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The IP version to advertise in the discovery protocol. Valid values are ipv4 or ipv6.
	IPDiscovery *string `json:"ipDiscovery,omitempty" tf:"ip_discovery,omitempty"`

	// Specifies the destination and format of Redis SLOWLOG or Redis Engine Log. See the documentation on Amazon ElastiCache. See Log Delivery Configuration below for more details.
	LogDeliveryConfiguration []LogDeliveryConfigurationInitParameters `json:"logDeliveryConfiguration,omitempty" tf:"log_delivery_configuration,omitempty"`

	// ddd:hh24:mi (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: sun:05:00-sun:09:00.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The IP versions for cache cluster connections. IPv6 is supported with Redis engine 6.2 onword or Memcached version 1.6.6 for all Nitro system instances. Valid values are ipv4, ipv6 or dual_stack.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// create the resource.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// east-1:012345678999:my_sns_topic.
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`

	// –  The initial number of cache nodes that the cache cluster will have. For Redis, this value must be 1. For Memcached, this value must be between 1 and 40. If this number is reduced on subsequent runs, the highest numbered nodes will be removed.
	NumCacheNodes *float64 `json:"numCacheNodes,omitempty" tf:"num_cache_nodes,omitempty"`

	// Specify the outpost mode that will apply to the cache cluster creation. Valid values are "single-outpost" and "cross-outpost", however AWS currently only supports "single-outpost" mode.
	OutpostMode *string `json:"outpostMode,omitempty" tf:"outpost_mode,omitempty"`

	// –  The name of the parameter group to associate with this cache cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta1.ParameterGroup
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// Reference to a ParameterGroup in elasticache to populate parameterGroupName.
	// +kubebuilder:validation:Optional
	ParameterGroupNameRef *v1.Reference `json:"parameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ParameterGroup in elasticache to populate parameterGroupName.
	// +kubebuilder:validation:Optional
	ParameterGroupNameSelector *v1.Selector `json:"parameterGroupNameSelector,omitempty" tf:"-"`

	// create the resource.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// List of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of num_cache_nodes. If you want all the nodes in the same Availability Zone, use availability_zone instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones []*string `json:"preferredAvailabilityZones,omitempty" tf:"preferred_availability_zones,omitempty"`

	// The outpost ARN in which the cache cluster will be created.
	PreferredOutpostArn *string `json:"preferredOutpostArn,omitempty" tf:"preferred_outpost_arn,omitempty"`

	// ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta2.ReplicationGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ReplicationGroupID *string `json:"replicationGroupId,omitempty" tf:"replication_group_id,omitempty"`

	// Reference to a ReplicationGroup in elasticache to populate replicationGroupId.
	// +kubebuilder:validation:Optional
	ReplicationGroupIDRef *v1.Reference `json:"replicationGroupIdRef,omitempty" tf:"-"`

	// Selector for a ReplicationGroup in elasticache to populate replicationGroupId.
	// +kubebuilder:validation:Optional
	ReplicationGroupIDSelector *v1.Selector `json:"replicationGroupIdSelector,omitempty" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// –  One or more VPC security groups associated with the cache cluster. Cannot be provided with replication_group_id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// element string list containing an Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3. The object name cannot contain any commas. Changing snapshot_arns forces a new resource.
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// Name of a snapshot from which to restore data into the new node group. Changing snapshot_name forces a new resource.
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off. Please note that setting a snapshot_retention_limit is not supported on cache.t1.micro cache nodes
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// create the resource. Cannot be provided with replication_group_id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta1.SubnetGroup
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// Reference to a SubnetGroup in elasticache to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameRef *v1.Reference `json:"subnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup in elasticache to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameSelector *v1.Selector `json:"subnetGroupNameSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enable encryption in-transit. Supported only with Memcached versions 1.6.12 and later, running in a VPC. See the ElastiCache in-transit encryption documentation for more details.
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
}

type ClusterObservation struct {

	// Whether any database modifications are applied immediately, or during the next maintenance window. Default is false. See Amazon ElastiCache Documentation for more information..
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// The ARN of the created ElastiCache Cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
	// Only supported for engine type "redis" and if the engine version is 6 or higher.
	// Defaults to true.
	AutoMinorVersionUpgrade *string `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use preferred_availability_zones instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are single-az or cross-az, default is single-az. If you want to choose cross-az, num_cache_nodes must be greater than 1.
	AzMode *string `json:"azMode,omitempty" tf:"az_mode,omitempty"`

	// List of node objects including id, address, port and availability_zone.
	CacheNodes []CacheNodesObservation `json:"cacheNodes,omitempty" tf:"cache_nodes,omitempty"`

	// (Memcached only) DNS name of the cache cluster without the port appended.
	ClusterAddress *string `json:"clusterAddress,omitempty" tf:"cluster_address,omitempty"`

	// (Memcached only) Configuration endpoint to allow host discovery.
	ConfigurationEndpoint *string `json:"configurationEndpoint,omitempty" tf:"configuration_endpoint,omitempty"`

	// –  Name of the cache engine to be used for this cache cluster. Valid values are memcached and redis.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// –  Version number of the cache engine to be used.
	// If not set, defaults to the latest version.
	// See Describe Cache Engine Versions in the AWS Documentation for supported versions.
	// When engine is redis and the version is 7 or higher, the major and minor version should be set, e.g., 7.2.
	// When the version is 6, the major and minor version can be set, e.g., 6.2,
	// or the minor version can be unspecified which will use the latest version at creation time, e.g., 6.x.
	// Otherwise, specify the full version desired, e.g., 5.0.6.
	// The actual engine version used is returned in the attribute engine_version_actual, see Attribute Reference below. Cannot be provided with replication_group_id.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Because ElastiCache pulls the latest minor or patch for a version, this attribute returns the running version of the cache engine.
	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	// Name of your final cluster snapshot. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP version to advertise in the discovery protocol. Valid values are ipv4 or ipv6.
	IPDiscovery *string `json:"ipDiscovery,omitempty" tf:"ip_discovery,omitempty"`

	// Specifies the destination and format of Redis SLOWLOG or Redis Engine Log. See the documentation on Amazon ElastiCache. See Log Delivery Configuration below for more details.
	LogDeliveryConfiguration []LogDeliveryConfigurationObservation `json:"logDeliveryConfiguration,omitempty" tf:"log_delivery_configuration,omitempty"`

	// ddd:hh24:mi (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: sun:05:00-sun:09:00.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The IP versions for cache cluster connections. IPv6 is supported with Redis engine 6.2 onword or Memcached version 1.6.6 for all Nitro system instances. Valid values are ipv4, ipv6 or dual_stack.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// create the resource.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// east-1:012345678999:my_sns_topic.
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`

	// –  The initial number of cache nodes that the cache cluster will have. For Redis, this value must be 1. For Memcached, this value must be between 1 and 40. If this number is reduced on subsequent runs, the highest numbered nodes will be removed.
	NumCacheNodes *float64 `json:"numCacheNodes,omitempty" tf:"num_cache_nodes,omitempty"`

	// Specify the outpost mode that will apply to the cache cluster creation. Valid values are "single-outpost" and "cross-outpost", however AWS currently only supports "single-outpost" mode.
	OutpostMode *string `json:"outpostMode,omitempty" tf:"outpost_mode,omitempty"`

	// –  The name of the parameter group to associate with this cache cluster.
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// create the resource.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// List of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of num_cache_nodes. If you want all the nodes in the same Availability Zone, use availability_zone instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones []*string `json:"preferredAvailabilityZones,omitempty" tf:"preferred_availability_zones,omitempty"`

	// The outpost ARN in which the cache cluster will be created.
	PreferredOutpostArn *string `json:"preferredOutpostArn,omitempty" tf:"preferred_outpost_arn,omitempty"`

	// ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupID *string `json:"replicationGroupId,omitempty" tf:"replication_group_id,omitempty"`

	// –  One or more VPC security groups associated with the cache cluster. Cannot be provided with replication_group_id.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// element string list containing an Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3. The object name cannot contain any commas. Changing snapshot_arns forces a new resource.
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// Name of a snapshot from which to restore data into the new node group. Changing snapshot_name forces a new resource.
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off. Please note that setting a snapshot_retention_limit is not supported on cache.t1.micro cache nodes
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// create the resource. Cannot be provided with replication_group_id.
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Enable encryption in-transit. Supported only with Memcached versions 1.6.12 and later, running in a VPC. See the ElastiCache in-transit encryption documentation for more details.
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
}

type ClusterParameters struct {

	// Whether any database modifications are applied immediately, or during the next maintenance window. Default is false. See Amazon ElastiCache Documentation for more information..
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
	// Only supported for engine type "redis" and if the engine version is 6 or higher.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *string `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use preferred_availability_zones instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are single-az or cross-az, default is single-az. If you want to choose cross-az, num_cache_nodes must be greater than 1.
	// +kubebuilder:validation:Optional
	AzMode *string `json:"azMode,omitempty" tf:"az_mode,omitempty"`

	// –  Name of the cache engine to be used for this cache cluster. Valid values are memcached and redis.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// –  Version number of the cache engine to be used.
	// If not set, defaults to the latest version.
	// See Describe Cache Engine Versions in the AWS Documentation for supported versions.
	// When engine is redis and the version is 7 or higher, the major and minor version should be set, e.g., 7.2.
	// When the version is 6, the major and minor version can be set, e.g., 6.2,
	// or the minor version can be unspecified which will use the latest version at creation time, e.g., 6.x.
	// Otherwise, specify the full version desired, e.g., 5.0.6.
	// The actual engine version used is returned in the attribute engine_version_actual, see Attribute Reference below. Cannot be provided with replication_group_id.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Name of your final cluster snapshot. If omitted, no final snapshot will be made.
	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The IP version to advertise in the discovery protocol. Valid values are ipv4 or ipv6.
	// +kubebuilder:validation:Optional
	IPDiscovery *string `json:"ipDiscovery,omitempty" tf:"ip_discovery,omitempty"`

	// Specifies the destination and format of Redis SLOWLOG or Redis Engine Log. See the documentation on Amazon ElastiCache. See Log Delivery Configuration below for more details.
	// +kubebuilder:validation:Optional
	LogDeliveryConfiguration []LogDeliveryConfigurationParameters `json:"logDeliveryConfiguration,omitempty" tf:"log_delivery_configuration,omitempty"`

	// ddd:hh24:mi (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: sun:05:00-sun:09:00.
	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The IP versions for cache cluster connections. IPv6 is supported with Redis engine 6.2 onword or Memcached version 1.6.6 for all Nitro system instances. Valid values are ipv4, ipv6 or dual_stack.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// create the resource.
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// east-1:012345678999:my_sns_topic.
	// +kubebuilder:validation:Optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`

	// –  The initial number of cache nodes that the cache cluster will have. For Redis, this value must be 1. For Memcached, this value must be between 1 and 40. If this number is reduced on subsequent runs, the highest numbered nodes will be removed.
	// +kubebuilder:validation:Optional
	NumCacheNodes *float64 `json:"numCacheNodes,omitempty" tf:"num_cache_nodes,omitempty"`

	// Specify the outpost mode that will apply to the cache cluster creation. Valid values are "single-outpost" and "cross-outpost", however AWS currently only supports "single-outpost" mode.
	// +kubebuilder:validation:Optional
	OutpostMode *string `json:"outpostMode,omitempty" tf:"outpost_mode,omitempty"`

	// –  The name of the parameter group to associate with this cache cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta1.ParameterGroup
	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// Reference to a ParameterGroup in elasticache to populate parameterGroupName.
	// +kubebuilder:validation:Optional
	ParameterGroupNameRef *v1.Reference `json:"parameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ParameterGroup in elasticache to populate parameterGroupName.
	// +kubebuilder:validation:Optional
	ParameterGroupNameSelector *v1.Selector `json:"parameterGroupNameSelector,omitempty" tf:"-"`

	// create the resource.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// List of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of num_cache_nodes. If you want all the nodes in the same Availability Zone, use availability_zone instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	// +kubebuilder:validation:Optional
	PreferredAvailabilityZones []*string `json:"preferredAvailabilityZones,omitempty" tf:"preferred_availability_zones,omitempty"`

	// The outpost ARN in which the cache cluster will be created.
	// +kubebuilder:validation:Optional
	PreferredOutpostArn *string `json:"preferredOutpostArn,omitempty" tf:"preferred_outpost_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta2.ReplicationGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ReplicationGroupID *string `json:"replicationGroupId,omitempty" tf:"replication_group_id,omitempty"`

	// Reference to a ReplicationGroup in elasticache to populate replicationGroupId.
	// +kubebuilder:validation:Optional
	ReplicationGroupIDRef *v1.Reference `json:"replicationGroupIdRef,omitempty" tf:"-"`

	// Selector for a ReplicationGroup in elasticache to populate replicationGroupId.
	// +kubebuilder:validation:Optional
	ReplicationGroupIDSelector *v1.Selector `json:"replicationGroupIdSelector,omitempty" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// –  One or more VPC security groups associated with the cache cluster. Cannot be provided with replication_group_id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// element string list containing an Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3. The object name cannot contain any commas. Changing snapshot_arns forces a new resource.
	// +kubebuilder:validation:Optional
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// Name of a snapshot from which to restore data into the new node group. Changing snapshot_name forces a new resource.
	// +kubebuilder:validation:Optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off. Please note that setting a snapshot_retention_limit is not supported on cache.t1.micro cache nodes
	// +kubebuilder:validation:Optional
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	// +kubebuilder:validation:Optional
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// create the resource. Cannot be provided with replication_group_id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta1.SubnetGroup
	// +kubebuilder:validation:Optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// Reference to a SubnetGroup in elasticache to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameRef *v1.Reference `json:"subnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup in elasticache to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameSelector *v1.Selector `json:"subnetGroupNameSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enable encryption in-transit. Supported only with Memcached versions 1.6.12 and later, running in a VPC. See the ElastiCache in-transit encryption documentation for more details.
	// +kubebuilder:validation:Optional
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
}

type LogDeliveryConfigurationInitParameters struct {

	// Name of either the CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// For CloudWatch Logs use cloudwatch-logs or for Kinesis Data Firehose use kinesis-firehose.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// Valid values are json or text
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`

	// Valid values are  slow-log or engine-log. Max 1 of each.
	LogType *string `json:"logType,omitempty" tf:"log_type,omitempty"`
}

type LogDeliveryConfigurationObservation struct {

	// Name of either the CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// For CloudWatch Logs use cloudwatch-logs or for Kinesis Data Firehose use kinesis-firehose.
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// Valid values are json or text
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`

	// Valid values are  slow-log or engine-log. Max 1 of each.
	LogType *string `json:"logType,omitempty" tf:"log_type,omitempty"`
}

type LogDeliveryConfigurationParameters struct {

	// Name of either the CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// For CloudWatch Logs use cloudwatch-logs or for Kinesis Data Firehose use kinesis-firehose.
	// +kubebuilder:validation:Optional
	DestinationType *string `json:"destinationType" tf:"destination_type,omitempty"`

	// Valid values are json or text
	// +kubebuilder:validation:Optional
	LogFormat *string `json:"logFormat" tf:"log_format,omitempty"`

	// Valid values are  slow-log or engine-log. Max 1 of each.
	// +kubebuilder:validation:Optional
	LogType *string `json:"logType" tf:"log_type,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cluster is the Schema for the Clusters API. Provides an ElastiCache Cluster resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
