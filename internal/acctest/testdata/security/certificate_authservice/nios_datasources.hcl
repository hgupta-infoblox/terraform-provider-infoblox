# Auto-generated datasource acceptance-test cases for CertificateAuthservice.
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - cacertificate : O="Infoblox",L="BLR",ST="KA",C="IN"
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.auto_populate_login", "nios.comment", "nios.disabled", "nios.enable_password_request", "nios.enable_remote_lookup", "nios.max_retries", "nios.name", "nios.ocsp_check", "nios.recovery_interval", "nios.remote_lookup_username", "nios.response_timeout", "nios.trust_model", "nios.user_match_type"]

  step {
    nios {
      name            = "{{random}}"
      ocsp_check      = "DISABLED"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

}
