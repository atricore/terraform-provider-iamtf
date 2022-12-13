package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdFacebook_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idpFacebook)
	config := mgr.GetFixtures("idp_facebook.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_facebook_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idpFacebook)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idpFacebook, createDoesIdFacebookExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp_fb", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp_fb", ri)),
				),
			},
		},
	})
}

func createDoesIdFacebookExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
