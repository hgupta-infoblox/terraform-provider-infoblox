package security

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// TacacsplusAuthserviceNIOSFieldMap maps infoblox model fields to NIOS struct fields
var TacacsplusAuthserviceNIOSFieldMap = map[string]string{
	"Id":               "Ref",
	"NIOS.AcctRetries": "AcctRetries",
	"NIOS.AcctTimeout": "AcctTimeout",
	"NIOS.AuthRetries": "AuthRetries",
	"NIOS.AuthTimeout": "AuthTimeout",
	"NIOS.Comment":     "Comment",
	"NIOS.Disable":     "Disable",
	"NIOS.Name":        "Name",
	"NIOS.Servers":     "Servers",
}

// TODO: only searchable fields should be included here
// TacacsplusAuthserviceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var TacacsplusAuthserviceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                "_ref",
		"nios.acct_retries": "acct_retries",
		"nios.acct_timeout": "acct_timeout",
		"nios.auth_retries": "auth_retries",
		"nios.auth_timeout": "auth_timeout",
		"nios.comment":      "comment",
		"nios.disable":      "disable",
		"nios.name":         "name",
		"nios.servers":      "servers",
	},
}
