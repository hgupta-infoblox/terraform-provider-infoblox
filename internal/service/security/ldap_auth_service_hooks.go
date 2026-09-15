package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateLdapAuthService validates the LdapAuthService configuration.
func ValidateLdapAuthService(ctx context.Context, data LdapAuthServiceModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSLdapAuthServiceModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateLdapAuthServiceNIOSConfig(ctx, nios, resp)
	}
}

func validateLdapAuthServiceNIOSConfig(ctx context.Context, m *NIOSLdapAuthServiceModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenLdapAuthServiceNIOS copies bind_password from the planned servers into the
// flattened response, because the API never echoes that field back.
func PostFlattenLdapAuthServiceNIOS(ctx context.Context, planned, flattened *NIOSLdapAuthServiceModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}
	if !planned.Servers.IsUnknown() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.Servers, flattened.Servers, "bind_password"); !d.HasError() {
			if resultList, ok := result.(basetypes.ListValue); ok {
				flattened.Servers = resultList
			}
		}
	}
}
