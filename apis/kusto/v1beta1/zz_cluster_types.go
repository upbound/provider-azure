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

type ClusterObservation struct {

	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionURI *string `json:"dataIngestionUri,omitempty" tf:"data_ingestion_uri,omitempty"`

	// The Kusto Cluster ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The FQDN of the Azure Kusto Cluster.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ClusterParameters struct {

	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	// +kubebuilder:validation:Optional
	AutoStopEnabled *bool `json:"autoStopEnabled,omitempty" tf:"auto_stop_enabled,omitempty"`

	// Specifies if the cluster's disks are encrypted.
	// +kubebuilder:validation:Optional
	DiskEncryptionEnabled *bool `json:"diskEncryptionEnabled,omitempty" tf:"disk_encryption_enabled,omitempty"`

	// Is the cluster's double encryption enabled? Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty" tf:"double_encryption_enabled,omitempty"`

	// . The engine type that should be used. Possible values are V2 and V3. Defaults to V2.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// An list of language_extensions to enable. Valid values are: PYTHON and R.
	// +kubebuilder:validation:Optional
	LanguageExtensions []*string `json:"languageExtensions,omitempty" tf:"language_extensions,omitempty"`

	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// An optimized_auto_scale block as defined below.
	// +kubebuilder:validation:Optional
	OptimizedAutoScale []OptimizedAutoScaleParameters `json:"optimizedAutoScale,omitempty" tf:"optimized_auto_scale,omitempty"`

	// Is the public network access enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies if the purge operations are enabled.
	// +kubebuilder:validation:Optional
	PurgeEnabled *bool `json:"purgeEnabled,omitempty" tf:"purge_enabled,omitempty"`

	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A sku block as defined below.
	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// Specifies if the streaming ingest is enabled.
	// +kubebuilder:validation:Optional
	StreamingIngestionEnabled *bool `json:"streamingIngestionEnabled,omitempty" tf:"streaming_ingestion_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use trusted_external_tenants = ["*"] to explicitly allow all other tenants, trusted_external_tenants = ["MyTenantOnly"] for only your tenant or trusted_external_tenants = ["<tenantId1>", "<tenantIdx>"] to allow specific other tenants.
	// +kubebuilder:validation:Optional
	TrustedExternalTenants []*string `json:"trustedExternalTenants,omitempty" tf:"trusted_external_tenants,omitempty"`

	// A virtual_network_configuration block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VirtualNetworkConfiguration []VirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// Specifies a list of Availability Zones in which this Kusto Cluster should be located. Changing this forces a new Kusto Cluster to be created.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type IdentityObservation struct {

	// The Principal ID associated with this System Assigned Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this System Assigned Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Kusto Cluster.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: SystemAssigned, UserAssigned and SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type OptimizedAutoScaleObservation struct {
}

type OptimizedAutoScaleParameters struct {

	// The maximum number of allowed instances. Must between 0 and 1000.
	// +kubebuilder:validation:Required
	MaximumInstances *float64 `json:"maximumInstances" tf:"maximum_instances,omitempty"`

	// The minimum number of allowed instances. Must between 0 and 1000.
	// +kubebuilder:validation:Required
	MinimumInstances *float64 `json:"minimumInstances" tf:"minimum_instances,omitempty"`
}

type SkuObservation struct {
}

type SkuParameters struct {

	// Specifies the node count for the cluster. Boundaries depend on the SKU name.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the SKU. Valid values are: Dev(No SLA)_Standard_D11_v2, Dev(No SLA)_Standard_E2a_v4, Standard_D11_v2, Standard_D12_v2, Standard_D13_v2, Standard_D14_v2, Standard_DS13_v2+1TB_PS, Standard_DS13_v2+2TB_PS, Standard_DS14_v2+3TB_PS, Standard_DS14_v2+4TB_PS, Standard_E16as_v4+3TB_PS, Standard_E16as_v4+4TB_PS, Standard_E16a_v4, Standard_E2a_v4, Standard_E4a_v4, Standard_E64i_v3, Standard_E8as_v4+1TB_PS, Standard_E8as_v4+2TB_PS, Standard_E8a_v4, Standard_L16s, Standard_L4s, Standard_L8s, Standard_L16s_v2 and Standard_L8s_v2.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type VirtualNetworkConfigurationObservation struct {
}

type VirtualNetworkConfigurationParameters struct {

	// Data management's service public IP address resource id.
	// +kubebuilder:validation:Required
	DataManagementPublicIPID *string `json:"dataManagementPublicIpId" tf:"data_management_public_ip_id,omitempty"`

	// Engine service's public IP address resource id.
	// +kubebuilder:validation:Required
	EnginePublicIPID *string `json:"enginePublicIpId" tf:"engine_public_ip_id,omitempty"`

	// The subnet resource id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Manages Kusto (also known as Azure Data Explorer) Cluster
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
