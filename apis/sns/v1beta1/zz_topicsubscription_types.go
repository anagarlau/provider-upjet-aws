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

type TopicSubscriptionInitParameters struct {

	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is 1.
	ConfirmationTimeoutInMinutes *float64 `json:"confirmationTimeoutInMinutes,omitempty" tf:"confirmation_timeout_in_minutes,omitempty"`

	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the SNS docs for more details.
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`

	// Endpoint to send data to. The contents vary with the protocol. See details below.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Whether the endpoint is capable of auto confirming subscription (e.g., PagerDuty). Default is false.
	EndpointAutoConfirms *bool `json:"endpointAutoConfirms,omitempty" tf:"endpoint_auto_confirms,omitempty"`

	// Reference to a Queue in sqs to populate endpoint.
	// +kubebuilder:validation:Optional
	EndpointRef *v1.Reference `json:"endpointRef,omitempty" tf:"-"`

	// Selector for a Queue in sqs to populate endpoint.
	// +kubebuilder:validation:Optional
	EndpointSelector *v1.Selector `json:"endpointSelector,omitempty" tf:"-"`

	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the SNS docs for more details.
	FilterPolicy *string `json:"filterPolicy,omitempty" tf:"filter_policy,omitempty"`

	// Whether the filter_policy applies to MessageAttributes (default) or MessageBody.
	FilterPolicyScope *string `json:"filterPolicyScope,omitempty" tf:"filter_policy_scope,omitempty"`

	// Protocol to use. Valid values are: sqs, sms, lambda, firehose, and application. Protocols email, email-json, http and https are also valid but partially supported. See details below.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is false.
	RawMessageDelivery *bool `json:"rawMessageDelivery,omitempty" tf:"raw_message_delivery,omitempty"`

	// JSON String with the redrive policy that will be used in the subscription. Refer to the SNS docs for more details.
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// JSON String with the archived message replay policy that will be used in the subscription. Refer to the SNS docs for more details.
	ReplayPolicy *string `json:"replayPolicy,omitempty" tf:"replay_policy,omitempty"`

	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to SNS docs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	SubscriptionRoleArn *string `json:"subscriptionRoleArn,omitempty" tf:"subscription_role_arn,omitempty"`

	// Reference to a Role in iam to populate subscriptionRoleArn.
	// +kubebuilder:validation:Optional
	SubscriptionRoleArnRef *v1.Reference `json:"subscriptionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate subscriptionRoleArn.
	// +kubebuilder:validation:Optional
	SubscriptionRoleArnSelector *v1.Selector `json:"subscriptionRoleArnSelector,omitempty" tf:"-"`

	// ARN of the SNS topic to subscribe to.
	// +crossplane:generate:reference:type=Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

type TopicSubscriptionObservation struct {

	// ARN of the subscription.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is 1.
	ConfirmationTimeoutInMinutes *float64 `json:"confirmationTimeoutInMinutes,omitempty" tf:"confirmation_timeout_in_minutes,omitempty"`

	// Whether the subscription confirmation request was authenticated.
	ConfirmationWasAuthenticated *bool `json:"confirmationWasAuthenticated,omitempty" tf:"confirmation_was_authenticated,omitempty"`

	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the SNS docs for more details.
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`

	// Endpoint to send data to. The contents vary with the protocol. See details below.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Whether the endpoint is capable of auto confirming subscription (e.g., PagerDuty). Default is false.
	EndpointAutoConfirms *bool `json:"endpointAutoConfirms,omitempty" tf:"endpoint_auto_confirms,omitempty"`

	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the SNS docs for more details.
	FilterPolicy *string `json:"filterPolicy,omitempty" tf:"filter_policy,omitempty"`

	// Whether the filter_policy applies to MessageAttributes (default) or MessageBody.
	FilterPolicyScope *string `json:"filterPolicyScope,omitempty" tf:"filter_policy_scope,omitempty"`

	// ARN of the subscription.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// AWS account ID of the subscription's owner.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Whether the subscription has not been confirmed.
	PendingConfirmation *bool `json:"pendingConfirmation,omitempty" tf:"pending_confirmation,omitempty"`

	// Protocol to use. Valid values are: sqs, sms, lambda, firehose, and application. Protocols email, email-json, http and https are also valid but partially supported. See details below.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is false.
	RawMessageDelivery *bool `json:"rawMessageDelivery,omitempty" tf:"raw_message_delivery,omitempty"`

	// JSON String with the redrive policy that will be used in the subscription. Refer to the SNS docs for more details.
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// JSON String with the archived message replay policy that will be used in the subscription. Refer to the SNS docs for more details.
	ReplayPolicy *string `json:"replayPolicy,omitempty" tf:"replay_policy,omitempty"`

	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to SNS docs.
	SubscriptionRoleArn *string `json:"subscriptionRoleArn,omitempty" tf:"subscription_role_arn,omitempty"`

	// ARN of the SNS topic to subscribe to.
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type TopicSubscriptionParameters struct {

	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is 1.
	// +kubebuilder:validation:Optional
	ConfirmationTimeoutInMinutes *float64 `json:"confirmationTimeoutInMinutes,omitempty" tf:"confirmation_timeout_in_minutes,omitempty"`

	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the SNS docs for more details.
	// +kubebuilder:validation:Optional
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy,omitempty"`

	// Endpoint to send data to. The contents vary with the protocol. See details below.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Whether the endpoint is capable of auto confirming subscription (e.g., PagerDuty). Default is false.
	// +kubebuilder:validation:Optional
	EndpointAutoConfirms *bool `json:"endpointAutoConfirms,omitempty" tf:"endpoint_auto_confirms,omitempty"`

	// Reference to a Queue in sqs to populate endpoint.
	// +kubebuilder:validation:Optional
	EndpointRef *v1.Reference `json:"endpointRef,omitempty" tf:"-"`

	// Selector for a Queue in sqs to populate endpoint.
	// +kubebuilder:validation:Optional
	EndpointSelector *v1.Selector `json:"endpointSelector,omitempty" tf:"-"`

	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the SNS docs for more details.
	// +kubebuilder:validation:Optional
	FilterPolicy *string `json:"filterPolicy,omitempty" tf:"filter_policy,omitempty"`

	// Whether the filter_policy applies to MessageAttributes (default) or MessageBody.
	// +kubebuilder:validation:Optional
	FilterPolicyScope *string `json:"filterPolicyScope,omitempty" tf:"filter_policy_scope,omitempty"`

	// Protocol to use. Valid values are: sqs, sms, lambda, firehose, and application. Protocols email, email-json, http and https are also valid but partially supported. See details below.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is false.
	// +kubebuilder:validation:Optional
	RawMessageDelivery *bool `json:"rawMessageDelivery,omitempty" tf:"raw_message_delivery,omitempty"`

	// JSON String with the redrive policy that will be used in the subscription. Refer to the SNS docs for more details.
	// +kubebuilder:validation:Optional
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// JSON String with the archived message replay policy that will be used in the subscription. Refer to the SNS docs for more details.
	// +kubebuilder:validation:Optional
	ReplayPolicy *string `json:"replayPolicy,omitempty" tf:"replay_policy,omitempty"`

	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to SNS docs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SubscriptionRoleArn *string `json:"subscriptionRoleArn,omitempty" tf:"subscription_role_arn,omitempty"`

	// Reference to a Role in iam to populate subscriptionRoleArn.
	// +kubebuilder:validation:Optional
	SubscriptionRoleArnRef *v1.Reference `json:"subscriptionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate subscriptionRoleArn.
	// +kubebuilder:validation:Optional
	SubscriptionRoleArnSelector *v1.Selector `json:"subscriptionRoleArnSelector,omitempty" tf:"-"`

	// ARN of the SNS topic to subscribe to.
	// +crossplane:generate:reference:type=Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

// TopicSubscriptionSpec defines the desired state of TopicSubscription
type TopicSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicSubscriptionParameters `json:"forProvider"`
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
	InitProvider TopicSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// TopicSubscriptionStatus defines the observed state of TopicSubscription.
type TopicSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TopicSubscription is the Schema for the TopicSubscriptions API. Provides a resource for subscribing to SNS topics.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TopicSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   TopicSubscriptionSpec   `json:"spec"`
	Status TopicSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicSubscriptionList contains a list of TopicSubscriptions
type TopicSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicSubscription `json:"items"`
}

// Repository type metadata.
var (
	TopicSubscription_Kind             = "TopicSubscription"
	TopicSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicSubscription_Kind}.String()
	TopicSubscription_KindAPIVersion   = TopicSubscription_Kind + "." + CRDGroupVersion.String()
	TopicSubscription_GroupVersionKind = CRDGroupVersion.WithKind(TopicSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicSubscription{}, &TopicSubscriptionList{})
}
