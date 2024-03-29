// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package botservice

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures authorization group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_bot_channel_sms", func(r *config.Resource) {
		r.Path = "botchannelsms"
	})
	p.AddResourceConfigurator("azurerm_bot_channel_web_chat", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"site_names", "site"},
		}
		r.MetaResource.ArgumentDocs["site_names"] = "Deprecated: siteNames will be removed in favour of the site code block."
	})
}
