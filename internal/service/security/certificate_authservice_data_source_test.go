package security_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCertificateAuthserviceDataSource(t *testing.T) {
	dsType := "infoblox_certificate_authservice"
	resourceType := "infoblox_certificate_authservice"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckCertificateAuthserviceExistsNIOS,
			Destroy: testAccCheckCertificateAuthserviceDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "security/certificate_authservice/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
