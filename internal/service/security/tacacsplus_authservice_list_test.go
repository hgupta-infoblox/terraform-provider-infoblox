package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccTacacsplusAuthserviceList(t *testing.T) {
	resourceType := "infoblox_tacacsplus_authservice"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckTacacsplusAuthserviceExistsNIOS,
			Destroy: testAccCheckTacacsplusAuthserviceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "security/tacacsplus_authservice/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
