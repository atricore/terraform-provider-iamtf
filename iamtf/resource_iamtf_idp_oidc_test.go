package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdpOidc_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idpOidc)
	config := mgr.GetFixtures("idp_oidc.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_oidc_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idpOidc)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idpOidc, createDoesIdpOidcExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp-oidc", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp-oidc", ri)),
				),
			},
		},
	})
}

func createDoesIdpOidcExist() func(string) (bool, error) {
	return func(id string) (bool, error) {
		return false, nil
	}
}
