// Note: RPZ zones must be pre-created in NIOS (infoblox_zone_rp is not managed by this provider).
// Create a Substitute (IPv4 Address) Rule with Basic Fields
resource "infoblox_record_rpz_a" "create_record_basic" {
  nios = {
    name     = "a.rpz.example.com"
    ipv4addr = "192.168.1.1"
    rp_zone  = "rpz.example.com"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Substitute (IPv4 Address) Rule with Additional Fields
resource "infoblox_record_rpz_a" "create_record_additional_fields" {
  nios = {
    // Basic Fields
    name     = "a1.rpz.example.com"
    ipv4addr = "192.168.1.2"
    rp_zone  = "rpz.example.com"

    // Additional Fields
    ttl     = 3600
    disable = false
    comment = "RPZ A record created by Terraform"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create DNS View (Required as Parent)
// Note: RPZ zone "custom-rpz.example.com" must be pre-created in NIOS in this view.
resource "infoblox_view" "custom_view" {
  nios = {
    name = "custom-view"
  }
}

// Create a Substitute (IPv4 Address) Rule in a Custom View
resource "infoblox_record_rpz_a" "create_record_custom_view" {
  nios = {
    name     = "a.custom-rpz.example.com"
    ipv4addr = "192.168.2.1"
    rp_zone  = "custom-rpz.example.com"
    view     = infoblox_view.custom_view.nios.name
  }
}
