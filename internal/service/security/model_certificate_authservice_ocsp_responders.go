package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// CertificateAuthserviceOcspRespondersModel is the Terraform model for CertificateAuthserviceOcspResponders
type CertificateAuthserviceOcspRespondersModel struct {
	FqdnOrIp         types.String `tfsdk:"fqdn_or_ip"`
	Port             types.Int64  `tfsdk:"port"`
	Comment          types.String `tfsdk:"comment"`
	Disabled         types.Bool   `tfsdk:"disabled"`
	Certificate      types.String `tfsdk:"certificate"`
	CertificateToken types.String `tfsdk:"certificate_token"`
}

// CertificateAuthserviceOcspRespondersAttrTypes contains the attribute types for CertificateAuthserviceOcspRespondersModel
var CertificateAuthserviceOcspRespondersAttrTypes = map[string]attr.Type{
	"fqdn_or_ip":        types.StringType,
	"port":              types.Int64Type,
	"comment":           types.StringType,
	"disabled":          types.BoolType,
	"certificate":       types.StringType,
	"certificate_token": types.StringType,
}

// CertificateAuthserviceOcspRespondersResourceSchemaAttributes contains the schema attributes for CertificateAuthserviceOcspRespondersModel
var CertificateAuthserviceOcspRespondersResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn_or_ip": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The FQDN (Fully Qualified Domain Name) or IP address of the server.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(80),
		MarkdownDescription: "The port used for connecting.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The descriptive comment for the OCSP authentication responder.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if this OCSP authentication responder is disabled.",
	},
	"certificate": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "",
	},
	"certificate_token": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop.",
	},
}

// ExpandCertificateAuthserviceOcspResponders converts a Terraform Object to SDK type
func ExpandCertificateAuthserviceOcspResponders(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niossecurity.CertificateAuthserviceOcspResponders {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CertificateAuthserviceOcspRespondersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *CertificateAuthserviceOcspRespondersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niossecurity.CertificateAuthserviceOcspResponders {
	if m == nil {
		return nil
	}
	to := &niossecurity.CertificateAuthserviceOcspResponders{
		FqdnOrIp:         flex.ExpandStringPointerNullAsEmpty(m.FqdnOrIp),
		Port:             flex.ExpandInt64Pointer(m.Port),
		Comment:          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disabled:         flex.ExpandBoolPointer(m.Disabled),
		CertificateToken: flex.ExpandStringPointerNullAsEmpty(m.CertificateToken),
	}
	return to
}

// FlattenCertificateAuthserviceOcspResponders converts an SDK type to Terraform Object
func FlattenCertificateAuthserviceOcspResponders(ctx context.Context, from *niossecurity.CertificateAuthserviceOcspResponders, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CertificateAuthserviceOcspRespondersAttrTypes)
	}
	m := &CertificateAuthserviceOcspRespondersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CertificateAuthserviceOcspRespondersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *CertificateAuthserviceOcspRespondersModel) Flatten(ctx context.Context, from *niossecurity.CertificateAuthserviceOcspResponders, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.FqdnOrIp = flex.FlattenStringPointerEmptyAsNull(from.FqdnOrIp)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.CertificateToken = flex.FlattenStringPointerEmptyAsNull(from.CertificateToken)
}
