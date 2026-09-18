package smartfolder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSmartfolderPersonal validates the SmartfolderPersonal configuration.
func ValidateSmartfolderPersonal(ctx context.Context, data SmartfolderPersonalModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSmartfolderPersonalModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSmartfolderPersonalNIOSConfig(ctx, nios, resp)
	}
}

func validateSmartfolderPersonalNIOSConfig(ctx context.Context, m *NIOSSmartfolderPersonalModel, resp *resource.ValidateConfigResponse) {
}
