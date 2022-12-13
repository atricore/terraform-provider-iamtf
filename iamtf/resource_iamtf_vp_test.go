package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccVP_crud0(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(vp)
	config := mgr.GetFixtures("vp.tf", ri, t)
	updatedConfig := mgr.GetFixtures("vp_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", vp)

	// TODO : Validate other fields ?
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(vp, createDoesVPExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("vp", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("vp", ri)),
				),
			},
		},
	})
}

func TestAccVP_crud1(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(vp)
	config := mgr.GetFixtures("vp_1.tf", ri, t)
	updatedConfig := mgr.GetFixtures("vp_1_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", vp)

	// TODO : Validate other fields ?
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(vp, createDoesVPExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("vp_1", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("vp_1", ri)),
				),
			},
		},
	})
}

func TestAccVP_Attrs1_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(vp)
	config := mgr.GetFixtures("vp_attrs_1.tf", ri, t)
	updatedConfig := mgr.GetFixtures("vp_attrs_1_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", vp)

	// TODO : Validate other fields ?
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(vp, createDoesVPExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("vp", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("vp", ri)),
				),
			},
		},
	})
}

func createDoesVPExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
