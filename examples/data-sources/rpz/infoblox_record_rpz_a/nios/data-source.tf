// Retrieve a specific Substitute (IPv4 Address) Rule by filters
data "infoblox_record_rpz_a" "get_record_using_filters" {
  filters = {
    name = "a.rpz.example.com"
  }
}

// Retrieve specific Substitute (IPv4 Address) Rules using Extensible Attributes
data "infoblox_record_rpz_a" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Substitute (IPv4 Address) Rules
data "infoblox_record_rpz_a" "get_all_record_rpz_a" {}
