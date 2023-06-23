package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdGoogle_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idpGoogle)
	config := mgr.GetFixtures("idp_google.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_google_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idpGoogle)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idpGoogle, createDoesIdGoogleExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp-google", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp-google", ri)),
				),
			},
		},
	})
}

func createDoesIdGoogleExist() func(string) (bool, error) {

	return func(id string) (bool, error) {
		return false, nil
	}
}
