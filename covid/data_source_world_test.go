package covid

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCovidWorld(t *testing.T) {
	casesRegex, _ := regexp.Compile(`^[0-9]+$`)
	countryRegex, _ := regexp.Compile(`[a-zA-Z]+$`)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWorld(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.covid_world.all", "countries.0.country"),
					resource.TestCheckResourceAttrSet(
						"data.covid_world.all", "countries.0.cases"),
					resource.TestMatchResourceAttr(
						"data.covid_world.all", "countries.0.country", countryRegex),
					resource.TestMatchResourceAttr(
						"data.covid_world.all", "countries.0.cases", casesRegex),
					resource.TestCheckResourceAttrSet(
						"data.covid_world.all", "countries.1.country"),
					resource.TestCheckResourceAttrSet(
						"data.covid_world.all", "countries.1.cases"),
					resource.TestMatchResourceAttr(
						"data.covid_world.all", "countries.1.country", countryRegex),
					resource.TestMatchResourceAttr(
						"data.covid_world.all", "countries.1.cases", casesRegex),
				),
			},
		},
	})
}

func testAccCheckWorld() string {
	return `data "covid_world" "all" {}`
}
