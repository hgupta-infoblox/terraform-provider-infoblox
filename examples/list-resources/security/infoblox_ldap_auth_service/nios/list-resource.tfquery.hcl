// List specific LDAP Auth Services using filters
list "infoblox_ldap_auth_service" "list_ldap_auth_services_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ldap_authservice"
    }
  }
}

// List LDAP Auth Services with resource details included
list "infoblox_ldap_auth_service" "list_ldap_auth_services_with_resource" {
  provider         = infoblox
  include_resource = true
}
