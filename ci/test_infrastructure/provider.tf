provider "linode" {
  region = "${var.linode_region}"
  token  = "${var.linode_token}"
}

provider "random" {
}
