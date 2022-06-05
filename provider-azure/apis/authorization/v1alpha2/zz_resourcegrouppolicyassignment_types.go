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

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NonComplianceMessageObservation struct {
}

type NonComplianceMessageParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type ResourceGroupPolicyAssignmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`
}

type ResourceGroupPolicyAssignmentParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	NonComplianceMessage []NonComplianceMessageParameters `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// +kubebuilder:validation:Optional
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	PolicyDefinitionID *string `json:"policyDefinitionId" tf:"policy_definition_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1alpha2.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupIDRef *v1.Reference `json:"resourceGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupIDSelector *v1.Selector `json:"resourceGroupIdSelector,omitempty" tf:"-"`
}

// ResourceGroupPolicyAssignmentSpec defines the desired state of ResourceGroupPolicyAssignment
type ResourceGroupPolicyAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupPolicyAssignmentParameters `json:"forProvider"`
}

// ResourceGroupPolicyAssignmentStatus defines the observed state of ResourceGroupPolicyAssignment.
type ResourceGroupPolicyAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignment is the Schema for the ResourceGroupPolicyAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroupPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupPolicyAssignmentSpec   `json:"spec"`
	Status            ResourceGroupPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignmentList contains a list of ResourceGroupPolicyAssignments
type ResourceGroupPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroupPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroupPolicyAssignment_Kind             = "ResourceGroupPolicyAssignment"
	ResourceGroupPolicyAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroupPolicyAssignment_Kind}.String()
	ResourceGroupPolicyAssignment_KindAPIVersion   = ResourceGroupPolicyAssignment_Kind + "." + CRDGroupVersion.String()
	ResourceGroupPolicyAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroupPolicyAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroupPolicyAssignment{}, &ResourceGroupPolicyAssignmentList{})
}
