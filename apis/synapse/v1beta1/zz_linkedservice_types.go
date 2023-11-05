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

type IntegrationRuntimeInitParameters struct {

	// A map of parameters to associate with the integration runtime.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type IntegrationRuntimeObservation struct {

	// The integration runtime reference to associate with the Synapse Linked Service.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the integration runtime.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type IntegrationRuntimeParameters struct {

	// The integration runtime reference to associate with the Synapse Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.IntegrationRuntimeAzure
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a IntegrationRuntimeAzure in synapse to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a IntegrationRuntimeAzure in synapse to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the integration runtime.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceInitParameters struct {

	// A map of additional properties to associate with the Synapse Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Synapse Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Synapse Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A integration_runtime block as defined below.
	IntegrationRuntime []IntegrationRuntimeInitParameters `json:"integrationRuntime,omitempty" tf:"integration_runtime,omitempty"`

	// A map of parameters to associate with the Synapse Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The type of data stores that will be connected to Synapse. Valid Values include AmazonMWS, AmazonRdsForOracle, AmazonRdsForSqlServer, AmazonRedshift, AmazonS3, AzureBatch. Changing this forces a new resource to be created.
	// AzureBlobFS, AzureBlobStorage, AzureDataExplorer, AzureDataLakeAnalytics, AzureDataLakeStore, AzureDatabricks, AzureDatabricksDeltaLake, AzureFileStorage, AzureFunction,
	// AzureKeyVault, AzureML, AzureMLService, AzureMariaDB, AzureMySql, AzurePostgreSql, AzureSqlDW, AzureSqlDatabase, AzureSqlMI, AzureSearch, AzureStorage,
	// AzureTableStorage, Cassandra, CommonDataServiceForApps, Concur, CosmosDb, CosmosDbMongoDbApi, Couchbase, CustomDataSource, Db2, Drill,
	// Dynamics, DynamicsAX, DynamicsCrm, Eloqua, FileServer, FtpServer, GoogleAdWords, GoogleBigQuery, GoogleCloudStorage, Greenplum, HBase, HDInsight,
	// HDInsightOnDemand, HttpServer, Hdfs, Hive, Hubspot, Impala, Informix, Jira, LinkedService, Magento, MariaDB, Marketo, MicrosoftAccess, MongoDb,
	// MongoDbAtlas, MongoDbV2, MySql, Netezza, OData, Odbc, Office365, Oracle, OracleServiceCloud, Paypal, Phoenix, PostgreSql, Presto, QuickBooks,
	// Responsys, RestService, SqlServer, Salesforce, SalesforceMarketingCloud, SalesforceServiceCloud, SapBW, SapCloudForCustomer, SapEcc, SapHana, SapOpenHub,
	// SapTable, ServiceNow, Sftp, SharePointOnlineList, Shopify, Snowflake, Spark, Square, Sybase, Teradata, Vertica, Web, Xero, Zoho.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A JSON object that contains the properties of the Synapse Linked Service.
	TypePropertiesJSON *string `json:"typePropertiesJson,omitempty" tf:"type_properties_json,omitempty"`
}

type LinkedServiceObservation struct {

	// A map of additional properties to associate with the Synapse Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Synapse Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Synapse Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Synapse Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A integration_runtime block as defined below.
	IntegrationRuntime []IntegrationRuntimeObservation `json:"integrationRuntime,omitempty" tf:"integration_runtime,omitempty"`

	// A map of parameters to associate with the Synapse Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The Synapse Workspace ID in which to associate the Linked Service with. Changing this forces a new Synapse Linked Service to be created.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// The type of data stores that will be connected to Synapse. Valid Values include AmazonMWS, AmazonRdsForOracle, AmazonRdsForSqlServer, AmazonRedshift, AmazonS3, AzureBatch. Changing this forces a new resource to be created.
	// AzureBlobFS, AzureBlobStorage, AzureDataExplorer, AzureDataLakeAnalytics, AzureDataLakeStore, AzureDatabricks, AzureDatabricksDeltaLake, AzureFileStorage, AzureFunction,
	// AzureKeyVault, AzureML, AzureMLService, AzureMariaDB, AzureMySql, AzurePostgreSql, AzureSqlDW, AzureSqlDatabase, AzureSqlMI, AzureSearch, AzureStorage,
	// AzureTableStorage, Cassandra, CommonDataServiceForApps, Concur, CosmosDb, CosmosDbMongoDbApi, Couchbase, CustomDataSource, Db2, Drill,
	// Dynamics, DynamicsAX, DynamicsCrm, Eloqua, FileServer, FtpServer, GoogleAdWords, GoogleBigQuery, GoogleCloudStorage, Greenplum, HBase, HDInsight,
	// HDInsightOnDemand, HttpServer, Hdfs, Hive, Hubspot, Impala, Informix, Jira, LinkedService, Magento, MariaDB, Marketo, MicrosoftAccess, MongoDb,
	// MongoDbAtlas, MongoDbV2, MySql, Netezza, OData, Odbc, Office365, Oracle, OracleServiceCloud, Paypal, Phoenix, PostgreSql, Presto, QuickBooks,
	// Responsys, RestService, SqlServer, Salesforce, SalesforceMarketingCloud, SalesforceServiceCloud, SapBW, SapCloudForCustomer, SapEcc, SapHana, SapOpenHub,
	// SapTable, ServiceNow, Sftp, SharePointOnlineList, Shopify, Snowflake, Spark, Square, Sybase, Teradata, Vertica, Web, Xero, Zoho.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A JSON object that contains the properties of the Synapse Linked Service.
	TypePropertiesJSON *string `json:"typePropertiesJson,omitempty" tf:"type_properties_json,omitempty"`
}

type LinkedServiceParameters struct {

	// A map of additional properties to associate with the Synapse Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Synapse Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Synapse Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A integration_runtime block as defined below.
	// +kubebuilder:validation:Optional
	IntegrationRuntime []IntegrationRuntimeParameters `json:"integrationRuntime,omitempty" tf:"integration_runtime,omitempty"`

	// A map of parameters to associate with the Synapse Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The Synapse Workspace ID in which to associate the Linked Service with. Changing this forces a new Synapse Linked Service to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`

	// The type of data stores that will be connected to Synapse. Valid Values include AmazonMWS, AmazonRdsForOracle, AmazonRdsForSqlServer, AmazonRedshift, AmazonS3, AzureBatch. Changing this forces a new resource to be created.
	// AzureBlobFS, AzureBlobStorage, AzureDataExplorer, AzureDataLakeAnalytics, AzureDataLakeStore, AzureDatabricks, AzureDatabricksDeltaLake, AzureFileStorage, AzureFunction,
	// AzureKeyVault, AzureML, AzureMLService, AzureMariaDB, AzureMySql, AzurePostgreSql, AzureSqlDW, AzureSqlDatabase, AzureSqlMI, AzureSearch, AzureStorage,
	// AzureTableStorage, Cassandra, CommonDataServiceForApps, Concur, CosmosDb, CosmosDbMongoDbApi, Couchbase, CustomDataSource, Db2, Drill,
	// Dynamics, DynamicsAX, DynamicsCrm, Eloqua, FileServer, FtpServer, GoogleAdWords, GoogleBigQuery, GoogleCloudStorage, Greenplum, HBase, HDInsight,
	// HDInsightOnDemand, HttpServer, Hdfs, Hive, Hubspot, Impala, Informix, Jira, LinkedService, Magento, MariaDB, Marketo, MicrosoftAccess, MongoDb,
	// MongoDbAtlas, MongoDbV2, MySql, Netezza, OData, Odbc, Office365, Oracle, OracleServiceCloud, Paypal, Phoenix, PostgreSql, Presto, QuickBooks,
	// Responsys, RestService, SqlServer, Salesforce, SalesforceMarketingCloud, SalesforceServiceCloud, SapBW, SapCloudForCustomer, SapEcc, SapHana, SapOpenHub,
	// SapTable, ServiceNow, Sftp, SharePointOnlineList, Shopify, Snowflake, Spark, Square, Sybase, Teradata, Vertica, Web, Xero, Zoho.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A JSON object that contains the properties of the Synapse Linked Service.
	// +kubebuilder:validation:Optional
	TypePropertiesJSON *string `json:"typePropertiesJson,omitempty" tf:"type_properties_json,omitempty"`
}

// LinkedServiceSpec defines the desired state of LinkedService
type LinkedServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceParameters `json:"forProvider"`
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
	InitProvider LinkedServiceInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceStatus defines the observed state of LinkedService.
type LinkedServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedService is the Schema for the LinkedServices API. Manages a Linked Service (connection) between a resource and Azure Synapse. This is a generic resource that supports all different Linked Service Types.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.typePropertiesJson) || (has(self.initProvider) && has(self.initProvider.typePropertiesJson))",message="spec.forProvider.typePropertiesJson is a required parameter"
	Spec   LinkedServiceSpec   `json:"spec"`
	Status LinkedServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceList contains a list of LinkedServices
type LinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedService `json:"items"`
}

// Repository type metadata.
var (
	LinkedService_Kind             = "LinkedService"
	LinkedService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedService_Kind}.String()
	LinkedService_KindAPIVersion   = LinkedService_Kind + "." + CRDGroupVersion.String()
	LinkedService_GroupVersionKind = CRDGroupVersion.WithKind(LinkedService_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedService{}, &LinkedServiceList{})
}
