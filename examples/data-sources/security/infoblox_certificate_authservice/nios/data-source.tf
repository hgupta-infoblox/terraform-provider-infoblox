// Retrieve a specific Certificate Authservice by filters
data "infoblox_certificate_authservice" "get_certificate_authservice_using_filters" {
  filters = {
    name = "example_certificate_authservice2"
  }
}

// Retrieve all Certificate Authservices
data "infoblox_certificate_authservice" "get_all_certificate_authservices" {}
