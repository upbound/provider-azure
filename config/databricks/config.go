// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package databricks

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures databricks group
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("azurerm_databricks_workspace_customer_managed_key", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			TerraformName: "azurerm_databricks_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["key_vault_key_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_key",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.MetaResource.Description = "Manages a Customer Managed Key for a Databricks Workspace root DBFS. This resource has been deprecated and will be removed in future versions of provider. Please use the WorkspaceRootDbfsCustomerManagedKey resource instead."
	})
}
