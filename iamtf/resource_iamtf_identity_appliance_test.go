package iamtf

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJossoIdentityAppliance_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(identityAppliance)
	config := mgr.GetFixtures("ida.tf", ri, t)
	updatedConfig := mgr.GetFixtures("ida_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", identityAppliance)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(identityAppliance, createDoesApplianceExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceName(ri)),
					resource.TestCheckResourceAttr(resourceName, "namespace", buildResourceNamespace(ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Appliance #"+strconv.Itoa(ri)),
					resource.TestCheckResourceAttr(resourceName, "location", testResourceURL),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceName(ri)),
					resource.TestCheckResourceAttr(resourceName, "namespace", buildResourceNamespace(ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Appliance #"+strconv.Itoa(ri)),
					resource.TestCheckResourceAttr(resourceName, "location", testResourceURL),
				),
			},
		},
	})
}

func createDoesApplianceExist() func(string) (bool, error) {
	return func(id string) (bool, error) {
		res, err := getJossoClient(testaccProvider.Meta()).GetAppliance(id)
		return res.Name != nil, err
	}
}

func deleteIdentityAppliances(client *cli.IdbusApiClient) error {

	client.Logger().Debug("deleting ALL appliances!")
	idams, err := client.GetAppliances()

	if err != nil {
		return fmt.Errorf("error getting appliances for deletion: %v", err)
	}

	for _, ida := range idams {

		if strings.HasPrefix(*ida.Name, testResourcePrefix) {
			client.Logger().Infof("deleting %s", *ida.Name)
			if _, err = client.DeleteAppliance(*ida.Name); err != nil {
				return fmt.Errorf("error deleting appliances %s: %v", *ida.Name, err)
			}
		}
	}
	return nil
}
