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

type Password1Observation struct {
}

type Password1Parameters struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type Password2Observation struct {
}

type Password2Parameters struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type TokenPasswordObservation struct {

	// The ID of the Container Registry Token Password.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TokenPasswordParameters struct {

	// The ID of the Container Registry Token that this Container Registry Token Password resides in. Changing this forces a new Container Registry Token Password to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerregistry/v1beta1.Token
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerRegistryTokenID *string `json:"containerRegistryTokenId,omitempty" tf:"container_registry_token_id,omitempty"`

	// Reference to a Token in containerregistry to populate containerRegistryTokenId.
	// +kubebuilder:validation:Optional
	ContainerRegistryTokenIDRef *v1.Reference `json:"containerRegistryTokenIdRef,omitempty" tf:"-"`

	// Selector for a Token in containerregistry to populate containerRegistryTokenId.
	// +kubebuilder:validation:Optional
	ContainerRegistryTokenIDSelector *v1.Selector `json:"containerRegistryTokenIdSelector,omitempty" tf:"-"`

	// One password block as defined below.
	// +kubebuilder:validation:Required
	Password1 []Password1Parameters `json:"password1" tf:"password1,omitempty"`

	// One password block as defined below.
	// +kubebuilder:validation:Optional
	Password2 []Password2Parameters `json:"password2,omitempty" tf:"password2,omitempty"`
}

// TokenPasswordSpec defines the desired state of TokenPassword
type TokenPasswordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenPasswordParameters `json:"forProvider"`
}

// TokenPasswordStatus defines the observed state of TokenPassword.
type TokenPasswordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenPasswordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TokenPassword is the Schema for the TokenPasswords API. Manages a Container Registry Token Password.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TokenPassword struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TokenPasswordSpec   `json:"spec"`
	Status            TokenPasswordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenPasswordList contains a list of TokenPasswords
type TokenPasswordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TokenPassword `json:"items"`
}

// Repository type metadata.
var (
	TokenPassword_Kind             = "TokenPassword"
	TokenPassword_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TokenPassword_Kind}.String()
	TokenPassword_KindAPIVersion   = TokenPassword_Kind + "." + CRDGroupVersion.String()
	TokenPassword_GroupVersionKind = CRDGroupVersion.WithKind(TokenPassword_Kind)
)

func init() {
	SchemeBuilder.Register(&TokenPassword{}, &TokenPasswordList{})
}
