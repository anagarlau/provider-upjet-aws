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

type VolumeAttachmentInitParameters struct {

	// The device name to expose to the instance (for
	// example, /dev/sdh or xvdh).  See Device Naming on Linux Instances and Device Naming on Windows Instances for more information.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Set to true if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in data loss. See
	// Detaching an Amazon EBS Volume from an Instance for more information.
	ForceDetach *bool `json:"forceDetach,omitempty" tf:"force_detach,omitempty"`

	// ID of the Instance to attach to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Set this to true to ensure that the target instance is stopped
	// before trying to detach the volume. Stops the instance, if it is not already stopped.
	StopInstanceBeforeDetaching *bool `json:"stopInstanceBeforeDetaching,omitempty" tf:"stop_instance_before_detaching,omitempty"`

	// ID of the Volume to be attached
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EBSVolume
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a EBSVolume in ec2 to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a EBSVolume in ec2 to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

type VolumeAttachmentObservation struct {

	// The device name to expose to the instance (for
	// example, /dev/sdh or xvdh).  See Device Naming on Linux Instances and Device Naming on Windows Instances for more information.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Set to true if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in data loss. See
	// Detaching an Amazon EBS Volume from an Instance for more information.
	ForceDetach *bool `json:"forceDetach,omitempty" tf:"force_detach,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the Instance to attach to
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Set this to true to ensure that the target instance is stopped
	// before trying to detach the volume. Stops the instance, if it is not already stopped.
	StopInstanceBeforeDetaching *bool `json:"stopInstanceBeforeDetaching,omitempty" tf:"stop_instance_before_detaching,omitempty"`

	// ID of the Volume to be attached
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type VolumeAttachmentParameters struct {

	// The device name to expose to the instance (for
	// example, /dev/sdh or xvdh).  See Device Naming on Linux Instances and Device Naming on Windows Instances for more information.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Set to true if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in data loss. See
	// Detaching an Amazon EBS Volume from an Instance for more information.
	// +kubebuilder:validation:Optional
	ForceDetach *bool `json:"forceDetach,omitempty" tf:"force_detach,omitempty"`

	// ID of the Instance to attach to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Set this to true to ensure that the target instance is stopped
	// before trying to detach the volume. Stops the instance, if it is not already stopped.
	// +kubebuilder:validation:Optional
	StopInstanceBeforeDetaching *bool `json:"stopInstanceBeforeDetaching,omitempty" tf:"stop_instance_before_detaching,omitempty"`

	// ID of the Volume to be attached
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EBSVolume
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a EBSVolume in ec2 to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a EBSVolume in ec2 to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

// VolumeAttachmentSpec defines the desired state of VolumeAttachment
type VolumeAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeAttachmentParameters `json:"forProvider"`
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
	InitProvider VolumeAttachmentInitParameters `json:"initProvider,omitempty"`
}

// VolumeAttachmentStatus defines the observed state of VolumeAttachment.
type VolumeAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VolumeAttachment is the Schema for the VolumeAttachments API. Provides an AWS EBS Volume Attachment
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VolumeAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.deviceName) || (has(self.initProvider) && has(self.initProvider.deviceName))",message="spec.forProvider.deviceName is a required parameter"
	Spec   VolumeAttachmentSpec   `json:"spec"`
	Status VolumeAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachmentList contains a list of VolumeAttachments
type VolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeAttachment `json:"items"`
}

// Repository type metadata.
var (
	VolumeAttachment_Kind             = "VolumeAttachment"
	VolumeAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeAttachment_Kind}.String()
	VolumeAttachment_KindAPIVersion   = VolumeAttachment_Kind + "." + CRDGroupVersion.String()
	VolumeAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VolumeAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeAttachment{}, &VolumeAttachmentList{})
}
