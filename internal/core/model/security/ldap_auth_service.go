package security

import (
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
)

// Infoblox LdapAuthService model
type LdapAuthService struct {
	Id   *string
	NIOS *NIOSLdapAuthServiceExt
}

// NIOSLdapAuthServiceExt - NIOS specific fields for LdapAuthService
type NIOSLdapAuthServiceExt struct {
	Comment                     *string
	Disable                     *bool
	EaMapping                   []niossecurity.LdapAuthServiceEaMapping
	LdapGroupAttribute          *string
	LdapGroupAuthenticationType *string
	LdapUserAttribute           *string
	Mode                        *string
	Name                        *string
	RecoveryInterval            *int64
	Retries                     *int64
	SearchScope                 *string
	Servers                     []niossecurity.LdapAuthServiceServers
	Timeout                     *int64
}
