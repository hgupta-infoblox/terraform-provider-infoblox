package security

import (
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
)

// Infoblox TacacsplusAuthservice model
type TacacsplusAuthservice struct {
	Id   *string
	NIOS *NIOSTacacsplusAuthserviceExt
}

// NIOSTacacsplusAuthserviceExt - NIOS specific fields for TacacsplusAuthservice
type NIOSTacacsplusAuthserviceExt struct {
	AcctRetries *int64
	AcctTimeout *int64
	AuthRetries *int64
	AuthTimeout *int64
	Comment     *string
	Disable     *bool
	Name        *string
	Servers     []niossecurity.TacacsplusAuthserviceServers
}
