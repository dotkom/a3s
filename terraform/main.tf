terraform {
  required_version = "~> 1.1.7"
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.8.0"
    }
  }
  backend "s3" {
    bucket = "a3s-dev"
    key = "terraform/a3s.tfstate"
    region = "eu-north-1"
  }
}

provider "aws" {
  region = "eu-north-1"
  default_tags {
  }
}
