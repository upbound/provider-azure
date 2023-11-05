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

type SpringCloudJavaDeploymentInitParameters struct {

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Specifies the jvm option of the Spring Cloud Deployment.
	JvmOptions *string `json:"jvmOptions,omitempty" tf:"jvm_options,omitempty"`

	// A quota block as defined below.
	Quota []SpringCloudJavaDeploymentQuotaInitParameters `json:"quota,omitempty" tf:"quota,omitempty"`

	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are Java_8, Java_11 and Java_17. Defaults to Java_8.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

type SpringCloudJavaDeploymentObservation struct {

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// The ID of the Spring Cloud Deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Specifies the jvm option of the Spring Cloud Deployment.
	JvmOptions *string `json:"jvmOptions,omitempty" tf:"jvm_options,omitempty"`

	// A quota block as defined below.
	Quota []SpringCloudJavaDeploymentQuotaObservation `json:"quota,omitempty" tf:"quota,omitempty"`

	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are Java_8, Java_11 and Java_17. Defaults to Java_8.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`

	// Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`
}

type SpringCloudJavaDeploymentParameters struct {

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	// +kubebuilder:validation:Optional
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Specifies the jvm option of the Spring Cloud Deployment.
	// +kubebuilder:validation:Optional
	JvmOptions *string `json:"jvmOptions,omitempty" tf:"jvm_options,omitempty"`

	// A quota block as defined below.
	// +kubebuilder:validation:Optional
	Quota []SpringCloudJavaDeploymentQuotaParameters `json:"quota,omitempty" tf:"quota,omitempty"`

	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are Java_8, Java_11 and Java_17. Defaults to Java_8.
	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`

	// Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudApp
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Reference to a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDRef *v1.Reference `json:"springCloudAppIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDSelector *v1.Selector `json:"springCloudAppIdSelector,omitempty" tf:"-"`
}

type SpringCloudJavaDeploymentQuotaInitParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudJavaDeploymentQuotaObservation struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudJavaDeploymentQuotaParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

// SpringCloudJavaDeploymentSpec defines the desired state of SpringCloudJavaDeployment
type SpringCloudJavaDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudJavaDeploymentParameters `json:"forProvider"`
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
	InitProvider SpringCloudJavaDeploymentInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudJavaDeploymentStatus defines the observed state of SpringCloudJavaDeployment.
type SpringCloudJavaDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudJavaDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudJavaDeployment is the Schema for the SpringCloudJavaDeployments API. Manages an Azure Spring Cloud Deployment with a Java runtime.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudJavaDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudJavaDeploymentSpec   `json:"spec"`
	Status            SpringCloudJavaDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudJavaDeploymentList contains a list of SpringCloudJavaDeployments
type SpringCloudJavaDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudJavaDeployment `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudJavaDeployment_Kind             = "SpringCloudJavaDeployment"
	SpringCloudJavaDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudJavaDeployment_Kind}.String()
	SpringCloudJavaDeployment_KindAPIVersion   = SpringCloudJavaDeployment_Kind + "." + CRDGroupVersion.String()
	SpringCloudJavaDeployment_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudJavaDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudJavaDeployment{}, &SpringCloudJavaDeploymentList{})
}
