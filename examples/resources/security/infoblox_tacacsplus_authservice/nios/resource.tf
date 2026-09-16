// Create a TACACS+ Authentication Service
// Note: Only one TACACS+ Authentication Service may be configured per grid.
resource "infoblox_tacacsplus_authservice" "example" {
  nios = {
    name    = "example_tacacsplus_authservice"
    comment = "Example TACACS+ Authentication Service"
    disable = false

    acct_retries = 2
    acct_timeout = 2300
    auth_retries = 2
    auth_timeout = 7000

    servers = [
      {
        address        = "192.168.1.10"
        port           = 49
        shared_secret  = "example_secret"
        auth_type      = "CHAP"
        disable        = false
        use_mgmt_port  = false
        use_accounting = false
        comment        = "Primary TACACS+ server"
      }
    ]
  }
}
