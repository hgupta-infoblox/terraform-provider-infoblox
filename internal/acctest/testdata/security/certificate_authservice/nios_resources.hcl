# Auto-generated resource acceptance-test cases for CertificateAuthservice.
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - cacertificate : O="Infoblox",L="BLR",ST="KA",C="IN"
#   - cacertificate : CN="dummy.example.com",O="TestOrg",C="US"
case "basic" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.name"                    = "{{random}}"
      "nios.ca_certificates.#"       = "1"
      "nios.ca_certificates.0"       = "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"
      "nios.disabled"                = "false"
      "nios.enable_password_request" = "true"
      "nios.enable_remote_lookup"    = "false"
      "nios.auto_populate_login"     = "S_DN_CN"
      "nios.max_retries"             = "0"
      "nios.recovery_interval"       = "30"
      "nios.response_timeout"        = "1000"
      "nios.trust_model"             = "DIRECT"
      "nios.user_match_type"         = "AUTO_MATCH"
    }
  }

}

case "disappears" {
  backend               = "nios"
  skip                  = true
  skip_reason           = "t.Skip: skipping "
  disappears            = true
  expect_non_empty_plan = true

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

}

case "auto_populate_login" {
  backend = "nios"

  step {
    nios {
      ocsp_check          = "DISABLED"
      name                = "{{random}}"
      ca_certificates     = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      auto_populate_login = "SAN_EMAIL"
    }
    check = {
      "nios.auto_populate_login" = "SAN_EMAIL"
    }
  }

  step {
    nios {
      ocsp_check          = "DISABLED"
      name                = "{{random}}"
      ca_certificates     = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      auto_populate_login = "SERIAL_NUMBER"
    }
    check = {
      "nios.auto_populate_login" = "SERIAL_NUMBER"
    }
  }

}

case "ca_certificates" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.ca_certificates.0" = "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"
    }
  }

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYzU0ODUyZmRkNjk3ODVhMTI3OWEwZGUxYmE0OTA5ZDFhMGY1ODFkMGYxZTVmMzhhYjk5OGJkZWYxNWVmM2Y0NGUyN2UzYTJmMWM3OTg5NjIyZjY2NzVmZTI2ZWNkMWM4ZmJlNDJkNzhmYzY1YTU2MWRmYWUyN2Q3MzJmNzdmZmI:CN%3D%22dummy.example.com%22%2CO%3D%22TestOrg%22%2CC%3D%22US%22"]
    }
    check = {
      "nios.ca_certificates.0" = "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYzU0ODUyZmRkNjk3ODVhMTI3OWEwZGUxYmE0OTA5ZDFhMGY1ODFkMGYxZTVmMzhhYjk5OGJkZWYxNWVmM2Y0NGUyN2UzYTJmMWM3OTg5NjIyZjY2NzVmZTI2ZWNkMWM4ZmJlNDJkNzhmYzY1YTU2MWRmYWUyN2Q3MzJmNzdmZmI:CN%3D%22dummy.example.com%22%2CO%3D%22TestOrg%22%2CC%3D%22US%22"
    }
  }

}

case "comment" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      comment         = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      comment         = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "disabled" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      disabled        = true
    }
    check = {
      "nios.disabled" = "true"
    }
  }

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      disabled        = false
    }
    check = {
      "nios.disabled" = "false"
    }
  }

}

case "enable_password_request" {
  backend = "nios"

  step {
    nios {
      ocsp_check              = "DISABLED"
      enable_password_request = true
      name                    = "{{random}}"
      ca_certificates         = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.enable_password_request" = "true"
    }
  }

  step {
    nios {
      ocsp_check              = "DISABLED"
      enable_password_request = false
      name                    = "{{random}}"
      ca_certificates         = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.enable_password_request" = "false"
    }
  }

}

case "enable_remote_lookup" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS"

  step {
    nios {
      ocsp_check           = "DISABLED"
      enable_remote_lookup = false
      name                 = "{{random}}"
      ca_certificates      = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.enable_remote_lookup" = "false"
    }
  }

  step {
    nios {
      ocsp_check              = "DISABLED"
      enable_remote_lookup    = true
      name                    = "{{random}}"
      ca_certificates         = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      remote_lookup_username  = "admin"
      remote_lookup_password  = "infoblox"
      enable_password_request = false
    }
    check = {
      "nios.enable_remote_lookup" = "true"
    }
  }

}

case "max_retries" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      max_retries     = 4
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.max_retries" = "4"
    }
  }

  step {
    nios {
      ocsp_check      = "DISABLED"
      max_retries     = 5
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.max_retries" = "5"
    }
  }

}

case "name" {
  backend = "nios"

  step {
    nios {
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      ocsp_check      = "DISABLED"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name            = "{{random2}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      ocsp_check      = "DISABLED"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "ocsp_check" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.ocsp_check" = "DISABLED"
    }
  }

  step {
    nios {
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      ocsp_check      = "AIA_ONLY"
    }
    check = {
      "nios.ocsp_check" = "AIA_ONLY"
    }
  }

}

