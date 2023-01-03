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

type MongoCollectionAutoscaleSettingsObservation struct {
}

type MongoCollectionAutoscaleSettingsParameters struct {

	// The maximum throughput of the MongoDB collection (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type MongoCollectionIndexObservation struct {
}

type MongoCollectionIndexParameters struct {

	// Specifies the list of user settable keys for each Cosmos DB Mongo Collection.
	// +kubebuilder:validation:Required
	Keys []*string `json:"keys" tf:"keys,omitempty"`

	// Is the index unique or not? Defaults to false.
	// +kubebuilder:validation:Optional
	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`
}

type MongoCollectionObservation struct {

	// The ID of the Cosmos DB Mongo Collection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more system_indexes blocks as defined below.
	SystemIndexes []SystemIndexesObservation `json:"systemIndexes,omitempty" tf:"system_indexes,omitempty"`
}

type MongoCollectionParameters struct {

	// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to -1, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number n – items will expire n seconds after their last modified time.
	// +kubebuilder:validation:Optional
	AnalyticalStorageTTL *float64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// An autoscale_settings block as defined below.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []MongoCollectionAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MongoDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a MongoDatabase to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a MongoDatabase to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The default Time To Live in seconds. If the value is -1, items are not automatically expired.
	// +kubebuilder:validation:Optional
	DefaultTTLSeconds *float64 `json:"defaultTtlSeconds,omitempty" tf:"default_ttl_seconds,omitempty"`

	// One or more index blocks as defined below.
	// +kubebuilder:validation:Optional
	Index []MongoCollectionIndexParameters `json:"index,omitempty" tf:"index,omitempty"`

	// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the key to partition on for sharding. There must not be any other unique index keys. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ShardKey *string `json:"shardKey,omitempty" tf:"shard_key,omitempty"`

	// The throughput of the MongoDB collection (RU/s). Must be set in increments of 100. The minimum value is 400.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type SystemIndexesObservation struct {

	// The list of system keys which are not settable for each Cosmos DB Mongo Collection.
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`

	// Identifies whether the table contains no duplicate values.
	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`
}

type SystemIndexesParameters struct {
}

// MongoCollectionSpec defines the desired state of MongoCollection
type MongoCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MongoCollectionParameters `json:"forProvider"`
}

// MongoCollectionStatus defines the observed state of MongoCollection.
type MongoCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MongoCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MongoCollection is the Schema for the MongoCollections API. Manages a Mongo Collection within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MongoCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongoCollectionSpec   `json:"spec"`
	Status            MongoCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoCollectionList contains a list of MongoCollections
type MongoCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoCollection `json:"items"`
}

// Repository type metadata.
var (
	MongoCollection_Kind             = "MongoCollection"
	MongoCollection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MongoCollection_Kind}.String()
	MongoCollection_KindAPIVersion   = MongoCollection_Kind + "." + CRDGroupVersion.String()
	MongoCollection_GroupVersionKind = CRDGroupVersion.WithKind(MongoCollection_Kind)
)

func init() {
	SchemeBuilder.Register(&MongoCollection{}, &MongoCollectionList{})
}
