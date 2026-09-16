package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// TacacsplusAuthserviceServersModel is the Terraform model for TacacsplusAuthserviceServers
type TacacsplusAuthserviceServersModel struct {
	Address       types.String `tfsdk:"address"`
	Port          types.Int64  `tfsdk:"port"`
	SharedSecret  types.String `tfsdk:"shared_secret"`
	AuthType      types.String `tfsdk:"auth_type"`
	Comment       types.String `tfsdk:"comment"`
	Disable       types.Bool   `tfsdk:"disable"`
	UseMgmtPort   types.Bool   `tfsdk:"use_mgmt_port"`
	UseAccounting types.Bool   `tfsdk:"use_accounting"`
}

// TacacsplusAuthserviceServersAttrTypes contains the attribute types for TacacsplusAuthserviceServersModel
var TacacsplusAuthserviceServersAttrTypes = map[string]attr.Type{
	"address":        types.StringType,
	"port":           types.Int64Type,
	"shared_secret":  types.StringType,
	"auth_type":      types.StringType,
	"comment":        types.StringType,
	"disable":        types.BoolType,
	"use_mgmt_port":  types.BoolType,
	"use_accounting": types.BoolType,
}

// TacacsplusAuthserviceServersResourceSchemaAttributes contains the schema attributes for TacacsplusAuthserviceServersModel
var TacacsplusAuthserviceServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The valid IP address or FQDN of the TACACS+ server.",
	},
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(49),
		Validators: []validator.Int64{
			int64validator.Between(1, 65535),
		},
		MarkdownDescription: "The TACACS+ server port.",
	},
	"shared_secret": schema.StringAttribute{
		Sensitive: true,
		Required:  true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The secret key with which to connect to the TACACS+ server.",
	},
	"auth_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ASCII", "PAP", "CHAP"),
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
		MarkdownDescription: "The TACACS+ descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the TACACS+ server is disabled.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the TACACS+ server is connected via the management interface.",
	},
	"use_accounting": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the TACACS+ accounting server is used.",
	},
}

// ExpandTacacsplusAuthserviceServers converts a Terraform Object to SDK type
func ExpandTacacsplusAuthserviceServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossecurity.TacacsplusAuthserviceServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TacacsplusAuthserviceServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TacacsplusAuthserviceServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossecurity.TacacsplusAuthserviceServers {
	if m == nil {
		return nil
	}
	to := &niossecurity.TacacsplusAuthserviceServers{
		Address:       flex.ExpandStringPointer(m.Address),
		Port:          flex.ExpandInt64Pointer(m.Port),
		SharedSecret:  flex.ExpandStringPointerNullAsEmpty(m.SharedSecret),
		AuthType:      flex.ExpandStringPointer(m.AuthType),
		Comment:       flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:       flex.ExpandBoolPointer(m.Disable),
		UseMgmtPort:   flex.ExpandBoolPointer(m.UseMgmtPort),
		UseAccounting: flex.ExpandBoolPointer(m.UseAccounting),
	}
	return to
}

// FlattenTacacsplusAuthserviceServers converts an SDK type to Terraform Object
func FlattenTacacsplusAuthserviceServers(ctx context.Context, from *niossecurity.TacacsplusAuthserviceServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TacacsplusAuthserviceServersAttrTypes)
	}
	m := &TacacsplusAuthserviceServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TacacsplusAuthserviceServersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TacacsplusAuthserviceServersModel) Flatten(ctx context.Context, from *niossecurity.TacacsplusAuthserviceServers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.SharedSecret = flex.FlattenStringPointerEmptyAsNull(from.SharedSecret)
	m.AuthType = flex.FlattenStringPointerEmptyAsNull(from.AuthType)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.UseMgmtPort = flex.FlattenBoolPointer(from.UseMgmtPort)
	m.UseAccounting = flex.FlattenBoolPointer(from.UseAccounting)
}
