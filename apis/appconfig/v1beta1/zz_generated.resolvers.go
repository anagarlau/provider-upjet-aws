// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

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

func (mg *ConfigurationProfile) ResolveReferences( // ResolveReferences of this ConfigurationProfile.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ApplicationIDRef,
			Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RetrievalRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RetrievalRoleArnRef,
			Selector:     mg.Spec.ForProvider.RetrievalRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RetrievalRoleArn")
	}
	mg.Spec.ForProvider.RetrievalRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RetrievalRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ApplicationIDRef,
			Selector:     mg.Spec.InitProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationID")
	}
	mg.Spec.InitProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RetrievalRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RetrievalRoleArnRef,
			Selector:     mg.Spec.InitProvider.RetrievalRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RetrievalRoleArn")
	}
	mg.Spec.InitProvider.RetrievalRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RetrievalRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Deployment.
func (mg *Deployment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ApplicationIDRef,
			Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "ConfigurationProfile", "ConfigurationProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationProfileID),
			Extract:      resource.ExtractParamPath("configuration_profile_id", true),
			Reference:    mg.Spec.ForProvider.ConfigurationProfileIDRef,
			Selector:     mg.Spec.ForProvider.ConfigurationProfileIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationProfileID")
	}
	mg.Spec.ForProvider.ConfigurationProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationProfileIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "HostedConfigurationVersion", "HostedConfigurationVersionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationVersion),
			Extract:      resource.ExtractParamPath("version_number", true),
			Reference:    mg.Spec.ForProvider.ConfigurationVersionRef,
			Selector:     mg.Spec.ForProvider.ConfigurationVersionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationVersion")
	}
	mg.Spec.ForProvider.ConfigurationVersion = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationVersionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "DeploymentStrategy", "DeploymentStrategyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeploymentStrategyID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DeploymentStrategyIDRef,
			Selector:     mg.Spec.ForProvider.DeploymentStrategyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeploymentStrategyID")
	}
	mg.Spec.ForProvider.DeploymentStrategyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeploymentStrategyIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Environment", "EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
			Extract:      resource.ExtractParamPath("environment_id", true),
			Reference:    mg.Spec.ForProvider.EnvironmentIDRef,
			Selector:     mg.Spec.ForProvider.EnvironmentIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvironmentID")
	}
	mg.Spec.ForProvider.EnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyIdentifier),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.KMSKeyIdentifierRef,
			Selector:     mg.Spec.ForProvider.KMSKeyIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyIdentifier")
	}
	mg.Spec.ForProvider.KMSKeyIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ApplicationIDRef,
			Selector:     mg.Spec.InitProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationID")
	}
	mg.Spec.InitProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "ConfigurationProfile", "ConfigurationProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationProfileID),
			Extract:      resource.ExtractParamPath("configuration_profile_id", true),
			Reference:    mg.Spec.InitProvider.ConfigurationProfileIDRef,
			Selector:     mg.Spec.InitProvider.ConfigurationProfileIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationProfileID")
	}
	mg.Spec.InitProvider.ConfigurationProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationProfileIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "HostedConfigurationVersion", "HostedConfigurationVersionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationVersion),
			Extract:      resource.ExtractParamPath("version_number", true),
			Reference:    mg.Spec.InitProvider.ConfigurationVersionRef,
			Selector:     mg.Spec.InitProvider.ConfigurationVersionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationVersion")
	}
	mg.Spec.InitProvider.ConfigurationVersion = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationVersionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "DeploymentStrategy", "DeploymentStrategyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DeploymentStrategyID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DeploymentStrategyIDRef,
			Selector:     mg.Spec.InitProvider.DeploymentStrategyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DeploymentStrategyID")
	}
	mg.Spec.InitProvider.DeploymentStrategyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DeploymentStrategyIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Environment", "EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EnvironmentID),
			Extract:      resource.ExtractParamPath("environment_id", true),
			Reference:    mg.Spec.InitProvider.EnvironmentIDRef,
			Selector:     mg.Spec.InitProvider.EnvironmentIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EnvironmentID")
	}
	mg.Spec.InitProvider.EnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EnvironmentIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyIdentifier),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.KMSKeyIdentifierRef,
			Selector:     mg.Spec.InitProvider.KMSKeyIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyIdentifier")
	}
	mg.Spec.InitProvider.KMSKeyIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Environment.
