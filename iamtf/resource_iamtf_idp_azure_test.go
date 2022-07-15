package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJossoIdAzure_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(azure)
	config := mgr.GetFixtures("idazure.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idazure_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", azure)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(azure, createDoesIdAzureExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idazure", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idazure", ri)),
				),
			},
		},
	})
}

func createDoesIdAzureExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
