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

type EmailTemplateInitParameters struct {

	// The body of the Email. Its format has to be a well-formed HTML document.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The subject of the Email.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type EmailTemplateObservation struct {

	// The name of the API Management Service in which the Email Template should exist. Changing this forces a new API Management Email Template to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The body of the Email. Its format has to be a well-formed HTML document.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The description of the Email Template.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the API Management Email Template.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the API Management Email Template should exist. Changing this forces a new API Management Email Template to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The subject of the Email.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The name of the Email Template. Possible values are AccountClosedDeveloper, ApplicationApprovedNotificationMessage, ConfirmSignUpIdentityDefault, EmailChangeIdentityDefault, InviteUserNotificationMessage, NewCommentNotificationMessage, NewDeveloperNotificationMessage, NewIssueNotificationMessage, PasswordResetByAdminNotificationMessage, PasswordResetIdentityDefault, PurchaseDeveloperNotificationMessage, QuotaLimitApproachingDeveloperNotificationMessage, RejectDeveloperNotificationMessage, RequestDeveloperNotificationMessage. Changing this forces a new API Management Email Template to be created.
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// The title of the Email Template.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type EmailTemplateParameters struct {

	// The name of the API Management Service in which the Email Template should exist. Changing this forces a new API Management Email Template to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The body of the Email. Its format has to be a well-formed HTML document.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The name of the Resource Group where the API Management Email Template should exist. Changing this forces a new API Management Email Template to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The subject of the Email.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The name of the Email Template. Possible values are AccountClosedDeveloper, ApplicationApprovedNotificationMessage, ConfirmSignUpIdentityDefault, EmailChangeIdentityDefault, InviteUserNotificationMessage, NewCommentNotificationMessage, NewDeveloperNotificationMessage, NewIssueNotificationMessage, PasswordResetByAdminNotificationMessage, PasswordResetIdentityDefault, PurchaseDeveloperNotificationMessage, QuotaLimitApproachingDeveloperNotificationMessage, RejectDeveloperNotificationMessage, RequestDeveloperNotificationMessage. Changing this forces a new API Management Email Template to be created.
	// +kubebuilder:validation:Required
	TemplateName *string `json:"templateName" tf:"template_name,omitempty"`
}

// EmailTemplateSpec defines the desired state of EmailTemplate
type EmailTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailTemplateParameters `json:"forProvider"`
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
	InitProvider EmailTemplateInitParameters `json:"initProvider,omitempty"`
}

// EmailTemplateStatus defines the observed state of EmailTemplate.
type EmailTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailTemplate is the Schema for the EmailTemplates API. Manages a API Management Email Template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EmailTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subject) || (has(self.initProvider) && has(self.initProvider.subject))",message="spec.forProvider.subject is a required parameter"
	Spec   EmailTemplateSpec   `json:"spec"`
	Status EmailTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailTemplateList contains a list of EmailTemplates
type EmailTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailTemplate `json:"items"`
}

// Repository type metadata.
var (
	EmailTemplate_Kind             = "EmailTemplate"
	EmailTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailTemplate_Kind}.String()
	EmailTemplate_KindAPIVersion   = EmailTemplate_Kind + "." + CRDGroupVersion.String()
	EmailTemplate_GroupVersionKind = CRDGroupVersion.WithKind(EmailTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailTemplate{}, &EmailTemplateList{})
}
