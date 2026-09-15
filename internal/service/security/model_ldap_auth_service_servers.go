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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// LdapAuthServiceServersModel is the Terraform model for LdapAuthServiceServers
type LdapAuthServiceServersModel struct {
	Address            types.String `tfsdk:"address"`
	AuthenticationType types.String `tfsdk:"authentication_type"`
	BaseDn             types.String `tfsdk:"base_dn"`
	BindPassword       types.String `tfsdk:"bind_password"`
	BindUserDn         types.String `tfsdk:"bind_user_dn"`
	Comment            types.String `tfsdk:"comment"`
	Disable            types.Bool   `tfsdk:"disable"`
	Encryption         types.String `tfsdk:"encryption"`
	Port               types.Int64  `tfsdk:"port"`
	UseMgmtPort        types.Bool   `tfsdk:"use_mgmt_port"`
	Version            types.String `tfsdk:"version"`
}

// LdapAuthServiceServersAttrTypes contains the attribute types for LdapAuthServiceServersModel
var LdapAuthServiceServersAttrTypes = map[string]attr.Type{
	"address":             types.StringType,
	"authentication_type": types.StringType,
	"base_dn":             types.StringType,
	"bind_password":       types.StringType,
	"bind_user_dn":        types.StringType,
	"comment":             types.StringType,
	"disable":             types.BoolType,
	"encryption":          types.StringType,
	"port":                types.Int64Type,
	"use_mgmt_port":       types.BoolType,
	"version":             types.StringType,
}

// LdapAuthServiceServersResourceSchemaAttributes contains the schema attributes for LdapAuthServiceServersModel
var LdapAuthServiceServersResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IP address or FQDN of the LDAP server.",
	},
	"authentication_type": schema.StringAttribute{
		Default: stringdefault.StaticString("ANONYMOUS"),
		Validators: []validator.String{
			stringvalidator.OneOf("ANONYMOUS", "AUTHENTICATED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The authentication type for the LDAP server.",
	},
	"base_dn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The base DN for the LDAP server.",
	},
	"bind_password": schema.StringAttribute{
		Sensitive:           true,
		Optional:            true,
		MarkdownDescription: "The user password for authentication.",
	},
	"bind_user_dn": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The user DN for authentication.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The LDAP descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the LDAP server is disabled.",
	},
	"encryption": schema.StringAttribute{
		Default: stringdefault.StaticString("SSL"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "SSL"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The LDAP server encryption type.",
	},
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(636),
		Validators: []validator.Int64{
			int64validator.Between(1, 65535),
		},
		MarkdownDescription: "The LDAP server port.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the connection via the MGMT interface is allowed.",
	},
	"version": schema.StringAttribute{
		Default: stringdefault.StaticString("V3"),
		Validators: []validator.String{
			stringvalidator.OneOf("V2", "V3"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The LDAP server version.",
	},
}

// ExpandLdapAuthServiceServers converts a Terraform Object to SDK type
func ExpandLdapAuthServiceServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossecurity.LdapAuthServiceServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LdapAuthServiceServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *LdapAuthServiceServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossecurity.LdapAuthServiceServers {
	if m == nil {
		return nil
	}
	to := &niossecurity.LdapAuthServiceServers{
		Address:            flex.ExpandStringPointerNullAsEmpty(m.Address),
		AuthenticationType: flex.ExpandStringPointerNullAsEmpty(m.AuthenticationType),
		BaseDn:             flex.ExpandStringPointerNullAsEmpty(m.BaseDn),
		BindPassword:       flex.ExpandStringPointerNullAsEmpty(m.BindPassword),
		BindUserDn:         flex.ExpandStringPointerNullAsEmpty(m.BindUserDn),
		Comment:            flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:            flex.ExpandBoolPointer(m.Disable),
		Encryption:         flex.ExpandStringPointerNullAsEmpty(m.Encryption),
		Port:               flex.ExpandInt64Pointer(m.Port),
		UseMgmtPort:        flex.ExpandBoolPointer(m.UseMgmtPort),
		Version:            flex.ExpandStringPointerNullAsEmpty(m.Version),
	}
	return to
}

// FlattenLdapAuthServiceServers converts an SDK type to Terraform Object
func FlattenLdapAuthServiceServers(ctx context.Context, from *niossecurity.LdapAuthServiceServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LdapAuthServiceServersAttrTypes)
	}
	m := &LdapAuthServiceServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LdapAuthServiceServersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *LdapAuthServiceServersModel) Flatten(ctx context.Context, from *niossecurity.LdapAuthServiceServers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.AuthenticationType = flex.FlattenStringPointerEmptyAsNull(from.AuthenticationType)
	m.BaseDn = flex.FlattenStringPointerEmptyAsNull(from.BaseDn)
	m.BindPassword = flex.FlattenStringPointerEmptyAsNull(from.BindPassword)
	m.BindUserDn = flex.FlattenStringPointerEmptyAsNull(from.BindUserDn)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.Encryption = flex.FlattenStringPointerEmptyAsNull(from.Encryption)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.UseMgmtPort = flex.FlattenBoolPointer(from.UseMgmtPort)
	m.Version = flex.FlattenStringPointerEmptyAsNull(from.Version)
}
