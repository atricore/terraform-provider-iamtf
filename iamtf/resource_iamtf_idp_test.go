package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdP_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idp)
	config := mgr.GetFixtures("idp.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idp)

	// TODO : Validate other fields ?
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idp, createDoesIdPExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp", ri)),
					resource.TestCheckResourceAttr(resourceName, "branding", "josso25-branding"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp", ri)),
					resource.TestCheckResourceAttr(resourceName, "branding", "josso2-default-branding"),
				),
			},
		},
	})
}

func TestAccIdP_crud1(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idp)
	config := mgr.GetFixtures("idp_1.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_1_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idp)

	// TODO : Validate other fields ?
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idp, createDoesIdPExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp_1", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp_1", ri)),
				),
			},
		},
	})
}

func TestAccIdP_AuthnWia_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idp)
	config := mgr.GetFixtures("idp_authn_wia.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_authn_wia_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idp)

	// TODO : Validate other fields ?
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idp, createDoesIdPExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp", ri)),
				),
			},
		},
	})
}

func TestAccIdP_Attrs1_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idp)
	config := mgr.GetFixtures("idp_attrs_1.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idp_attrs_1_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idp)

	// TODO : Validate other fields ?
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idp, createDoesIdPExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp", ri)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idp", ri)),
				),
			},
		},
	})
}

func createDoesIdPExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
