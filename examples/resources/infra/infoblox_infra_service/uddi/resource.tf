resource "infoblox_infra_service" "example" {
  uddi = {
    name         = "example_service"
    pool_id      = "infra/pool/<pool-id>"
    service_type = "dhcp"

    // Other Optional fields
    description     = "DHCP service"
    desired_version = "3.5.0"
    desired_state   = "start"
    tags = {
      Site = "location-1"
    }
  }
}
