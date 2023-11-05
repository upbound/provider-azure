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

type ReferenceInputBlobInitParameters struct {

	// The authentication mode for the Stream Analytics Reference Input. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The date format. Wherever {date} appears in path_pattern, the value of this property is used as the date format instead.
	DateFormat *string `json:"dateFormat,omitempty" tf:"date_format,omitempty"`

	// The name of the Reference Input Blob. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern *string `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// A serialization block as defined below.
	Serialization []ReferenceInputBlobSerializationInitParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The time format. Wherever {time} appears in path_pattern, the value of this property is used as the time format instead.
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`
}

type ReferenceInputBlobObservation struct {

	// The authentication mode for the Stream Analytics Reference Input. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The date format. Wherever {date} appears in path_pattern, the value of this property is used as the date format instead.
	DateFormat *string `json:"dateFormat,omitempty" tf:"date_format,omitempty"`

	// The ID of the Stream Analytics Reference Input Blob.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Reference Input Blob. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern *string `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A serialization block as defined below.
	Serialization []ReferenceInputBlobSerializationObservation `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The name of the Storage Account that has the blob container with reference data.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// The name of the Container within the Storage Account.
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// The time format. Wherever {time} appears in path_pattern, the value of this property is used as the time format instead.
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`
}

type ReferenceInputBlobParameters struct {

	// The authentication mode for the Stream Analytics Reference Input. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	// +kubebuilder:validation:Optional
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The date format. Wherever {date} appears in path_pattern, the value of this property is used as the date format instead.
	// +kubebuilder:validation:Optional
	DateFormat *string `json:"dateFormat,omitempty" tf:"date_format,omitempty"`

	// The name of the Reference Input Blob. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	// +kubebuilder:validation:Optional
	PathPattern *string `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	// +kubebuilder:validation:Optional
	Serialization []ReferenceInputBlobSerializationParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The Access Key which should be used to connect to this Storage Account. Required if authentication_mode is ConnectionString.
	// +kubebuilder:validation:Optional
	StorageAccountKeySecretRef *v1.SecretKeySelector `json:"storageAccountKeySecretRef,omitempty" tf:"-"`

	// The name of the Storage Account that has the blob container with reference data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// The name of the Container within the Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	StorageContainerName *string `json:"storageContainerName,omitempty" tf:"storage_container_name,omitempty"`

	// Reference to a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameRef *v1.Reference `json:"storageContainerNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate storageContainerName.
	// +kubebuilder:validation:Optional
	StorageContainerNameSelector *v1.Selector `json:"storageContainerNameSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// The time format. Wherever {time} appears in path_pattern, the value of this property is used as the time format instead.
	// +kubebuilder:validation:Optional
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`
}

type ReferenceInputBlobSerializationInitParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for the reference data. Possible values are Avro, Csv and Json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ReferenceInputBlobSerializationObservation struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for the reference data. Possible values are Avro, Csv and Json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ReferenceInputBlobSerializationParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for the reference data. Possible values are Avro, Csv and Json.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// ReferenceInputBlobSpec defines the desired state of ReferenceInputBlob
type ReferenceInputBlobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReferenceInputBlobParameters `json:"forProvider"`
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
	InitProvider ReferenceInputBlobInitParameters `json:"initProvider,omitempty"`
}

// ReferenceInputBlobStatus defines the observed state of ReferenceInputBlob.
type ReferenceInputBlobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReferenceInputBlobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceInputBlob is the Schema for the ReferenceInputBlobs API. Manages a Stream Analytics Reference Input Blob.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ReferenceInputBlob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dateFormat) || (has(self.initProvider) && has(self.initProvider.dateFormat))",message="spec.forProvider.dateFormat is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pathPattern) || (has(self.initProvider) && has(self.initProvider.pathPattern))",message="spec.forProvider.pathPattern is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serialization) || (has(self.initProvider) && has(self.initProvider.serialization))",message="spec.forProvider.serialization is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeFormat) || (has(self.initProvider) && has(self.initProvider.timeFormat))",message="spec.forProvider.timeFormat is a required parameter"
	Spec   ReferenceInputBlobSpec   `json:"spec"`
	Status ReferenceInputBlobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceInputBlobList contains a list of ReferenceInputBlobs
type ReferenceInputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceInputBlob `json:"items"`
}

// Repository type metadata.
var (
	ReferenceInputBlob_Kind             = "ReferenceInputBlob"
	ReferenceInputBlob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReferenceInputBlob_Kind}.String()
	ReferenceInputBlob_KindAPIVersion   = ReferenceInputBlob_Kind + "." + CRDGroupVersion.String()
	ReferenceInputBlob_GroupVersionKind = CRDGroupVersion.WithKind(ReferenceInputBlob_Kind)
)

func init() {
	SchemeBuilder.Register(&ReferenceInputBlob{}, &ReferenceInputBlobList{})
}
