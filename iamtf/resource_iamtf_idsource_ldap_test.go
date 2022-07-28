package iamtf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdSourceLdap_crud(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(idSourceLdap)
	config := mgr.GetFixtures("idsource_ldap.tf", ri, t)
	updatedConfig := mgr.GetFixtures("idsource_ldap_updated.tf", ri, t)
	resourceName := fmt.Sprintf("%s.test", idSourceLdap)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testaccPreCheck(t) },
		ProviderFactories: testaccProvidersFactories,
		CheckDestroy:      createCheckResourceDestroy(idSourceLdap, createDoesIdSourceLdapExist()),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idvault", ri)),
					resource.TestCheckResourceAttr(resourceName, "provider_url", "ldap://localhost:10389"),
					resource.TestCheckResourceAttr(resourceName, "username", "uid=admin,ou=system"),
					resource.TestCheckResourceAttr(resourceName, "password", "secret"),
					resource.TestCheckResourceAttr(resourceName, "users_ctx_dn", "dc=example,dc=com,ou=IAM,ou=People"),
					resource.TestCheckResourceAttr(resourceName, "userid_attr", "uid"),
					resource.TestCheckResourceAttr(resourceName, "groups_ctx_dn", "dc=example,dc=com,ou=IAM,ou=Groups"),
					resource.TestCheckResourceAttr(resourceName, "groupid_attr", "cn"),
					resource.TestCheckResourceAttr(resourceName, "groupmember_attr", "uniquemember"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", buildResourceNameForPrefix("idvault", ri)),
					resource.TestCheckResourceAttr(resourceName, "provider_url", "ldaps://localhost:10636"),
					resource.TestCheckResourceAttr(resourceName, "username", "uid=manager,ou=system"),
					resource.TestCheckResourceAttr(resourceName, "password", "changeme"),
					resource.TestCheckResourceAttr(resourceName, "users_ctx_dn", "dc=example1,dc=com,ou=IAM,ou=People"),
					resource.TestCheckResourceAttr(resourceName, "userid_attr", "uid1"),
					resource.TestCheckResourceAttr(resourceName, "groups_ctx_dn", "dc=example1,dc=com,ou=IAM,ou=Groups"),
					resource.TestCheckResourceAttr(resourceName, "groupid_attr", "cn1"),
					resource.TestCheckResourceAttr(resourceName, "groupmember_attr", "uniquemember1"),
				),
			},
		},
	})
}

func createDoesIdSourceLdapExist() func(string) (bool, error) {
	// TODO : infer appliance name and lookup for resource
	return func(id string) (bool, error) {
		return false, nil
	}
}
