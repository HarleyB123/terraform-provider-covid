package covid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "https://coronavirus-19-api.herokuapp.com/countries",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"covid_world": dataSourceWorld(),
//			"country":   dataSourceCountry(), - coming soon
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	providerConfig := make(map[string]interface{})
	providerConfig["url"] = d.Get("url").(string)
	return providerConfig, diags
}
