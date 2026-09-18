// List specific Personal Smart Folders using filters
list "infoblox_smartfolder_personal" "list_personal_smart_folders_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "My Personal Folder"
    }
  }
  limit = 10
}

// List all Personal Smart Folders with resource details
list "infoblox_smartfolder_personal" "list_all_personal_smart_folders" {
  provider         = infoblox
  include_resource = true
}
