package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRadiusAuthserviceDataSource(t *testing.T) {
	dsType := "infoblox_radius_authservice"
	resourceType := "infoblox_radius_authservice"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRadiusAuthserviceExistsNIOS,
			Destroy: testAccCheckRadiusAuthserviceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "security/radius_authservice/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
