output "public_ip" {
  value   = aws_lightsail_static_ip.this.ip_address
}

output "instance_arn" {
  value   = aws_lightsail_instance.this.arn
}

output "instance_id" {
  value   = aws_lightsail_instance.this.id
}