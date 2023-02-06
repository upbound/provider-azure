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

type CriteriaDimensionObservation struct {
}

type CriteriaDimensionParameters struct {

	// Specifies the name which should be used for this Monitor Scheduled Query Rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Operator for dimension values. Possible values are Exclude,and Include.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// List of dimension values. Use a wildcard * to collect all.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type FailingPeriodsObservation struct {
}

type FailingPeriodsParameters struct {

	// Specifies the number of violations to trigger an alert. Should be smaller or equal to number_of_evaluation_periods. Possible value is integer between 1 and 6.
	// +kubebuilder:validation:Required
	MinimumFailingPeriodsToTriggerAlert *float64 `json:"minimumFailingPeriodsToTriggerAlert" tf:"minimum_failing_periods_to_trigger_alert,omitempty"`

	// Specifies the number of aggregated look-back points. The look-back time window is calculated based on the aggregation granularity window_duration and the selected number of aggregated points. Possible value is integer between 1 and 6.
	// +kubebuilder:validation:Required
	NumberOfEvaluationPeriods *float64 `json:"numberOfEvaluationPeriods" tf:"number_of_evaluation_periods,omitempty"`
}

type MonitorScheduledQueryRulesAlertV2ActionObservation struct {
}

type MonitorScheduledQueryRulesAlertV2ActionParameters struct {

	// List of Action Group resource ids to invoke when the alert fires.
	// +kubebuilder:validation:Optional
	ActionGroups []*string `json:"actionGroups,omitempty" tf:"action_groups,omitempty"`

	// Specifies the properties of an alert payload.
	// +kubebuilder:validation:Optional
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`
}

type MonitorScheduledQueryRulesAlertV2CriteriaObservation struct {
}

type MonitorScheduledQueryRulesAlertV2CriteriaParameters struct {

	// A dimension block as defined below.
	// +kubebuilder:validation:Optional
	Dimension []CriteriaDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A failing_periods block as defined below.
	// +kubebuilder:validation:Optional
	FailingPeriods []FailingPeriodsParameters `json:"failingPeriods,omitempty" tf:"failing_periods,omitempty"`

	// Specifies the column containing the metric measure number.
	// +kubebuilder:validation:Optional
	MetricMeasureColumn *string `json:"metricMeasureColumn,omitempty" tf:"metric_measure_column,omitempty"`

	// Specifies the criteria operator. Possible values are Equal, GreaterThan, GreaterThanOrEqual, LessThan,and LessThanOrEqual.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// The query to run on logs. The results returned by this query are used to populate the alert.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// Specifies the column containing the resource id. The content of the column must be an uri formatted as resource id.
	// +kubebuilder:validation:Optional
	ResourceIDColumn *string `json:"resourceIdColumn,omitempty" tf:"resource_id_column,omitempty"`

	// Specifies the criteria threshold value that activates the alert.
	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// The type of aggregation to apply to the data points in aggregation granularity. Possible values are Average, Count, Maximum, Minimum,and Total.
	// +kubebuilder:validation:Required
	TimeAggregationMethod *string `json:"timeAggregationMethod" tf:"time_aggregation_method,omitempty"`
}

type MonitorScheduledQueryRulesAlertV2Observation struct {

	// The api-version used when creating this alert rule.
	CreatedWithAPIVersion *string `json:"createdWithApiVersion,omitempty" tf:"created_with_api_version,omitempty"`

	// The id of the Monitor Scheduled Query Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// True if this alert rule is a legacy Log Analytic Rule.
	IsALegacyLogAnalyticsRule *bool `json:"isALegacyLogAnalyticsRule,omitempty" tf:"is_a_legacy_log_analytics_rule,omitempty"`

	// The flag indicates whether this Scheduled Query Rule has been configured to be stored in the customer's storage.
	IsWorkspaceAlertsStorageConfigured *bool `json:"isWorkspaceAlertsStorageConfigured,omitempty" tf:"is_workspace_alerts_storage_configured,omitempty"`
}

type MonitorScheduledQueryRulesAlertV2Parameters struct {

	// An action block as defined below.
	// +kubebuilder:validation:Optional
	Action []MonitorScheduledQueryRulesAlertV2ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the flag that indicates whether the alert should be automatically resolved or not. Value should be true or false. The default is false.
	// +kubebuilder:validation:Optional
	AutoMitigationEnabled *bool `json:"autoMitigationEnabled,omitempty" tf:"auto_mitigation_enabled,omitempty"`

	// A criteria block as defined below.
	// +kubebuilder:validation:Required
	Criteria []MonitorScheduledQueryRulesAlertV2CriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// Specifies the description of the scheduled query rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the display name of the alert rule.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies the flag which indicates whether this scheduled query rule is enabled. Value should be true or false. The default is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// How often the scheduled query rule is evaluated, represented in ISO 8601 duration format. Possible values are PT1M, PT5M, PT10M, PT15M, PT30M, PT45M, PT1H, PT2H, PT3H, PT4H, PT5H, PT6H, P1D.
	// +kubebuilder:validation:Optional
	EvaluationFrequency *string `json:"evaluationFrequency,omitempty" tf:"evaluation_frequency,omitempty"`

	// Specifies the Azure Region where the Monitor Scheduled Query Rule should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Mute actions for the chosen period of time in ISO 8601 duration format after the alert is fired. Possible values are PT5M, PT10M, PT15M, PT30M, PT45M, PT1H, PT2H, PT3H, PT4H, PT5H, PT6H, P1D and P2D.
	// +kubebuilder:validation:Optional
	MuteActionsAfterAlertDuration *string `json:"muteActionsAfterAlertDuration,omitempty" tf:"mute_actions_after_alert_duration,omitempty"`

	// Set this if the alert evaluation period is different from the query time range. If not specified, the value is window_duration*number_of_evaluation_periods. Possible values are PT5M, PT10M, PT15M, PT20M, PT30M, PT45M, PT1H, PT2H, PT3H, PT4H, PT5H, PT6H, P1D and P2D.
	// +kubebuilder:validation:Optional
	QueryTimeRangeOverride *string `json:"queryTimeRangeOverride,omitempty" tf:"query_time_range_override,omitempty"`

	// Specifies the name of the Resource Group where the Monitor Scheduled Query Rule should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the list of resource ids that this scheduled query rule is scoped to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// References to ApplicationInsights to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesRefs []v1.Reference `json:"scopesRefs,omitempty" tf:"-"`

	// Selector for a list of ApplicationInsights to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesSelector *v1.Selector `json:"scopesSelector,omitempty" tf:"-"`

	// Severity of the alert. Should be an integer between 0 and 4. Value of 0 is severest.
	// +kubebuilder:validation:Required
	Severity *float64 `json:"severity" tf:"severity,omitempty"`

	// Specifies the flag which indicates whether the provided query should be validated or not. The default is false.
	// +kubebuilder:validation:Optional
	SkipQueryValidation *bool `json:"skipQueryValidation,omitempty" tf:"skip_query_validation,omitempty"`

	// A mapping of tags which should be assigned to the Monitor Scheduled Query Rule.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria.
	// +kubebuilder:validation:Optional
	TargetResourceTypes []*string `json:"targetResourceTypes,omitempty" tf:"target_resource_types,omitempty"`

	// Specifies the period of time in ISO 8601 duration format on which the Scheduled Query Rule will be executed (bin size). If evaluation_frequency is PT1M, possible values are PT1M, PT5M, PT10M, PT15M, PT30M, PT45M, PT1H, PT2H, PT3H, PT4H, PT5H, and PT6H. Otherwise, possible values are PT5M, PT10M, PT15M, PT30M, PT45M, PT1H, PT2H, PT3H, PT4H, PT5H, PT6H, P1D, and P2D.
	// +kubebuilder:validation:Required
	WindowDuration *string `json:"windowDuration" tf:"window_duration,omitempty"`

	// Specifies the flag which indicates whether this scheduled query rule check if storage is configured. Value should be true or false. The default is false.
	// +kubebuilder:validation:Optional
	WorkspaceAlertsStorageEnabled *bool `json:"workspaceAlertsStorageEnabled,omitempty" tf:"workspace_alerts_storage_enabled,omitempty"`
}

// MonitorScheduledQueryRulesAlertV2Spec defines the desired state of MonitorScheduledQueryRulesAlertV2
type MonitorScheduledQueryRulesAlertV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorScheduledQueryRulesAlertV2Parameters `json:"forProvider"`
}

// MonitorScheduledQueryRulesAlertV2Status defines the observed state of MonitorScheduledQueryRulesAlertV2.
type MonitorScheduledQueryRulesAlertV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorScheduledQueryRulesAlertV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesAlertV2 is the Schema for the MonitorScheduledQueryRulesAlertV2s API. Manages an AlertingAction Scheduled Query Rules Version 2 resource within Azure Monitor
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorScheduledQueryRulesAlertV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorScheduledQueryRulesAlertV2Spec   `json:"spec"`
	Status            MonitorScheduledQueryRulesAlertV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesAlertV2List contains a list of MonitorScheduledQueryRulesAlertV2s
type MonitorScheduledQueryRulesAlertV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorScheduledQueryRulesAlertV2 `json:"items"`
}

// Repository type metadata.
var (
	MonitorScheduledQueryRulesAlertV2_Kind             = "MonitorScheduledQueryRulesAlertV2"
	MonitorScheduledQueryRulesAlertV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorScheduledQueryRulesAlertV2_Kind}.String()
	MonitorScheduledQueryRulesAlertV2_KindAPIVersion   = MonitorScheduledQueryRulesAlertV2_Kind + "." + CRDGroupVersion.String()
	MonitorScheduledQueryRulesAlertV2_GroupVersionKind = CRDGroupVersion.WithKind(MonitorScheduledQueryRulesAlertV2_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorScheduledQueryRulesAlertV2{}, &MonitorScheduledQueryRulesAlertV2List{})
}
