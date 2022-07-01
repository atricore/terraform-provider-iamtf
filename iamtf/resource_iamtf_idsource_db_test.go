package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDbIdSource_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(dbidsource)
	config := mgr.GetFixtures("idsource_db.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idsource_db_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", dbidsource)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(dbidsource, createDoesDbIdSourceExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("dbid", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("dbid", ri)),
				),
			},
		},
	})
}

func createDoesDbIdSourceExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
