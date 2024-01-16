provider "aws" {
  
}

resource "aws_instance" "nginx_ec2" {
    availability_zone = "*"
    count = 2
}

provider "local-exec" {
    # command = "ansible-playbook -i ${self.network_interface.0.access_config.0.nat_ip}, --private-key ${local.private_key_path} nginx.yaml"
    command = "ansible-playbook -i ${aws_instance.nginx.public_ip}, --private-key ${local.private_key_path} play_nginx.yml"
  
}

