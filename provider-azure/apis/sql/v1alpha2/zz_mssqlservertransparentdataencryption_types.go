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

type MSSQLServerTransparentDataEncryptionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MSSQLServerTransparentDataEncryptionParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/keyvault/v1alpha2.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// MSSQLServerTransparentDataEncryptionSpec defines the desired state of MSSQLServerTransparentDataEncryption
type MSSQLServerTransparentDataEncryptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLServerTransparentDataEncryptionParameters `json:"forProvider"`
}

// MSSQLServerTransparentDataEncryptionStatus defines the observed state of MSSQLServerTransparentDataEncryption.
type MSSQLServerTransparentDataEncryptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLServerTransparentDataEncryptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerTransparentDataEncryption is the Schema for the MSSQLServerTransparentDataEncryptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLServerTransparentDataEncryption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLServerTransparentDataEncryptionSpec   `json:"spec"`
	Status            MSSQLServerTransparentDataEncryptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerTransparentDataEncryptionList contains a list of MSSQLServerTransparentDataEncryptions
type MSSQLServerTransparentDataEncryptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServerTransparentDataEncryption `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServerTransparentDataEncryption_Kind             = "MSSQLServerTransparentDataEncryption"
	MSSQLServerTransparentDataEncryption_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServerTransparentDataEncryption_Kind}.String()
	MSSQLServerTransparentDataEncryption_KindAPIVersion   = MSSQLServerTransparentDataEncryption_Kind + "." + CRDGroupVersion.String()
	MSSQLServerTransparentDataEncryption_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServerTransparentDataEncryption_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServerTransparentDataEncryption{}, &MSSQLServerTransparentDataEncryptionList{})
}
