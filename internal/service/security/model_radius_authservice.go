package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RadiusAuthserviceModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RadiusAuthserviceAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRadiusAuthserviceAttrTypes},
}

type NIOSRadiusAuthserviceModel struct {
	AcctRetries      types.Int64  `tfsdk:"acct_retries"`
	AcctTimeout      types.Int64  `tfsdk:"acct_timeout"`
	AuthRetries      types.Int64  `tfsdk:"auth_retries"`
	AuthTimeout      types.Int64  `tfsdk:"auth_timeout"`
	CacheTtl         types.Int64  `tfsdk:"cache_ttl"`
	Comment          types.String `tfsdk:"comment"`
	Disable          types.Bool   `tfsdk:"disable"`
	EnableCache      types.Bool   `tfsdk:"enable_cache"`
	Mode             types.String `tfsdk:"mode"`
	Name             types.String `tfsdk:"name"`
	RecoveryInterval types.Int64  `tfsdk:"recovery_interval"`
	Servers          types.List   `tfsdk:"servers"`
}

var NIOSRadiusAuthserviceAttrTypes = map[string]attr.Type{
	"acct_retries":      types.Int64Type,
	"acct_timeout":      types.Int64Type,
	"auth_retries":      types.Int64Type,
	"auth_timeout":      types.Int64Type,
	"cache_ttl":         types.Int64Type,
	"comment":           types.StringType,
	"disable":           types.BoolType,
	"enable_cache":      types.BoolType,
	"mode":              types.StringType,
	"name":              types.StringType,
	"recovery_interval": types.Int64Type,
	"servers":           types.ListType{ElemType: types.ObjectType{AttrTypes: RadiusAuthserviceServersAttrTypes}},
}

const (
	RadiusAuthserviceReturnFields = "acct_retries,acct_timeout,auth_retries,auth_timeout,cache_ttl,comment,disable,enable_cache,mode,name,recovery_interval,servers"
)

var RadiusAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RadiusAuthserviceResourceNiosSchemaAttributes,
	},
}

var RadiusAuthserviceResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"acct_retries": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1000),
		MarkdownDescription: "The number of times to attempt to contact an accounting RADIUS server.",
	},
	"acct_timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5000),
		MarkdownDescription: "The number of seconds to wait for a response from the RADIUS server.",
	},
	"auth_retries": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(6),
		Validators: []validator.Int64{
			int64validator.Between(1, 10),
		},
		MarkdownDescription: "The number of times to attempt to contact an authentication RADIUS server.",
	},
	"auth_timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5000),
		MarkdownDescription: "The number of seconds to wait for a response from the RADIUS server.",
	},
	"cache_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(3600),
		MarkdownDescription: "The TTL of cached authentication data in seconds.",
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
		MarkdownDescription: "Determines whether the RADIUS authentication service is disabled.",
	},
	"enable_cache": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the authentication cache is enabled.",
	},
	"mode": schema.StringAttribute{
		Default: stringdefault.StaticString("HUNT_GROUP"),
		Validators: []validator.String{
			stringvalidator.OneOf("HUNT_GROUP", "ROUND_ROBIN"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The way to contact the RADIUS server.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The RADIUS authentication service name.",
	},
	"recovery_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(30),
		MarkdownDescription: "The time period to wait before retrying a server that has been marked as down.",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RadiusAuthserviceServersResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The ordered list of RADIUS authentication servers.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RadiusAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.RadiusAuthservice {
	if m == nil {
		return nil
	}

	obj := &coremodel.RadiusAuthservice{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRadiusAuthserviceModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRadiusAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRadiusAuthserviceExt {
	return &coremodel.NIOSRadiusAuthserviceExt{
		AcctRetries:      flex.ExpandInt64Pointer(m.AcctRetries),
		AcctTimeout:      flex.ExpandInt64Pointer(m.AcctTimeout),
		AuthRetries:      flex.ExpandInt64Pointer(m.AuthRetries),
		AuthTimeout:      flex.ExpandInt64Pointer(m.AuthTimeout),
		CacheTtl:         flex.ExpandInt64Pointer(m.CacheTtl),
		Comment:          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:          flex.ExpandBoolPointer(m.Disable),
		EnableCache:      flex.ExpandBoolPointer(m.EnableCache),
		Mode:             flex.ExpandStringPointerNullAsEmpty(m.Mode),
		Name:             flex.ExpandStringPointerNullAsEmpty(m.Name),
		RecoveryInterval: flex.ExpandInt64Pointer(m.RecoveryInterval),
		Servers:          flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandRadiusAuthserviceServers),
	}
}

// Flatten populates the TF model from a core response.
func (m *RadiusAuthserviceModel) Flatten(ctx context.Context, resp *coremodel.RadiusAuthservice, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRadiusAuthserviceModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRadiusAuthserviceModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSRadiusAuthserviceModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenRadiusAuthserviceNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRadiusAuthserviceAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRadiusAuthserviceAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRadiusAuthserviceModel) Flatten(ctx context.Context, from *coremodel.NIOSRadiusAuthserviceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AcctRetries = flex.FlattenInt64Pointer(from.AcctRetries)
	m.AcctTimeout = flex.FlattenInt64Pointer(from.AcctTimeout)
	m.AuthRetries = flex.FlattenInt64Pointer(from.AuthRetries)
	m.AuthTimeout = flex.FlattenInt64Pointer(from.AuthTimeout)
	m.CacheTtl = flex.FlattenInt64Pointer(from.CacheTtl)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.EnableCache = flex.FlattenBoolPointer(from.EnableCache)
	m.Mode = flex.FlattenStringPointerEmptyAsNull(from.Mode)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.RecoveryInterval = flex.FlattenInt64Pointer(from.RecoveryInterval)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, RadiusAuthserviceServersAttrTypes, diags, FlattenRadiusAuthserviceServers)
}
