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

type AuthenticationObservation struct {
}

type AuthenticationParameters struct {

	// Service principal certificate for servicePrincipal auth. Should be specified when type is set to servicePrincipalCertificate.
	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// Client ID for userAssignedIdentity or servicePrincipal auth. Should be specified when type is set to servicePrincipalSecret or servicePrincipalCertificate. When type is set to userAssignedIdentity, client_id and subscription_id should be either both specified or both not specified.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Username or account name for secret auth. name and secret should be either both specified or both not specified when type is set to secret.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Principal ID for servicePrincipal auth. Should be specified when type is set to servicePrincipalSecret or servicePrincipalCertificate.
	// +kubebuilder:validation:Optional
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// Password or account key for secret auth. secret and name should be either both specified or both not specified when type is set to secret.
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// Subscription ID for userAssignedIdentity. subscription_id and client_id should be either both specified or both not specified.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The authentication type. Possible values are systemAssignedIdentity, userAssignedIdentity, servicePrincipalSecret, servicePrincipalCertificate, secret. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SecretStoreObservation struct {
}

type SecretStoreParameters struct {

	// The key vault id to store secret.
	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`
}

type SpringCloudConnectionObservation struct {

	// The ID of the service connector.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SpringCloudConnectionParameters struct {

	// The authentication info. An authentication block as defined below.
	// +kubebuilder:validation:Required
	Authentication []AuthenticationParameters `json:"authentication" tf:"authentication,omitempty"`

	// The application client type. Possible values are none, dotnet, java, python, go, php, ruby, django, nodejs and springBoot.
	// +kubebuilder:validation:Optional
	ClientType *string `json:"clientType,omitempty" tf:"client_type,omitempty"`

	// The name of the service connection. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// An option to store secret value in secure place. An secret_store block as defined below.
	// +kubebuilder:validation:Optional
	SecretStore []SecretStoreParameters `json:"secretStore,omitempty" tf:"secret_store,omitempty"`

	// The ID of the data source spring cloud. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudJavaDeployment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudID *string `json:"springCloudId,omitempty" tf:"spring_cloud_id,omitempty"`

	// Reference to a SpringCloudJavaDeployment in appplatform to populate springCloudId.
	// +kubebuilder:validation:Optional
	SpringCloudIDRef *v1.Reference `json:"springCloudIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudJavaDeployment in appplatform to populate springCloudId.
	// +kubebuilder:validation:Optional
	SpringCloudIDSelector *v1.Selector `json:"springCloudIdSelector,omitempty" tf:"-"`

	// The ID of the target resource. Changing this forces a new resource to be created. Possible values are Postgres, PostgresFlexible, Mysql, Sql, Redis, RedisEnterprise, CosmosCassandra, CosmosGremlin, CosmosMongo, CosmosSql, CosmosTable, StorageBlob, StorageQueue, StorageFile, StorageTable, AppConfig, EventHub, ServiceBus, SignalR, WebPubSub, ConfluentKafka.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.SQLDatabase
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a SQLDatabase in cosmosdb to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a SQLDatabase in cosmosdb to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`

	// The type of the VNet solution. Possible values are serviceEndpoint, privateLink.
	// +kubebuilder:validation:Optional
	VnetSolution *string `json:"vnetSolution,omitempty" tf:"vnet_solution,omitempty"`
}

// SpringCloudConnectionSpec defines the desired state of SpringCloudConnection
type SpringCloudConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudConnectionParameters `json:"forProvider"`
}

// SpringCloudConnectionStatus defines the observed state of SpringCloudConnection.
type SpringCloudConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudConnection is the Schema for the SpringCloudConnections API. Manages a service connector for spring cloud app.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudConnectionSpec   `json:"spec"`
	Status            SpringCloudConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudConnectionList contains a list of SpringCloudConnections
type SpringCloudConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudConnection `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudConnection_Kind             = "SpringCloudConnection"
	SpringCloudConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudConnection_Kind}.String()
	SpringCloudConnection_KindAPIVersion   = SpringCloudConnection_Kind + "." + CRDGroupVersion.String()
	SpringCloudConnection_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudConnection{}, &SpringCloudConnectionList{})
}
