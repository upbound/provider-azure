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

type ResourceGuardObservation struct {

	// The ID of the Resource Guard.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResourceGuardParameters struct {

	// The Azure Region where the Resource Guard should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the Resource Group where the Resource Guard should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Resource Guard.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of the critical operations which are not protected by this Resource Guard.
	// +kubebuilder:validation:Optional
	VaultCriticalOperationExclusionList []*string `json:"vaultCriticalOperationExclusionList,omitempty" tf:"vault_critical_operation_exclusion_list,omitempty"`
}

// ResourceGuardSpec defines the desired state of ResourceGuard
type ResourceGuardSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGuardParameters `json:"forProvider"`
}

// ResourceGuardStatus defines the observed state of ResourceGuard.
type ResourceGuardStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGuardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGuard is the Schema for the ResourceGuards API. Manages a Resource Guard.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGuard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGuardSpec   `json:"spec"`
	Status            ResourceGuardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGuardList contains a list of ResourceGuards
type ResourceGuardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGuard `json:"items"`
}

// Repository type metadata.
var (
	ResourceGuard_Kind             = "ResourceGuard"
	ResourceGuard_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGuard_Kind}.String()
	ResourceGuard_KindAPIVersion   = ResourceGuard_Kind + "." + CRDGroupVersion.String()
	ResourceGuard_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGuard_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGuard{}, &ResourceGuardList{})
}
