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

type ResourceGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResourceGroupParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupParameters `json:"forProvider"`
}

// ResourceGroupStatus defines the observed state of ResourceGroup.
type ResourceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroup is the Schema for the ResourceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupSpec   `json:"spec"`
	Status            ResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupList contains a list of ResourceGroups
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroup_Kind             = "ResourceGroup"
	ResourceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroup_Kind}.String()
	ResourceGroup_KindAPIVersion   = ResourceGroup_Kind + "." + CRDGroupVersion.String()
	ResourceGroup_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
