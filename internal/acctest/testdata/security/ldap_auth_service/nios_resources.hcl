# Auto-generated resource acceptance-test cases for LdapAuthService.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
    check = {
      "nios.name"                           = "{{random}}"
      "nios.servers.#"                      = "1"
      "nios.servers.0.address"              = "2.2.2.2"
      "nios.servers.0.authentication_type"  = "ANONYMOUS"
      "nios.servers.0.base_dn"              = "ou=People,dc=example,dc=com"
      "nios.servers.0.disable"              = "false"
      "nios.servers.0.encryption"           = "SSL"
      "nios.servers.0.port"                 = "636"
      "nios.servers.0.use_mgmt_port"        = "false"
      "nios.servers.0.version"              = "V3"
      "nios.search_scope"                   = "ONELEVEL"
      "nios.disable"                        = "false"
      "nios.ea_mapping.#"                   = "0"
      "nios.ldap_group_attribute"           = "memberOf"
      "nios.ldap_group_authentication_type" = "GROUP_ATTRIBUTE"
      "nios.mode"                           = "ORDERED_LIST"
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
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      comment             = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      comment             = "This is an updated comment"
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
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      disable             = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      disable             = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "ea_mapping" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ea_mapping          = [{ mapped_ea = "Availability zone", name = "ldapfield" }]
    }
    check = {
      "nios.ea_mapping.0.mapped_ea" = "Availability zone"
      "nios.ea_mapping.0.name"      = "ldapfield"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ea_mapping          = [{ mapped_ea = "Subnet Name", name = "ldapfield12" }]
    }
    check = {
      "nios.ea_mapping.0.mapped_ea" = "Subnet Name"
      "nios.ea_mapping.0.name"      = "ldapfield12"
    }
  }

}

case "ldap_group_attribute" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      ldap_user_attribute  = "adminID"
      recovery_interval    = 30
      retries              = 5
      timeout              = 5
      servers              = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ldap_group_attribute = "namecn"
    }
    check = {
      "nios.ldap_group_attribute" = "namecn"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      ldap_user_attribute  = "adminID"
      recovery_interval    = 30
      retries              = 5
      timeout              = 5
      servers              = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ldap_group_attribute = "namecnid"
    }
    check = {
      "nios.ldap_group_attribute" = "namecnid"
    }
  }

}

case "ldap_group_authentication_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                           = "{{random}}"
      ldap_user_attribute            = "adminID"
      recovery_interval              = 30
      retries                        = 5
      timeout                        = 5
      servers                        = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ldap_group_authentication_type = "POSIX_GROUP"
    }
    check = {
      "nios.ldap_group_authentication_type" = "POSIX_GROUP"
    }
  }

  step {
    nios {
      name                           = "{{random}}"
      ldap_user_attribute            = "adminID"
      recovery_interval              = 30
      retries                        = 5
      timeout                        = 5
      servers                        = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ldap_group_authentication_type = "GROUP_ATTRIBUTE"
    }
    check = {
      "nios.ldap_group_authentication_type" = "GROUP_ATTRIBUTE"
    }
  }

}

case "ldap_user_attribute" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ldap_user_attribute = "adminID"
    }
    check = {
      "nios.ldap_user_attribute" = "adminID"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      ldap_user_attribute = "adminID12"
    }
    check = {
      "nios.ldap_user_attribute" = "adminID12"
    }
  }

}

case "mode" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      mode                = "ORDERED_LIST"
    }
    check = {
      "nios.mode" = "ORDERED_LIST"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      mode                = "ROUND_ROBIN"
    }
    check = {
      "nios.mode" = "ROUND_ROBIN"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                = "{{random2}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
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
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
    check = {
      "nios.recovery_interval" = "30"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 60
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
    check = {
      "nios.recovery_interval" = "60"
    }
  }

}

case "retries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      retries             = 2
    }
    check = {
      "nios.retries" = "2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      retries             = 5
    }
    check = {
      "nios.retries" = "5"
    }
  }

}

case "search_scope" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      search_scope        = "BASE"
    }
    check = {
      "nios.search_scope" = "BASE"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      search_scope        = "SUBTREE"
    }
    check = {
      "nios.search_scope" = "SUBTREE"
    }
  }

}

