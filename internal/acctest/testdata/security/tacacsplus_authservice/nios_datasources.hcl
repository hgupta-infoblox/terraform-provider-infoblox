# Auto-generated datasource acceptance-test cases for TacacsplusAuthservice.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.acct_retries", "nios.acct_timeout", "nios.auth_retries", "nios.auth_timeout", "nios.comment", "nios.disable", "nios.name"]

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
  }

}
