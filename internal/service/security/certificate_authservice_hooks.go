package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateCertificateAuthservice validates the CertificateAuthservice configuration.
func ValidateCertificateAuthservice(ctx context.Context, data CertificateAuthserviceModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSCertificateAuthserviceModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateCertificateAuthserviceNIOSConfig(ctx, nios, resp)
	}
}

func validateCertificateAuthserviceNIOSConfig(ctx context.Context, m *NIOSCertificateAuthserviceModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	ocspCheck := m.OcspCheck
	ocspResponders := m.OcspResponders

	// Guard unknown values — plan-time values from variables are unknown; erroring on unknown rejects valid configs.
	if !ocspCheck.IsUnknown() && !ocspResponders.IsUnknown() {
		isManualCheck := ocspCheck.IsNull() || ocspCheck.ValueString() == "MANUAL" || ocspCheck.ValueString() == "AIA_AND_MANUAL"
		if isManualCheck && ocspResponders.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("ocsp_responders"),
				"Invalid Configuration",
				"At least one `ocsp_responders` must be specified when `ocsp_check` is set to `MANUAL` or `AIA_AND_MANUAL`, else set the ocsp_check to 'DISABLED'.",
			)
		}
	}

	if m.EnableRemoteLookup.IsUnknown() {
		return
	}
	if !m.EnableRemoteLookup.IsNull() && m.EnableRemoteLookup.ValueBool() {
		missingService := m.RemoteLookupService.IsNull()
		missingUsername := m.RemoteLookupUsername.IsNull()
		missingPassword := m.RemoteLookupPassword.IsNull()
		if missingService || missingUsername || missingPassword {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("enable_remote_lookup"),
				"Invalid Configuration",
				"When `enable_remote_lookup` is set to `true`, all fields `remote_lookup_service`, `remote_lookup_username`, and `remote_lookup_password` must be provided.",
			)
		}

		if !m.EnablePasswordRequest.IsUnknown() && (m.EnablePasswordRequest.IsNull() || m.EnablePasswordRequest.ValueBool()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("enable_password_request"),
				"Invalid Configuration",
				"When `enable_remote_lookup` is set to `true`, `enable_password_request` must be set to `false`.",
			)
		}

		if !m.UserMatchType.IsNull() && !m.UserMatchType.IsUnknown() && m.UserMatchType.ValueString() != "AUTO_MATCH" {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("user_match_type"),
				"Invalid Configuration",
				"`user_match_type` must be set to \"AUTO_MATCH\" to use remote lookup services.",
			)
		}
	}
}

// PostFlattenCertificateAuthserviceNIOS preserves write-only and oneOf fields from plan to state.
// remote_lookup_service is a oneOf SDK union type that codegen cannot round-trip; skip_expand/
// skip_flatten remove the broken generated lines.
// remote_lookup_password is write-only: NIOS never echoes it back in GET responses.
// Both fields must be copied from the planned model to avoid perpetual diffs. When the planned
// value is unknown (Optional+Computed and user did not set it), resolve to null so that Terraform
// receives a known value after apply.
func PostFlattenCertificateAuthserviceNIOS(ctx context.Context, planned, flattened *NIOSCertificateAuthserviceModel, diags *diag.Diagnostics) {
	if planned != nil && !planned.RemoteLookupService.IsUnknown() {
		flattened.RemoteLookupService = planned.RemoteLookupService
	} else {
		flattened.RemoteLookupService = types.StringNull()
	}
	if planned != nil && !planned.RemoteLookupPassword.IsUnknown() {
		flattened.RemoteLookupPassword = planned.RemoteLookupPassword
	} else {
		flattened.RemoteLookupPassword = types.StringNull()
	}
}
