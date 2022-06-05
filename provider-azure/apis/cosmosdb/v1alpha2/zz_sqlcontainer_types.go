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

type CompositeIndexIndexObservation struct {
}

type CompositeIndexIndexParameters struct {

	// +kubebuilder:validation:Required
	Order *string `json:"order" tf:"order,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type ExcludedPathObservation struct {
}

type ExcludedPathParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type IncludedPathObservation struct {
}

type IncludedPathParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type IndexingPolicyCompositeIndexObservation struct {
}

type IndexingPolicyCompositeIndexParameters struct {

	// +kubebuilder:validation:Required
	Index []CompositeIndexIndexParameters `json:"index" tf:"index,omitempty"`
}

type IndexingPolicyObservation struct {
	SpatialIndex []IndexingPolicySpatialIndexObservation `json:"spatialIndex,omitempty" tf:"spatial_index,omitempty"`
}

type IndexingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	CompositeIndex []IndexingPolicyCompositeIndexParameters `json:"compositeIndex,omitempty" tf:"composite_index,omitempty"`

	// +kubebuilder:validation:Optional
	ExcludedPath []ExcludedPathParameters `json:"excludedPath,omitempty" tf:"excluded_path,omitempty"`

	// +kubebuilder:validation:Optional
	IncludedPath []IncludedPathParameters `json:"includedPath,omitempty" tf:"included_path,omitempty"`

	// +kubebuilder:validation:Optional
	IndexingMode *string `json:"indexingMode,omitempty" tf:"indexing_mode,omitempty"`

	// +kubebuilder:validation:Optional
	SpatialIndex []IndexingPolicySpatialIndexParameters `json:"spatialIndex,omitempty" tf:"spatial_index,omitempty"`
}

type IndexingPolicySpatialIndexObservation struct {
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

type IndexingPolicySpatialIndexParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type SQLContainerAutoscaleSettingsObservation struct {
}

type SQLContainerAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type SQLContainerConflictResolutionPolicyObservation struct {
}

type SQLContainerConflictResolutionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty" tf:"conflict_resolution_path,omitempty"`

	// +kubebuilder:validation:Optional
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty" tf:"conflict_resolution_procedure,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type SQLContainerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IndexingPolicy []IndexingPolicyObservation `json:"indexingPolicy,omitempty" tf:"indexing_policy,omitempty"`
}

type SQLContainerParameters struct {

	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AnalyticalStorageTTL *float64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []SQLContainerAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// +kubebuilder:validation:Optional
	ConflictResolutionPolicy []SQLContainerConflictResolutionPolicyParameters `json:"conflictResolutionPolicy,omitempty" tf:"conflict_resolution_policy,omitempty"`

	// +crossplane:generate:reference:type=SQLDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	IndexingPolicy []IndexingPolicyParameters `json:"indexingPolicy,omitempty" tf:"indexing_policy,omitempty"`

	// +kubebuilder:validation:Required
	PartitionKeyPath *string `json:"partitionKeyPath" tf:"partition_key_path,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionKeyVersion *float64 `json:"partitionKeyVersion,omitempty" tf:"partition_key_version,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	UniqueKey []SQLContainerUniqueKeyParameters `json:"uniqueKey,omitempty" tf:"unique_key,omitempty"`
}

type SQLContainerUniqueKeyObservation struct {
}

type SQLContainerUniqueKeyParameters struct {

	// +kubebuilder:validation:Required
	Paths []*string `json:"paths" tf:"paths,omitempty"`
}

// SQLContainerSpec defines the desired state of SQLContainer
type SQLContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLContainerParameters `json:"forProvider"`
}

// SQLContainerStatus defines the observed state of SQLContainer.
type SQLContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLContainer is the Schema for the SQLContainers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLContainerSpec   `json:"spec"`
	Status            SQLContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLContainerList contains a list of SQLContainers
type SQLContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLContainer `json:"items"`
}

// Repository type metadata.
var (
	SQLContainer_Kind             = "SQLContainer"
	SQLContainer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLContainer_Kind}.String()
	SQLContainer_KindAPIVersion   = SQLContainer_Kind + "." + CRDGroupVersion.String()
	SQLContainer_GroupVersionKind = CRDGroupVersion.WithKind(SQLContainer_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLContainer{}, &SQLContainerList{})
}
