package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCertificateAuthserviceList(t *testing.T) {
	resourceType := "infoblox_certificate_authservice"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckCertificateAuthserviceExistsNIOS,
			Destroy: testAccCheckCertificateAuthserviceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "security/certificate_authservice/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
