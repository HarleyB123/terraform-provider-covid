terraform {
  required_providers {
    covid = {
      source  = "local/provider/covid"
    }
  }
}

provider "covid" {}
