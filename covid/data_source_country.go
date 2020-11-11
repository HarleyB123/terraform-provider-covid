package covid

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	"io/ioutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCountry() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCountryRead,
		Schema: map[string]*schema.Schema{
			"country": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cases": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourceCountryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	providerConfig := m.(map[string]interface{})
	url := providerConfig["url"].(string)
	userChosenCountry := d.Get("country").(string)

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

	content, _ := ioutil.ReadAll(r.Body)
	// Unmarshal data
	var covid []covid
	err = json.Unmarshal([]byte(content), &covid)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("cases", -1); err != nil {
		return diag.FromErr(err)
	}

	for _, v := range covid {
		if strings.EqualFold(v.Country, userChosenCountry) {
			if err := d.Set("cases", v.Cases); err != nil {
				return diag.FromErr(err)
			}
			break
		}
	}
	// Change ID every run to force update
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

