package security

import (
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
)

// Infoblox CertificateAuthservice model
type CertificateAuthservice struct {
	Id   *string
	NIOS *NIOSCertificateAuthserviceExt
}

// NIOSCertificateAuthserviceExt - NIOS specific fields for CertificateAuthservice
type NIOSCertificateAuthserviceExt struct {
	AutoPopulateLogin     *string
	CaCertificates        []string
	Comment               *string
	Disabled              *bool
	EnablePasswordRequest *bool
	EnableRemoteLookup    *bool
	MaxRetries            *int64
	Name                  *string
	OcspCheck             *string
	OcspResponders        []niossecurity.CertificateAuthserviceOcspResponders
	RecoveryInterval      *int64
	RemoteLookupPassword  *string
	RemoteLookupService   *niossecurity.CertificateAuthserviceRemoteLookupService
	RemoteLookupUsername  *string
	ResponseTimeout       *int64
	TrustModel            *string
	UserMatchType         *string
}
