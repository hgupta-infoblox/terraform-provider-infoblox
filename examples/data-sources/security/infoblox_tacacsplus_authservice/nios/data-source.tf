// Retrieve a specific TACACS+ Authservice by filters
data "infoblox_tacacsplus_authservice" "get_tacacsplus_authservice_using_filters" {
  filters = {
    name = "tacacsplus_authservice1"
  }
}

// Retrieve all TACACS+ Authservices
data "infoblox_tacacsplus_authservice" "get_all_tacacsplus_authservices" {}
