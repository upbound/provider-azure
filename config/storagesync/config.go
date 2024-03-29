// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package storagesync

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures storagesync group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_storage_sync", func(r *config.Resource) {
		r.Kind = "StorageSync"
	})
}
