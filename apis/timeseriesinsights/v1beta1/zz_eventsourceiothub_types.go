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

type EventSourceIOTHubObservation struct {

	// The ID of the IoT Time Series Insights IoTHub Event Source.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EventSourceIOTHubParameters struct {

	// Specifies the name of the IotHub Consumer Group that holds the partitions from which events will be read.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHubConsumerGroup
	// +kubebuilder:validation:Optional
	ConsumerGroupName *string `json:"consumerGroupName,omitempty" tf:"consumer_group_name,omitempty"`

	// Reference to a IOTHubConsumerGroup in devices to populate consumerGroupName.
	// +kubebuilder:validation:Optional
	ConsumerGroupNameRef *v1.Reference `json:"consumerGroupNameRef,omitempty" tf:"-"`

	// Selector for a IOTHubConsumerGroup in devices to populate consumerGroupName.
	// +kubebuilder:validation:Optional
	ConsumerGroupNameSelector *v1.Selector `json:"consumerGroupNameSelector,omitempty" tf:"-"`

	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/timeseriesinsights/v1beta1.Gen2Environment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// Reference to a Gen2Environment in timeseriesinsights to populate environmentId.
	// +kubebuilder:validation:Optional
	EnvironmentIDRef *v1.Reference `json:"environmentIdRef,omitempty" tf:"-"`

	// Selector for a Gen2Environment in timeseriesinsights to populate environmentId.
	// +kubebuilder:validation:Optional
	EnvironmentIDSelector *v1.Selector `json:"environmentIdSelector,omitempty" tf:"-"`

	// Specifies the resource id where events will be coming from.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EventSourceResourceID *string `json:"eventSourceResourceId,omitempty" tf:"event_source_resource_id,omitempty"`

	// Reference to a IOTHub in devices to populate eventSourceResourceId.
	// +kubebuilder:validation:Optional
	EventSourceResourceIDRef *v1.Reference `json:"eventSourceResourceIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate eventSourceResourceId.
	// +kubebuilder:validation:Optional
	EventSourceResourceIDSelector *v1.Selector `json:"eventSourceResourceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the IotHub which will be associated with this resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the name of the Shared Access key that grants the Event Source access to the IotHub.
	// +kubebuilder:validation:Required
	SharedAccessKeyName *string `json:"sharedAccessKeyName" tf:"shared_access_key_name,omitempty"`

	// Specifies the value of the Shared Access Policy key that grants the Time Series Insights service read access to the IotHub.
	// +kubebuilder:validation:Required
	SharedAccessKeySecretRef v1.SecretKeySelector `json:"sharedAccessKeySecretRef" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	// +kubebuilder:validation:Optional
	TimestampPropertyName *string `json:"timestampPropertyName,omitempty" tf:"timestamp_property_name,omitempty"`
}

// EventSourceIOTHubSpec defines the desired state of EventSourceIOTHub
type EventSourceIOTHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSourceIOTHubParameters `json:"forProvider"`
}

// EventSourceIOTHubStatus defines the observed state of EventSourceIOTHub.
type EventSourceIOTHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSourceIOTHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceIOTHub is the Schema for the EventSourceIOTHubs API. Manages an Azure IoT Time Series Insights IoTHub Event Source.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventSourceIOTHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSourceIOTHubSpec   `json:"spec"`
	Status            EventSourceIOTHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceIOTHubList contains a list of EventSourceIOTHubs
type EventSourceIOTHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSourceIOTHub `json:"items"`
}

// Repository type metadata.
var (
	EventSourceIOTHub_Kind             = "EventSourceIOTHub"
	EventSourceIOTHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSourceIOTHub_Kind}.String()
	EventSourceIOTHub_KindAPIVersion   = EventSourceIOTHub_Kind + "." + CRDGroupVersion.String()
	EventSourceIOTHub_GroupVersionKind = CRDGroupVersion.WithKind(EventSourceIOTHub_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSourceIOTHub{}, &EventSourceIOTHubList{})
}