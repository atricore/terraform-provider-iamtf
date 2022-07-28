package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWebLogic_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(weblogic)
	config := mgr.GetFixtures("app_weblogic.tf", ri, t)
	updatedConfig := mgr.GetFixtures("app_weblogic_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", weblogic)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(weblogic, createDoesWebLogicExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("wl", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("wl", ri)),
				),
			},
		},
	})
}

func createDoesWebLogicExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
