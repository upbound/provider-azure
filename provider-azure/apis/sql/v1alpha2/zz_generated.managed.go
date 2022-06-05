/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this MSSQLServer.
func (mg *MSSQLServer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLServer.
func (mg *MSSQLServer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLServer.
func (mg *MSSQLServer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLServer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLServer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLServer.
func (mg *MSSQLServer) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLServer.
func (mg *MSSQLServer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLServer.
func (mg *MSSQLServer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLServer.
func (mg *MSSQLServer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLServer.
func (mg *MSSQLServer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLServer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLServer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLServer.
func (mg *MSSQLServer) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLServer.
func (mg *MSSQLServer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLServerTransparentDataEncryption.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLServerTransparentDataEncryption) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLServerTransparentDataEncryption.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLServerTransparentDataEncryption) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
