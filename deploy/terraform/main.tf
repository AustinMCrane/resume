terraform {
  required_providers {
    linode = {
      source = "linode/linode"
      version = "2.5.2"
    }
  }
}

provider "linode" {}

resource "linode_instance" "austinmcrane" {
    region = "us-central"
    label = "austinmcrane"
    image = "linode/ubuntu22.04"
    type = "g6-nanode-1"

    root_pass = var.root_password
    #authorized_keys = [linode_sshkey.main_key.ssh_key]
}

output "ip_address" {
  value = linode_instance.austinmcrane.ip_address
}
