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

type HostPoolRegistrationInfoInitParameters struct {

	// A valid RFC3339Time for the expiration of the token..
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`
}

type HostPoolRegistrationInfoObservation struct {

	// A valid RFC3339Time for the expiration of the token..
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The ID of the Virtual Desktop Host Pool to link the Registration Info to. Changing this forces a new Registration Info resource to be created. Only a single virtual_desktop_host_pool_registration_info resource should be associated with a given hostpool. Assigning multiple resources will produce inconsistent results.
	HostpoolID *string `json:"hostpoolId,omitempty" tf:"hostpool_id,omitempty"`

	// The ID of the Virtual Desktop Host Pool Registration Info resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostPoolRegistrationInfoParameters struct {

	// A valid RFC3339Time for the expiration of the token..
	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The ID of the Virtual Desktop Host Pool to link the Registration Info to. Changing this forces a new Registration Info resource to be created. Only a single virtual_desktop_host_pool_registration_info resource should be associated with a given hostpool. Assigning multiple resources will produce inconsistent results.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/virtualdesktop/v1beta1.HostPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	HostpoolID *string `json:"hostpoolId,omitempty" tf:"hostpool_id,omitempty"`

	// Reference to a HostPool in virtualdesktop to populate hostpoolId.
	// +kubebuilder:validation:Optional
	HostpoolIDRef *v1.Reference `json:"hostpoolIdRef,omitempty" tf:"-"`

	// Selector for a HostPool in virtualdesktop to populate hostpoolId.
	// +kubebuilder:validation:Optional
	HostpoolIDSelector *v1.Selector `json:"hostpoolIdSelector,omitempty" tf:"-"`
}

// HostPoolRegistrationInfoSpec defines the desired state of HostPoolRegistrationInfo
type HostPoolRegistrationInfoSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostPoolRegistrationInfoParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HostPoolRegistrationInfoInitParameters `json:"initProvider,omitempty"`
}

// HostPoolRegistrationInfoStatus defines the observed state of HostPoolRegistrationInfo.
type HostPoolRegistrationInfoStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostPoolRegistrationInfoObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostPoolRegistrationInfo is the Schema for the HostPoolRegistrationInfos API. Manages a Virtual Desktop Host Pool Registration Info.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HostPoolRegistrationInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.expirationDate) || has(self.initProvider.expirationDate)",message="expirationDate is a required parameter"
	Spec   HostPoolRegistrationInfoSpec   `json:"spec"`
	Status HostPoolRegistrationInfoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostPoolRegistrationInfoList contains a list of HostPoolRegistrationInfos
type HostPoolRegistrationInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostPoolRegistrationInfo `json:"items"`
}

// Repository type metadata.
var (
	HostPoolRegistrationInfo_Kind             = "HostPoolRegistrationInfo"
	HostPoolRegistrationInfo_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostPoolRegistrationInfo_Kind}.String()
	HostPoolRegistrationInfo_KindAPIVersion   = HostPoolRegistrationInfo_Kind + "." + CRDGroupVersion.String()
	HostPoolRegistrationInfo_GroupVersionKind = CRDGroupVersion.WithKind(HostPoolRegistrationInfo_Kind)
)

func init() {
	SchemeBuilder.Register(&HostPoolRegistrationInfo{}, &HostPoolRegistrationInfoList{})
}