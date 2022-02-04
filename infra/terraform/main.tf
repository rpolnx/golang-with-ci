resource "aws_lightsail_static_ip" "this" {
  name = var.private_ip_name
}

resource "aws_lightsail_instance" "this" {
  name              = lookup(var.basic_instance_info, "instance_name")
  availability_zone = lookup(var.basic_instance_info, "availability_zone")
  blueprint_id      = lookup(var.basic_instance_info, "blueprint_id")
  bundle_id         = lookup(var.basic_instance_info, "bundle_id")
  key_pair_name     = lookup(var.basic_instance_info, "key_name")
  user_data = local.user_data
  tags = var.tags
}

resource "aws_lightsail_static_ip_attachment" "this" {
  static_ip_name = aws_lightsail_static_ip.this.name
  instance_name  = aws_lightsail_instance.this.name
}

resource "aws_lightsail_instance_public_ports" "this" {
  instance_name = aws_lightsail_instance.this.name

  port_info {
    protocol  = "tcp"
    from_port = 22
    to_port   = 22
  }

  port_info {
    protocol  = "tcp"
    from_port = 80
    to_port   = 80
  }

  port_info {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
  }

    port_info {
    protocol  = "tcp"
    from_port = 8080
    to_port   = 8080
  }
}