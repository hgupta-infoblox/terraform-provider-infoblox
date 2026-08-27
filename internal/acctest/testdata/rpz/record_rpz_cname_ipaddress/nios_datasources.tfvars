# Auto-generated datasource acceptance-test cases for RecordRpzCnameIpaddress.
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - RPZ zone : tf-acc-rpz.com        (view: default)
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = "11.0.0.30.tf-acc-rpz.com"
      canonical = "11.0.0.30"
      rp_zone   = "tf-acc-rpz.com"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = "11.0.0.31.tf-acc-rpz.com"
      canonical = "11.0.0.31"
      rp_zone   = "tf-acc-rpz.com"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
