package covid

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWorld() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorldRead,
		Schema: map[string]*schema.Schema{
			"countries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cases": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWorldRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	providerConfig := m.(map[string]interface{})
	url := providerConfig["url"].(string)

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// Unmarshal data
	allCountries := &covid{}
	err = json.NewDecoder(r.Body).Decode(&allCountries)
	if err != nil {
		return diag.FromErr(err)
	}

	countries := make([]map[string]interface{}, 0)
	// Set broken values for all currently available city data
	for _, v := range allCountries.Countries {
		country := make(map[string]interface{})
		country["country"] = v.Country
		country["cases"] = v.Cases
		countries = append(countries, country)
	}
	if err := d.Set("countries", countries); err != nil {
		return diag.FromErr(err)
	}

	// Change ID every run to force update
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
