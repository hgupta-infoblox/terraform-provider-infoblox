package security

import (
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
)

// Infoblox RadiusAuthservice model
type RadiusAuthservice struct {
	Id   *string
	NIOS *NIOSRadiusAuthserviceExt
}

// NIOSRadiusAuthserviceExt - NIOS specific fields for RadiusAuthservice
type NIOSRadiusAuthserviceExt struct {
	AcctRetries      *int64
	AcctTimeout      *int64
	AuthRetries      *int64
	AuthTimeout      *int64
	CacheTtl         *int64
	Comment          *string
	Disable          *bool
	EnableCache      *bool
	Mode             *string
	Name             *string
	RecoveryInterval *int64
	Servers          []niossecurity.RadiusAuthserviceServers
}
