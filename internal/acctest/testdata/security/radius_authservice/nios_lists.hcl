# Auto-generated list acceptance-test cases for RadiusAuthservice.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name = "{{random}}"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
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

  step {
    nios {
      name = "{{random}}"
      servers = [
        {
          address       = "2.2.3.1"
          shared_secret = "test"
        }
      ]
    }
  }

  step {
    query    = true
    provider = infoblox
    filter {
      name = "nios.name"
    }
  }

}
