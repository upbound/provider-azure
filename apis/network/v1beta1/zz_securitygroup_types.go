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

type SecurityGroupObservation struct {

	// The ID of the Network Security Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityGroupParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// List of objects representing security rules, as defined below.
	// +kubebuilder:validation:Optional
	SecurityRule []SecurityRuleParameters `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityRuleObservation struct {
}

type SecurityRuleParameters struct {

	// Specifies whether network traffic is allowed or denied. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	Access *string `json:"access,omitempty" tf:"access"`

	// A description for this rule. Restricted to 140 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if destination_address_prefixes is not specified.
	// +kubebuilder:validation:Optional
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix"`

	// List of destination address prefixes. Tags may not be used. This is required if destination_address_prefix is not specified.
	// +kubebuilder:validation:Optional
	DestinationAddressPrefixes []*string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes"`

	// A List of destination Application Security Group IDs
	// +kubebuilder:validation:Optional
	DestinationApplicationSecurityGroupIds []*string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids"`

	// Destination Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if destination_port_ranges is not specified.
	// +kubebuilder:validation:Optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range"`

	// List of destination ports or port ranges. This is required if destination_port_range is not specified.
	// +kubebuilder:validation:Optional
	DestinationPortRanges []*string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges"`

	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are Inbound and Outbound.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction"`

	// The name of the security rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority"`

	// Network protocol this rule applies to. Possible values include Tcp, Udp, Icmp, Esp, Ah or * (which matches all).
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if source_address_prefixes is not specified.
	// +kubebuilder:validation:Optional
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix"`

	// List of source address prefixes. Tags may not be used. This is required if source_address_prefix is not specified.
	// +kubebuilder:validation:Optional
	SourceAddressPrefixes []*string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes"`

	// A List of source Application Security Group IDs
	// +kubebuilder:validation:Optional
	SourceApplicationSecurityGroupIds []*string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids"`

	// Source Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if source_port_ranges is not specified.
	// +kubebuilder:validation:Optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range"`

	// List of source ports or port ranges. This is required if source_port_range is not specified.
	// +kubebuilder:validation:Optional
	SourcePortRanges []*string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API. Manages a network security group that contains a list of network security rules. Network security groups enable inbound or outbound traffic to be enabled or denied.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
