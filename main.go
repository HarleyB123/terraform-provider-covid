package main
  
import (
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
        "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

        "github.com/HarleyB123/terraform-provider-covid/covid"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() *schema.Provider {
                        return covid.Provider()
                },
        })
}
