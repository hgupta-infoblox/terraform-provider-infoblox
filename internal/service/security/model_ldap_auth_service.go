package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type LdapAuthServiceModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var LdapAuthServiceAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSLdapAuthServiceAttrTypes},
}

type NIOSLdapAuthServiceModel struct {
	Comment                     types.String `tfsdk:"comment"`
	Disable                     types.Bool   `tfsdk:"disable"`
	EaMapping                   types.List   `tfsdk:"ea_mapping"`
	LdapGroupAttribute          types.String `tfsdk:"ldap_group_attribute"`
	LdapGroupAuthenticationType types.String `tfsdk:"ldap_group_authentication_type"`
	LdapUserAttribute           types.String `tfsdk:"ldap_user_attribute"`
	Mode                        types.String `tfsdk:"mode"`
	Name                        types.String `tfsdk:"name"`
	RecoveryInterval            types.Int64  `tfsdk:"recovery_interval"`
	Retries                     types.Int64  `tfsdk:"retries"`
	SearchScope                 types.String `tfsdk:"search_scope"`
	Servers                     types.List   `tfsdk:"servers"`
	Timeout                     types.Int64  `tfsdk:"timeout"`
}

var NIOSLdapAuthServiceAttrTypes = map[string]attr.Type{
	"comment":                        types.StringType,
	"disable":                        types.BoolType,
	"ea_mapping":                     types.ListType{ElemType: types.ObjectType{AttrTypes: LdapAuthServiceEaMappingAttrTypes}},
	"ldap_group_attribute":           types.StringType,
	"ldap_group_authentication_type": types.StringType,
	"ldap_user_attribute":            types.StringType,
	"mode":                           types.StringType,
	"name":                           types.StringType,
	"recovery_interval":              types.Int64Type,
	"retries":                        types.Int64Type,
	"search_scope":                   types.StringType,
	"servers":                        types.ListType{ElemType: types.ObjectType{AttrTypes: LdapAuthServiceServersAttrTypes}},
	"timeout":                        types.Int64Type,
}

const (
	LdapAuthServiceReturnFields = "comment,disable,ea_mapping,ldap_group_attribute,ldap_group_authentication_type,ldap_user_attribute,mode,name,recovery_interval,retries,search_scope,servers,timeout"
)

var LdapAuthServiceResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          LdapAuthServiceResourceNiosSchemaAttributes,
	},
}

var LdapAuthServiceResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The LDAP descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the LDAP authentication service is disabled.",
	},
	"ea_mapping": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: LdapAuthServiceEaMappingResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The mapping LDAP fields to extensible attributes.",
	},
	"ldap_group_attribute": schema.StringAttribute{
		Default:  stringdefault.StaticString("memberOf"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the LDAP attribute that defines group membership.",
	},
	"ldap_group_authentication_type": schema.StringAttribute{
		Default: stringdefault.StaticString("GROUP_ATTRIBUTE"),
		Validators: []validator.String{
			stringvalidator.OneOf("GROUP_ATTRIBUTE", "POSIX_GROUP"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The LDAP group authentication type.",
	},
	"ldap_user_attribute": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The LDAP userid attribute that is used for search.",
	},
	"mode": schema.StringAttribute{
		Default: stringdefault.StaticString("ORDERED_LIST"),
		Validators: []validator.String{
			stringvalidator.OneOf("ORDERED_LIST", "ROUND_ROBIN"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The LDAP authentication mode.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The LDAP authentication service name.",
	},
	"recovery_interval": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 600),
		},
		MarkdownDescription: "The period of time in seconds to wait before trying to contact a LDAP server that has been marked as 'DOWN'.",
	},
	"retries": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 5),
		},
		MarkdownDescription: "The maximum number of LDAP authentication attempts.",
	},
	"search_scope": schema.StringAttribute{
		Default: stringdefault.StaticString("ONELEVEL"),
		Validators: []validator.String{
			stringvalidator.OneOf("BASE", "ONELEVEL", "SUBTREE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The starting point of the LDAP search.",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: LdapAuthServiceServersResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of LDAP servers used for authentication.",
	},
	"timeout": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 600),
		},
		MarkdownDescription: "The LDAP authentication timeout in seconds.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *LdapAuthServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.LdapAuthService {
	if m == nil {
		return nil
	}

	obj := &coremodel.LdapAuthService{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSLdapAuthServiceModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSLdapAuthServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSLdapAuthServiceExt {
	return &coremodel.NIOSLdapAuthServiceExt{
		Comment:                     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:                     flex.ExpandBoolPointer(m.Disable),
		EaMapping:                   flex.ExpandFrameworkListNestedBlock(ctx, m.EaMapping, diags, ExpandLdapAuthServiceEaMapping),
		LdapGroupAttribute:          flex.ExpandStringPointerNullAsEmpty(m.LdapGroupAttribute),
		LdapGroupAuthenticationType: flex.ExpandStringPointerNullAsEmpty(m.LdapGroupAuthenticationType),
		LdapUserAttribute:           flex.ExpandStringPointerNullAsEmpty(m.LdapUserAttribute),
		Mode:                        flex.ExpandStringPointerNullAsEmpty(m.Mode),
		Name:                        flex.ExpandStringPointerNullAsEmpty(m.Name),
		RecoveryInterval:            flex.ExpandInt64Pointer(m.RecoveryInterval),
		Retries:                     flex.ExpandInt64Pointer(m.Retries),
		SearchScope:                 flex.ExpandStringPointerNullAsEmpty(m.SearchScope),
		Servers:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandLdapAuthServiceServers),
		Timeout:                     flex.ExpandInt64Pointer(m.Timeout),
	}
}

// Flatten populates the TF model from a core response.
func (m *LdapAuthServiceModel) Flatten(ctx context.Context, resp *coremodel.LdapAuthService, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSLdapAuthServiceModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSLdapAuthServiceModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSLdapAuthServiceModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenLdapAuthServiceNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSLdapAuthServiceAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSLdapAuthServiceAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSLdapAuthServiceModel) Flatten(ctx context.Context, from *coremodel.NIOSLdapAuthServiceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.EaMapping = flex.FlattenFrameworkListNestedBlock(ctx, from.EaMapping, LdapAuthServiceEaMappingAttrTypes, diags, FlattenLdapAuthServiceEaMapping)
	m.LdapGroupAttribute = flex.FlattenStringPointerEmptyAsNull(from.LdapGroupAttribute)
	m.LdapGroupAuthenticationType = flex.FlattenStringPointerEmptyAsNull(from.LdapGroupAuthenticationType)
	m.LdapUserAttribute = flex.FlattenStringPointerEmptyAsNull(from.LdapUserAttribute)
	m.Mode = flex.FlattenStringPointerEmptyAsNull(from.Mode)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.RecoveryInterval = flex.FlattenInt64Pointer(from.RecoveryInterval)
	m.Retries = flex.FlattenInt64Pointer(from.Retries)
	m.SearchScope = flex.FlattenStringPointerEmptyAsNull(from.SearchScope)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, LdapAuthServiceServersAttrTypes, diags, FlattenLdapAuthServiceServers)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
}
