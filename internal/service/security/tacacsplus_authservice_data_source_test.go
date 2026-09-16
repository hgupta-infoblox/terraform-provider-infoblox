package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccTacacsplusAuthserviceDataSource(t *testing.T) {
	dsType := "infoblox_tacacsplus_authservice"
	resourceType := "infoblox_tacacsplus_authservice"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckTacacsplusAuthserviceExistsNIOS,
			Destroy: testAccCheckTacacsplusAuthserviceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "security/tacacsplus_authservice/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
