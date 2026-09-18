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
	niossmartfolder "github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// SmartfolderPersonalGroupBysModel is the Terraform model for SmartfolderPersonalGroupBys
type SmartfolderPersonalGroupBysModel struct {
	Value          types.String `tfsdk:"value"`
	ValueType      types.String `tfsdk:"value_type"`
	EnableGrouping types.Bool   `tfsdk:"enable_grouping"`
}

// SmartfolderPersonalGroupBysAttrTypes contains the attribute types for SmartfolderPersonalGroupBysModel
var SmartfolderPersonalGroupBysAttrTypes = map[string]attr.Type{
	"value":           types.StringType,
	"value_type":      types.StringType,
	"enable_grouping": types.BoolType,
}

// SmartfolderPersonalGroupBysResourceSchemaAttributes contains the schema attributes for SmartfolderPersonalGroupBysModel
var SmartfolderPersonalGroupBysResourceSchemaAttributes = map[string]schema.Attribute{
	"value": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the Smart Folder grouping attribute.",
	},
	"value_type": schema.StringAttribute{
		Default: stringdefault.StaticString("NORMAL"),
		Validators: []validator.String{
			stringvalidator.OneOf("EXTATTR", "NORMAL"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of the Smart Folder grouping attribute value.",
	},
	"enable_grouping": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the grouping is enabled.",
	},
}

// ExpandSmartfolderPersonalGroupBys converts a Terraform Object to SDK type
func ExpandSmartfolderPersonalGroupBys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossmartfolder.SmartfolderPersonalGroupBys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SmartfolderPersonalGroupBysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SmartfolderPersonalGroupBysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossmartfolder.SmartfolderPersonalGroupBys {
	if m == nil {
		return nil
	}
	to := &niossmartfolder.SmartfolderPersonalGroupBys{
		Value:          flex.ExpandStringPointerNullAsEmpty(m.Value),
		ValueType:      flex.ExpandStringPointerNullAsEmpty(m.ValueType),
		EnableGrouping: flex.ExpandBoolPointer(m.EnableGrouping),
	}
	return to
}

// FlattenSmartfolderPersonalGroupBys converts an SDK type to Terraform Object
func FlattenSmartfolderPersonalGroupBys(ctx context.Context, from *niossmartfolder.SmartfolderPersonalGroupBys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SmartfolderPersonalGroupBysAttrTypes)
	}
	m := &SmartfolderPersonalGroupBysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SmartfolderPersonalGroupBysAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SmartfolderPersonalGroupBysModel) Flatten(ctx context.Context, from *niossmartfolder.SmartfolderPersonalGroupBys, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Value = flex.FlattenStringPointerEmptyAsNull(from.Value)
	m.ValueType = flex.FlattenStringPointerEmptyAsNull(from.ValueType)
	m.EnableGrouping = flex.FlattenBoolPointer(from.EnableGrouping)
}
