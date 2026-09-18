package smartfolder

import (
	niossmartfolder "github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
)

// Infoblox SmartfolderPersonal model
type SmartfolderPersonal struct {
	Id   *string
	NIOS *NIOSSmartfolderPersonalExt
}

// NIOSSmartfolderPersonalExt - NIOS specific fields for SmartfolderPersonal
type NIOSSmartfolderPersonalExt struct {
	Comment    *string
	GroupBys   []niossmartfolder.SmartfolderPersonalGroupBys
	Name       *string
	QueryItems []niossmartfolder.SmartfolderPersonalQueryItems
}
