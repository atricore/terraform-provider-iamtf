package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJossoIss_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(iss)
	config := mgr.GetFixtures("iss.tf", ri, t)
	updatedConfig := mgr.GetFixtures("iss_update.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", iss)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(iss, createDoesIssExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("iss", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Iss Iss-Exect-Env"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("iss", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Iss Iss-Exect-Env-updated"),
				),
			},
		},
	})
}

func createDoesIssExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
