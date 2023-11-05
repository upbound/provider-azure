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

type MarketplaceAgreementInitParameters struct {

	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type MarketplaceAgreementObservation struct {

	// The ID of the Marketplace Agreement.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LicenseTextLink *string `json:"licenseTextLink,omitempty" tf:"license_text_link,omitempty"`

	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	PrivacyPolicyLink *string `json:"privacyPolicyLink,omitempty" tf:"privacy_policy_link,omitempty"`

	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type MarketplaceAgreementParameters struct {

	// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

// MarketplaceAgreementSpec defines the desired state of MarketplaceAgreement
type MarketplaceAgreementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MarketplaceAgreementParameters `json:"forProvider"`
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
	InitProvider MarketplaceAgreementInitParameters `json:"initProvider,omitempty"`
}

// MarketplaceAgreementStatus defines the observed state of MarketplaceAgreement.
type MarketplaceAgreementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MarketplaceAgreementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MarketplaceAgreement is the Schema for the MarketplaceAgreements API. Allows accepting the Legal Terms for a Marketplace Image.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MarketplaceAgreement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.offer) || (has(self.initProvider) && has(self.initProvider.offer))",message="spec.forProvider.offer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plan) || (has(self.initProvider) && has(self.initProvider.plan))",message="spec.forProvider.plan is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publisher) || (has(self.initProvider) && has(self.initProvider.publisher))",message="spec.forProvider.publisher is a required parameter"
	Spec   MarketplaceAgreementSpec   `json:"spec"`
	Status MarketplaceAgreementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MarketplaceAgreementList contains a list of MarketplaceAgreements
type MarketplaceAgreementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MarketplaceAgreement `json:"items"`
}

// Repository type metadata.
var (
	MarketplaceAgreement_Kind             = "MarketplaceAgreement"
	MarketplaceAgreement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MarketplaceAgreement_Kind}.String()
	MarketplaceAgreement_KindAPIVersion   = MarketplaceAgreement_Kind + "." + CRDGroupVersion.String()
	MarketplaceAgreement_GroupVersionKind = CRDGroupVersion.WithKind(MarketplaceAgreement_Kind)
)

func init() {
	SchemeBuilder.Register(&MarketplaceAgreement{}, &MarketplaceAgreementList{})
}