case "ocsp_responders" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Requires a certificate_token from a prior fileop upload; certificate upload via token is not yet automated in acceptance tests."

  step {
    nios {
      name            = "{{random}}"
      ocsp_responders = [{ fqdn_or_ip = "3.3.3.3", certificate_token = "placeholder_token_1" }, { fqdn_or_ip = "3.3.32.3", certificate_token = "placeholder_token_2" }]
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.ocsp_responders.0.fqdn_or_ip" = "3.3.3.3"
      "nios.ocsp_responders.#"            = "2"
      "nios.ocsp_responders.0.disabled"   = "false"
      "nios.ocsp_responders.0.port"       = "80"
    }
  }

  step {
    nios {
      name            = "{{random}}"
      ocsp_responders = [{ fqdn_or_ip = "3.3.32.3", certificate_token = "placeholder_token_2" }]
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.ocsp_responders.0.fqdn_or_ip" = "3.3.32.3"
      "nios.ocsp_responders.0.disabled"   = "false"
      "nios.ocsp_responders.#"            = "1"
      "nios.ocsp_responders.0.port"       = "80"
    }
  }

}

case "recovery_interval" {
  backend = "nios"

  step {
    nios {
      ocsp_check        = "DISABLED"
      name              = "{{random}}"
      ca_certificates   = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      recovery_interval = 3
    }
    check = {
      "nios.recovery_interval" = "3"
    }
  }

  step {
    nios {
      ocsp_check        = "DISABLED"
      name              = "{{random}}"
      ca_certificates   = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      recovery_interval = 5
    }
    check = {
      "nios.recovery_interval" = "5"
    }
  }

}

case "remote_lookup_service" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: remote_lookup_service is a oneOf SDK union type; skip_expand/skip_flatten prevent full round-trip through the API. Field is preserved in state via PostFlattenNIOS but cannot be functionally tested without a live remote lookup service on the grid."

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

}

case "remote_lookup_username" {
  backend = "nios"

  step {
    nios {
      ocsp_check             = "DISABLED"
      name                   = "{{random}}"
      remote_lookup_username = "username1"
      ca_certificates        = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.remote_lookup_username" = "username1"
    }
  }

  step {
    nios {
      ocsp_check             = "DISABLED"
      name                   = "{{random}}"
      remote_lookup_username = "username2"
      ca_certificates        = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.remote_lookup_username" = "username2"
    }
  }

}

case "remote_lookup_user_password" {
  backend = "nios"

  step {
    nios {
      ocsp_check             = "DISABLED"
      name                   = "{{random}}"
      remote_lookup_username = "username1"
      remote_lookup_password = "password1"
      ca_certificates        = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

  step {
    nios {
      ocsp_check             = "DISABLED"
      name                   = "{{random}}"
      remote_lookup_username = "username1"
      remote_lookup_password = "password2"
      ca_certificates        = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
  }

}

case "response_timeout" {
  backend = "nios"

  step {
    nios {
      ocsp_check       = "DISABLED"
      response_timeout = 3000
      name             = "{{random}}"
      ca_certificates  = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.response_timeout" = "3000"
    }
  }

  step {
    nios {
      ocsp_check       = "DISABLED"
      response_timeout = 5000
      name             = "{{random}}"
      ca_certificates  = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
    }
    check = {
      "nios.response_timeout" = "5000"
    }
  }

}

case "trust_model" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      trust_model     = "DELEGATED"
    }
    check = {
      "nios.trust_model" = "DELEGATED"
    }
  }

  step {
    nios {
      ocsp_check      = "DISABLED"
      name            = "{{random}}"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      trust_model     = "DIRECT"
    }
    check = {
      "nios.trust_model" = "DIRECT"
    }
  }

}

case "user_match_type" {
  backend = "nios"

  step {
    nios {
      ocsp_check      = "DISABLED"
      user_match_type = "DIRECT_MATCH"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      name            = "{{random}}"
    }
    check = {
      "nios.user_match_type" = "DIRECT_MATCH"
    }
  }

  step {
    nios {
      ocsp_check      = "DISABLED"
      user_match_type = "AUTO_MATCH"
      ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuYmQzMTM3NzEwZDQ4ZDRhZWI3ZWI2YTE0M2RiOGEyZThkZDQ2MTVkZmFkOTBiZWQ1MDdiYmY3ZWM4MGUyMDRhMTM5NjYzYzlkODdhOGQyNjY5OTY4ODhmMDc3NjViNmVkNDllNmNmYjg3NTJiOWVhYWQ5ZjkyZTRmZjkzMzMwMWI:O%3D%22Infoblox%22%2CL%3D%22BLR%22%2CST%3D%22KA%22%2CC%3D%22IN%22"]
      name            = "{{random}}"
    }
    check = {
      "nios.user_match_type" = "AUTO_MATCH"
    }
  }

}
