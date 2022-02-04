provider "aws" {
  region = var.region
}

terraform {
  required_version = ">= 1.0.0" 

  required_providers {
    aws = ">= 3.24"
  }

  backend "s3" {
    bucket = "rpolnx-tfstate"
    key    = "lightsail-intance/terraform.tfstate"
    region = "sa-east-1"
  }
}
