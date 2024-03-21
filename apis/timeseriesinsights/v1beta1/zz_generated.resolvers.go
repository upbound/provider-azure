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
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *EventSourceEventHub) ResolveReferences( // ResolveReferences of this EventSourceEventHub.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "ConsumerGroup", "ConsumerGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConsumerGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ConsumerGroupNameRef,
			Selector:     mg.Spec.ForProvider.ConsumerGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConsumerGroupName")
	}
	mg.Spec.ForProvider.ConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConsumerGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("timeseriesinsights.azure.upbound.io", "v1beta1", "Gen2Environment", "Gen2EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
			Extract:      resource.ExtractResourceID(),
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
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "EventHub", "EventHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventHubNameRef,
			Selector:     mg.Spec.ForProvider.EventHubNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubName")
	}
	mg.Spec.ForProvider.EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "EventHub", "EventHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventSourceResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EventSourceResourceIDRef,
			Selector:     mg.Spec.ForProvider.EventSourceResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventSourceResourceID")
	}
	mg.Spec.ForProvider.EventSourceResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventSourceResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "EventHubNamespace", "EventHubNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NamespaceNameRef,
			Selector:     mg.Spec.ForProvider.NamespaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceName")
	}
	mg.Spec.ForProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "AuthorizationRule", "AuthorizationRuleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SharedAccessKeyName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SharedAccessKeyNameRef,
			Selector:     mg.Spec.ForProvider.SharedAccessKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SharedAccessKeyName")
	}
	mg.Spec.ForProvider.SharedAccessKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SharedAccessKeyNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "ConsumerGroup", "ConsumerGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConsumerGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ConsumerGroupNameRef,
			Selector:     mg.Spec.InitProvider.ConsumerGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConsumerGroupName")
	}
	mg.Spec.InitProvider.ConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConsumerGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "EventHub", "EventHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventHubNameRef,
			Selector:     mg.Spec.InitProvider.EventHubNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventHubName")
	}
	mg.Spec.InitProvider.EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventHubNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "EventHub", "EventHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventSourceResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EventSourceResourceIDRef,
			Selector:     mg.Spec.InitProvider.EventSourceResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventSourceResourceID")
	}
	mg.Spec.InitProvider.EventSourceResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventSourceResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "EventHubNamespace", "EventHubNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NamespaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NamespaceNameRef,
			Selector:     mg.Spec.InitProvider.NamespaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NamespaceName")
	}
	mg.Spec.InitProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NamespaceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "AuthorizationRule", "AuthorizationRuleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SharedAccessKeyName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.SharedAccessKeyNameRef,
			Selector:     mg.Spec.InitProvider.SharedAccessKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SharedAccessKeyName")
	}
	mg.Spec.InitProvider.SharedAccessKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SharedAccessKeyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventSourceIOTHub.
func (mg *EventSourceIOTHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHubConsumerGroup", "IOTHubConsumerGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConsumerGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ConsumerGroupNameRef,
			Selector:     mg.Spec.ForProvider.ConsumerGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConsumerGroupName")
	}
	mg.Spec.ForProvider.ConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConsumerGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("timeseriesinsights.azure.upbound.io", "v1beta1", "Gen2Environment", "Gen2EnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
			Extract:      resource.ExtractResourceID(),
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
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventSourceResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EventSourceResourceIDRef,
			Selector:     mg.Spec.ForProvider.EventSourceResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventSourceResourceID")
	}
	mg.Spec.ForProvider.EventSourceResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventSourceResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IOTHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.IOTHubNameRef,
			Selector:     mg.Spec.ForProvider.IOTHubNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IOTHubName")
	}
	mg.Spec.ForProvider.IOTHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IOTHubNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHubConsumerGroup", "IOTHubConsumerGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConsumerGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ConsumerGroupNameRef,
			Selector:     mg.Spec.InitProvider.ConsumerGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConsumerGroupName")
	}
	mg.Spec.InitProvider.ConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConsumerGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventSourceResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EventSourceResourceIDRef,
			Selector:     mg.Spec.InitProvider.EventSourceResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventSourceResourceID")
	}
	mg.Spec.InitProvider.EventSourceResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventSourceResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IOTHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.IOTHubNameRef,
			Selector:     mg.Spec.InitProvider.IOTHubNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IOTHubName")
	}
	mg.Spec.InitProvider.IOTHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IOTHubNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Gen2Environment.
func (mg *Gen2Environment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Storage); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Storage[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Storage[i3].NameRef,
				Selector:     mg.Spec.ForProvider.Storage[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Storage[i3].Name")
		}
		mg.Spec.ForProvider.Storage[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Storage[i3].NameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Storage); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Storage[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.Storage[i3].NameRef,
				Selector:     mg.Spec.InitProvider.Storage[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Storage[i3].Name")
		}
		mg.Spec.InitProvider.Storage[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Storage[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ReferenceDataSet.
func (mg *ReferenceDataSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("timeseriesinsights.azure.upbound.io", "v1beta1", "StandardEnvironment", "StandardEnvironmentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TimeSeriesInsightsEnvironmentID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TimeSeriesInsightsEnvironmentIDRef,
			Selector:     mg.Spec.ForProvider.TimeSeriesInsightsEnvironmentIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TimeSeriesInsightsEnvironmentID")
	}
	mg.Spec.ForProvider.TimeSeriesInsightsEnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TimeSeriesInsightsEnvironmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StandardEnvironment.
func (mg *StandardEnvironment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
