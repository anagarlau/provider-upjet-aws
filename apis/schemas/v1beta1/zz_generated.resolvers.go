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
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Discoverer) ResolveReferences( // ResolveReferences of this Discoverer.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.SourceArnRef,
			Selector:     mg.Spec.ForProvider.SourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceArn")
	}
	mg.Spec.ForProvider.SourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.SourceArnRef,
			Selector:     mg.Spec.InitProvider.SourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SourceArn")
	}
	mg.Spec.InitProvider.SourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Schema.
func (mg *Schema) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("schemas.aws.upbound.io", "v1beta1", "Registry", "RegistryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RegistryName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RegistryNameRef,
			Selector:     mg.Spec.ForProvider.RegistryNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RegistryName")
	}
	mg.Spec.ForProvider.RegistryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RegistryNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("schemas.aws.upbound.io", "v1beta1", "Registry", "RegistryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RegistryName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RegistryNameRef,
			Selector:     mg.Spec.InitProvider.RegistryNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RegistryName")
	}
	mg.Spec.InitProvider.RegistryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RegistryNameRef = rsp.ResolvedReference

	return nil
}
