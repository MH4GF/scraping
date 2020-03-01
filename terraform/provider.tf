provider "aws" {
  region  = "ap-northeast-1"
  profile = "ort"
}

/* 使用時の AWS Account ID */
data "aws_caller_identity" "self" {}
