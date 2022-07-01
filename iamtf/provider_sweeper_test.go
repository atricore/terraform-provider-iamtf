package iamtf

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var testResourcePrefix = "testacc"
var testResourceURL = "http://localhost:8081"

// TestMain overridden main testing function. Package level BeforeAll and AfterAll.
// It also delineates between acceptance tests and unit tests
func TestMain(m *testing.M) {
	// Acceptance test sweepers necessary to prevent dangling resources
	setupSweeper(identityAppliance, deleteIdentityAppliances)

	// add zones sweeper
	resource.TestMain(m)
}

// Sets up sweeper to clean up dangling resources
func setupSweeper(resourceType string, del func(*cli.IdbusApiClient) error) {
	resource.AddTestSweepers(resourceType, &resource.Sweeper{
		Name: resourceType,
		F: func(_ string) error {
			client, err := sharedClient()
			if err != nil {
				return err
			}
			return del(client)
		},
	})
}

// Builds test specific resource name
func buildResourceFQN(resourceType string, testID int) string {
	return resourceType + "." + buildResourceName(testID)
}

func buildResourceNameForPrefix(prefix string, testID int) string {
	return prefix + "-" + strconv.Itoa(testID)
}

func buildResourceName(testID int) string {
	return buildResourceNameForPrefix(testResourcePrefix, testID)
}

func buildResourceNamespace(testID int) string {
	return "com.atricore.idbus.test." + strings.ToLower(testResourcePrefix) + "." + strconv.Itoa(testID)
}

func buildResourceLocation(baseURL string, testID int) string {
	return baseURL + "/" + strings.ToUpper(buildResourceName(testID))
}

// sharedClient returns a common JOSSO Client for sweepers
func sharedClient() (*cli.IdbusApiClient, error) {
	err := accPreCheck()
	if err != nil {
		return nil, err
	}
	c, err := jossoConfig()
	if err != nil {
		return nil, fmt.Errorf("cannot create API client %v", err)
	}

	return c.apiClient, nil
}
