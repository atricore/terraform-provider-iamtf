package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppAgentRe_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(josso1Rs)
	config := mgr.GetFixtures("app_agent.tf", ri, t)
	updatedConfig := mgr.GetFixtures("app_agent_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", josso1Rs)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(josso1Rs, createDoesJosso1ReExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "app_slo_location", fmt.Sprintf("http://myapp-%d:8080/partnerapp/slo", ri)),
					resource.TestCheckResourceAttr(resourceName, "app_location", fmt.Sprintf("http://myapp-%d:8080/partnerapp", ri)),
					// TODO user array? resource.TestCheckResourceAttr(resourceName, "ignored_web_resources", ("*.ico")),
					resource.TestCheckResourceAttr(resourceName, "default_resource", fmt.Sprintf("http://myapp-%d:8080/partnerapp/home", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "desc app-a"),
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-agent", ri)),
					resource.TestCheckResourceAttr(resourceName, "dashboard_url", fmt.Sprintf("http://myapp-%d:8080/partnerapp/dashboard", ri)),
					resource.TestCheckResourceAttr(resourceName, "error_binding", "JSON"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "app_slo_location", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp/slo", ri)),
					resource.TestCheckResourceAttr(resourceName, "app_location", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp", ri)),
					// TODO user array? : resource.TestCheckResourceAttr(resourceName, "ignored_web_resources", ("*.jpg")),
					resource.TestCheckResourceAttr(resourceName, "default_resource", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp/home", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "desc app-a updated"),
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-agent", ri)),
					resource.TestCheckResourceAttr(resourceName, "dashboard_url", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp/dashboard", ri)),
					resource.TestCheckResourceAttr(resourceName, "error_binding", "ARTIFACT"),
				),
			},
		},
	})
}

func TestAcc1Re_crud1(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(josso1Rs)
	config := mgr.GetFixtures("app_agent_1.tf", ri, t)
	updatedConfig := mgr.GetFixtures("app_agent_1_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", josso1Rs)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(josso1Rs, createDoesJosso1ReExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "app_slo_location", fmt.Sprintf("http://myapp-%d:8080/partnerapp/slo", ri)),
					resource.TestCheckResourceAttr(resourceName, "app_location", fmt.Sprintf("http://myapp-%d:8080/partnerapp", ri)),
					// TODO user array? resource.TestCheckResourceAttr(resourceName, "ignored_web_resources", ("*.ico")),
					resource.TestCheckResourceAttr(resourceName, "default_resource", fmt.Sprintf("http://myapp-%d:8080/partnerapp/home", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "desc app-a"),
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-agent", ri)),
					resource.TestCheckResourceAttr(resourceName, "dashboard_url", fmt.Sprintf("http://myapp-%d:8080/partnerapp/dashboard", ri)),
					resource.TestCheckResourceAttr(resourceName, "error_binding", "JSON"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "app_slo_location", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp/slo", ri)),
					resource.TestCheckResourceAttr(resourceName, "app_location", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp", ri)),
					// TODO user array? : resource.TestCheckResourceAttr(resourceName, "ignored_web_resources", ("*.jpg")),
					resource.TestCheckResourceAttr(resourceName, "default_resource", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp/home", ri)),
					resource.TestCheckResourceAttr(resourceName, "description", "desc app-a updated"),
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("app-agent", ri)),
					resource.TestCheckResourceAttr(resourceName, "dashboard_url", fmt.Sprintf("http://myapp-updated-%d:8080/partnerapp/dashboard", ri)),
					resource.TestCheckResourceAttr(resourceName, "error_binding", "ARTIFACT"),
				),
			},
		},
	})
}

func createDoesJosso1ReExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
