// Create Record RPZ CNAME IP Address with Basic Fields
resource "infoblox_record_rpz_cname_ipaddress" "basic" {
  nios = {
    name      = "11.0.0.0.rpzip.example.com"
    canonical = "11.0.0.0"
    rp_zone   = "rpzip.example.com"
  }
}

// Create Record RPZ CNAME IP Address with Additional Fields
resource "infoblox_record_rpz_cname_ipaddress" "additional_fields" {
  nios = {
    name      = "11.0.0.1.rpzip.example.com"
    canonical = "11.0.0.1"
    rp_zone   = "rpzip.example.com"
    view      = "default"
    ttl       = 10
    comment   = "Example RPZ CNAME IP address record"
    disable   = false
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Block IP Address (No Data) Rule
resource "infoblox_record_rpz_cname_ipaddress" "block_no_data" {
  nios = {
    name      = "11.0.0.3.rpzip.example.com"
    canonical = "*"
    rp_zone   = "rpzip.example.com"
  }
}
