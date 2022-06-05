/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VirtualNetworkRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualNetworkRuleParameters struct {

	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// VirtualNetworkRuleSpec defines the desired state of VirtualNetworkRule
type VirtualNetworkRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkRuleParameters `json:"forProvider"`
}

// VirtualNetworkRuleStatus defines the observed state of VirtualNetworkRule.
type VirtualNetworkRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkRule is the Schema for the VirtualNetworkRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkRuleSpec   `json:"spec"`
	Status            VirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkRuleList contains a list of VirtualNetworkRules
type VirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkRule_Kind             = "VirtualNetworkRule"
	VirtualNetworkRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetworkRule_Kind}.String()
	VirtualNetworkRule_KindAPIVersion   = VirtualNetworkRule_Kind + "." + CRDGroupVersion.String()
	VirtualNetworkRule_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetworkRule_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkRule{}, &VirtualNetworkRuleList{})
}
