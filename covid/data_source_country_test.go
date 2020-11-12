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
			//incorrect spelling country
			{
				Config: testAccCheckCountry("Brasil"),
				ExpectError: regexp.MustCompile("Unable to find Country Brasil, did you mean Brazil?"),
			},
			// invalid country
			{
				Config: testAccCheckCountry("not_found"),
				ExpectError: regexp.MustCompile("Unable to find Country not_found"),
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
