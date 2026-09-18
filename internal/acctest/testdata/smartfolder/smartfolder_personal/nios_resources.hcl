# Auto-generated resource acceptance-test cases for SmartfolderPersonal.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
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
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "Updated comment"
    }
    check = {
      "nios.comment" = "Updated comment"
    }
  }

}

case "group_bys" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      group_bys = [{ enable_grouping = true, value = "Availability zone", value_type = "NORMAL" }]
    }
    check = {
      "nios.group_bys.#"                 = "1"
      "nios.group_bys.0.enable_grouping" = "true"
      "nios.group_bys.0.value"           = "Availability zone"
      "nios.group_bys.0.value_type"      = "NORMAL"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      group_bys = [{ enable_grouping = true, value = "Site", value_type = "EXTATTR" }]
    }
    check = {
      "nios.group_bys.#"                 = "1"
      "nios.group_bys.0.enable_grouping" = "true"
      "nios.group_bys.0.value"           = "Site"
      "nios.group_bys.0.value_type"      = "EXTATTR"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      group_bys = [{ enable_grouping = false, value = "Site", value_type = "EXTATTR" }]
    }
    check = {
      "nios.group_bys.#"                 = "1"
      "nios.group_bys.0.enable_grouping" = "false"
      "nios.group_bys.0.value"           = "Site"
      "nios.group_bys.0.value_type"      = "EXTATTR"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "query_items" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      query_items = [{ field_type = "NORMAL", name = "type", op_match = true, operator = "EQ", value = { value_string = "Network" }, value_type = "ENUM" }]
    }
    check = {
      "nios.query_items.#"                    = "1"
      "nios.query_items.0.field_type"         = "NORMAL"
      "nios.query_items.0.name"               = "type"
      "nios.query_items.0.op_match"           = "true"
      "nios.query_items.0.operator"           = "EQ"
      "nios.query_items.0.value.value_string" = "Network"
      "nios.query_items.0.value_type"         = "ENUM"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      query_items = [{ field_type = "NORMAL", name = "last_discovered_timestamp", op_match = true, operator = "EQ", value = { value_date = 1757469600 }, value_type = "DATE" }]
    }
    check = {
      "nios.query_items.#"                  = "1"
      "nios.query_items.0.field_type"       = "NORMAL"
      "nios.query_items.0.name"             = "last_discovered_timestamp"
      "nios.query_items.0.op_match"         = "true"
      "nios.query_items.0.operator"         = "EQ"
      "nios.query_items.0.value.value_date" = "1757469600"
      "nios.query_items.0.value_type"       = "DATE"
    }
  }

}
