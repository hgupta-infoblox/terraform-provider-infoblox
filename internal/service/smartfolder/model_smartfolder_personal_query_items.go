package smartfolder

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	niossmartfolder "github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// SmartfolderPersonalQueryItemsModel is the Terraform model for SmartfolderPersonalQueryItems
type SmartfolderPersonalQueryItemsModel struct {
	Name      types.String `tfsdk:"name"`
	FieldType types.String `tfsdk:"field_type"`
	Operator  types.String `tfsdk:"operator"`
	OpMatch   types.Bool   `tfsdk:"op_match"`
	ValueType types.String `tfsdk:"value_type"`
	Value     types.Object `tfsdk:"value"`
}

// SmartfolderPersonalQueryItemsAttrTypes contains the attribute types for SmartfolderPersonalQueryItemsModel
var SmartfolderPersonalQueryItemsAttrTypes = map[string]attr.Type{
	"name":       types.StringType,
	"field_type": types.StringType,
	"operator":   types.StringType,
	"op_match":   types.BoolType,
	"value_type": types.StringType,
	"value":      types.ObjectType{AttrTypes: SmartfolderpersonalqueryitemsValueAttrTypes},
}

// SmartfolderPersonalQueryItemsResourceSchemaAttributes contains the schema attributes for SmartfolderPersonalQueryItemsModel
var SmartfolderPersonalQueryItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Default:  stringdefault.StaticString("type"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The Smart Folder query name.",
	},
	"field_type": schema.StringAttribute{
		Default: stringdefault.StaticString("NORMAL"),
		Validators: []validator.String{
			stringvalidator.OneOf("EXTATTR", "NORMAL"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Smart Folder query field type.",
	},
	"operator": schema.StringAttribute{
		Default: stringdefault.StaticString("EQ"),
		Validators: []validator.String{
			stringvalidator.OneOf("BEGINS_WITH", "CONTAINS", "DROPS_BY", "ENDS_WITH", "EQ", "GEQ", "GT", "HAS_VALUE", "INHERITANCE_STATE_EQUALS", "IP_ADDR_WITHIN", "LEQ", "LT", "MATCH_EXPR", "RELATIVE_DATE", "RISES_BY", "SUFFIX_MATCH"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Smart Folder operator used in query.",
	},
	"op_match": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether the query operator should match.",
	},
	"value_type": schema.StringAttribute{
		Default: stringdefault.StaticString("ENUM"),
		Validators: []validator.String{
			stringvalidator.OneOf("BOOLEAN", "DATE", "EMAIL", "ENUM", "INTEGER", "OBJTYPE", "STRING", "URL"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Smart Folder query value type.",
	},
	"value": schema.SingleNestedAttribute{
		Attributes: SmartfolderpersonalqueryitemsValueResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(SmartfolderpersonalqueryitemsValueAttrTypes, map[string]attr.Value{
			"value_string":  types.StringValue("Network/Zone/Range/Member"),
			"value_integer": types.Int64Null(),
			"value_date":    types.Int64Null(),
			"value_boolean": types.BoolNull(),
		})),
		MarkdownDescription: "The Smart Folder query value.",
	},
}

// ExpandSmartfolderPersonalQueryItems converts a Terraform Object to SDK type
func ExpandSmartfolderPersonalQueryItems(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossmartfolder.SmartfolderPersonalQueryItems {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SmartfolderPersonalQueryItemsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SmartfolderPersonalQueryItemsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossmartfolder.SmartfolderPersonalQueryItems {
	if m == nil {
		return nil
	}
	to := &niossmartfolder.SmartfolderPersonalQueryItems{
		Name:      flex.ExpandStringPointerNullAsEmpty(m.Name),
		FieldType: flex.ExpandStringPointerNullAsEmpty(m.FieldType),
		Operator:  flex.ExpandStringPointerNullAsEmpty(m.Operator),
		OpMatch:   flex.ExpandBoolPointer(m.OpMatch),
		ValueType: flex.ExpandStringPointerNullAsEmpty(m.ValueType),
		Value:     ExpandSmartfolderpersonalqueryitemsValue(ctx, m.Value, diags),
	}
	return to
}

// FlattenSmartfolderPersonalQueryItems converts an SDK type to Terraform Object
func FlattenSmartfolderPersonalQueryItems(ctx context.Context, from *niossmartfolder.SmartfolderPersonalQueryItems, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderPersonalQueryItemsAttrTypes)
	}
	m := &SmartfolderPersonalQueryItemsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderPersonalQueryItemsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SmartfolderPersonalQueryItemsModel) Flatten(ctx context.Context, from *niossmartfolder.SmartfolderPersonalQueryItems, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.FieldType = flex.FlattenStringPointerEmptyAsNull(from.FieldType)
	m.Operator = flex.FlattenStringPointerEmptyAsNull(from.Operator)
	m.OpMatch = flex.FlattenBoolPointer(from.OpMatch)
	m.ValueType = flex.FlattenStringPointerEmptyAsNull(from.ValueType)
	m.Value = FlattenSmartfolderpersonalqueryitemsValue(ctx, from.Value, diags)
}
