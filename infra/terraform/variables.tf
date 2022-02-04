variable "env" {
  description = "Instance name"
  type        = string
  default     = "DEV"
}

variable "tags" {
  description = "Instance tags"
  type        = map(string)
  default     = {
      Environment           = "DEV"
      Terraform             = true
  }
}

variable "private_ip_name" {
  description = "Private ip name"
  type        = string
  default     = "dev-ip"
}

variable "basic_instance_info" {
  description = "Instance name"
  type        = map(string)
  default     = {
      key_name                  = "aws_lightsail_key"
      availability_zone         = "us-east-1b"
      blueprint_id              = "ubuntu_20_04"
      bundle_id                 = "micro_2_0"
      instance_name             = "instance-dev"
  }
}

variable "region" {
  description = "Region"
  type        = string
  default     = "us-east-1"
}