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

type BackupInstanceDiskInitParameters struct {

	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type BackupInstanceDiskObservation struct {

	// The ID of the Backup Policy.
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// The ID of the source Disk. Changing this forces a new Backup Instance Disk to be created.
	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// The ID of the Backup Instance Disk.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where snapshots are stored. Changing this forces a new Backup Instance Disk to be created.
	SnapshotResourceGroupName *string `json:"snapshotResourceGroupName,omitempty" tf:"snapshot_resource_group_name,omitempty"`

	// The ID of the Backup Vault within which the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type BackupInstanceDiskParameters struct {

	// The ID of the Backup Policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dataprotection/v1beta1.BackupPolicyDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// Reference to a BackupPolicyDisk in dataprotection to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDRef *v1.Reference `json:"backupPolicyIdRef,omitempty" tf:"-"`

	// Selector for a BackupPolicyDisk in dataprotection to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDSelector *v1.Selector `json:"backupPolicyIdSelector,omitempty" tf:"-"`

	// The ID of the source Disk. Changing this forces a new Backup Instance Disk to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// Reference to a ManagedDisk in compute to populate diskId.
	// +kubebuilder:validation:Optional
	DiskIDRef *v1.Reference `json:"diskIdRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate diskId.
	// +kubebuilder:validation:Optional
	DiskIDSelector *v1.Selector `json:"diskIdSelector,omitempty" tf:"-"`

	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where snapshots are stored. Changing this forces a new Backup Instance Disk to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	SnapshotResourceGroupName *string `json:"snapshotResourceGroupName,omitempty" tf:"snapshot_resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate snapshotResourceGroupName.
	// +kubebuilder:validation:Optional
	SnapshotResourceGroupNameRef *v1.Reference `json:"snapshotResourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate snapshotResourceGroupName.
	// +kubebuilder:validation:Optional
	SnapshotResourceGroupNameSelector *v1.Selector `json:"snapshotResourceGroupNameSelector,omitempty" tf:"-"`

	// The ID of the Backup Vault within which the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dataprotection/v1beta1.BackupVault
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`

	// Reference to a BackupVault in dataprotection to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDRef *v1.Reference `json:"vaultIdRef,omitempty" tf:"-"`

	// Selector for a BackupVault in dataprotection to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDSelector *v1.Selector `json:"vaultIdSelector,omitempty" tf:"-"`
}

// BackupInstanceDiskSpec defines the desired state of BackupInstanceDisk
type BackupInstanceDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupInstanceDiskParameters `json:"forProvider"`
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
	InitProvider BackupInstanceDiskInitParameters `json:"initProvider,omitempty"`
}

// BackupInstanceDiskStatus defines the observed state of BackupInstanceDisk.
type BackupInstanceDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupInstanceDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupInstanceDisk is the Schema for the BackupInstanceDisks API. Manages a Backup Instance to back up Disk.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupInstanceDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   BackupInstanceDiskSpec   `json:"spec"`
	Status BackupInstanceDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupInstanceDiskList contains a list of BackupInstanceDisks
type BackupInstanceDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupInstanceDisk `json:"items"`
}

// Repository type metadata.
var (
	BackupInstanceDisk_Kind             = "BackupInstanceDisk"
	BackupInstanceDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupInstanceDisk_Kind}.String()
	BackupInstanceDisk_KindAPIVersion   = BackupInstanceDisk_Kind + "." + CRDGroupVersion.String()
	BackupInstanceDisk_GroupVersionKind = CRDGroupVersion.WithKind(BackupInstanceDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupInstanceDisk{}, &BackupInstanceDiskList{})
}
