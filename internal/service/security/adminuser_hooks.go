package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAdminuser validates the Adminuser configuration.
func ValidateAdminuser(ctx context.Context, data AdminuserModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSAdminuserModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateAdminuserNIOSConfig(ctx, nios, resp)
	}
}

func validateAdminuserNIOSConfig(ctx context.Context, m *NIOSAdminuserModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenAdminuserNIOS(ctx context.Context, planned, flattened *NIOSAdminuserModel, diags *diag.Diagnostics) {
	// password is write-only: NIOS never returns it. Copy from plan so state stays
	// consistent with config; on import (planned == nil) leave it null.
	if planned != nil {
		flattened.Password = planned.Password
	} else {
		flattened.Password = types.StringNull()
	}

	// NIOS may return ssh_keys even when use_ssh_keys=false and the user did not
	// configure the field. Restore the planned null to prevent a spurious diff.
	if planned != nil && planned.SshKeys.IsNull() {
		flattened.SshKeys = planned.SshKeys
	}
}
