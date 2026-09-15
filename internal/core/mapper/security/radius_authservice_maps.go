package security

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RadiusAuthserviceNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RadiusAuthserviceNIOSFieldMap = map[string]string{
	"Id":                    "Ref",
	"NIOS.AcctRetries":      "AcctRetries",
	"NIOS.AcctTimeout":      "AcctTimeout",
	"NIOS.AuthRetries":      "AuthRetries",
	"NIOS.AuthTimeout":      "AuthTimeout",
	"NIOS.CacheTtl":         "CacheTtl",
	"NIOS.Comment":          "Comment",
	"NIOS.Disable":          "Disable",
	"NIOS.EnableCache":      "EnableCache",
	"NIOS.Mode":             "Mode",
	"NIOS.Name":             "Name",
	"NIOS.RecoveryInterval": "RecoveryInterval",
	"NIOS.Servers":          "Servers",
}

// TODO: only searchable fields should be included here
// RadiusAuthserviceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RadiusAuthserviceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                     "_ref",
		"nios.acct_retries":      "acct_retries",
		"nios.acct_timeout":      "acct_timeout",
		"nios.auth_retries":      "auth_retries",
		"nios.auth_timeout":      "auth_timeout",
		"nios.cache_ttl":         "cache_ttl",
		"nios.comment":           "comment",
		"nios.disable":           "disable",
		"nios.enable_cache":      "enable_cache",
		"nios.mode":              "mode",
		"nios.name":              "name",
		"nios.recovery_interval": "recovery_interval",
		"nios.servers":           "servers",
	},
}
