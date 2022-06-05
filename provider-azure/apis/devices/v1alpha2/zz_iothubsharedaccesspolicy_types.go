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

type IOTHubSharedAccessPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubSharedAccessPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DeviceConnect *bool `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`

	// +crossplane:generate:reference:type=IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iotHubNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iotHubNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RegistryRead *bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryWrite *bool `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceConnect *bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

// IOTHubSharedAccessPolicySpec defines the desired state of IOTHubSharedAccessPolicy
type IOTHubSharedAccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubSharedAccessPolicyParameters `json:"forProvider"`
}

// IOTHubSharedAccessPolicyStatus defines the observed state of IOTHubSharedAccessPolicy.
type IOTHubSharedAccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubSharedAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubSharedAccessPolicy is the Schema for the IOTHubSharedAccessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubSharedAccessPolicySpec   `json:"spec"`
	Status            IOTHubSharedAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubSharedAccessPolicyList contains a list of IOTHubSharedAccessPolicys
type IOTHubSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubSharedAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	IOTHubSharedAccessPolicy_Kind             = "IOTHubSharedAccessPolicy"
	IOTHubSharedAccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubSharedAccessPolicy_Kind}.String()
	IOTHubSharedAccessPolicy_KindAPIVersion   = IOTHubSharedAccessPolicy_Kind + "." + CRDGroupVersion.String()
	IOTHubSharedAccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubSharedAccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubSharedAccessPolicy{}, &IOTHubSharedAccessPolicyList{})
}
