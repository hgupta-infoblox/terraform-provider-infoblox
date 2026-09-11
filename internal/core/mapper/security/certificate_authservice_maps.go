package security

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// CertificateAuthserviceNIOSFieldMap maps infoblox model fields to NIOS struct fields
var CertificateAuthserviceNIOSFieldMap = map[string]string{
	"Id":                         "Ref",
	"NIOS.AutoPopulateLogin":     "AutoPopulateLogin",
	"NIOS.CaCertificates":        "CaCertificates",
	"NIOS.Comment":               "Comment",
	"NIOS.Disabled":              "Disabled",
	"NIOS.EnablePasswordRequest": "EnablePasswordRequest",
	"NIOS.EnableRemoteLookup":    "EnableRemoteLookup",
	"NIOS.MaxRetries":            "MaxRetries",
	"NIOS.Name":                  "Name",
	"NIOS.OcspCheck":             "OcspCheck",
	"NIOS.OcspResponders":        "OcspResponders",
	"NIOS.RecoveryInterval":      "RecoveryInterval",
	"NIOS.RemoteLookupPassword":  "RemoteLookupPassword",
	"NIOS.RemoteLookupService":   "RemoteLookupService",
	"NIOS.RemoteLookupUsername":  "RemoteLookupUsername",
	"NIOS.ResponseTimeout":       "ResponseTimeout",
	"NIOS.TrustModel":            "TrustModel",
	"NIOS.UserMatchType":         "UserMatchType",
}

// TODO: only searchable fields should be included here
// CertificateAuthserviceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var CertificateAuthserviceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                           "_ref",
		"nios.auto_populate_login":     "auto_populate_login",
		"nios.ca_certificates":         "ca_certificates",
		"nios.comment":                 "comment",
		"nios.disabled":                "disabled",
		"nios.enable_password_request": "enable_password_request",
		"nios.enable_remote_lookup":    "enable_remote_lookup",
		"nios.max_retries":             "max_retries",
		"nios.name":                    "name",
		"nios.ocsp_check":              "ocsp_check",
		"nios.ocsp_responders":         "ocsp_responders",
		"nios.recovery_interval":       "recovery_interval",
		"nios.remote_lookup_password":  "remote_lookup_password",
		"nios.remote_lookup_service":   "remote_lookup_service",
		"nios.remote_lookup_username":  "remote_lookup_username",
		"nios.response_timeout":        "response_timeout",
		"nios.trust_model":             "trust_model",
		"nios.user_match_type":         "user_match_type",
	},
}