func (mg *Environment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ApplicationIDRef,
			Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Monitor); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudwatch.aws.upbound.io", "v1beta1", "MetricAlarm", "MetricAlarmList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Monitor[i3].AlarmArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Monitor[i3].AlarmArnRef,
				Selector:     mg.Spec.ForProvider.Monitor[i3].AlarmArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Monitor[i3].AlarmArn")
		}
		mg.Spec.ForProvider.Monitor[i3].AlarmArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Monitor[i3].AlarmArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Monitor); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Monitor[i3].AlarmRoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Monitor[i3].AlarmRoleArnRef,
				Selector:     mg.Spec.ForProvider.Monitor[i3].AlarmRoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Monitor[i3].AlarmRoleArn")
		}
		mg.Spec.ForProvider.Monitor[i3].AlarmRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Monitor[i3].AlarmRoleArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ApplicationIDRef,
			Selector:     mg.Spec.InitProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationID")
	}
	mg.Spec.InitProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Monitor); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudwatch.aws.upbound.io", "v1beta1", "MetricAlarm", "MetricAlarmList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Monitor[i3].AlarmArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Monitor[i3].AlarmArnRef,
				Selector:     mg.Spec.InitProvider.Monitor[i3].AlarmArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Monitor[i3].AlarmArn")
		}
		mg.Spec.InitProvider.Monitor[i3].AlarmArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Monitor[i3].AlarmArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Monitor); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Monitor[i3].AlarmRoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Monitor[i3].AlarmRoleArnRef,
				Selector:     mg.Spec.InitProvider.Monitor[i3].AlarmRoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Monitor[i3].AlarmRoleArn")
		}
		mg.Spec.InitProvider.Monitor[i3].AlarmRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Monitor[i3].AlarmRoleArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Extension.
func (mg *Extension) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ActionPoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ActionPoint[i3].Action); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArnRef,
					Selector:     mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArn")
			}
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ActionPoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ActionPoint[i3].Action); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URI),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URIRef,
					Selector:     mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URISelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URI")
			}
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URI = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URIRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ActionPoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ActionPoint[i3].Action); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ActionPoint[i3].Action[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.ActionPoint[i3].Action[i4].RoleArnRef,
					Selector:     mg.Spec.InitProvider.ActionPoint[i3].Action[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ActionPoint[i3].Action[i4].RoleArn")
			}
			mg.Spec.InitProvider.ActionPoint[i3].Action[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ActionPoint[i3].Action[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ActionPoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ActionPoint[i3].Action); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ActionPoint[i3].Action[i4].URI),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.ActionPoint[i3].Action[i4].URIRef,
					Selector:     mg.Spec.InitProvider.ActionPoint[i3].Action[i4].URISelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ActionPoint[i3].Action[i4].URI")
			}
			mg.Spec.InitProvider.ActionPoint[i3].Action[i4].URI = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ActionPoint[i3].Action[i4].URIRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this ExtensionAssociation.
func (mg *ExtensionAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Extension", "ExtensionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExtensionArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ExtensionArnRef,
			Selector:     mg.Spec.ForProvider.ExtensionArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExtensionArn")
	}
	mg.Spec.ForProvider.ExtensionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExtensionArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ResourceArnRef,
			Selector:     mg.Spec.ForProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Extension", "ExtensionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ExtensionArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ExtensionArnRef,
			Selector:     mg.Spec.InitProvider.ExtensionArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ExtensionArn")
	}
	mg.Spec.InitProvider.ExtensionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ExtensionArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ResourceArnRef,
			Selector:     mg.Spec.InitProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceArn")
	}
	mg.Spec.InitProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedConfigurationVersion.
func (mg *HostedConfigurationVersion) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ApplicationIDRef,
			Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "ConfigurationProfile", "ConfigurationProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationProfileID),
			Extract:      resource.ExtractParamPath("configuration_profile_id", true),
			Reference:    mg.Spec.ForProvider.ConfigurationProfileIDRef,
			Selector:     mg.Spec.ForProvider.ConfigurationProfileIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationProfileID")
	}
	mg.Spec.ForProvider.ConfigurationProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationProfileIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ApplicationIDRef,
			Selector:     mg.Spec.InitProvider.ApplicationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationID")
	}
	mg.Spec.InitProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appconfig.aws.upbound.io", "v1beta1", "ConfigurationProfile", "ConfigurationProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationProfileID),
			Extract:      resource.ExtractParamPath("configuration_profile_id", true),
			Reference:    mg.Spec.InitProvider.ConfigurationProfileIDRef,
			Selector:     mg.Spec.InitProvider.ConfigurationProfileIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationProfileID")
	}
	mg.Spec.InitProvider.ConfigurationProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationProfileIDRef = rsp.ResolvedReference

	return nil
}
