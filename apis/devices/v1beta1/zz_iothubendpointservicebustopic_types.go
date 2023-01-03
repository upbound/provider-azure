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

type IOTHubEndpointServiceBusTopicObservation struct {

	// The ID of the IoTHub ServiceBus Topic Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubEndpointServiceBusTopicParameters struct {

	// Type used to authenticate against the Service Bus Topic endpoint. Possible values are keyBased and identityBased. Defaults to keyBased.
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The connection string for the endpoint. This attribute can only be specified and is mandatory when authentication_type is keyBased.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef *v1.SecretKeySelector `json:"connectionStringSecretRef,omitempty" tf:"-"`

	// URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	// +kubebuilder:validation:Optional
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// Name of the Service Bus Topic. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	// +kubebuilder:validation:Optional
	EntityPath *string `json:"entityPath,omitempty" tf:"entity_path,omitempty"`

	// The IoTHub ID for the endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// ID of the User Managed Identity used to authenticate against the Service Bus Topic endpoint.
	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// The name of the resource group under which the Service Bus Topic has been created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IOTHubEndpointServiceBusTopicSpec defines the desired state of IOTHubEndpointServiceBusTopic
type IOTHubEndpointServiceBusTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubEndpointServiceBusTopicParameters `json:"forProvider"`
}

// IOTHubEndpointServiceBusTopicStatus defines the observed state of IOTHubEndpointServiceBusTopic.
type IOTHubEndpointServiceBusTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubEndpointServiceBusTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEndpointServiceBusTopic is the Schema for the IOTHubEndpointServiceBusTopics API. Manages an IotHub ServiceBus Topic Endpoint
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubEndpointServiceBusTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubEndpointServiceBusTopicSpec   `json:"spec"`
	Status            IOTHubEndpointServiceBusTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEndpointServiceBusTopicList contains a list of IOTHubEndpointServiceBusTopics
type IOTHubEndpointServiceBusTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubEndpointServiceBusTopic `json:"items"`
}

// Repository type metadata.
var (
	IOTHubEndpointServiceBusTopic_Kind             = "IOTHubEndpointServiceBusTopic"
	IOTHubEndpointServiceBusTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubEndpointServiceBusTopic_Kind}.String()
	IOTHubEndpointServiceBusTopic_KindAPIVersion   = IOTHubEndpointServiceBusTopic_Kind + "." + CRDGroupVersion.String()
	IOTHubEndpointServiceBusTopic_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubEndpointServiceBusTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubEndpointServiceBusTopic{}, &IOTHubEndpointServiceBusTopicList{})
}
