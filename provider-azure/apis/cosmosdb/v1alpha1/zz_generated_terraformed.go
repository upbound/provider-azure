/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this CassandraCluster
func (mg *CassandraCluster) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_cassandra_cluster"
}

// GetConnectionDetailsMapping for this CassandraCluster
func (tr *CassandraCluster) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"default_admin_password": "spec.forProvider.defaultAdminPasswordSecretRef"}
}

// GetObservation of this CassandraCluster
func (tr *CassandraCluster) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CassandraCluster
func (tr *CassandraCluster) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CassandraCluster
func (tr *CassandraCluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CassandraCluster
func (tr *CassandraCluster) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CassandraCluster
func (tr *CassandraCluster) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CassandraCluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CassandraCluster) LateInitialize(attrs []byte) (bool, error) {
	params := &CassandraClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CassandraCluster) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CassandraDatacenter
func (mg *CassandraDatacenter) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_cassandra_datacenter"
}

// GetConnectionDetailsMapping for this CassandraDatacenter
func (tr *CassandraDatacenter) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CassandraDatacenter
func (tr *CassandraDatacenter) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CassandraDatacenter
func (tr *CassandraDatacenter) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CassandraDatacenter
func (tr *CassandraDatacenter) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CassandraDatacenter
func (tr *CassandraDatacenter) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CassandraDatacenter
func (tr *CassandraDatacenter) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CassandraDatacenter using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CassandraDatacenter) LateInitialize(attrs []byte) (bool, error) {
	params := &CassandraDatacenterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CassandraDatacenter) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SQLRoleAssignment
func (mg *SQLRoleAssignment) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_sql_role_assignment"
}

// GetConnectionDetailsMapping for this SQLRoleAssignment
func (tr *SQLRoleAssignment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SQLRoleAssignment
func (tr *SQLRoleAssignment) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SQLRoleAssignment
func (tr *SQLRoleAssignment) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SQLRoleAssignment
func (tr *SQLRoleAssignment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SQLRoleAssignment
func (tr *SQLRoleAssignment) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SQLRoleAssignment
func (tr *SQLRoleAssignment) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SQLRoleAssignment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SQLRoleAssignment) LateInitialize(attrs []byte) (bool, error) {
	params := &SQLRoleAssignmentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SQLRoleAssignment) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SQLRoleDefinition
func (mg *SQLRoleDefinition) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_sql_role_definition"
}

// GetConnectionDetailsMapping for this SQLRoleDefinition
func (tr *SQLRoleDefinition) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SQLRoleDefinition
func (tr *SQLRoleDefinition) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SQLRoleDefinition
func (tr *SQLRoleDefinition) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SQLRoleDefinition
func (tr *SQLRoleDefinition) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SQLRoleDefinition
func (tr *SQLRoleDefinition) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SQLRoleDefinition
func (tr *SQLRoleDefinition) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SQLRoleDefinition using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SQLRoleDefinition) LateInitialize(attrs []byte) (bool, error) {
	params := &SQLRoleDefinitionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SQLRoleDefinition) GetTerraformSchemaVersion() int {
	return 0
}
