package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateTacacsplusAuthservice validates the TacacsplusAuthservice configuration.
func ValidateTacacsplusAuthservice(ctx context.Context, data TacacsplusAuthserviceModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSTacacsplusAuthserviceModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateTacacsplusAuthserviceNIOSConfig(ctx, nios, resp)
	}
}

func validateTacacsplusAuthserviceNIOSConfig(ctx context.Context, m *NIOSTacacsplusAuthserviceModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenTacacsplusAuthserviceNIOS(ctx context.Context, planned, flattened *NIOSTacacsplusAuthserviceModel, diags *diag.Diagnostics) {
	if flattened == nil {
		return
	}

	if flattened.Comment.IsNull() {
		flattened.Comment = types.StringValue("")
	}

	if planned == nil {
		return
	}

	// NIOS does not echo shared_secret back in API responses; copy it from plan so state stays consistent.
	if !planned.Servers.IsUnknown() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.Servers, flattened.Servers, "shared_secret"); !d.HasError() {
			if resultList, ok := result.(basetypes.ListValue); ok {
				flattened.Servers = resultList
			}
		}
	}
}
