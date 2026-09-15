package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RadiusAuthserviceServersModel is the Terraform model for RadiusAuthserviceServers
type RadiusAuthserviceServersModel struct {
	AcctPort      types.Int64  `tfsdk:"acct_port"`
	AuthPort      types.Int64  `tfsdk:"auth_port"`
	AuthType      types.String `tfsdk:"auth_type"`
	Comment       types.String `tfsdk:"comment"`
	Disable       types.Bool   `tfsdk:"disable"`
	Address       types.String `tfsdk:"address"`
	SharedSecret  types.String `tfsdk:"shared_secret"`
	UseAccounting types.Bool   `tfsdk:"use_accounting"`
	UseMgmtPort   types.Bool   `tfsdk:"use_mgmt_port"`
}

// RadiusAuthserviceServersAttrTypes contains the attribute types for RadiusAuthserviceServersModel
var RadiusAuthserviceServersAttrTypes = map[string]attr.Type{
	"acct_port":      types.Int64Type,
	"auth_port":      types.Int64Type,
	"auth_type":      types.StringType,
	"comment":        types.StringType,
	"disable":        types.BoolType,
	"address":        types.StringType,
	"shared_secret":  types.StringType,
	"use_accounting": types.BoolType,
	"use_mgmt_port":  types.BoolType,
}

// RadiusAuthserviceServersResourceSchemaAttributes contains the schema attributes for RadiusAuthserviceServersModel
var RadiusAuthserviceServersResourceSchemaAttributes = map[string]schema.Attribute{
	"acct_port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1813),
		MarkdownDescription: "The accounting port.",
	},
	"auth_port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1812),
		MarkdownDescription: "The authorization port.",
	},
	"auth_type": schema.StringAttribute{
		Default: stringdefault.StaticString("PAP"),
		Validators: []validator.String{
			stringvalidator.OneOf("CHAP", "PAP"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The authentication protocol.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The RADIUS descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the RADIUS server is disabled.",
	},
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The FQDN or the IP address of the RADIUS server that is used for authentication.",
	},
	"shared_secret": schema.StringAttribute{
		Sensitive: true,
		Required:  true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The shared secret that the NIOS appliance and the RADIUS server use to encrypt and decrypt their messages.",
	},
	"use_accounting": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether RADIUS accounting is enabled.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether connection via the management interface is allowed.",
	},
}

// ExpandRadiusAuthserviceServers converts a Terraform Object to SDK type
func ExpandRadiusAuthserviceServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossecurity.RadiusAuthserviceServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RadiusAuthserviceServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RadiusAuthserviceServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossecurity.RadiusAuthserviceServers {
	if m == nil {
		return nil
	}
	to := &niossecurity.RadiusAuthserviceServers{
		AcctPort:      flex.ExpandInt64Pointer(m.AcctPort),
		AuthPort:      flex.ExpandInt64Pointer(m.AuthPort),
		AuthType:      flex.ExpandStringPointerNullAsEmpty(m.AuthType),
		Comment:       flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:       flex.ExpandBoolPointer(m.Disable),
		Address:       flex.ExpandStringPointerNullAsEmpty(m.Address),
		SharedSecret:  flex.ExpandStringPointerNullAsEmpty(m.SharedSecret),
		UseAccounting: flex.ExpandBoolPointer(m.UseAccounting),
		UseMgmtPort:   flex.ExpandBoolPointer(m.UseMgmtPort),
	}
	return to
}

// FlattenRadiusAuthserviceServers converts an SDK type to Terraform Object
func FlattenRadiusAuthserviceServers(ctx context.Context, from *niossecurity.RadiusAuthserviceServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RadiusAuthserviceServersAttrTypes)
	}
	m := &RadiusAuthserviceServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RadiusAuthserviceServersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RadiusAuthserviceServersModel) Flatten(ctx context.Context, from *niossecurity.RadiusAuthserviceServers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AcctPort = flex.FlattenInt64Pointer(from.AcctPort)
	m.AuthPort = flex.FlattenInt64Pointer(from.AuthPort)
	m.AuthType = flex.FlattenStringPointerEmptyAsNull(from.AuthType)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.SharedSecret = flex.FlattenStringPointerEmptyAsNull(from.SharedSecret)
	m.UseAccounting = flex.FlattenBoolPointer(from.UseAccounting)
	m.UseMgmtPort = flex.FlattenBoolPointer(from.UseMgmtPort)
}
