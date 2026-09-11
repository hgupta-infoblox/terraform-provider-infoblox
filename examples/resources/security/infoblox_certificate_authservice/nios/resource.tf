// Create Certificate Authservice with Basic Fields
resource "infoblox_certificate_authservice" "certificate_authservice_basic" {
  nios = {
    name            = "example_certificate_authservice"
    ca_certificates = ["cacertificate/<ca_certificate_ref>"]
    ocsp_check      = "DISABLED"
  }
}

// Create Certificate Authservice with OCSP check enabled via AIA
resource "infoblox_certificate_authservice" "certificate_authservice_aia" {
  nios = {
    name             = "example_certificate_authservice_aia"
    ca_certificates  = ["cacertificate/<ca_certificate_ref>"]
    ocsp_check       = "AIA_ONLY"
    comment          = "Certificate authservice using AIA for OCSP"
    trust_model      = "DELEGATED"
    response_timeout = 2000
    max_retries      = 2
  }
}

// Create Certificate Authservice with OCSP responders
// Note: certificate_token is obtained by uploading a certificate via the NIOS fileop API prior to applying.
resource "infoblox_certificate_authservice" "certificate_authservice_manual_ocsp" {
  nios = {
    name            = "example_certificate_authservice_manual"
    ca_certificates = ["cacertificate/<ca_certificate_ref>"]
    ocsp_check      = "MANUAL"
    ocsp_responders = [
      {
        fqdn_or_ip        = "ocsp.example.com"
        port              = 443
        certificate_token = "<token_from_fileop_uploadinit>"
      }
    ]
    recovery_interval = 20
    response_timeout  = 2000
    max_retries       = 2
    disabled          = false
  }
}
