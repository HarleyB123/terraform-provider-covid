terraform {
  required_providers {
    covid = {
      source = "local/provider/covid"
    }
  }
}

provider "covid" {}

data "covid_world" "all" {}

data "covid_country" "Cyprus" {
  country = "Cyprus"
}


// Get covid cases in each country
output "all_available_countries" {
  value = data.covid_world.all.countries
}

// Get covid cases in a given country
output "cyprus_cases" {
  value = data.covid_country.Cyprus.cases
}
