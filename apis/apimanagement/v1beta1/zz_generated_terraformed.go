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

// GetTerraformResourceType returns Terraform resource type for this Management
func (mg *Management) GetTerraformResourceType() string {
	return "azurerm_api_management"
}

// GetConnectionDetailsMapping for this Management
func (tr *Management) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"certificate[*].certificate_password": "spec.forProvider.certificate[*].certificatePasswordSecretRef", "certificate[*].encoded_certificate": "spec.forProvider.certificate[*].encodedCertificateSecretRef", "hostname_configuration[*].developer_portal[*].certificate": "status.atProvider.hostnameConfiguration[*].developerPortal[*].certificate", "hostname_configuration[*].developer_portal[*].certificate_password": "status.atProvider.hostnameConfiguration[*].developerPortal[*].certificatePassword", "hostname_configuration[*].management[*].certificate": "status.atProvider.hostnameConfiguration[*].management[*].certificate", "hostname_configuration[*].management[*].certificate_password": "status.atProvider.hostnameConfiguration[*].management[*].certificatePassword", "hostname_configuration[*].portal[*].certificate": "status.atProvider.hostnameConfiguration[*].portal[*].certificate", "hostname_configuration[*].portal[*].certificate_password": "status.atProvider.hostnameConfiguration[*].portal[*].certificatePassword", "hostname_configuration[*].proxy[*].certificate": "status.atProvider.hostnameConfiguration[*].proxy[*].certificate", "hostname_configuration[*].proxy[*].certificate_password": "status.atProvider.hostnameConfiguration[*].proxy[*].certificatePassword", "hostname_configuration[*].scm[*].certificate": "status.atProvider.hostnameConfiguration[*].scm[*].certificate", "hostname_configuration[*].scm[*].certificate_password": "status.atProvider.hostnameConfiguration[*].scm[*].certificatePassword", "tenant_access[*].primary_key": "status.atProvider.tenantAccess[*].primaryKey", "tenant_access[*].secondary_key": "status.atProvider.tenantAccess[*].secondaryKey"}
}

// GetObservation of this Management
func (tr *Management) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Management
func (tr *Management) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Management
func (tr *Management) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Management
func (tr *Management) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Management
func (tr *Management) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Management using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Management) LateInitialize(attrs []byte) (bool, error) {
	params := &ManagementParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Management) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this API
func (mg *API) GetTerraformResourceType() string {
	return "azurerm_api_management_api"
}

// GetConnectionDetailsMapping for this API
func (tr *API) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this API
func (tr *API) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this API
func (tr *API) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this API
func (tr *API) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this API
func (tr *API) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this API
func (tr *API) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this API using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *API) LateInitialize(attrs []byte) (bool, error) {
	params := &APIParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *API) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this APIOperation
func (mg *APIOperation) GetTerraformResourceType() string {
	return "azurerm_api_management_api_operation"
}

