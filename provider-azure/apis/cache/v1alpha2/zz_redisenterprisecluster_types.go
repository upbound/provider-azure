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

type RedisEnterpriseClusterObservation struct {
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RedisEnterpriseClusterParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// RedisEnterpriseClusterSpec defines the desired state of RedisEnterpriseCluster
type RedisEnterpriseClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisEnterpriseClusterParameters `json:"forProvider"`
}

// RedisEnterpriseClusterStatus defines the observed state of RedisEnterpriseCluster.
type RedisEnterpriseClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisEnterpriseClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseCluster is the Schema for the RedisEnterpriseClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisEnterpriseCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseClusterSpec   `json:"spec"`
	Status            RedisEnterpriseClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseClusterList contains a list of RedisEnterpriseClusters
type RedisEnterpriseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseCluster `json:"items"`
}

// Repository type metadata.
var (
	RedisEnterpriseCluster_Kind             = "RedisEnterpriseCluster"
	RedisEnterpriseCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisEnterpriseCluster_Kind}.String()
	RedisEnterpriseCluster_KindAPIVersion   = RedisEnterpriseCluster_Kind + "." + CRDGroupVersion.String()
	RedisEnterpriseCluster_GroupVersionKind = CRDGroupVersion.WithKind(RedisEnterpriseCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisEnterpriseCluster{}, &RedisEnterpriseClusterList{})
}
