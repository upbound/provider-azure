/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this FirewallRule.
func (mg *FirewallRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FirewallRule.
func (mg *FirewallRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FirewallRule.
func (mg *FirewallRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FirewallRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FirewallRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FirewallRule.
func (mg *FirewallRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FirewallRule.
func (mg *FirewallRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FirewallRule.
func (mg *FirewallRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FirewallRule.
func (mg *FirewallRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FirewallRule.
func (mg *FirewallRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FirewallRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FirewallRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FirewallRule.
func (mg *FirewallRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FirewallRule.
func (mg *FirewallRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IntegrationRuntimeAzure.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IntegrationRuntimeAzure) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IntegrationRuntimeAzure.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IntegrationRuntimeAzure) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IntegrationRuntimeAzure.
func (mg *IntegrationRuntimeAzure) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this IntegrationRuntimeSelfHosted.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *IntegrationRuntimeSelfHosted) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this IntegrationRuntimeSelfHosted.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *IntegrationRuntimeSelfHosted) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this IntegrationRuntimeSelfHosted.
func (mg *IntegrationRuntimeSelfHosted) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LinkedService.
func (mg *LinkedService) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LinkedService.
func (mg *LinkedService) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LinkedService.
func (mg *LinkedService) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LinkedService.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LinkedService) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LinkedService.
func (mg *LinkedService) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LinkedService.
func (mg *LinkedService) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LinkedService.
func (mg *LinkedService) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LinkedService.
func (mg *LinkedService) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LinkedService.
func (mg *LinkedService) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LinkedService.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LinkedService) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LinkedService.
func (mg *LinkedService) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LinkedService.
func (mg *LinkedService) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ManagedPrivateEndpoint.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ManagedPrivateEndpoint) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ManagedPrivateEndpoint.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ManagedPrivateEndpoint) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PrivateLinkHub.
func (mg *PrivateLinkHub) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PrivateLinkHub.
func (mg *PrivateLinkHub) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PrivateLinkHub.
func (mg *PrivateLinkHub) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PrivateLinkHub.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PrivateLinkHub) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PrivateLinkHub.
func (mg *PrivateLinkHub) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PrivateLinkHub.
func (mg *PrivateLinkHub) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PrivateLinkHub.
func (mg *PrivateLinkHub) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PrivateLinkHub.
func (mg *PrivateLinkHub) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PrivateLinkHub.
func (mg *PrivateLinkHub) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PrivateLinkHub.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PrivateLinkHub) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PrivateLinkHub.
func (mg *PrivateLinkHub) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PrivateLinkHub.
func (mg *PrivateLinkHub) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SQLPool.
func (mg *SQLPool) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SQLPool.
func (mg *SQLPool) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SQLPool.
func (mg *SQLPool) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SQLPool.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SQLPool) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SQLPool.
func (mg *SQLPool) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SQLPool.
func (mg *SQLPool) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SQLPool.
func (mg *SQLPool) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SQLPool.
func (mg *SQLPool) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SQLPool.
func (mg *SQLPool) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SQLPool.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SQLPool) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SQLPool.
func (mg *SQLPool) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SQLPool.
func (mg *SQLPool) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SQLPoolExtendedAuditingPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SQLPoolExtendedAuditingPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SQLPoolExtendedAuditingPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SQLPoolExtendedAuditingPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SQLPoolExtendedAuditingPolicy.
func (mg *SQLPoolExtendedAuditingPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SQLPoolSecurityAlertPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SQLPoolSecurityAlertPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SQLPoolSecurityAlertPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SQLPoolSecurityAlertPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SQLPoolSecurityAlertPolicy.
func (mg *SQLPoolSecurityAlertPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SQLPoolWorkloadClassifier.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SQLPoolWorkloadClassifier) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SQLPoolWorkloadClassifier.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SQLPoolWorkloadClassifier) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SQLPoolWorkloadClassifier.
func (mg *SQLPoolWorkloadClassifier) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SQLPoolWorkloadGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SQLPoolWorkloadGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SQLPoolWorkloadGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SQLPoolWorkloadGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SQLPoolWorkloadGroup.
func (mg *SQLPoolWorkloadGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SparkPool.
func (mg *SparkPool) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SparkPool.
func (mg *SparkPool) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SparkPool.
func (mg *SparkPool) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SparkPool.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SparkPool) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SparkPool.
func (mg *SparkPool) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SparkPool.
func (mg *SparkPool) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SparkPool.
func (mg *SparkPool) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SparkPool.
func (mg *SparkPool) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SparkPool.
func (mg *SparkPool) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SparkPool.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SparkPool) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SparkPool.
func (mg *SparkPool) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SparkPool.
func (mg *SparkPool) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Workspace.
func (mg *Workspace) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Workspace.
func (mg *Workspace) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Workspace.
func (mg *Workspace) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Workspace.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Workspace) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this Workspace.
func (mg *Workspace) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Workspace.
func (mg *Workspace) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Workspace.
func (mg *Workspace) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Workspace.
func (mg *Workspace) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Workspace.
func (mg *Workspace) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Workspace.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Workspace) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this Workspace.
func (mg *Workspace) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Workspace.
func (mg *Workspace) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this WorkspaceAADAdmin.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *WorkspaceAADAdmin) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this WorkspaceAADAdmin.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *WorkspaceAADAdmin) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this WorkspaceAADAdmin.
func (mg *WorkspaceAADAdmin) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this WorkspaceExtendedAuditingPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *WorkspaceExtendedAuditingPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this WorkspaceExtendedAuditingPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *WorkspaceExtendedAuditingPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this WorkspaceExtendedAuditingPolicy.
func (mg *WorkspaceExtendedAuditingPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this WorkspaceSecurityAlertPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *WorkspaceSecurityAlertPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this WorkspaceSecurityAlertPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *WorkspaceSecurityAlertPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this WorkspaceSecurityAlertPolicy.
func (mg *WorkspaceSecurityAlertPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
