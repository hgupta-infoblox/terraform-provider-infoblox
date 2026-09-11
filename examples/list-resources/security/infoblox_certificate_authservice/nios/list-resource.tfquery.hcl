// List specific Certificate Authservices using filters
list "infoblox_certificate_authservice" "list_certificate_authservices_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_certificate_authservice"
    }
  }
  limit = 10
}

// List Certificate Authservices with resource details included
list "infoblox_certificate_authservice" "list_certificate_authservices_with_resource" {
  provider         = infoblox
  include_resource = true
}
