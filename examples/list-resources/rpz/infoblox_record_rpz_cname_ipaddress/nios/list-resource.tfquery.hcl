// List specific RPZ CNAME IP address records using filters
list "infoblox_record_rpz_cname_ipaddress" "by_view" {
  provider = infoblox
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}

// List RPZ CNAME IP address records using Extensible Attributes
list "infoblox_record_rpz_cname_ipaddress" "by_ext_attr" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List RPZ CNAME IP address records with resource details included
list "infoblox_record_rpz_cname_ipaddress" "with_resource" {
  provider         = infoblox
  include_resource = true
}
