package smartfolder_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSmartfolderPersonalList(t *testing.T) {
	resourceType := "infoblox_smartfolder_personal"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSmartfolderPersonalExistsNIOS,
			Destroy: testAccCheckSmartfolderPersonalDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "smartfolder/smartfolder_personal/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
