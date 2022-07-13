package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJossoIdFacebook_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idFacebook)
	config := mgr.GetFixtures("idfacebook.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idfacebook_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idFacebook)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idFacebook, createDoesIdFacebookExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idfacebook", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idfacebook", ri)),
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
