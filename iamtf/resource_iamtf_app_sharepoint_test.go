package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSharePoint_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(spoint)
	config := mgr.GetFixtures("app_sharepoint.tf", ri, t)
	updatedConfig := mgr.GetFixtures("app_sharepoint_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", spoint)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(spoint, createDoesSharePointExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-sharePoint", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-sharePoint", ri)),
				),
			},
		},
	})
}

func createDoesSharePointExist() func(string) (bool, error) {

	return func(id string) (bool, error) {
		return false, nil
	}
}
