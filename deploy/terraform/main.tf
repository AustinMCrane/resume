terraform {
  required_providers {
    linode = {
      source = "linode/linode"
      version = "2.5.2"
    }
  }
}

provider "linode" {
  token = var.linode_token
}

resource "linode_sshkey" "main_key" {
    label = "main"
    ssh_key = chomp(file("~/.ssh/id_rsa.pub"))
}

resource "linode_instance" "web" {
    region = "us-central"
    label = "austinmcrane"
    image = "linode/ubuntu22.04"
    type = "g6-nanode-1"

    root_pass = var.root_password
    authorized_keys = [linode_sshkey.main_key.ssh_key]
}

output "ip_address" {
  value = linode_instance.web.ip_address
}
