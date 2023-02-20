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

type SpringCloudAPIPortalObservation struct {

	// The ID of the Spring Cloud API Portal.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// TODO.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SpringCloudAPIPortalParameters struct {

	// Specifies a list of Spring Cloud Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GatewayIds []*string `json:"gatewayIds,omitempty" tf:"gateway_ids,omitempty"`

	// References to SpringCloudGateway in appplatform to populate gatewayIds.
	// +kubebuilder:validation:Optional
	GatewayIdsRefs []v1.Reference `json:"gatewayIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SpringCloudGateway in appplatform to populate gatewayIds.
	// +kubebuilder:validation:Optional
	GatewayIdsSelector *v1.Selector `json:"gatewayIdsSelector,omitempty" tf:"-"`

	// is only https is allowed?
	// +kubebuilder:validation:Optional
	HTTPSOnlyEnabled *bool `json:"httpsOnlyEnabled,omitempty" tf:"https_only_enabled,omitempty"`

	// Specifies the required instance count of the Spring Cloud API Portal. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Is the public network access enabled?
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud API Portal to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`

	// A sso block as defined below.
	// +kubebuilder:validation:Optional
	Sso []SsoParameters `json:"sso,omitempty" tf:"sso,omitempty"`
}

type SsoObservation struct {
}

type SsoParameters struct {

	// The public identifier for the application.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The secret known only to the application and the authorization server.
	// +kubebuilder:validation:Optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// The URI of Issuer Identifier.
	// +kubebuilder:validation:Optional
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// It defines the specific actions applications can be allowed to do on a user's behalf.
	// +kubebuilder:validation:Optional
	Scope []*string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// SpringCloudAPIPortalSpec defines the desired state of SpringCloudAPIPortal
type SpringCloudAPIPortalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudAPIPortalParameters `json:"forProvider"`
}

// SpringCloudAPIPortalStatus defines the observed state of SpringCloudAPIPortal.
type SpringCloudAPIPortalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudAPIPortalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAPIPortal is the Schema for the SpringCloudAPIPortals API. Manages a Spring Cloud API Portal.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudAPIPortal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudAPIPortalSpec   `json:"spec"`
	Status            SpringCloudAPIPortalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAPIPortalList contains a list of SpringCloudAPIPortals
type SpringCloudAPIPortalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudAPIPortal `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudAPIPortal_Kind             = "SpringCloudAPIPortal"
	SpringCloudAPIPortal_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudAPIPortal_Kind}.String()
	SpringCloudAPIPortal_KindAPIVersion   = SpringCloudAPIPortal_Kind + "." + CRDGroupVersion.String()
	SpringCloudAPIPortal_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudAPIPortal_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudAPIPortal{}, &SpringCloudAPIPortalList{})
}
