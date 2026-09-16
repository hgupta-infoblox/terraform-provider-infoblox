# Auto-generated resource acceptance-test cases for TacacsplusAuthservice.
case "basic" {
  backend = "nios"

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
    check = {
      "nios.name"                = "{{random}}"
      "nios.servers.#"           = "1"
      "nios.servers.0.address"   = "2.2.3.3"
      "nios.servers.0.auth_type" = "CHAP"
      "nios.servers.0.disable"   = "false"
      "nios.servers.0.port"      = "49"
      "nios.acct_retries"        = "0"
      "nios.acct_timeout"        = "1000"
      "nios.auth_retries"        = "0"
      "nios.auth_timeout"        = "5000"
      "nios.comment"             = ""
      "nios.disable"             = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
  }

}

case "acct_retries" {
  backend = "nios"

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      acct_retries = 1
    }
    check = {
      "nios.acct_retries" = "1"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      acct_retries = 3
    }
    check = {
      "nios.acct_retries" = "3"
    }
  }

}

case "acct_timeout" {
  backend = "nios"

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      acct_timeout = 3600
    }
    check = {
      "nios.acct_timeout" = "3600"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      acct_timeout = 7200
    }
    check = {
      "nios.acct_timeout" = "7200"
    }
  }

}

case "auth_retries" {
  backend = "nios"

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      auth_retries = 2
    }
    check = {
      "nios.auth_retries" = "2"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      auth_retries = 4
    }
    check = {
      "nios.auth_retries" = "4"
    }
  }

}

case "auth_timeout" {
  backend = "nios"

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      auth_timeout = 7000
    }
    check = {
      "nios.auth_timeout" = "7000"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      servers      = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      auth_timeout = 6000
    }
    check = {
      "nios.auth_timeout" = "6000"
    }
  }

}

case "comment" {
  backend = "nios"

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      comment = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "disable" {
  backend = "nios"

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      disable = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
      disable = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "name" {
  backend = "nios"

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name    = "{{random2}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "servers" {
  backend = "nios"

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
    check = {
      "nios.servers.#"                = "1"
      "nios.servers.0.address"        = "2.2.3.3"
      "nios.servers.0.auth_type"      = "CHAP"
      "nios.servers.0.disable"        = "false"
      "nios.servers.0.port"           = "49"
      "nios.servers.0.use_accounting" = "false"
      "nios.servers.0.use_mgmt_port"  = "false"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.1.3", auth_type = "PAP", disable = true, port = 49, use_accounting = false, use_mgmt_port = true, shared_secret = "testing_key" }]
    }
    check = {
      "nios.servers.#"                = "1"
      "nios.servers.0.address"        = "2.2.1.3"
      "nios.servers.0.auth_type"      = "PAP"
      "nios.servers.0.disable"        = "true"
      "nios.servers.0.port"           = "49"
      "nios.servers.0.use_accounting" = "false"
      "nios.servers.0.use_mgmt_port"  = "true"
    }
  }

}
