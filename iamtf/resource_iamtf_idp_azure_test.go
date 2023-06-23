package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdAzure_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idpAzure)
	config := mgr.GetFixtures("idp_azure.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_azure_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idpAzure)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idpAzure, createDoesIdAzureExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp-az", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp-az", ri)),
				),
			},
		},
	})
}

func createDoesIdAzureExist() func(string) (bool, error) {

	return func(id string) (bool, error) {
		return false, nil
	}
}
