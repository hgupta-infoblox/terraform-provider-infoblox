package smartfolder

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SmartfolderPersonalNIOSFieldMap maps infoblox model fields to NIOS struct fields
var SmartfolderPersonalNIOSFieldMap = map[string]string{
	"Id":              "Ref",
	"NIOS.Comment":    "Comment",
	"NIOS.GroupBys":   "GroupBys",
	"NIOS.Name":       "Name",
	"NIOS.QueryItems": "QueryItems",
}

// TODO: only searchable fields should be included here
// SmartfolderPersonalFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SmartfolderPersonalFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":               "_ref",
		"nios.comment":     "comment",
		"nios.group_bys":   "group_bys",
		"nios.name":        "name",
		"nios.query_items": "query_items",
	},
}
