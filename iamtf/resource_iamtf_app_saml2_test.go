package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJossoExtSaml2Sp_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(extSaml2Sp)
	config := mgr.GetFixtures("app_saml2.tf", ri, t)
	updatedConfig := mgr.GetFixtures("app_saml2_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", extSaml2Sp)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(extSaml2Sp, createDoesExtSaml2SpExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("sp", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("sp", ri)),
				),
			},
		},
	})
}

func createDoesExtSaml2SpExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
