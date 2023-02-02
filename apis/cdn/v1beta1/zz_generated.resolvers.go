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
	v1beta12 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Endpoint.
func (mg *Endpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProfileName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProfileNameRef,
		Selector:     mg.Spec.ForProvider.ProfileNameSelector,
		To: reference.To{
			List:    &ProfileList{},
			Managed: &Profile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProfileName")
	}
	mg.Spec.ForProvider.ProfileName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProfileNameRef = rsp.ResolvedReference

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

// ResolveReferences of this FrontdoorCustomDomain.
func (mg *FrontdoorCustomDomain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorProfileID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorProfileIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorProfileIDSelector,
		To: reference.To{
			List:    &FrontdoorProfileList{},
			Managed: &FrontdoorProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorProfileID")
	}
	mg.Spec.ForProvider.CdnFrontdoorProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorProfileIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DNSZoneID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DNSZoneIDRef,
		Selector:     mg.Spec.ForProvider.DNSZoneIDSelector,
		To: reference.To{
			List:    &v1beta11.DNSZoneList{},
			Managed: &v1beta11.DNSZone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DNSZoneID")
	}
	mg.Spec.ForProvider.DNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DNSZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FrontdoorCustomDomainAssociation.
func (mg *FrontdoorCustomDomainAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorCustomDomainID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorCustomDomainIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorCustomDomainIDSelector,
		To: reference.To{
			List:    &FrontdoorCustomDomainList{},
			Managed: &FrontdoorCustomDomain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorCustomDomainID")
	}
	mg.Spec.ForProvider.CdnFrontdoorCustomDomainID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorCustomDomainIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CdnFrontdoorRouteIds),
		Extract:       resource.ExtractResourceID(),
		References:    mg.Spec.ForProvider.CdnFrontdoorRouteIdsRefs,
		Selector:      mg.Spec.ForProvider.CdnFrontdoorRouteIdsSelector,
		To: reference.To{
			List:    &FrontdoorRouteList{},
			Managed: &FrontdoorRoute{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorRouteIds")
	}
	mg.Spec.ForProvider.CdnFrontdoorRouteIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CdnFrontdoorRouteIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this FrontdoorEndpoint.
func (mg *FrontdoorEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorProfileID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorProfileIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorProfileIDSelector,
		To: reference.To{
			List:    &FrontdoorProfileList{},
			Managed: &FrontdoorProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorProfileID")
	}
	mg.Spec.ForProvider.CdnFrontdoorProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorProfileIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FrontdoorOrigin.
func (mg *FrontdoorOrigin) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorOriginGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorOriginGroupIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorOriginGroupIDSelector,
		To: reference.To{
			List:    &FrontdoorOriginGroupList{},
			Managed: &FrontdoorOriginGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorOriginGroupID")
	}
	mg.Spec.ForProvider.CdnFrontdoorOriginGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorOriginGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostName),
		Extract:      resource.ExtractParamPath("primary_blob_host", true),
		Reference:    mg.Spec.ForProvider.HostNameRef,
		Selector:     mg.Spec.ForProvider.HostNameSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostName")
	}
	mg.Spec.ForProvider.HostName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OriginHostHeader),
		Extract:      resource.ExtractParamPath("primary_blob_host", true),
		Reference:    mg.Spec.ForProvider.OriginHostHeaderRef,
		Selector:     mg.Spec.ForProvider.OriginHostHeaderSelector,
		To: reference.To{
			List:    &v1beta12.AccountList{},
			Managed: &v1beta12.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OriginHostHeader")
	}
	mg.Spec.ForProvider.OriginHostHeader = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OriginHostHeaderRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PrivateLink); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateLink[i3].Location),
			Extract:      resource.ExtractParamPath("location", false),
			Reference:    mg.Spec.ForProvider.PrivateLink[i3].LocationRef,
			Selector:     mg.Spec.ForProvider.PrivateLink[i3].LocationSelector,
			To: reference.To{
				List:    &v1beta12.AccountList{},
				Managed: &v1beta12.Account{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PrivateLink[i3].Location")
		}
		mg.Spec.ForProvider.PrivateLink[i3].Location = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PrivateLink[i3].LocationRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.PrivateLink); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateLink[i3].PrivateLinkTargetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PrivateLink[i3].PrivateLinkTargetIDRef,
			Selector:     mg.Spec.ForProvider.PrivateLink[i3].PrivateLinkTargetIDSelector,
			To: reference.To{
				List:    &v1beta12.AccountList{},
				Managed: &v1beta12.Account{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PrivateLink[i3].PrivateLinkTargetID")
		}
		mg.Spec.ForProvider.PrivateLink[i3].PrivateLinkTargetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PrivateLink[i3].PrivateLinkTargetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this FrontdoorOriginGroup.
func (mg *FrontdoorOriginGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorProfileID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorProfileIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorProfileIDSelector,
		To: reference.To{
			List:    &FrontdoorProfileList{},
			Managed: &FrontdoorProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorProfileID")
	}
	mg.Spec.ForProvider.CdnFrontdoorProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorProfileIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FrontdoorProfile.
func (mg *FrontdoorProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
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

// ResolveReferences of this FrontdoorRoute.
func (mg *FrontdoorRoute) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CdnFrontdoorCustomDomainIds),
		Extract:       resource.ExtractResourceID(),
		References:    mg.Spec.ForProvider.CdnFrontdoorCustomDomainIdsRefs,
		Selector:      mg.Spec.ForProvider.CdnFrontdoorCustomDomainIdsSelector,
		To: reference.To{
			List:    &FrontdoorCustomDomainList{},
			Managed: &FrontdoorCustomDomain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorCustomDomainIds")
	}
	mg.Spec.ForProvider.CdnFrontdoorCustomDomainIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CdnFrontdoorCustomDomainIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorEndpointID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorEndpointIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorEndpointIDSelector,
		To: reference.To{
			List:    &FrontdoorEndpointList{},
			Managed: &FrontdoorEndpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorEndpointID")
	}
	mg.Spec.ForProvider.CdnFrontdoorEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorEndpointIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorOriginGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorOriginGroupIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorOriginGroupIDSelector,
		To: reference.To{
			List:    &FrontdoorOriginGroupList{},
			Managed: &FrontdoorOriginGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorOriginGroupID")
	}
	mg.Spec.ForProvider.CdnFrontdoorOriginGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorOriginGroupIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CdnFrontdoorOriginIds),
		Extract:       resource.ExtractResourceID(),
		References:    mg.Spec.ForProvider.CdnFrontdoorOriginIdsRefs,
		Selector:      mg.Spec.ForProvider.CdnFrontdoorOriginIdsSelector,
		To: reference.To{
			List:    &FrontdoorOriginList{},
			Managed: &FrontdoorOrigin{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorOriginIds")
	}
	mg.Spec.ForProvider.CdnFrontdoorOriginIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CdnFrontdoorOriginIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CdnFrontdoorRuleSetIds),
		Extract:       resource.ExtractResourceID(),
		References:    mg.Spec.ForProvider.CdnFrontdoorRuleSetIdsRefs,
		Selector:      mg.Spec.ForProvider.CdnFrontdoorRuleSetIdsSelector,
		To: reference.To{
			List:    &FrontdoorRuleSetList{},
			Managed: &FrontdoorRuleSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorRuleSetIds")
	}
	mg.Spec.ForProvider.CdnFrontdoorRuleSetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CdnFrontdoorRuleSetIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this FrontdoorRule.
func (mg *FrontdoorRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Actions[i3].RouteConfigurationOverrideAction); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].RouteConfigurationOverrideAction[i4].CdnFrontdoorOriginGroupID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Actions[i3].RouteConfigurationOverrideAction[i4].CdnFrontdoorOriginGroupIDRef,
				Selector:     mg.Spec.ForProvider.Actions[i3].RouteConfigurationOverrideAction[i4].CdnFrontdoorOriginGroupIDSelector,
				To: reference.To{
					List:    &FrontdoorOriginGroupList{},
					Managed: &FrontdoorOriginGroup{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].RouteConfigurationOverrideAction[i4].CdnFrontdoorOriginGroupID")
			}
			mg.Spec.ForProvider.Actions[i3].RouteConfigurationOverrideAction[i4].CdnFrontdoorOriginGroupID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Actions[i3].RouteConfigurationOverrideAction[i4].CdnFrontdoorOriginGroupIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorRuleSetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorRuleSetIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorRuleSetIDSelector,
		To: reference.To{
			List:    &FrontdoorRuleSetList{},
			Managed: &FrontdoorRuleSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorRuleSetID")
	}
	mg.Spec.ForProvider.CdnFrontdoorRuleSetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorRuleSetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FrontdoorRuleSet.
func (mg *FrontdoorRuleSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CdnFrontdoorProfileID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CdnFrontdoorProfileIDRef,
		Selector:     mg.Spec.ForProvider.CdnFrontdoorProfileIDSelector,
		To: reference.To{
			List:    &FrontdoorProfileList{},
			Managed: &FrontdoorProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CdnFrontdoorProfileID")
	}
	mg.Spec.ForProvider.CdnFrontdoorProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CdnFrontdoorProfileIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Profile.
func (mg *Profile) ResolveReferences(ctx context.Context, c client.Reader) error {
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
