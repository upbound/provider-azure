/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this SentinelAutomationRule
func (mg *SentinelAutomationRule) GetTerraformResourceType() string {
	return "azurerm_sentinel_automation_rule"
}

// GetConnectionDetailsMapping for this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAutomationRule
func (tr *SentinelAutomationRule) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAutomationRule
func (tr *SentinelAutomationRule) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelAutomationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAutomationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAutomationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAutomationRule) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this SentinelWatchlist
func (mg *SentinelWatchlist) GetTerraformResourceType() string {
	return "azurerm_sentinel_watchlist"
}

// GetConnectionDetailsMapping for this SentinelWatchlist
func (tr *SentinelWatchlist) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelWatchlist
func (tr *SentinelWatchlist) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelWatchlist
func (tr *SentinelWatchlist) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelWatchlist
func (tr *SentinelWatchlist) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelWatchlist
func (tr *SentinelWatchlist) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelWatchlist
func (tr *SentinelWatchlist) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SentinelWatchlist using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelWatchlist) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelWatchlistParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelWatchlist) GetTerraformSchemaVersion() int {
	return 0
}
