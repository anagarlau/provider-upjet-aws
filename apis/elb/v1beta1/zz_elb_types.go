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

type AccessLogsInitParameters struct {

	// The S3 bucket name to store the logs in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The S3 bucket prefix. Logs are stored in the root if not configured.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// Boolean to enable / disable access_logs. Default is true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`
}

type AccessLogsObservation struct {

	// The S3 bucket name to store the logs in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The S3 bucket prefix. Logs are stored in the root if not configured.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// Boolean to enable / disable access_logs. Default is true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`
}

type AccessLogsParameters struct {

	// The S3 bucket name to store the logs in.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// The S3 bucket prefix. Logs are stored in the root if not configured.
	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// Boolean to enable / disable access_logs. Default is true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`
}

type ELBInitParameters struct {

	// An Access Logs block. Access Logs documented below.
	AccessLogs []AccessLogsInitParameters `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`

	// The AZ's to serve traffic in.
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Boolean to enable connection draining. Default: false
	ConnectionDraining *bool `json:"connectionDraining,omitempty" tf:"connection_draining,omitempty"`

	// The time in seconds to allow for connections to drain. Default: 300
	ConnectionDrainingTimeout *float64 `json:"connectionDrainingTimeout,omitempty" tf:"connection_draining_timeout,omitempty"`

	// Enable cross-zone load balancing. Default: true
	CrossZoneLoadBalancing *bool `json:"crossZoneLoadBalancing,omitempty" tf:"cross_zone_load_balancing,omitempty"`

	// Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are monitor, defensive (default), strictest.
	DesyncMitigationMode *string `json:"desyncMitigationMode,omitempty" tf:"desync_mitigation_mode,omitempty"`

	// A health_check block. Health Check documented below.
	HealthCheck []HealthCheckInitParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The time in seconds that the connection is allowed to be idle. Default: 60
	IdleTimeout *float64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// A list of instance ids to place in the ELB pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +listType=set
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// References to Instance in ec2 to populate instances.
	// +kubebuilder:validation:Optional
	InstancesRefs []v1.Reference `json:"instancesRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in ec2 to populate instances.
	// +kubebuilder:validation:Optional
	InstancesSelector *v1.Selector `json:"instancesSelector,omitempty" tf:"-"`

	// If true, ELB will be an internal ELB.
	Internal *bool `json:"internal,omitempty" tf:"internal,omitempty"`

	// A list of listener blocks. Listeners documented below.
	Listener []ListenerInitParameters `json:"listener,omitempty" tf:"listener,omitempty"`

	// A list of security group IDs to assign to the ELB.
	// Only valid if creating an ELB within a VPC
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The name of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Use this for Classic or Default VPC only.
	SourceSecurityGroup *string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`

	// A list of subnet IDs to attach to the ELB. When an update to subnets will remove all current subnets, this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +listType=set
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// References to Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsRefs []v1.Reference `json:"subnetsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsSelector *v1.Selector `json:"subnetsSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ELBObservation struct {

	// An Access Logs block. Access Logs documented below.
	AccessLogs []AccessLogsObservation `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`

	// The ARN of the ELB
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AZ's to serve traffic in.
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Boolean to enable connection draining. Default: false
	ConnectionDraining *bool `json:"connectionDraining,omitempty" tf:"connection_draining,omitempty"`

	// The time in seconds to allow for connections to drain. Default: 300
	ConnectionDrainingTimeout *float64 `json:"connectionDrainingTimeout,omitempty" tf:"connection_draining_timeout,omitempty"`

	// Enable cross-zone load balancing. Default: true
	CrossZoneLoadBalancing *bool `json:"crossZoneLoadBalancing,omitempty" tf:"cross_zone_load_balancing,omitempty"`

	// The DNS name of the ELB
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are monitor, defensive (default), strictest.
	DesyncMitigationMode *string `json:"desyncMitigationMode,omitempty" tf:"desync_mitigation_mode,omitempty"`

	// A health_check block. Health Check documented below.
	HealthCheck []HealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The name of the ELB
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The time in seconds that the connection is allowed to be idle. Default: 60
	IdleTimeout *float64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// A list of instance ids to place in the ELB pool.
	// +listType=set
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// If true, ELB will be an internal ELB.
	Internal *bool `json:"internal,omitempty" tf:"internal,omitempty"`

	// A list of listener blocks. Listeners documented below.
	Listener []ListenerObservation `json:"listener,omitempty" tf:"listener,omitempty"`

	// A list of security group IDs to assign to the ELB.
	// Only valid if creating an ELB within a VPC
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The name of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Use this for Classic or Default VPC only.
	SourceSecurityGroup *string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`

	// The ID of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Only available on ELBs launched in a VPC.
	SourceSecurityGroupID *string `json:"sourceSecurityGroupId,omitempty" tf:"source_security_group_id,omitempty"`

	// A list of subnet IDs to attach to the ELB. When an update to subnets will remove all current subnets, this will force a new resource.
	// +listType=set
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ELBParameters struct {

	// An Access Logs block. Access Logs documented below.
	// +kubebuilder:validation:Optional
	AccessLogs []AccessLogsParameters `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`

	// The AZ's to serve traffic in.
	// +kubebuilder:validation:Optional
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Boolean to enable connection draining. Default: false
	// +kubebuilder:validation:Optional
	ConnectionDraining *bool `json:"connectionDraining,omitempty" tf:"connection_draining,omitempty"`

	// The time in seconds to allow for connections to drain. Default: 300
	// +kubebuilder:validation:Optional
	ConnectionDrainingTimeout *float64 `json:"connectionDrainingTimeout,omitempty" tf:"connection_draining_timeout,omitempty"`

	// Enable cross-zone load balancing. Default: true
	// +kubebuilder:validation:Optional
	CrossZoneLoadBalancing *bool `json:"crossZoneLoadBalancing,omitempty" tf:"cross_zone_load_balancing,omitempty"`

	// Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are monitor, defensive (default), strictest.
	// +kubebuilder:validation:Optional
	DesyncMitigationMode *string `json:"desyncMitigationMode,omitempty" tf:"desync_mitigation_mode,omitempty"`

	// A health_check block. Health Check documented below.
	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The time in seconds that the connection is allowed to be idle. Default: 60
	// +kubebuilder:validation:Optional
	IdleTimeout *float64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// A list of instance ids to place in the ELB pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +kubebuilder:validation:Optional
	// +listType=set
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// References to Instance in ec2 to populate instances.
	// +kubebuilder:validation:Optional
	InstancesRefs []v1.Reference `json:"instancesRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in ec2 to populate instances.
	// +kubebuilder:validation:Optional
	InstancesSelector *v1.Selector `json:"instancesSelector,omitempty" tf:"-"`

	// If true, ELB will be an internal ELB.
	// +kubebuilder:validation:Optional
	Internal *bool `json:"internal,omitempty" tf:"internal,omitempty"`

	// A list of listener blocks. Listeners documented below.
	// +kubebuilder:validation:Optional
	Listener []ListenerParameters `json:"listener,omitempty" tf:"listener,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of security group IDs to assign to the ELB.
	// Only valid if creating an ELB within a VPC
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The name of the security group that you can use as
	// part of your inbound rules for your load balancer's back-end application
	// instances. Use this for Classic or Default VPC only.
	// +kubebuilder:validation:Optional
	SourceSecurityGroup *string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`

	// A list of subnet IDs to attach to the ELB. When an update to subnets will remove all current subnets, this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	// +listType=set
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// References to Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsRefs []v1.Reference `json:"subnetsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsSelector *v1.Selector `json:"subnetsSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthCheckInitParameters struct {

	// The number of checks before the instance is declared healthy.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
	// values are:
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The length of time before the check times out.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The number of checks before the instance is declared unhealthy.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckObservation struct {

	// The number of checks before the instance is declared healthy.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
	// values are:
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The length of time before the check times out.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The number of checks before the instance is declared unhealthy.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckParameters struct {

	// The number of checks before the instance is declared healthy.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold" tf:"healthy_threshold,omitempty"`

	// The publishing interval in minutes. Valid values: 5 and 60. Default: 60
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval" tf:"interval,omitempty"`

	// The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
	// values are:
	// +kubebuilder:validation:Optional
	Target *string `json:"target" tf:"target,omitempty"`

	// The length of time before the check times out.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`

	// The number of checks before the instance is declared unhealthy.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" tf:"unhealthy_threshold,omitempty"`
}

type ListenerInitParameters struct {

	// The port on the instance to route to
	InstancePort *float64 `json:"instancePort,omitempty" tf:"instance_port,omitempty"`

	// The protocol to use to the instance. Valid
	// values are HTTP, HTTPS, TCP, or SSL
	InstanceProtocol *string `json:"instanceProtocol,omitempty" tf:"instance_protocol,omitempty"`

	// The port to listen on for the load balancer
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The protocol to listen on. Valid values are HTTP,
	// HTTPS, TCP, or SSL
	LBProtocol *string `json:"lbProtocol,omitempty" tf:"lb_protocol,omitempty"`

	// The ARN of an SSL certificate you have
	// uploaded to AWS IAM. Note ECDSA-specific restrictions below.  Only valid when
	SSLCertificateID *string `json:"sslCertificateId,omitempty" tf:"ssl_certificate_id,omitempty"`
}

type ListenerObservation struct {

	// The port on the instance to route to
	InstancePort *float64 `json:"instancePort,omitempty" tf:"instance_port,omitempty"`

	// The protocol to use to the instance. Valid
	// values are HTTP, HTTPS, TCP, or SSL
	InstanceProtocol *string `json:"instanceProtocol,omitempty" tf:"instance_protocol,omitempty"`

	// The port to listen on for the load balancer
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The protocol to listen on. Valid values are HTTP,
	// HTTPS, TCP, or SSL
	LBProtocol *string `json:"lbProtocol,omitempty" tf:"lb_protocol,omitempty"`

	// The ARN of an SSL certificate you have
	// uploaded to AWS IAM. Note ECDSA-specific restrictions below.  Only valid when
	SSLCertificateID *string `json:"sslCertificateId,omitempty" tf:"ssl_certificate_id,omitempty"`
}

type ListenerParameters struct {

	// The port on the instance to route to
	// +kubebuilder:validation:Optional
	InstancePort *float64 `json:"instancePort" tf:"instance_port,omitempty"`

	// The protocol to use to the instance. Valid
	// values are HTTP, HTTPS, TCP, or SSL
	// +kubebuilder:validation:Optional
	InstanceProtocol *string `json:"instanceProtocol" tf:"instance_protocol,omitempty"`

	// The port to listen on for the load balancer
	// +kubebuilder:validation:Optional
	LBPort *float64 `json:"lbPort" tf:"lb_port,omitempty"`

	// The protocol to listen on. Valid values are HTTP,
	// HTTPS, TCP, or SSL
	// +kubebuilder:validation:Optional
	LBProtocol *string `json:"lbProtocol" tf:"lb_protocol,omitempty"`

	// The ARN of an SSL certificate you have
	// uploaded to AWS IAM. Note ECDSA-specific restrictions below.  Only valid when
	// +kubebuilder:validation:Optional
	SSLCertificateID *string `json:"sslCertificateId,omitempty" tf:"ssl_certificate_id,omitempty"`
}

// ELBSpec defines the desired state of ELB
type ELBSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ELBParameters `json:"forProvider"`
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
	InitProvider ELBInitParameters `json:"initProvider,omitempty"`
}

// ELBStatus defines the observed state of ELB.
type ELBStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ELBObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ELB is the Schema for the ELBs API. Provides an Elastic Load Balancer resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ELB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.listener) || (has(self.initProvider) && has(self.initProvider.listener))",message="spec.forProvider.listener is a required parameter"
	Spec   ELBSpec   `json:"spec"`
	Status ELBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ELBList contains a list of ELBs
type ELBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ELB `json:"items"`
}

// Repository type metadata.
var (
	ELB_Kind             = "ELB"
	ELB_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ELB_Kind}.String()
	ELB_KindAPIVersion   = ELB_Kind + "." + CRDGroupVersion.String()
	ELB_GroupVersionKind = CRDGroupVersion.WithKind(ELB_Kind)
)

func init() {
	SchemeBuilder.Register(&ELB{}, &ELBList{})
}
