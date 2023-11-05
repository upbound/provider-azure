/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ManagementGroup.
func (mg *ManagementGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParentManagementGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ParentManagementGroupIDRef,
		Selector:     mg.Spec.ForProvider.ParentManagementGroupIDSelector,
		To: reference.To{
			List:    &ManagementGroupList{},
			Managed: &ManagementGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParentManagementGroupID")
	}
	mg.Spec.ForProvider.ParentManagementGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentManagementGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ManagementGroupSubscriptionAssociation.
func (mg *ManagementGroupSubscriptionAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagementGroupID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ManagementGroupIDRef,
		Selector:     mg.Spec.ForProvider.ManagementGroupIDSelector,
		To: reference.To{
			List:    &ManagementGroupList{},
			Managed: &ManagementGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagementGroupID")
	}
	mg.Spec.ForProvider.ManagementGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManagementGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubscriptionID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubscriptionIDRef,
		Selector:     mg.Spec.ForProvider.SubscriptionIDSelector,
		To: reference.To{
			List:    &v1beta1.SubscriptionList{},
			Managed: &v1beta1.Subscription{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubscriptionID")
	}
	mg.Spec.ForProvider.SubscriptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubscriptionIDRef = rsp.ResolvedReference

	return nil
}
