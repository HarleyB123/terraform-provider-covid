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


// Uncomment these if you want to test failure cases

/*
data "covid_country" "slightly_mispelled" {
   country = "Brasil" // should be Brazil
}

output "mispell_error" {
   value = data.covid_country.slightly_mispelled.cases
}

data "covid_country" "not_a_country" {
   country = "blahblahland"
}

output "not_a_country_error" {
   value = data.covid_country.not_a_country.cases
}
*/
