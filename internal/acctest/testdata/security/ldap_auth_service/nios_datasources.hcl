# Auto-generated datasource acceptance-test cases for LdapAuthService.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ldap_group_attribute", "nios.ldap_group_authentication_type", "nios.ldap_user_attribute", "nios.mode", "nios.name", "nios.recovery_interval", "nios.retries", "nios.search_scope", "nios.timeout"]

  step {
    nios {
      name                = "{{random}}"
      ldap_user_attribute = "adminID"
      recovery_interval   = 30
      retries             = 5
      timeout             = 15
      servers             = [{ address = "2.2.2.2", base_dn = "ou=People,dc=example,dc=com" }]
    }
  }

}
