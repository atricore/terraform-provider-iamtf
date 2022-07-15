package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJossoIdGoogle_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(google)
	config := mgr.GetFixtures("idgoogle.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idgoogle_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", google)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(google, createDoesIdGoogleExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idgoogle", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idgoogle", ri)),
				),
			},
		},
	})
}

func createDoesIdGoogleExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