case "servers" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.2"
      "nios.servers.0.authentication_type" = "ANONYMOUS"
      "nios.servers.0.base_dn"             = "ou=People,dc=example,dc=com"
      "nios.servers.0.disable"             = "false"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "636"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V3"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.4", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.4"
      "nios.servers.0.authentication_type" = "ANONYMOUS"
      "nios.servers.0.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.0.disable"             = "true"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "631"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.4", authentication_type = "AUTHENTICATED", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2", bind_password = "test", bind_user_dn = "cn=ldapbind,ou=People,dc=example,dc=com" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.4"
      "nios.servers.0.authentication_type" = "AUTHENTICATED"
      "nios.servers.0.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.0.disable"             = "true"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "631"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.4", authentication_type = "AUTHENTICATED", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2", bind_password = "test123", bind_user_dn = "cn=ldapbind,ou=People,dc=example,dc=com" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.4"
      "nios.servers.0.authentication_type" = "AUTHENTICATED"
      "nios.servers.0.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.0.disable"             = "true"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "631"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.4", authentication_type = "AUTHENTICATED", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2", bind_password = "test123", bind_user_dn = "cn=ldapbind,ou=People,dc=example,dc=com" }, { address = "2.2.2.5", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.4"
      "nios.servers.0.authentication_type" = "AUTHENTICATED"
      "nios.servers.0.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.0.disable"             = "true"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "631"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V2"
      "nios.servers.1.address"             = "2.2.2.5"
      "nios.servers.1.authentication_type" = "ANONYMOUS"
      "nios.servers.1.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.1.disable"             = "true"
      "nios.servers.1.encryption"          = "SSL"
      "nios.servers.1.port"                = "631"
      "nios.servers.1.use_mgmt_port"       = "false"
      "nios.servers.1.version"             = "V2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.4", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2" }, { address = "2.2.2.5", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.4"
      "nios.servers.0.authentication_type" = "ANONYMOUS"
      "nios.servers.0.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.0.disable"             = "true"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "631"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V2"
      "nios.servers.1.address"             = "2.2.2.5"
      "nios.servers.1.authentication_type" = "ANONYMOUS"
      "nios.servers.1.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.1.disable"             = "true"
      "nios.servers.1.encryption"          = "SSL"
      "nios.servers.1.port"                = "631"
      "nios.servers.1.use_mgmt_port"       = "false"
      "nios.servers.1.version"             = "V2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.2"
      "nios.servers.0.authentication_type" = "ANONYMOUS"
      "nios.servers.0.base_dn"             = "ou=People,dc=example,dc=com"
      "nios.servers.0.disable"             = "false"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "636"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V3"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.2"
      "nios.servers.0.authentication_type" = "ANONYMOUS"
      "nios.servers.0.base_dn"             = "ou=People,dc=example,dc=com"
      "nios.servers.0.disable"             = "false"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "636"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V3"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 5
      servers             = [{ address = "2.2.2.4", authentication_type = "AUTHENTICATED", base_dn = "ou=People,dc=example1,dc=com", disable = true, encryption = "SSL", port = 631, use_mgmt_port = false, version = "V2", bind_password = "test", bind_user_dn = "cn=ldapbind,ou=People,dc=example,dc=com" }]
    }
    check = {
      "nios.servers.0.address"             = "2.2.2.4"
      "nios.servers.0.authentication_type" = "AUTHENTICATED"
      "nios.servers.0.base_dn"             = "ou=People,dc=example1,dc=com"
      "nios.servers.0.disable"             = "true"
      "nios.servers.0.encryption"          = "SSL"
      "nios.servers.0.port"                = "631"
      "nios.servers.0.use_mgmt_port"       = "false"
      "nios.servers.0.version"             = "V2"
    }
  }

}

case "timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      timeout             = 15
    }
    check = {
      "nios.timeout" = "15"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      servers             = [{ address = "2.2.2.2", authentication_type = "ANONYMOUS", base_dn = "ou=People,dc=example,dc=com", disable = false, encryption = "SSL", port = 636, use_mgmt_port = false, version = "V3" }]
      timeout             = 25
    }
    check = {
      "nios.timeout" = "25"
    }
  }

}
