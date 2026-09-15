# Auto-generated datasource acceptance-test cases for RadiusAuthservice.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.acct_retries", "nios.acct_timeout", "nios.auth_retries", "nios.auth_timeout", "nios.cache_ttl", "nios.comment", "nios.disable", "nios.enable_cache", "nios.mode", "nios.name", "nios.recovery_interval"]

  step {
    nios {
      name = "{{random}}"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
  }

}
