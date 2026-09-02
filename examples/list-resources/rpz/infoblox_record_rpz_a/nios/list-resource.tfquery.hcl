// List specific Substitute (IPv4 Address) Rules using filters
list "infoblox_record_rpz_a" "list_record_rpz_a_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "a.rpz.example.com"
    }
  }
}

// List specific Substitute (IPv4 Address) Rules using Extensible Attributes
list "infoblox_record_rpz_a" "list_record_rpz_a_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Substitute (IPv4 Address) Rules with resource details included
list "infoblox_record_rpz_a" "list_record_rpz_a_with_resource" {
  provider         = infoblox
  include_resource = true
}
