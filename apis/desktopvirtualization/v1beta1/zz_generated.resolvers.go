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
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/virtual/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this VirtualDesktopApplication.
func (mg *VirtualDesktopApplication) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ApplicationGroupIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationGroupIDSelector,
		To: reference.To{
			List:    &VirtualDesktopApplicationGroupList{},
			Managed: &VirtualDesktopApplicationGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationGroupID")
	}
	mg.Spec.ForProvider.ApplicationGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualDesktopApplicationGroup.
func (mg *VirtualDesktopApplicationGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.HostPoolIDRef,
		Selector:     mg.Spec.ForProvider.HostPoolIDSelector,
		To: reference.To{
			List:    &v1beta1.DesktopHostPoolList{},
			Managed: &v1beta1.DesktopHostPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostPoolID")
	}
	mg.Spec.ForProvider.HostPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostPoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualDesktopHostPoolRegistrationInfo.
func (mg *VirtualDesktopHostPoolRegistrationInfo) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostpoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.HostpoolIDRef,
		Selector:     mg.Spec.ForProvider.HostpoolIDSelector,
		To: reference.To{
			List:    &v1beta1.DesktopHostPoolList{},
			Managed: &v1beta1.DesktopHostPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostpoolID")
	}
	mg.Spec.ForProvider.HostpoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostpoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualDesktopScalingPlan.
func (mg *VirtualDesktopScalingPlan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.HostPool); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostPool[i3].HostpoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.HostPool[i3].HostpoolIDRef,
			Selector:     mg.Spec.ForProvider.HostPool[i3].HostpoolIDSelector,
			To: reference.To{
				List:    &v1beta1.DesktopHostPoolList{},
				Managed: &v1beta1.DesktopHostPool{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.HostPool[i3].HostpoolID")
		}
		mg.Spec.ForProvider.HostPool[i3].HostpoolID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.HostPool[i3].HostpoolIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualDesktopWorkspace.
func (mg *VirtualDesktopWorkspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualDesktopWorkspaceApplicationGroupAssociation.
func (mg *VirtualDesktopWorkspaceApplicationGroupAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ApplicationGroupIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationGroupIDSelector,
		To: reference.To{
			List:    &VirtualDesktopApplicationGroupList{},
			Managed: &VirtualDesktopApplicationGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationGroupID")
	}
	mg.Spec.ForProvider.ApplicationGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
		To: reference.To{
			List:    &VirtualDesktopWorkspaceList{},
			Managed: &VirtualDesktopWorkspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}
