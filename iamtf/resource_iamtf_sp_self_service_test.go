package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSelfService_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(selfService)
	config := mgr.GetFixtures("self_service.tf", ri, t)
	updatedConfig := mgr.GetFixtures("self_service_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", selfService)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(selfService, createDoesSelfServiceExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("self-svc-rs", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("self-svc-rs", ri)),
				),
			},
		},
	})
}

func createDoesSelfServiceExist() func(string) (bool, error) {

	return func(id string) (bool, error) {
		return false, nil
	}
}
