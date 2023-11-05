// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type FlexibleServerDatabaseInitParameters struct {

	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs to be a valid PostgreSQL Charset. Defaults to UTF8. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs to be a valid PostgreSQL Collation. Defaults to en_US.utf8. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`
}

type FlexibleServerDatabaseObservation struct {

	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs to be a valid PostgreSQL Charset. Defaults to UTF8. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs to be a valid PostgreSQL Collation. Defaults to en_US.utf8. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The ID of the Azure PostgreSQL Flexible Server Database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Azure PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type FlexibleServerDatabaseParameters struct {

	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs to be a valid PostgreSQL Charset. Defaults to UTF8. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	// +kubebuilder:validation:Optional
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs to be a valid PostgreSQL Collation. Defaults to en_US.utf8. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The ID of the Azure PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	// +crossplane:generate:reference:type=FlexibleServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a FlexibleServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a FlexibleServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// FlexibleServerDatabaseSpec defines the desired state of FlexibleServerDatabase
type FlexibleServerDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerDatabaseParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FlexibleServerDatabaseInitParameters `json:"initProvider,omitempty"`
}

// FlexibleServerDatabaseStatus defines the observed state of FlexibleServerDatabase.
type FlexibleServerDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerDatabase is the Schema for the FlexibleServerDatabases API. Manages a PostgreSQL Flexible Server Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServerDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServerDatabaseSpec   `json:"spec"`
	Status            FlexibleServerDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerDatabaseList contains a list of FlexibleServerDatabases
type FlexibleServerDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerDatabase `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerDatabase_Kind             = "FlexibleServerDatabase"
	FlexibleServerDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServerDatabase_Kind}.String()
	FlexibleServerDatabase_KindAPIVersion   = FlexibleServerDatabase_Kind + "." + CRDGroupVersion.String()
	FlexibleServerDatabase_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServerDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerDatabase{}, &FlexibleServerDatabaseList{})
}
