package iamtf

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIdentityAppliance_read(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(identityAppliance)
	identityApplianceConfig := mgr.GetFixtures("generic_ida.tf", ri, t)
	config := mgr.GetFixtures("datasource.tf", ri, t)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		Steps: []resource.TestStep{
			{
				Config: identityApplianceConfig,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.iamtf_identity_appliance.test", "name"),
				),
			},
		},
	})
}
