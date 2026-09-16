// List specific TACACS+ Authservices using filters
list "infoblox_tacacsplus_authservice" "list_tacacsplus_authservices_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_tacacsplus_authservice"
    }
  }
  limit = 10
}

// List TACACS+ Authservices with resource details included
list "infoblox_tacacsplus_authservice" "list_tacacsplus_authservices_with_resource" {
  provider         = infoblox
  include_resource = true
}
