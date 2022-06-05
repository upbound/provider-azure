//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataOptionsObservation) DeepCopyInto(out *ExportDataOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataOptionsObservation.
func (in *ExportDataOptionsObservation) DeepCopy() *ExportDataOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(ExportDataOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataOptionsParameters) DeepCopyInto(out *ExportDataOptionsParameters) {
	*out = *in
	if in.TimeFrame != nil {
		in, out := &in.TimeFrame, &out.TimeFrame
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataOptionsParameters.
func (in *ExportDataOptionsParameters) DeepCopy() *ExportDataOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(ExportDataOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataStorageLocationObservation) DeepCopyInto(out *ExportDataStorageLocationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataStorageLocationObservation.
func (in *ExportDataStorageLocationObservation) DeepCopy() *ExportDataStorageLocationObservation {
	if in == nil {
		return nil
	}
	out := new(ExportDataStorageLocationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportDataStorageLocationParameters) DeepCopyInto(out *ExportDataStorageLocationParameters) {
	*out = *in
	if in.ContainerID != nil {
		in, out := &in.ContainerID, &out.ContainerID
		*out = new(string)
		**out = **in
	}
	if in.RootFolderPath != nil {
		in, out := &in.RootFolderPath, &out.RootFolderPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportDataStorageLocationParameters.
func (in *ExportDataStorageLocationParameters) DeepCopy() *ExportDataStorageLocationParameters {
	if in == nil {
		return nil
	}
	out := new(ExportDataStorageLocationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupCostManagementExport) DeepCopyInto(out *GroupCostManagementExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupCostManagementExport.
func (in *GroupCostManagementExport) DeepCopy() *GroupCostManagementExport {
	if in == nil {
		return nil
	}
	out := new(GroupCostManagementExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupCostManagementExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupCostManagementExportList) DeepCopyInto(out *GroupCostManagementExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupCostManagementExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupCostManagementExportList.
func (in *GroupCostManagementExportList) DeepCopy() *GroupCostManagementExportList {
	if in == nil {
		return nil
	}
	out := new(GroupCostManagementExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupCostManagementExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupCostManagementExportObservation) DeepCopyInto(out *GroupCostManagementExportObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupCostManagementExportObservation.
func (in *GroupCostManagementExportObservation) DeepCopy() *GroupCostManagementExportObservation {
	if in == nil {
		return nil
	}
	out := new(GroupCostManagementExportObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupCostManagementExportParameters) DeepCopyInto(out *GroupCostManagementExportParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.ExportDataOptions != nil {
		in, out := &in.ExportDataOptions, &out.ExportDataOptions
		*out = make([]ExportDataOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExportDataStorageLocation != nil {
		in, out := &in.ExportDataStorageLocation, &out.ExportDataStorageLocation
		*out = make([]ExportDataStorageLocationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecurrencePeriodEndDate != nil {
		in, out := &in.RecurrencePeriodEndDate, &out.RecurrencePeriodEndDate
		*out = new(string)
		**out = **in
	}
	if in.RecurrencePeriodStartDate != nil {
		in, out := &in.RecurrencePeriodStartDate, &out.RecurrencePeriodStartDate
		*out = new(string)
		**out = **in
	}
	if in.RecurrenceType != nil {
		in, out := &in.RecurrenceType, &out.RecurrenceType
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupIDRef != nil {
		in, out := &in.ResourceGroupIDRef, &out.ResourceGroupIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupIDSelector != nil {
		in, out := &in.ResourceGroupIDSelector, &out.ResourceGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupCostManagementExportParameters.
func (in *GroupCostManagementExportParameters) DeepCopy() *GroupCostManagementExportParameters {
	if in == nil {
		return nil
	}
	out := new(GroupCostManagementExportParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupCostManagementExportSpec) DeepCopyInto(out *GroupCostManagementExportSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupCostManagementExportSpec.
func (in *GroupCostManagementExportSpec) DeepCopy() *GroupCostManagementExportSpec {
	if in == nil {
		return nil
	}
	out := new(GroupCostManagementExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupCostManagementExportStatus) DeepCopyInto(out *GroupCostManagementExportStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupCostManagementExportStatus.
func (in *GroupCostManagementExportStatus) DeepCopy() *GroupCostManagementExportStatus {
	if in == nil {
		return nil
	}
	out := new(GroupCostManagementExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyExemption) DeepCopyInto(out *GroupPolicyExemption) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyExemption.
func (in *GroupPolicyExemption) DeepCopy() *GroupPolicyExemption {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyExemption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupPolicyExemption) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyExemptionList) DeepCopyInto(out *GroupPolicyExemptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupPolicyExemption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyExemptionList.
func (in *GroupPolicyExemptionList) DeepCopy() *GroupPolicyExemptionList {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyExemptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupPolicyExemptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyExemptionObservation) DeepCopyInto(out *GroupPolicyExemptionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyExemptionObservation.
func (in *GroupPolicyExemptionObservation) DeepCopy() *GroupPolicyExemptionObservation {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyExemptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyExemptionParameters) DeepCopyInto(out *GroupPolicyExemptionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ExemptionCategory != nil {
		in, out := &in.ExemptionCategory, &out.ExemptionCategory
		*out = new(string)
		**out = **in
	}
	if in.ExpiresOn != nil {
		in, out := &in.ExpiresOn, &out.ExpiresOn
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyAssignmentID != nil {
		in, out := &in.PolicyAssignmentID, &out.PolicyAssignmentID
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionReferenceIds != nil {
		in, out := &in.PolicyDefinitionReferenceIds, &out.PolicyDefinitionReferenceIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupIDRef != nil {
		in, out := &in.ResourceGroupIDRef, &out.ResourceGroupIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupIDSelector != nil {
		in, out := &in.ResourceGroupIDSelector, &out.ResourceGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyExemptionParameters.
func (in *GroupPolicyExemptionParameters) DeepCopy() *GroupPolicyExemptionParameters {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyExemptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyExemptionSpec) DeepCopyInto(out *GroupPolicyExemptionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyExemptionSpec.
func (in *GroupPolicyExemptionSpec) DeepCopy() *GroupPolicyExemptionSpec {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyExemptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyExemptionStatus) DeepCopyInto(out *GroupPolicyExemptionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyExemptionStatus.
func (in *GroupPolicyExemptionStatus) DeepCopy() *GroupPolicyExemptionStatus {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyExemptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyRemediation) DeepCopyInto(out *GroupPolicyRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyRemediation.
func (in *GroupPolicyRemediation) DeepCopy() *GroupPolicyRemediation {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupPolicyRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyRemediationList) DeepCopyInto(out *GroupPolicyRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupPolicyRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyRemediationList.
func (in *GroupPolicyRemediationList) DeepCopy() *GroupPolicyRemediationList {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupPolicyRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyRemediationObservation) DeepCopyInto(out *GroupPolicyRemediationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyRemediationObservation.
func (in *GroupPolicyRemediationObservation) DeepCopy() *GroupPolicyRemediationObservation {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyRemediationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyRemediationParameters) DeepCopyInto(out *GroupPolicyRemediationParameters) {
	*out = *in
	if in.LocationFilters != nil {
		in, out := &in.LocationFilters, &out.LocationFilters
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyAssignmentID != nil {
		in, out := &in.PolicyAssignmentID, &out.PolicyAssignmentID
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionID != nil {
		in, out := &in.PolicyDefinitionID, &out.PolicyDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.ResourceDiscoveryMode != nil {
		in, out := &in.ResourceDiscoveryMode, &out.ResourceDiscoveryMode
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupIDRef != nil {
		in, out := &in.ResourceGroupIDRef, &out.ResourceGroupIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupIDSelector != nil {
		in, out := &in.ResourceGroupIDSelector, &out.ResourceGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyRemediationParameters.
func (in *GroupPolicyRemediationParameters) DeepCopy() *GroupPolicyRemediationParameters {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyRemediationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyRemediationSpec) DeepCopyInto(out *GroupPolicyRemediationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyRemediationSpec.
func (in *GroupPolicyRemediationSpec) DeepCopy() *GroupPolicyRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupPolicyRemediationStatus) DeepCopyInto(out *GroupPolicyRemediationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupPolicyRemediationStatus.
func (in *GroupPolicyRemediationStatus) DeepCopy() *GroupPolicyRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(GroupPolicyRemediationStatus)
	in.DeepCopyInto(out)
	return out
}
