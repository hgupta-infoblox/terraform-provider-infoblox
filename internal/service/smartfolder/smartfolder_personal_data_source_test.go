package smartfolder_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSmartfolderPersonalDataSource(t *testing.T) {
	dsType := "infoblox_smartfolder_personal"
	resourceType := "infoblox_smartfolder_personal"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSmartfolderPersonalExistsNIOS,
			Destroy: testAccCheckSmartfolderPersonalDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "smartfolder/smartfolder_personal/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
