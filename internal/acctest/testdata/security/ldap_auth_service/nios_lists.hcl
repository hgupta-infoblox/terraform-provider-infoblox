# LdapAuthService — nios list cases
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 3
      timeout             = 5
      servers             = [{ address = "2.2.2.2", base_dn = "ou=People,dc=example,dc=com" }]
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }
}

case "filters" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 3
      timeout             = 5
      servers             = [{ address = "2.2.2.2", base_dn = "ou=People,dc=example,dc=com" }]
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "nios.name"
      }
    }
  }
}

