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

type HostObservation struct {

	// The CodeStar Host ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The CodeStar Host ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The CodeStar Host status. Possible values are PENDING, AVAILABLE, VPC_CONFIG_DELETING, VPC_CONFIG_INITIALIZING, and VPC_CONFIG_FAILED_INITIALIZATION.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type HostParameters struct {

	// The name of the host to be created. The name must be unique in the calling AWS account.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The endpoint of the infrastructure to be represented by the host after it is created.
	// +kubebuilder:validation:Required
	ProviderEndpoint *string `json:"providerEndpoint" tf:"provider_endpoint,omitempty"`

	// The name of the external provider where your third-party code repository is configured.
	// +kubebuilder:validation:Required
	ProviderType *string `json:"providerType" tf:"provider_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
	// +kubebuilder:validation:Optional
	VPCConfiguration []VPCConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type VPCConfigurationObservation struct {
}

type VPCConfigurationParameters struct {

	// he ID of the security group or security groups associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// The ID of the subnet or subnets associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.
	// +kubebuilder:validation:Optional
	TLSCertificate *string `json:"tlsCertificate,omitempty" tf:"tls_certificate,omitempty"`

	// The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

// HostSpec defines the desired state of Host
type HostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostParameters `json:"forProvider"`
}

// HostStatus defines the observed state of Host.
type HostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Host is the Schema for the Hosts API. Provides a CodeStar Host
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Host struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSpec   `json:"spec"`
	Status            HostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostList contains a list of Hosts
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Host `json:"items"`
}

// Repository type metadata.
var (
	Host_Kind             = "Host"
	Host_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Host_Kind}.String()
	Host_KindAPIVersion   = Host_Kind + "." + CRDGroupVersion.String()
	Host_GroupVersionKind = CRDGroupVersion.WithKind(Host_Kind)
)

func init() {
	SchemeBuilder.Register(&Host{}, &HostList{})
}