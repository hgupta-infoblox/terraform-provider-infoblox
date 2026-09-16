# TacacsplusAuthservice list cases.
# Note: NIOS allows only one TACACS+ auth service per grid (singleton). Cases run sequentially.
case "basic" {
  backend = "nios"
  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
  }
  step {
    query    = true
    provider = infoblox
    limit    = 5
  }
}
case "filters" {
  backend = "nios"
  step {
    nios {
      name    = "{{random}}"
      servers = [{ address = "2.2.3.3", auth_type = "CHAP", disable = false, port = 49, use_accounting = false, use_mgmt_port = false, shared_secret = "test" }]
    }
  }
  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = { name = "nios.name" }
    }
  }
}
