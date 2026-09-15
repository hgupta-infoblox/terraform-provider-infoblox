package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// LdapAuthServiceEaMappingModel is the Terraform model for LdapAuthServiceEaMapping
type LdapAuthServiceEaMappingModel struct {
	Name     types.String `tfsdk:"name"`
	MappedEa types.String `tfsdk:"mapped_ea"`
}

// LdapAuthServiceEaMappingAttrTypes contains the attribute types for LdapAuthServiceEaMappingModel
var LdapAuthServiceEaMappingAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"mapped_ea": types.StringType,
}

// LdapAuthServiceEaMappingResourceSchemaAttributes contains the schema attributes for LdapAuthServiceEaMappingModel
var LdapAuthServiceEaMappingResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The LDAP attribute name.",
	},
	"mapped_ea": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the extensible attribute definition object to which the LDAP attribute is mapped.",
	},
}

// ExpandLdapAuthServiceEaMapping converts a Terraform Object to SDK type
func ExpandLdapAuthServiceEaMapping(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossecurity.LdapAuthServiceEaMapping {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LdapAuthServiceEaMappingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *LdapAuthServiceEaMappingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossecurity.LdapAuthServiceEaMapping {
	if m == nil {
		return nil
	}
	to := &niossecurity.LdapAuthServiceEaMapping{
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
		MappedEa: flex.ExpandStringPointerNullAsEmpty(m.MappedEa),
	}
	return to
}

// FlattenLdapAuthServiceEaMapping converts an SDK type to Terraform Object
func FlattenLdapAuthServiceEaMapping(ctx context.Context, from *niossecurity.LdapAuthServiceEaMapping, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LdapAuthServiceEaMappingAttrTypes)
	}
	m := &LdapAuthServiceEaMappingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LdapAuthServiceEaMappingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *LdapAuthServiceEaMappingModel) Flatten(ctx context.Context, from *niossecurity.LdapAuthServiceEaMapping, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.MappedEa = flex.FlattenStringPointerEmptyAsNull(from.MappedEa)
}
