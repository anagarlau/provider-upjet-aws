//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPerformanceMetricSubscription) DeepCopyInto(out *NetworkPerformanceMetricSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPerformanceMetricSubscription.
func (in *NetworkPerformanceMetricSubscription) DeepCopy() *NetworkPerformanceMetricSubscription {
	if in == nil {
		return nil
	}
	out := new(NetworkPerformanceMetricSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPerformanceMetricSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPerformanceMetricSubscriptionInitParameters) DeepCopyInto(out *NetworkPerformanceMetricSubscriptionInitParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPerformanceMetricSubscriptionInitParameters.
func (in *NetworkPerformanceMetricSubscriptionInitParameters) DeepCopy() *NetworkPerformanceMetricSubscriptionInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkPerformanceMetricSubscriptionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPerformanceMetricSubscriptionList) DeepCopyInto(out *NetworkPerformanceMetricSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPerformanceMetricSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPerformanceMetricSubscriptionList.
func (in *NetworkPerformanceMetricSubscriptionList) DeepCopy() *NetworkPerformanceMetricSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(NetworkPerformanceMetricSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPerformanceMetricSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPerformanceMetricSubscriptionObservation) DeepCopyInto(out *NetworkPerformanceMetricSubscriptionObservation) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPerformanceMetricSubscriptionObservation.
func (in *NetworkPerformanceMetricSubscriptionObservation) DeepCopy() *NetworkPerformanceMetricSubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkPerformanceMetricSubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPerformanceMetricSubscriptionParameters) DeepCopyInto(out *NetworkPerformanceMetricSubscriptionParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPerformanceMetricSubscriptionParameters.
func (in *NetworkPerformanceMetricSubscriptionParameters) DeepCopy() *NetworkPerformanceMetricSubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkPerformanceMetricSubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPerformanceMetricSubscriptionSpec) DeepCopyInto(out *NetworkPerformanceMetricSubscriptionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPerformanceMetricSubscriptionSpec.
func (in *NetworkPerformanceMetricSubscriptionSpec) DeepCopy() *NetworkPerformanceMetricSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPerformanceMetricSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPerformanceMetricSubscriptionStatus) DeepCopyInto(out *NetworkPerformanceMetricSubscriptionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPerformanceMetricSubscriptionStatus.
func (in *NetworkPerformanceMetricSubscriptionStatus) DeepCopy() *NetworkPerformanceMetricSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkPerformanceMetricSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
