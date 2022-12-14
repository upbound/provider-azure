/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ResourceGroupTemplateDeployment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ResourceGroupTemplateDeployment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ResourceGroupTemplateDeployment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ResourceGroupTemplateDeployment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
