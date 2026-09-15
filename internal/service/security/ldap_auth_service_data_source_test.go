package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccLdapAuthServiceDataSource(t *testing.T) {
	dsType := "infoblox_ldap_auth_service"
	resourceType := "infoblox_ldap_auth_service"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckLdapAuthServiceExistsNIOS,
			Destroy: testAccCheckLdapAuthServiceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "security/ldap_auth_service/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
