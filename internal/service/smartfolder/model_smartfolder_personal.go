package smartfolder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/smartfolder"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SmartfolderPersonalModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var SmartfolderPersonalAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSSmartfolderPersonalAttrTypes},
}

type NIOSSmartfolderPersonalModel struct {
	Comment    types.String `tfsdk:"comment"`
	GroupBys   types.List   `tfsdk:"group_bys"`
	Name       types.String `tfsdk:"name"`
	QueryItems types.List   `tfsdk:"query_items"`
}

var NIOSSmartfolderPersonalAttrTypes = map[string]attr.Type{
	"comment":     types.StringType,
	"group_bys":   types.ListType{ElemType: types.ObjectType{AttrTypes: SmartfolderPersonalGroupBysAttrTypes}},
	"name":        types.StringType,
	"query_items": types.ListType{ElemType: types.ObjectType{AttrTypes: SmartfolderPersonalQueryItemsAttrTypes}},
}

const (
	SmartfolderPersonalReturnFields = "comment,group_bys,is_shortcut,name,query_items"
)

var SmartfolderPersonalResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          SmartfolderPersonalResourceNiosSchemaAttributes,
	},
}

var SmartfolderPersonalResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The personal Smart Folder descriptive comment.",
	},
	"group_bys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SmartfolderPersonalGroupBysResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The personal Smart Folder groupping rules.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The personal Smart Folder name.",
	},
	"query_items": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SmartfolderPersonalQueryItemsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: SmartfolderPersonalQueryItemsAttrTypes}, []attr.Value{types.ObjectValueMust(SmartfolderPersonalQueryItemsAttrTypes, map[string]attr.Value{
			"name":       types.StringValue("type"),
			"field_type": types.StringValue("NORMAL"),
			"operator":   types.StringValue("EQ"),
			"op_match":   types.BoolValue(true),
			"value_type": types.StringValue("ENUM"),
			"value":      types.ObjectValueMust(SmartfolderpersonalqueryitemsValueAttrTypes, map[string]attr.Value{"value_string": types.StringValue("Network/Zone/Range/Member"), "value_integer": types.Int64Null(), "value_date": types.Int64Null(), "value_boolean": types.BoolNull()}),
		})})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The personal Smart Folder filter queries.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *SmartfolderPersonalModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.SmartfolderPersonal {
	if m == nil {
		return nil
	}

	obj := &coremodel.SmartfolderPersonal{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSSmartfolderPersonalModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSSmartfolderPersonalModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSSmartfolderPersonalExt {
	return &coremodel.NIOSSmartfolderPersonalExt{
		Comment:    flex.ExpandStringPointerNullAsEmpty(m.Comment),
		GroupBys:   flex.ExpandFrameworkListNestedBlock(ctx, m.GroupBys, diags, ExpandSmartfolderPersonalGroupBys),
		Name:       flex.ExpandStringPointerNullAsEmpty(m.Name),
		QueryItems: flex.ExpandFrameworkListNestedBlock(ctx, m.QueryItems, diags, ExpandSmartfolderPersonalQueryItems),
	}
}

// Flatten populates the TF model from a core response.
func (m *SmartfolderPersonalModel) Flatten(ctx context.Context, resp *coremodel.SmartfolderPersonal, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSSmartfolderPersonalModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSSmartfolderPersonalModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSSmartfolderPersonalAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSSmartfolderPersonalAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSSmartfolderPersonalModel) Flatten(ctx context.Context, from *coremodel.NIOSSmartfolderPersonalExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.GroupBys = flex.FlattenFrameworkListNestedBlock(ctx, from.GroupBys, SmartfolderPersonalGroupBysAttrTypes, diags, FlattenSmartfolderPersonalGroupBys)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.QueryItems = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryItems, SmartfolderPersonalQueryItemsAttrTypes, diags, FlattenSmartfolderPersonalQueryItems)
}
