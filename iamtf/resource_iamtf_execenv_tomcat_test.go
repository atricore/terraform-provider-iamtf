package iamtf

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJossoTomcat_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(tomcat)
	config := mgr.GetFixtures("tomcat.tf", ri, t)
	updatedConfig := mgr.GetFixtures("tomcat_update.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", tomcat)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(tomcat, createDoesTomcatExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("tc", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Tomcat Tomcat-Exect-Env"),
					resource.TestCheckResourceAttr(resourceName, "version", "8.5"),
					resource.TestCheckResourceAttr(resourceName, "activation_install_samples", strconv.FormatBool(true)),
					resource.TestCheckResourceAttr(resourceName, "activation_path", "/opt/atricore/josso-ee-2/Tomcat-Exect-Env"),
					resource.TestCheckResourceAttr(resourceName, "activation_remote_target", "http://remote-josso:8081"),
					resource.TestCheckResourceAttr(resourceName, "activation_override_setup", strconv.FormatBool(true)),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("tc", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "Tomcat Tomcat-Exect-Env-updated"),
					resource.TestCheckResourceAttr(resourceName, "version", "9"),
					resource.TestCheckResourceAttr(resourceName, "activation_install_samples", strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(resourceName, "activation_path", "/opt/atricore/josso-ee-2/Tomcat-Exect-Env/updated"),
					resource.TestCheckResourceAttr(resourceName, "activation_remote_target", "http://remote-josso:8082"),
					resource.TestCheckResourceAttr(resourceName, "activation_override_setup", strconv.FormatBool(false)),
				),
			},
		},
	})
}

func createDoesTomcatExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
