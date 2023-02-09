/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	v1beta12 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AppServicePlan.
func (mg *AppServicePlan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FunctionApp.
func (mg *FunctionApp) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppServicePlanID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.AppServicePlanIDRef,
		Selector:     mg.Spec.ForProvider.AppServicePlanIDSelector,
		To: reference.To{
			List:    &AppServicePlanList{},
			Managed: &AppServicePlan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AppServicePlanID")
	}
	mg.Spec.ForProvider.AppServicePlanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AppServicePlanIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageAccountNameRef,
		Selector:     mg.Spec.ForProvider.StorageAccountNameSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountName")
	}
	mg.Spec.ForProvider.StorageAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FunctionAppActiveSlot.
func (mg *FunctionAppActiveSlot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SlotID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SlotIDRef,
		Selector:     mg.Spec.ForProvider.SlotIDSelector,
		To: reference.To{
			List:    &LinuxFunctionAppSlotList{},
			Managed: &LinuxFunctionAppSlot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SlotID")
	}
	mg.Spec.ForProvider.SlotID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SlotIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FunctionAppFunction.
func (mg *FunctionAppFunction) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionAppID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.FunctionAppIDRef,
		Selector:     mg.Spec.ForProvider.FunctionAppIDSelector,
		To: reference.To{
			List:    &LinuxFunctionAppList{},
			Managed: &LinuxFunctionApp{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionAppID")
	}
	mg.Spec.ForProvider.FunctionAppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionAppIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FunctionAppSlot.
func (mg *FunctionAppSlot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppServicePlanID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.AppServicePlanIDRef,
		Selector:     mg.Spec.ForProvider.AppServicePlanIDSelector,
		To: reference.To{
			List:    &AppServicePlanList{},
			Managed: &AppServicePlan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AppServicePlanID")
	}
	mg.Spec.ForProvider.AppServicePlanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AppServicePlanIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionAppName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.FunctionAppNameRef,
		Selector:     mg.Spec.ForProvider.FunctionAppNameSelector,
		To: reference.To{
			List:    &FunctionAppList{},
			Managed: &FunctionApp{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionAppName")
	}
	mg.Spec.ForProvider.FunctionAppName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionAppNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageAccountNameRef,
		Selector:     mg.Spec.ForProvider.StorageAccountNameSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountName")
	}
	mg.Spec.ForProvider.StorageAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LinuxFunctionApp.
func (mg *LinuxFunctionApp) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServicePlanID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServicePlanIDRef,
		Selector:     mg.Spec.ForProvider.ServicePlanIDSelector,
		To: reference.To{
			List:    &ServicePlanList{},
			Managed: &ServicePlan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServicePlanID")
	}
	mg.Spec.ForProvider.ServicePlanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServicePlanIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageAccountNameRef,
		Selector:     mg.Spec.ForProvider.StorageAccountNameSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountName")
	}
	mg.Spec.ForProvider.StorageAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkSubnetIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkSubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkSubnetID")
	}
	mg.Spec.ForProvider.VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LinuxFunctionAppSlot.
func (mg *LinuxFunctionAppSlot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionAppID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.FunctionAppIDRef,
		Selector:     mg.Spec.ForProvider.FunctionAppIDSelector,
		To: reference.To{
			List:    &LinuxFunctionAppList{},
			Managed: &LinuxFunctionApp{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionAppID")
	}
	mg.Spec.ForProvider.FunctionAppID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionAppIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageAccountNameRef,
		Selector:     mg.Spec.ForProvider.StorageAccountNameSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountName")
	}
	mg.Spec.ForProvider.StorageAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkSubnetIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkSubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkSubnetID")
	}
	mg.Spec.ForProvider.VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LinuxWebApp.
func (mg *LinuxWebApp) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServicePlanID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServicePlanIDRef,
		Selector:     mg.Spec.ForProvider.ServicePlanIDSelector,
		To: reference.To{
			List:    &ServicePlanList{},
			Managed: &ServicePlan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServicePlanID")
	}
	mg.Spec.ForProvider.ServicePlanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServicePlanIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkSubnetIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkSubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkSubnetID")
	}
	mg.Spec.ForProvider.VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LinuxWebAppSlot.
func (mg *LinuxWebAppSlot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppServiceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.AppServiceIDRef,
		Selector:     mg.Spec.ForProvider.AppServiceIDSelector,
		To: reference.To{
			List:    &LinuxWebAppList{},
			Managed: &LinuxWebApp{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AppServiceID")
	}
	mg.Spec.ForProvider.AppServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AppServiceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].IPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SiteConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef,
				Selector:     mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID")
			}
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SiteConfig[i3].ScmIPRestriction[i4].VirtualNetworkSubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkSubnetIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkSubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkSubnetID")
	}
	mg.Spec.ForProvider.VirtualNetworkSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServicePlan.
func (mg *ServicePlan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
