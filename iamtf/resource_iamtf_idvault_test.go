package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdVault_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idVault)
	config := mgr.GetFixtures("idvault.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idvault_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idVault)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idVault, createDoesIdVaultExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idvault", ri)),
					resource.TestCheckResourceAttr(resourceName, "connector", "connector-default"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idvault", ri)),
					resource.TestCheckResourceAttr(resourceName, "connector", "connector-2"),
				),
			},
		},
	})
}

func createDoesIdVaultExist() func(string) (bool, error) {

	return func(id string) (bool, error) {
		return false, nil
	}
}
