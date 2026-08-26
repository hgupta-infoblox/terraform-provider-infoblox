# RecordRpzCnameIpaddress — nios resource test cases
# TODO : must exist before running tests
#   Default view:  zone_rp "tf-acc-rpz.com" in view "default"
#   Custom view:   view "tf-acc-rpz-view" + zone_rp "tf-acc-rpz.com" in that view
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.1.tf-acc-rpz.com"
      canonical = "11.0.0.1"
      rp_zone   = "tf-acc-rpz.com"
    }
    check = {
      "nios.name"      = "11.0.0.1.tf-acc-rpz.com"
      "nios.canonical" = "11.0.0.1"
      "nios.rp_zone"   = "tf-acc-rpz.com"
      "nios.view"      = "default"
      "nios.disable"   = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name      = "11.0.0.2.tf-acc-rpz.com"
      canonical = "11.0.0.2"
      rp_zone   = "tf-acc-rpz.com"
    }
  }

}

case "canonical" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.3.tf-acc-rpz.com"
      canonical = "11.0.0.3"
      rp_zone   = "tf-acc-rpz.com"
    }
    check = {
      "nios.canonical" = "11.0.0.3"
    }
  }

  step {
    nios {
      name      = "11.0.0.3.tf-acc-rpz.com"
      canonical = "*"
      rp_zone   = "tf-acc-rpz.com"
    }
    check = {
      "nios.canonical" = "*"
    }
  }

  step {
    nios {
      name      = "11.0.0.3.tf-acc-rpz.com"
      canonical = "11.0.0.3"
      rp_zone   = "tf-acc-rpz.com"
    }
    check = {
      "nios.canonical" = "11.0.0.3"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.4.tf-acc-rpz.com"
      canonical = "11.0.0.4"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      comment   = "This is a new rpz cname record"
    }
    check = {
      "nios.comment" = "This is a new rpz cname record"
    }
  }

  step {
    nios {
      name      = "11.0.0.4.tf-acc-rpz.com"
      canonical = "11.0.0.4"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      comment   = "This is an updated rpz cname record"
    }
    check = {
      "nios.comment" = "This is an updated rpz cname record"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.5.tf-acc-rpz.com"
      canonical = "11.0.0.5"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "11.0.0.5.tf-acc-rpz.com"
      canonical = "11.0.0.5"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      disable   = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.6.tf-acc-rpz.com"
      canonical = "11.0.0.6"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "11.0.0.6.tf-acc-rpz.com"
      canonical = "11.0.0.6"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.7.tf-acc-rpz.com"
      canonical = "11.0.0.7"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
    }
  }

  step {
    nios {
      name      = "11.0.0.8.tf-acc-rpz.com"
      canonical = "11.0.0.8"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
    }
  }

}

case "rp_zone" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.9.tf-acc-rpz.com"
      canonical = "11.0.0.9"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
    }
    check = {
      "nios.rp_zone" = "tf-acc-rpz.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.10.tf-acc-rpz.com"
      canonical = "11.0.0.10"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "11.0.0.10.tf-acc-rpz.com"
      canonical = "11.0.0.10"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
      ttl       = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

  step {
    nios {
      name      = "11.0.0.10.tf-acc-rpz.com"
      canonical = "11.0.0.10"
      rp_zone   = "tf-acc-rpz.com"
      view      = "default"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "11.0.0.12.tf-acc-rpz.com"
      canonical = "11.0.0.12"
      rp_zone   = "tf-acc-rpz.com"
      view      = "tf-acc-rpz-view"
    }
    check = {
      "nios.view" = "tf-acc-rpz-view"
    }
  }

}
