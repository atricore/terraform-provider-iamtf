package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPhp_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(php)
	config := mgr.GetFixtures("php.tf", ri, t)
	updatedConfig := mgr.GetFixtures("php_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", php)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(php, createDoesPhpExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("php", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Php Php-Exect-Env"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("php", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Php Php-Exect-Env-updated"),
				),
			},
		},
	})
}

func createDoesPhpExist() func(string) (bool, error) {

	return func(id string) (bool, error) {
		return false, nil
	}
}