// GetConnectionDetailsMapping for this APIOperation
func (tr *APIOperation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this APIOperation
func (tr *APIOperation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this APIOperation
func (tr *APIOperation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this APIOperation
func (tr *APIOperation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this APIOperation
func (tr *APIOperation) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this APIOperation
func (tr *APIOperation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this APIOperation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *APIOperation) LateInitialize(attrs []byte) (bool, error) {
	params := &APIOperationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *APIOperation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this APIOperationPolicy
func (mg *APIOperationPolicy) GetTerraformResourceType() string {
	return "azurerm_api_management_api_operation_policy"
}

// GetConnectionDetailsMapping for this APIOperationPolicy
func (tr *APIOperationPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this APIOperationPolicy
func (tr *APIOperationPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this APIOperationPolicy
func (tr *APIOperationPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this APIOperationPolicy
func (tr *APIOperationPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this APIOperationPolicy
func (tr *APIOperationPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this APIOperationPolicy
func (tr *APIOperationPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this APIOperationPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *APIOperationPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &APIOperationPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *APIOperationPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this APIOperationTag
func (mg *APIOperationTag) GetTerraformResourceType() string {
	return "azurerm_api_management_api_operation_tag"
}

// GetConnectionDetailsMapping for this APIOperationTag
func (tr *APIOperationTag) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this APIOperationTag
func (tr *APIOperationTag) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this APIOperationTag
func (tr *APIOperationTag) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this APIOperationTag
func (tr *APIOperationTag) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this APIOperationTag
func (tr *APIOperationTag) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this APIOperationTag
func (tr *APIOperationTag) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this APIOperationTag using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *APIOperationTag) LateInitialize(attrs []byte) (bool, error) {
	params := &APIOperationTagParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *APIOperationTag) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this APIPolicy
func (mg *APIPolicy) GetTerraformResourceType() string {
	return "azurerm_api_management_api_policy"
}

// GetConnectionDetailsMapping for this APIPolicy
func (tr *APIPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this APIPolicy
func (tr *APIPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this APIPolicy
func (tr *APIPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this APIPolicy
func (tr *APIPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this APIPolicy
func (tr *APIPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this APIPolicy
func (tr *APIPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this APIPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *APIPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &APIPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *APIPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this APIRelease
func (mg *APIRelease) GetTerraformResourceType() string {
	return "azurerm_api_management_api_release"
}

// GetConnectionDetailsMapping for this APIRelease
func (tr *APIRelease) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this APIRelease
func (tr *APIRelease) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this APIRelease
func (tr *APIRelease) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this APIRelease
func (tr *APIRelease) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this APIRelease
func (tr *APIRelease) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this APIRelease
func (tr *APIRelease) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this APIRelease using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *APIRelease) LateInitialize(attrs []byte) (bool, error) {
	params := &APIReleaseParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *APIRelease) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this APISchema
func (mg *APISchema) GetTerraformResourceType() string {
	return "azurerm_api_management_api_schema"
}

// GetConnectionDetailsMapping for this APISchema
func (tr *APISchema) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this APISchema
func (tr *APISchema) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this APISchema
func (tr *APISchema) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this APISchema
func (tr *APISchema) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this APISchema
func (tr *APISchema) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this APISchema
func (tr *APISchema) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this APISchema using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *APISchema) LateInitialize(attrs []byte) (bool, error) {
	params := &APISchemaParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *APISchema) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this APIVersionSet
func (mg *APIVersionSet) GetTerraformResourceType() string {
	return "azurerm_api_management_api_version_set"
}

// GetConnectionDetailsMapping for this APIVersionSet
func (tr *APIVersionSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this APIVersionSet
func (tr *APIVersionSet) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this APIVersionSet
func (tr *APIVersionSet) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this APIVersionSet
func (tr *APIVersionSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this APIVersionSet
func (tr *APIVersionSet) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this APIVersionSet
func (tr *APIVersionSet) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this APIVersionSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *APIVersionSet) LateInitialize(attrs []byte) (bool, error) {
	params := &APIVersionSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *APIVersionSet) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this AuthorizationServer
func (mg *AuthorizationServer) GetTerraformResourceType() string {
	return "azurerm_api_management_authorization_server"
}

// GetConnectionDetailsMapping for this AuthorizationServer
func (tr *AuthorizationServer) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_secret": "spec.forProvider.clientSecretSecretRef", "resource_owner_password": "spec.forProvider.resourceOwnerPasswordSecretRef"}
}

// GetObservation of this AuthorizationServer
func (tr *AuthorizationServer) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AuthorizationServer
func (tr *AuthorizationServer) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AuthorizationServer
func (tr *AuthorizationServer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AuthorizationServer
func (tr *AuthorizationServer) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AuthorizationServer
func (tr *AuthorizationServer) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AuthorizationServer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AuthorizationServer) LateInitialize(attrs []byte) (bool, error) {
	params := &AuthorizationServerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AuthorizationServer) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Backend
func (mg *Backend) GetTerraformResourceType() string {
	return "azurerm_api_management_backend"
}

// GetConnectionDetailsMapping for this Backend
func (tr *Backend) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"proxy[*].password": "spec.forProvider.proxy[*].passwordSecretRef"}
}

// GetObservation of this Backend
func (tr *Backend) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Backend
func (tr *Backend) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Backend
func (tr *Backend) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Backend
func (tr *Backend) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Backend
func (tr *Backend) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Backend using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Backend) LateInitialize(attrs []byte) (bool, error) {
	params := &BackendParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Backend) GetTerraformSchemaVersion() int {
	return 0
}
