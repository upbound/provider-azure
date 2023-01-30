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

type BackendRequestDataMaskingHeadersObservation struct {
}

type BackendRequestDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendRequestDataMaskingObservation struct {
}

type BackendRequestDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []BackendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []BackendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type BackendRequestDataMaskingQueryParamsObservation struct {
}

type BackendRequestDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendResponseDataMaskingHeadersObservation struct {
}

type BackendResponseDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendResponseDataMaskingQueryParamsObservation struct {
}

type BackendResponseDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticBackendRequestObservation struct {
}

type DiagnosticBackendRequestParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []BackendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticBackendResponseDataMaskingObservation struct {
}

type DiagnosticBackendResponseDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []BackendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []BackendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticBackendResponseObservation struct {
}

type DiagnosticBackendResponseParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []DiagnosticBackendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingHeadersObservation struct {
}

type DiagnosticFrontendRequestDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingObservation struct {
}

type DiagnosticFrontendRequestDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []DiagnosticFrontendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []DiagnosticFrontendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingQueryParamsObservation struct {
}

type DiagnosticFrontendRequestDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestObservation struct {
}

type DiagnosticFrontendRequestParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []DiagnosticFrontendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingHeadersObservation struct {
}

type DiagnosticFrontendResponseDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingObservation struct {
}

type DiagnosticFrontendResponseDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []DiagnosticFrontendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []DiagnosticFrontendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingQueryParamsObservation struct {
}

type DiagnosticFrontendResponseDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseObservation struct {
}

type DiagnosticFrontendResponseParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking []DiagnosticFrontendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticObservation struct {

	// The ID of the API Management Diagnostic.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiagnosticParameters struct {

	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Logger
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementLoggerID *string `json:"apiManagementLoggerId,omitempty" tf:"api_management_logger_id,omitempty"`

	// Reference to a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDRef *v1.Reference `json:"apiManagementLoggerIdRef,omitempty" tf:"-"`

	// Selector for a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDSelector *v1.Selector `json:"apiManagementLoggerIdSelector,omitempty" tf:"-"`

	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	// +kubebuilder:validation:Optional
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors,omitempty"`

	// A backend_request block as defined below.
	// +kubebuilder:validation:Optional
	BackendRequest []DiagnosticBackendRequestParameters `json:"backendRequest,omitempty" tf:"backend_request,omitempty"`

	// A backend_response block as defined below.
	// +kubebuilder:validation:Optional
	BackendResponse []DiagnosticBackendResponseParameters `json:"backendResponse,omitempty" tf:"backend_response,omitempty"`

	// A frontend_request block as defined below.
	// +kubebuilder:validation:Optional
	FrontendRequest []DiagnosticFrontendRequestParameters `json:"frontendRequest,omitempty" tf:"frontend_request,omitempty"`

	// A frontend_response block as defined below.
	// +kubebuilder:validation:Optional
	FrontendResponse []DiagnosticFrontendResponseParameters `json:"frontendResponse,omitempty" tf:"frontend_response,omitempty"`

	// The HTTP Correlation Protocol to use. Possible values are None, Legacy or W3C.
	// +kubebuilder:validation:Optional
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol,omitempty"`

	// Log client IP address.
	// +kubebuilder:validation:Optional
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip,omitempty"`

	// The format of the Operation Name for Application Insights telemetries. Possible values are Name, and Url. Defaults to Name.
	// +kubebuilder:validation:Optional
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Sampling (%). For high traffic APIs, please read this documentation to understand performance implications and log sampling. Valid values are between 0.0 and 100.0.
	// +kubebuilder:validation:Optional
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// Logging verbosity. Possible values are verbose, information or error.
	// +kubebuilder:validation:Optional
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity,omitempty"`
}

// DiagnosticSpec defines the desired state of Diagnostic
type DiagnosticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiagnosticParameters `json:"forProvider"`
}

// DiagnosticStatus defines the observed state of Diagnostic.
type DiagnosticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiagnosticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Diagnostic is the Schema for the Diagnostics API. Manages an API Management Service Diagnostic.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Diagnostic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiagnosticSpec   `json:"spec"`
	Status            DiagnosticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiagnosticList contains a list of Diagnostics
type DiagnosticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Diagnostic `json:"items"`
}

// Repository type metadata.
var (
	Diagnostic_Kind             = "Diagnostic"
	Diagnostic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Diagnostic_Kind}.String()
	Diagnostic_KindAPIVersion   = Diagnostic_Kind + "." + CRDGroupVersion.String()
	Diagnostic_GroupVersionKind = CRDGroupVersion.WithKind(Diagnostic_Kind)
)

func init() {
	SchemeBuilder.Register(&Diagnostic{}, &DiagnosticList{})
}
