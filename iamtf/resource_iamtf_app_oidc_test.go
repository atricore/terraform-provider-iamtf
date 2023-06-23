package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOidcRp_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(oidcRp)
	config := mgr.GetFixtures("app_oidc.tf", ri, t)
	updatedConfig := mgr.GetFixtures("app_oidc_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", oidcRp)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(oidcRp, createDoesOidcRpExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-oidc", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-oidc", ri)),
				),
			},
		},
	})
}

func TestAccOidcRp_min_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(oidcRp)
	config := mgr.GetFixtures("app_oidc_min.tf", ri, t)
	updatedConfig := mgr.GetFixtures("app_oidc_min_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", oidcRp)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(oidcRp, createDoesOidcRpExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-oidc", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-oidc", ri)),
				),
			},
		},
	})
}

func createDoesOidcRpExist() func(string) (bool, error) {

	return func(id string) (bool, error) {
		return false, nil
	}
}
