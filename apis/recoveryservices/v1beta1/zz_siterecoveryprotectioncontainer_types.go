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

type SiteRecoveryProtectionContainerObservation struct {

	// The ID of the Site Recovery Protection Container.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SiteRecoveryProtectionContainerParameters struct {

	// Name of fabric that should contain this protection container. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.SiteRecoveryFabric
	// +kubebuilder:validation:Optional
	RecoveryFabricName *string `json:"recoveryFabricName,omitempty" tf:"recovery_fabric_name,omitempty"`

	// Reference to a SiteRecoveryFabric in recoveryservices to populate recoveryFabricName.
	// +kubebuilder:validation:Optional
	RecoveryFabricNameRef *v1.Reference `json:"recoveryFabricNameRef,omitempty" tf:"-"`

	// Selector for a SiteRecoveryFabric in recoveryservices to populate recoveryFabricName.
	// +kubebuilder:validation:Optional
	RecoveryFabricNameSelector *v1.Selector `json:"recoveryFabricNameSelector,omitempty" tf:"-"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
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

// SiteRecoveryProtectionContainerSpec defines the desired state of SiteRecoveryProtectionContainer
type SiteRecoveryProtectionContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryProtectionContainerParameters `json:"forProvider"`
}

// SiteRecoveryProtectionContainerStatus defines the observed state of SiteRecoveryProtectionContainer.
type SiteRecoveryProtectionContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryProtectionContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainer is the Schema for the SiteRecoveryProtectionContainers API. Manages a site recovery services protection container on Azure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SiteRecoveryProtectionContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteRecoveryProtectionContainerSpec   `json:"spec"`
	Status            SiteRecoveryProtectionContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainerList contains a list of SiteRecoveryProtectionContainers
type SiteRecoveryProtectionContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryProtectionContainer `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryProtectionContainer_Kind             = "SiteRecoveryProtectionContainer"
	SiteRecoveryProtectionContainer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryProtectionContainer_Kind}.String()
	SiteRecoveryProtectionContainer_KindAPIVersion   = SiteRecoveryProtectionContainer_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryProtectionContainer_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryProtectionContainer_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryProtectionContainer{}, &SiteRecoveryProtectionContainerList{})
}
