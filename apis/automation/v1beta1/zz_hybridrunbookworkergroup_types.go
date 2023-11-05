// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HybridRunBookWorkerGroupInitParameters struct {

	// The name of the Automation Account in which the Runbook Worker Group is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The name of resource type azurerm_automation_credential to use for hybrid worker.
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	// The name which should be used for this Automation Account Runbook Worker Group. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type HybridRunBookWorkerGroupObservation struct {

	// The name of the Automation Account in which the Runbook Worker Group is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The name of resource type azurerm_automation_credential to use for hybrid worker.
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	// The ID of the Automation Hybrid Runbook Worker Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name which should be used for this Automation Account Runbook Worker Group. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Resource Group where the Automation should exist. Changing this forces a new Automation to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type HybridRunBookWorkerGroupParameters struct {

	// The name of the Automation Account in which the Runbook Worker Group is created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The name of resource type azurerm_automation_credential to use for hybrid worker.
	// +kubebuilder:validation:Optional
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	// The name which should be used for this Automation Account Runbook Worker Group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Resource Group where the Automation should exist. Changing this forces a new Automation to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// HybridRunBookWorkerGroupSpec defines the desired state of HybridRunBookWorkerGroup
type HybridRunBookWorkerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HybridRunBookWorkerGroupParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HybridRunBookWorkerGroupInitParameters `json:"initProvider,omitempty"`
}

// HybridRunBookWorkerGroupStatus defines the observed state of HybridRunBookWorkerGroup.
type HybridRunBookWorkerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HybridRunBookWorkerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HybridRunBookWorkerGroup is the Schema for the HybridRunBookWorkerGroups API. Manages a Automation Account Runbook Worker Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HybridRunBookWorkerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.automationAccountName) || (has(self.initProvider) && has(self.initProvider.automationAccountName))",message="spec.forProvider.automationAccountName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   HybridRunBookWorkerGroupSpec   `json:"spec"`
	Status HybridRunBookWorkerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HybridRunBookWorkerGroupList contains a list of HybridRunBookWorkerGroups
type HybridRunBookWorkerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HybridRunBookWorkerGroup `json:"items"`
}

// Repository type metadata.
var (
	HybridRunBookWorkerGroup_Kind             = "HybridRunBookWorkerGroup"
	HybridRunBookWorkerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HybridRunBookWorkerGroup_Kind}.String()
	HybridRunBookWorkerGroup_KindAPIVersion   = HybridRunBookWorkerGroup_Kind + "." + CRDGroupVersion.String()
	HybridRunBookWorkerGroup_GroupVersionKind = CRDGroupVersion.WithKind(HybridRunBookWorkerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&HybridRunBookWorkerGroup{}, &HybridRunBookWorkerGroupList{})
}
