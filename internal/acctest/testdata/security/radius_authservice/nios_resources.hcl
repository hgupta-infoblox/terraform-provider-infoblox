# Auto-generated resource acceptance-test cases for RadiusAuthservice.
case "basic" {
  backend  = "nios"
  parallel = true

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
    check = {
      "nios.name"              = "{{random}}"
      "nios.acct_retries"      = "1000"
      "nios.acct_timeout"      = "5000"
      "nios.auth_retries"      = "6"
      "nios.auth_timeout"      = "5000"
      "nios.cache_ttl"         = "3600"
      "nios.disable"           = "false"
      "nios.enable_cache"      = "false"
      "nios.mode"              = "HUNT_GROUP"
      "nios.recovery_interval" = "30"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

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

case "acct_retries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      acct_retries = 20
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.acct_retries" = "20"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      acct_retries = 30
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.acct_retries" = "30"
    }
  }

}

case "acct_timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      acct_timeout = 3600
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.acct_timeout" = "3600"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      acct_timeout = 7200
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.acct_timeout" = "7200"
    }
  }

}

case "auth_retries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      auth_retries = 10
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.auth_retries" = "10"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      auth_retries = 7
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.auth_retries" = "7"
    }
  }

}

case "auth_timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      auth_timeout = 4000
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.auth_timeout" = "4000"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      auth_timeout = 4500
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.auth_timeout" = "4500"
    }
  }

}

case "cache_ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      cache_ttl = 4000
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.cache_ttl" = "4000"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      cache_ttl = 4500
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.cache_ttl" = "4500"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      disable = true
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      disable = false
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "enable_cache" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      enable_cache = true
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.enable_cache" = "true"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      enable_cache = false
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.enable_cache" = "false"
    }
  }

}

case "mode" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      mode = "ROUND_ROBIN"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.mode" = "ROUND_ROBIN"
    }
  }

  step {
    nios {
      name = "{{random}}"
      mode = "HUNT_GROUP"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.mode" = "HUNT_GROUP"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

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
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "recovery_interval" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name              = "{{random}}"
      recovery_interval = 45
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.recovery_interval" = "45"
    }
  }

  step {
    nios {
      name              = "{{random}}"
      recovery_interval = 60
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
    check = {
      "nios.recovery_interval" = "60"
    }
  }

}

case "servers" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
          auth_type     = "PAP"
        }
      ]
    }
    check = {
      "nios.servers.#"                = "1"
      "nios.servers.0.address"        = "2.2.3.1"
      "nios.servers.0.auth_port"      = "1812"
      "nios.servers.0.acct_port"      = "1813"
      "nios.servers.0.auth_type"      = "PAP"
      "nios.servers.0.disable"        = "false"
      "nios.servers.0.use_accounting" = "false"
      "nios.servers.0.use_mgmt_port"  = "false"
    }
  }

  step {
    nios {
      name = "{{random}}"
      servers = [
        {
          address        = "2.2.3.2"
          shared_secret  = "test"
          auth_type      = "CHAP"
          use_accounting = true
        }
      ]
    }
    check = {
      "nios.servers.#"                = "1"
      "nios.servers.0.address"        = "2.2.3.2"
      "nios.servers.0.auth_port"      = "1812"
      "nios.servers.0.acct_port"      = "1813"
      "nios.servers.0.auth_type"      = "CHAP"
      "nios.servers.0.disable"        = "false"
      "nios.servers.0.use_accounting" = "true"
      "nios.servers.0.use_mgmt_port"  = "false"
    }
  }

}
