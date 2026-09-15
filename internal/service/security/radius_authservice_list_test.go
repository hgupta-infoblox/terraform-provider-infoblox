package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRadiusAuthserviceList(t *testing.T) {
	resourceType := "infoblox_radius_authservice"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRadiusAuthserviceExistsNIOS,
			Destroy: testAccCheckRadiusAuthserviceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "security/radius_authservice/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
