# CertificateAuthservice — nios list cases
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - cacertificate : O="Infoblox",L="BLR",ST="KA",C="IN"
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}
