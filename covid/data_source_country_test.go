package covid

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCovidCountry(t *testing.T) {
	casesRegex, _ := regexp.Compile(`^[0-9]+$`)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// valid country
			{
				Config: testAccCheckCountry("Cyprus"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.covid_country.Cyprus", "country", "Cyprus"),
					resource.TestCheckResourceAttrSet(
						"data.covid_country.Cyprus", "cases"),
					resource.TestMatchResourceAttr(
						"data.covid_country.Cyprus", "cases", casesRegex),
				),
			},
			// invalid country
			{
				Config: testAccCheckCountry("not_found"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.covid_country.not_found", "country", "not_found"),
					resource.TestCheckResourceAttr(
						"data.covid_country.not_found", "cases", "-1"),
				),
			},
		},
	})
}

func testAccCheckCountry(country string) string {
	return fmt.Sprintf(`
data "covid_country" "%[1]v" {
  country = "%[1]v"
}
`, country)
}
