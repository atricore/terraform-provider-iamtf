package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppliance_001(t *testing.T) {
	runAccTest(t, "ida01")
}

func runAccTest(t *testing.T, n string) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(fmt.Sprintf("appliances/%s", n))
	identityApplianceConfig := mgr.GetFixtures(fmt.Sprintf("%s.tf", n), ri, t)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		Steps: []resource.TestStep{
			{
				Config: identityApplianceConfig,
			},
		},
	})
}
