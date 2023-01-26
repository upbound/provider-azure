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

type TriggerCustomEventObservation struct {

	// The ID of the Data Factory Custom Event Trigger.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TriggerCustomEventParameters struct {

	// Specifies if the Data Factory Custom Event Trigger is activated. Defaults to true.
	// +kubebuilder:validation:Optional
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// A map of additional properties to associate with the Data Factory Custom Event Trigger.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Custom Event Trigger.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Custom Event Trigger.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of Event Grid Topic in which event will be listened. Changing this forces a new resource.
	// +kubebuilder:validation:Required
	EventGridTopicID *string `json:"eventgridTopicId" tf:"eventgrid_topic_id,omitempty"`

	// List of events that will fire this trigger. At least one event must be specified.
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// One or more pipeline blocks as defined below.
	// +kubebuilder:validation:Required
	Pipeline []TriggerCustomEventPipelineParameters `json:"pipeline" tf:"pipeline,omitempty"`

	// The pattern that event subject starts with for trigger to fire.
	// +kubebuilder:validation:Optional
	SubjectBeginsWith *string `json:"subjectBeginsWith,omitempty" tf:"subject_begins_with,omitempty"`

	// The pattern that event subject ends with for trigger to fire.
	// +kubebuilder:validation:Optional
	SubjectEndsWith *string `json:"subjectEndsWith,omitempty" tf:"subject_ends_with,omitempty"`
}

type TriggerCustomEventPipelineObservation struct {
}

type TriggerCustomEventPipelineParameters struct {

	// The Data Factory Pipeline name that the trigger will act on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Pipeline
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Pipeline in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Pipeline in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The Data Factory Pipeline parameters that the trigger will act on.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// TriggerCustomEventSpec defines the desired state of TriggerCustomEvent
type TriggerCustomEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerCustomEventParameters `json:"forProvider"`
}

// TriggerCustomEventStatus defines the observed state of TriggerCustomEvent.
type TriggerCustomEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerCustomEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerCustomEvent is the Schema for the TriggerCustomEvents API. Manages a Custom Event Trigger inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TriggerCustomEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerCustomEventSpec   `json:"spec"`
	Status            TriggerCustomEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerCustomEventList contains a list of TriggerCustomEvents
type TriggerCustomEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerCustomEvent `json:"items"`
}

// Repository type metadata.
var (
	TriggerCustomEvent_Kind             = "TriggerCustomEvent"
	TriggerCustomEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerCustomEvent_Kind}.String()
	TriggerCustomEvent_KindAPIVersion   = TriggerCustomEvent_Kind + "." + CRDGroupVersion.String()
	TriggerCustomEvent_GroupVersionKind = CRDGroupVersion.WithKind(TriggerCustomEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerCustomEvent{}, &TriggerCustomEventList{})
}
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

type TriggerCustomEventObservation struct {

	// The ID of the Data Factory Custom Event Trigger.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TriggerCustomEventParameters struct {

	// Specifies if the Data Factory Custom Event Trigger is activated. Defaults to true.
	// +kubebuilder:validation:Optional
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// A map of additional properties to associate with the Data Factory Custom Event Trigger.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Custom Event Trigger.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Custom Event Trigger.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of Event Grid Topic in which event will be listened. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventgrid/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EventGridTopicID *string `json:"eventgridTopicId,omitempty" tf:"eventgrid_topic_id,omitempty"`

	// Reference to a Topic in eventgrid to populate eventgridTopicId.
	// +kubebuilder:validation:Optional
	EventGridTopicIDRef *v1.Reference `json:"eventgridTopicIdRef,omitempty" tf:"-"`

	// Selector for a Topic in eventgrid to populate eventgridTopicId.
	// +kubebuilder:validation:Optional
	EventGridTopicIDSelector *v1.Selector `json:"eventgridTopicIdSelector,omitempty" tf:"-"`

	// List of events that will fire this trigger. At least one event must be specified.
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// One or more pipeline blocks as defined below.
	// +kubebuilder:validation:Required
	Pipeline []TriggerCustomEventPipelineParameters `json:"pipeline" tf:"pipeline,omitempty"`

	// The pattern that event subject starts with for trigger to fire.
	// +kubebuilder:validation:Optional
	SubjectBeginsWith *string `json:"subjectBeginsWith,omitempty" tf:"subject_begins_with,omitempty"`

	// The pattern that event subject ends with for trigger to fire.
	// +kubebuilder:validation:Optional
	SubjectEndsWith *string `json:"subjectEndsWith,omitempty" tf:"subject_ends_with,omitempty"`
}

type TriggerCustomEventPipelineObservation struct {
}

type TriggerCustomEventPipelineParameters struct {

	// The Data Factory Pipeline name that the trigger will act on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Pipeline
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Pipeline in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Pipeline in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The Data Factory Pipeline parameters that the trigger will act on.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// TriggerCustomEventSpec defines the desired state of TriggerCustomEvent
type TriggerCustomEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerCustomEventParameters `json:"forProvider"`
}

// TriggerCustomEventStatus defines the observed state of TriggerCustomEvent.
type TriggerCustomEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerCustomEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerCustomEvent is the Schema for the TriggerCustomEvents API. Manages a Custom Event Trigger inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TriggerCustomEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerCustomEventSpec   `json:"spec"`
	Status            TriggerCustomEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerCustomEventList contains a list of TriggerCustomEvents
type TriggerCustomEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerCustomEvent `json:"items"`
}

// Repository type metadata.
var (
	TriggerCustomEvent_Kind             = "TriggerCustomEvent"
	TriggerCustomEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerCustomEvent_Kind}.String()
	TriggerCustomEvent_KindAPIVersion   = TriggerCustomEvent_Kind + "." + CRDGroupVersion.String()
	TriggerCustomEvent_GroupVersionKind = CRDGroupVersion.WithKind(TriggerCustomEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerCustomEvent{}, &TriggerCustomEventList{})
}
