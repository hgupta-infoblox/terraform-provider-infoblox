// Get services filtered by an attribute
data "infoblox_infra_service" "example_by_attribute" {
  filters = {
    "name" = "example_service"
  }
}

// Get services filtered by tag
data "infoblox_infra_service" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all services
data "infoblox_infra_service" "example_all" {}
