# RecordRpzCnameIpaddress — nios list cases
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - RPZ zone : tf-acc-rpz.com        (view: default)
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true

  step {
    nios {
      name      = "11.0.0.40.tf-acc-rpz.com"
      canonical = "11.0.0.40"
      rp_zone   = "tf-acc-rpz.com"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true

  step {
    nios {
      name      = "11.0.0.41.tf-acc-rpz.com"
      canonical = "11.0.0.41"
      rp_zone   = "tf-acc-rpz.com"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true

  step {
    nios {
      name      = "11.0.0.42.tf-acc-rpz.com"
      canonical = "11.0.0.42"
      rp_zone   = "tf-acc-rpz.com"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
