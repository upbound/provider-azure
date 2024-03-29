// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package automation

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures automation group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_automation_account", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"encryption"},
		}
	})
}
