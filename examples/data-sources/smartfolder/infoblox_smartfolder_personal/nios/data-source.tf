// Retrieve a specific Personal Smart Folder by name filter
data "infoblox_smartfolder_personal" "get_personal_smart_folder_by_name" {
  filters = {
    name = "My Personal Folder"
  }
}

// Retrieve all Personal Smart Folders
data "infoblox_smartfolder_personal" "get_all_personal_smart_folders" {}
