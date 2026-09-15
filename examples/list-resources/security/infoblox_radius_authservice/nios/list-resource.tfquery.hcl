// List specific Radius Authentication Services using filters
list "infoblox_radius_authservice" "list_radius_authservices_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_radius_authservice"
    }
  }
  limit = 10
}

// List all Radius Authentication Services
list "infoblox_radius_authservice" "list_all_radius_authservices" {
  provider = infoblox
}

// List Radius Authentication Services with resource details included
list "infoblox_radius_authservice" "list_radius_authservices_with_resource" {
  provider         = infoblox
  include_resource = true
}
