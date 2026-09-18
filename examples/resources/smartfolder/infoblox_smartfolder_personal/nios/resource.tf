// Create a Personal Smart Folder with basic fields
resource "infoblox_smartfolder_personal" "personal_smart_folder_basic" {
  nios = {
    name = "My Personal Folder"
  }
}

// Create a Personal Smart Folder with a comment
resource "infoblox_smartfolder_personal" "personal_smart_folder_with_comment" {
  nios = {
    name    = "My Commented Folder"
    comment = "Folder for tracking network objects by type"
  }
}

// Create a Personal Smart Folder with custom query items
resource "infoblox_smartfolder_personal" "personal_smart_folder_with_query" {
  nios = {
    name = "Network Type Folder"
    query_items = [
      {
        name       = "type"
        field_type = "NORMAL"
        operator   = "EQ"
        op_match   = true
        value_type = "ENUM"
        value = {
          value_string = "Network"
        }
      }
    ]
  }
}

// Create a Personal Smart Folder with grouping rules
resource "infoblox_smartfolder_personal" "personal_smart_folder_with_grouping" {
  nios = {
    name = "Grouped Folder"
    group_bys = [
      {
        value           = "Site"
        value_type      = "EXTATTR"
        enable_grouping = true
      }
    ]
  }
}
