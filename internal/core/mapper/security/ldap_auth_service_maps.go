package security

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// LdapAuthServiceNIOSFieldMap maps infoblox model fields to NIOS struct fields
var LdapAuthServiceNIOSFieldMap = map[string]string{
	"Id":                               "Ref",
	"NIOS.Comment":                     "Comment",
	"NIOS.Disable":                     "Disable",
	"NIOS.EaMapping":                   "EaMapping",
	"NIOS.LdapGroupAttribute":          "LdapGroupAttribute",
	"NIOS.LdapGroupAuthenticationType": "LdapGroupAuthenticationType",
	"NIOS.LdapUserAttribute":           "LdapUserAttribute",
	"NIOS.Mode":                        "Mode",
	"NIOS.Name":                        "Name",
	"NIOS.RecoveryInterval":            "RecoveryInterval",
	"NIOS.Retries":                     "Retries",
	"NIOS.SearchScope":                 "SearchScope",
	"NIOS.Servers":                     "Servers",
	"NIOS.Timeout":                     "Timeout",
}

// TODO: only searchable fields should be included here
// LdapAuthServiceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var LdapAuthServiceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                                  "_ref",
		"nios.comment":                        "comment",
		"nios.disable":                        "disable",
		"nios.ea_mapping":                     "ea_mapping",
		"nios.ldap_group_attribute":           "ldap_group_attribute",
		"nios.ldap_group_authentication_type": "ldap_group_authentication_type",
		"nios.ldap_user_attribute":            "ldap_user_attribute",
		"nios.mode":                           "mode",
		"nios.name":                           "name",
		"nios.recovery_interval":              "recovery_interval",
		"nios.retries":                        "retries",
		"nios.search_scope":                   "search_scope",
		"nios.servers":                        "servers",
		"nios.timeout":                        "timeout",
	},
}
