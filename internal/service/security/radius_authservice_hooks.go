package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateRadiusAuthservice validates the RadiusAuthservice configuration.
func ValidateRadiusAuthservice(ctx context.Context, data RadiusAuthserviceModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRadiusAuthserviceModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRadiusAuthserviceNIOSConfig(ctx, nios, resp)
	}
}

func validateRadiusAuthserviceNIOSConfig(ctx context.Context, m *NIOSRadiusAuthserviceModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenRadiusAuthserviceNIOS copies shared_secret from the prior state/plan into
// the freshly-flattened model. The NIOS API never echoes shared_secret back; without
// this copy, every read would null the field and trigger a spurious diff.
func PostFlattenRadiusAuthserviceNIOS(ctx context.Context, planned, flattened *NIOSRadiusAuthserviceModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}
	if !planned.Servers.IsUnknown() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.Servers, flattened.Servers, "shared_secret"); !d.HasError() {
			if resultList, ok := result.(basetypes.ListValue); ok {
				flattened.Servers = resultList
			}
		}
	}
}
