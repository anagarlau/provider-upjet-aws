// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Table.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Table) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyspaces.aws.upbound.io", "v1beta1", "Keyspace", "KeyspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyspaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KeyspaceNameRef,
			Selector:     mg.Spec.ForProvider.KeyspaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyspaceName")
	}
	mg.Spec.ForProvider.KeyspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyspaceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyspaces.aws.upbound.io", "v1beta1", "Keyspace", "KeyspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyspaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.KeyspaceNameRef,
			Selector:     mg.Spec.InitProvider.KeyspaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyspaceName")
	}
	mg.Spec.InitProvider.KeyspaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyspaceNameRef = rsp.ResolvedReference

	return nil
}
