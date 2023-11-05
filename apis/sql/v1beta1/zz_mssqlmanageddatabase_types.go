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

type MSSQLManagedDatabaseInitParameters struct {

	// A long_term_retention_policy block as defined below.
	LongTermRetentionPolicy []MSSQLManagedDatabaseLongTermRetentionPolicyInitParameters `json:"longTermRetentionPolicy,omitempty" tf:"long_term_retention_policy,omitempty"`

	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	ShortTermRetentionDays *float64 `json:"shortTermRetentionDays,omitempty" tf:"short_term_retention_days,omitempty"`
}

type MSSQLManagedDatabaseLongTermRetentionPolicyInitParameters struct {

	// The monthly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 120 months. e.g. P1Y, P1M, P4W or P30D.
	MonthlyRetention *string `json:"monthlyRetention,omitempty" tf:"monthly_retention,omitempty"`

	// The week of year to take the yearly backup. Value has to be between 1 and 52.
	WeekOfYear *float64 `json:"weekOfYear,omitempty" tf:"week_of_year,omitempty"`

	// The weekly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 520 weeks. e.g. P1Y, P1M, P1W or P7D.
	WeeklyRetention *string `json:"weeklyRetention,omitempty" tf:"weekly_retention,omitempty"`

	// The yearly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 10 years. e.g. P1Y, P12M, P52W or P365D.
	YearlyRetention *string `json:"yearlyRetention,omitempty" tf:"yearly_retention,omitempty"`
}

type MSSQLManagedDatabaseLongTermRetentionPolicyObservation struct {

	// The monthly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 120 months. e.g. P1Y, P1M, P4W or P30D.
	MonthlyRetention *string `json:"monthlyRetention,omitempty" tf:"monthly_retention,omitempty"`

	// The week of year to take the yearly backup. Value has to be between 1 and 52.
	WeekOfYear *float64 `json:"weekOfYear,omitempty" tf:"week_of_year,omitempty"`

	// The weekly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 520 weeks. e.g. P1Y, P1M, P1W or P7D.
	WeeklyRetention *string `json:"weeklyRetention,omitempty" tf:"weekly_retention,omitempty"`

	// The yearly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 10 years. e.g. P1Y, P12M, P52W or P365D.
	YearlyRetention *string `json:"yearlyRetention,omitempty" tf:"yearly_retention,omitempty"`
}

type MSSQLManagedDatabaseLongTermRetentionPolicyParameters struct {

	// The monthly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 120 months. e.g. P1Y, P1M, P4W or P30D.
	// +kubebuilder:validation:Optional
	MonthlyRetention *string `json:"monthlyRetention,omitempty" tf:"monthly_retention,omitempty"`

	// The week of year to take the yearly backup. Value has to be between 1 and 52.
	// +kubebuilder:validation:Optional
	WeekOfYear *float64 `json:"weekOfYear,omitempty" tf:"week_of_year,omitempty"`

	// The weekly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 520 weeks. e.g. P1Y, P1M, P1W or P7D.
	// +kubebuilder:validation:Optional
	WeeklyRetention *string `json:"weeklyRetention,omitempty" tf:"weekly_retention,omitempty"`

	// The yearly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 10 years. e.g. P1Y, P12M, P52W or P365D.
	// +kubebuilder:validation:Optional
	YearlyRetention *string `json:"yearlyRetention,omitempty" tf:"yearly_retention,omitempty"`
}

type MSSQLManagedDatabaseObservation struct {

	// The Azure SQL Managed Database ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A long_term_retention_policy block as defined below.
	LongTermRetentionPolicy []MSSQLManagedDatabaseLongTermRetentionPolicyObservation `json:"longTermRetentionPolicy,omitempty" tf:"long_term_retention_policy,omitempty"`

	// The ID of the Azure SQL Managed Instance on which to create this Managed Database. Changing this forces a new resource to be created.
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	ShortTermRetentionDays *float64 `json:"shortTermRetentionDays,omitempty" tf:"short_term_retention_days,omitempty"`
}

type MSSQLManagedDatabaseParameters struct {

	// A long_term_retention_policy block as defined below.
	// +kubebuilder:validation:Optional
	LongTermRetentionPolicy []MSSQLManagedDatabaseLongTermRetentionPolicyParameters `json:"longTermRetentionPolicy,omitempty" tf:"long_term_retention_policy,omitempty"`

	// The ID of the Azure SQL Managed Instance on which to create this Managed Database. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDRef *v1.Reference `json:"managedInstanceIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDSelector *v1.Selector `json:"managedInstanceIdSelector,omitempty" tf:"-"`

	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	// +kubebuilder:validation:Optional
	ShortTermRetentionDays *float64 `json:"shortTermRetentionDays,omitempty" tf:"short_term_retention_days,omitempty"`
}

// MSSQLManagedDatabaseSpec defines the desired state of MSSQLManagedDatabase
type MSSQLManagedDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLManagedDatabaseParameters `json:"forProvider"`
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
	InitProvider MSSQLManagedDatabaseInitParameters `json:"initProvider,omitempty"`
}

// MSSQLManagedDatabaseStatus defines the observed state of MSSQLManagedDatabase.
type MSSQLManagedDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLManagedDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedDatabase is the Schema for the MSSQLManagedDatabases API. Manages an Azure SQL Azure Managed Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLManagedDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLManagedDatabaseSpec   `json:"spec"`
	Status            MSSQLManagedDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedDatabaseList contains a list of MSSQLManagedDatabases
type MSSQLManagedDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLManagedDatabase `json:"items"`
}

// Repository type metadata.
var (
	MSSQLManagedDatabase_Kind             = "MSSQLManagedDatabase"
	MSSQLManagedDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLManagedDatabase_Kind}.String()
	MSSQLManagedDatabase_KindAPIVersion   = MSSQLManagedDatabase_Kind + "." + CRDGroupVersion.String()
	MSSQLManagedDatabase_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLManagedDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLManagedDatabase{}, &MSSQLManagedDatabaseList{})
}
