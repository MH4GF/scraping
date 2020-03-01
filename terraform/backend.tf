terraform {
  backend "s3" {
    bucket  = "ort-terraform"
    key     = "scraing_moneyforward/terraform.tfstate"
    region  = "ap-northeast-1"
    profile = "ort"
  }
}
