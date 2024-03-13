// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Attachment) ResolveReferences( // ResolveReferences of this Attachment.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta2", "AutoscalingGroup", "AutoscalingGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
			Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io", "v1beta1", "ELB", "ELBList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ELB),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ELBRef,
			Selector:     mg.Spec.ForProvider.ELBSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ELB")
	}
	mg.Spec.ForProvider.ELB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ELBRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta1", "LBTargetGroup", "LBTargetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBTargetGroupArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.LBTargetGroupArnRef,
			Selector:     mg.Spec.ForProvider.LBTargetGroupArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBTargetGroupArn")
	}
	mg.Spec.ForProvider.LBTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBTargetGroupArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta2", "AutoscalingGroup", "AutoscalingGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AutoscalingGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.AutoscalingGroupNameRef,
			Selector:     mg.Spec.InitProvider.AutoscalingGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AutoscalingGroupName")
	}
	mg.Spec.InitProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AutoscalingGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elb.aws.upbound.io", "v1beta1", "ELB", "ELBList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ELB),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ELBRef,
			Selector:     mg.Spec.InitProvider.ELBSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ELB")
	}
	mg.Spec.InitProvider.ELB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ELBRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta1", "LBTargetGroup", "LBTargetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LBTargetGroupArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.LBTargetGroupArnRef,
			Selector:     mg.Spec.InitProvider.LBTargetGroupArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LBTargetGroupArn")
	}
	mg.Spec.InitProvider.LBTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LBTargetGroupArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AutoscalingGroup.
func (mg *AutoscalingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta1", "LaunchConfiguration", "LaunchConfigurationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LaunchConfiguration),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LaunchConfigurationRef,
			Selector:     mg.Spec.ForProvider.LaunchConfigurationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LaunchConfiguration")
	}
	mg.Spec.ForProvider.LaunchConfiguration = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LaunchConfigurationRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LaunchTemplate); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "LaunchTemplate", "LaunchTemplateList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LaunchTemplate[i3].ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.LaunchTemplate[i3].IDRef,
				Selector:     mg.Spec.ForProvider.LaunchTemplate[i3].IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LaunchTemplate[i3].ID")
		}
		mg.Spec.ForProvider.LaunchTemplate[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LaunchTemplate[i3].IDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "LaunchTemplate", "LaunchTemplateList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef,
						Selector:     mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID")
				}
				mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "LaunchTemplate", "LaunchTemplateList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef,
							Selector:     mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID")
					}
					mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "PlacementGroup", "PlacementGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PlacementGroup),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PlacementGroupRef,
			Selector:     mg.Spec.ForProvider.PlacementGroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PlacementGroup")
	}
	mg.Spec.ForProvider.PlacementGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PlacementGroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceLinkedRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceLinkedRoleArnRef,
			Selector:     mg.Spec.ForProvider.ServiceLinkedRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceLinkedRoleArn")
	}
	mg.Spec.ForProvider.ServiceLinkedRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceLinkedRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCZoneIdentifier),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCZoneIdentifierRefs,
			Selector:      mg.Spec.ForProvider.VPCZoneIdentifierSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCZoneIdentifier")
	}
	mg.Spec.ForProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("autoscaling.aws.upbound.io", "v1beta1", "LaunchConfiguration", "LaunchConfigurationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LaunchConfiguration),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LaunchConfigurationRef,
			Selector:     mg.Spec.InitProvider.LaunchConfigurationSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LaunchConfiguration")
	}
	mg.Spec.InitProvider.LaunchConfiguration = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LaunchConfigurationRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.LaunchTemplate); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "LaunchTemplate", "LaunchTemplateList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LaunchTemplate[i3].ID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.LaunchTemplate[i3].IDRef,
				Selector:     mg.Spec.InitProvider.LaunchTemplate[i3].IDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LaunchTemplate[i3].ID")
		}
		mg.Spec.InitProvider.LaunchTemplate[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LaunchTemplate[i3].IDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "LaunchTemplate", "LaunchTemplateList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef,
						Selector:     mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID")
				}
				mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "LaunchTemplate", "LaunchTemplateList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID),
							Extract:      resource.ExtractResourceID(),
							Reference:    mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef,
							Selector:     mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID")
					}
					mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "PlacementGroup", "PlacementGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PlacementGroup),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PlacementGroupRef,
			Selector:     mg.Spec.InitProvider.PlacementGroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PlacementGroup")
	}
	mg.Spec.InitProvider.PlacementGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PlacementGroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceLinkedRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceLinkedRoleArnRef,
			Selector:     mg.Spec.InitProvider.ServiceLinkedRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceLinkedRoleArn")
	}
	mg.Spec.InitProvider.ServiceLinkedRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceLinkedRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCZoneIdentifier),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.VPCZoneIdentifierRefs,
			Selector:      mg.Spec.InitProvider.VPCZoneIdentifierSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCZoneIdentifier")
	}
	mg.Spec.InitProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences

	return nil
}
