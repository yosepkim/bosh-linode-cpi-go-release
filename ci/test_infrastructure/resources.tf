resource "random_string" "account_suffix" {
  length  = 4
  upper   = false
  special = false
  lower   = true
  number  = true
}

