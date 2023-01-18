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

type ApplicationSecurityGroupObservation struct {

	// The ID of the Application Security Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ApplicationSecurityGroupParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the Application Security Group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ApplicationSecurityGroupSpec defines the desired state of ApplicationSecurityGroup
type ApplicationSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationSecurityGroupParameters `json:"forProvider"`
}

// ApplicationSecurityGroupStatus defines the observed state of ApplicationSecurityGroup.
type ApplicationSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationSecurityGroup is the Schema for the ApplicationSecurityGroups API. Manages an Application Security Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSecurityGroupSpec   `json:"spec"`
	Status            ApplicationSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationSecurityGroupList contains a list of ApplicationSecurityGroups
type ApplicationSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	ApplicationSecurityGroup_Kind             = "ApplicationSecurityGroup"
	ApplicationSecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationSecurityGroup_Kind}.String()
	ApplicationSecurityGroup_KindAPIVersion   = ApplicationSecurityGroup_Kind + "." + CRDGroupVersion.String()
	ApplicationSecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationSecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationSecurityGroup{}, &ApplicationSecurityGroupList{})
}
