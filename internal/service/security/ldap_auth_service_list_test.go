package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccLdapAuthServiceList(t *testing.T) {
	resourceType := "infoblox_ldap_auth_service"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckLdapAuthServiceExistsNIOS,
			Destroy: testAccCheckLdapAuthServiceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "security/ldap_auth_service/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
