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

type AppActionHTTPInitParameters struct {

	// Specifies the HTTP Body that should be sent to the uri when this HTTP Action is triggered.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Specifies a Map of Key-Value Pairs that should be sent to the uri when this HTTP Action is triggered.
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include DELETE, GET, PATCH, POST and PUT.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Specifies a Map of Key-Value Pairs that should be sent to the uri when this HTTP Action is triggered.
	Queries map[string]*string `json:"queries,omitempty" tf:"queries,omitempty"`

	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A run_after block is as defined below.
	RunAfter []RunAfterInitParameters `json:"runAfter,omitempty" tf:"run_after,omitempty"`

	// Specifies the URI which will be called when this HTTP Action is triggered.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type AppActionHTTPObservation struct {

	// Specifies the HTTP Body that should be sent to the uri when this HTTP Action is triggered.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Specifies a Map of Key-Value Pairs that should be sent to the uri when this HTTP Action is triggered.
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// The ID of the HTTP Action within the Logic App Workflow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`

	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include DELETE, GET, PATCH, POST and PUT.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Specifies a Map of Key-Value Pairs that should be sent to the uri when this HTTP Action is triggered.
	Queries map[string]*string `json:"queries,omitempty" tf:"queries,omitempty"`

	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A run_after block is as defined below.
	RunAfter []RunAfterObservation `json:"runAfter,omitempty" tf:"run_after,omitempty"`

	// Specifies the URI which will be called when this HTTP Action is triggered.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type AppActionHTTPParameters struct {

	// Specifies the HTTP Body that should be sent to the uri when this HTTP Action is triggered.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Specifies a Map of Key-Value Pairs that should be sent to the uri when this HTTP Action is triggered.
	// +kubebuilder:validation:Optional
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logic/v1beta1.AppWorkflow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`

	// Reference to a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDRef *v1.Reference `json:"logicAppIdRef,omitempty" tf:"-"`

	// Selector for a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDSelector *v1.Selector `json:"logicAppIdSelector,omitempty" tf:"-"`

	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include DELETE, GET, PATCH, POST and PUT.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Specifies a Map of Key-Value Pairs that should be sent to the uri when this HTTP Action is triggered.
	// +kubebuilder:validation:Optional
	Queries map[string]*string `json:"queries,omitempty" tf:"queries,omitempty"`

	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A run_after block is as defined below.
	// +kubebuilder:validation:Optional
	RunAfter []RunAfterParameters `json:"runAfter,omitempty" tf:"run_after,omitempty"`

	// Specifies the URI which will be called when this HTTP Action is triggered.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type RunAfterInitParameters struct {

	// Specifies the name of the precedent HTTP Action.
	ActionName *string `json:"actionName,omitempty" tf:"action_name,omitempty"`

	// Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include Succeeded, Failed, Skipped and TimedOut.
	ActionResult *string `json:"actionResult,omitempty" tf:"action_result,omitempty"`
}

type RunAfterObservation struct {

	// Specifies the name of the precedent HTTP Action.
	ActionName *string `json:"actionName,omitempty" tf:"action_name,omitempty"`

	// Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include Succeeded, Failed, Skipped and TimedOut.
	ActionResult *string `json:"actionResult,omitempty" tf:"action_result,omitempty"`
}

type RunAfterParameters struct {

	// Specifies the name of the precedent HTTP Action.
	// +kubebuilder:validation:Optional
	ActionName *string `json:"actionName" tf:"action_name,omitempty"`

	// Specifies the expected result of the precedent HTTP Action, only after which the current HTTP Action will be triggered. Possible values include Succeeded, Failed, Skipped and TimedOut.
	// +kubebuilder:validation:Optional
	ActionResult *string `json:"actionResult" tf:"action_result,omitempty"`
}

// AppActionHTTPSpec defines the desired state of AppActionHTTP
type AppActionHTTPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppActionHTTPParameters `json:"forProvider"`
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
	InitProvider AppActionHTTPInitParameters `json:"initProvider,omitempty"`
}

// AppActionHTTPStatus defines the observed state of AppActionHTTP.
type AppActionHTTPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppActionHTTPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppActionHTTP is the Schema for the AppActionHTTPs API. Manages an HTTP Action within a Logic App Workflow
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppActionHTTP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.method) || (has(self.initProvider) && has(self.initProvider.method))",message="spec.forProvider.method is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.uri) || (has(self.initProvider) && has(self.initProvider.uri))",message="spec.forProvider.uri is a required parameter"
	Spec   AppActionHTTPSpec   `json:"spec"`
	Status AppActionHTTPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppActionHTTPList contains a list of AppActionHTTPs
type AppActionHTTPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppActionHTTP `json:"items"`
}

// Repository type metadata.
var (
	AppActionHTTP_Kind             = "AppActionHTTP"
	AppActionHTTP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppActionHTTP_Kind}.String()
	AppActionHTTP_KindAPIVersion   = AppActionHTTP_Kind + "." + CRDGroupVersion.String()
	AppActionHTTP_GroupVersionKind = CRDGroupVersion.WithKind(AppActionHTTP_Kind)
)

func init() {
	SchemeBuilder.Register(&AppActionHTTP{}, &AppActionHTTPList{})
}
