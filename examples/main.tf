terraform {
  required_providers {
    covid = {
      source  = "local/provider/covid"
    }
  }
}

provider "covid" {}

data "covid_world" "all" {}

output "all_available_countries" {
    value = data.covid_world.all
}
