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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type CertificateAuthserviceModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var CertificateAuthserviceAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSCertificateAuthserviceAttrTypes},
}

type NIOSCertificateAuthserviceModel struct {
	AutoPopulateLogin     types.String `tfsdk:"auto_populate_login"`
	CaCertificates        types.List   `tfsdk:"ca_certificates"`
	Comment               types.String `tfsdk:"comment"`
	Disabled              types.Bool   `tfsdk:"disabled"`
	EnablePasswordRequest types.Bool   `tfsdk:"enable_password_request"`
	EnableRemoteLookup    types.Bool   `tfsdk:"enable_remote_lookup"`
	MaxRetries            types.Int64  `tfsdk:"max_retries"`
	Name                  types.String `tfsdk:"name"`
	OcspCheck             types.String `tfsdk:"ocsp_check"`
	OcspResponders        types.List   `tfsdk:"ocsp_responders"`
	RecoveryInterval      types.Int64  `tfsdk:"recovery_interval"`
	RemoteLookupPassword  types.String `tfsdk:"remote_lookup_password"`
	RemoteLookupService   types.String `tfsdk:"remote_lookup_service"`
	RemoteLookupUsername  types.String `tfsdk:"remote_lookup_username"`
	ResponseTimeout       types.Int64  `tfsdk:"response_timeout"`
	TrustModel            types.String `tfsdk:"trust_model"`
	UserMatchType         types.String `tfsdk:"user_match_type"`
}

var NIOSCertificateAuthserviceAttrTypes = map[string]attr.Type{
	"auto_populate_login":     types.StringType,
	"ca_certificates":         types.ListType{ElemType: types.StringType},
	"comment":                 types.StringType,
	"disabled":                types.BoolType,
	"enable_password_request": types.BoolType,
	"enable_remote_lookup":    types.BoolType,
	"max_retries":             types.Int64Type,
	"name":                    types.StringType,
	"ocsp_check":              types.StringType,
	"ocsp_responders":         types.ListType{ElemType: types.ObjectType{AttrTypes: CertificateAuthserviceOcspRespondersAttrTypes}},
	"recovery_interval":       types.Int64Type,
	"remote_lookup_password":  types.StringType,
	"remote_lookup_service":   types.StringType,
	"remote_lookup_username":  types.StringType,
	"response_timeout":        types.Int64Type,
	"trust_model":             types.StringType,
	"user_match_type":         types.StringType,
}

const (
	CertificateAuthserviceReturnFields = "auto_populate_login,ca_certificates,comment,disabled,enable_password_request,enable_remote_lookup,max_retries,name,ocsp_check,ocsp_responders,recovery_interval,remote_lookup_service,remote_lookup_username,response_timeout,trust_model,user_match_type"
)

var CertificateAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          CertificateAuthserviceResourceNiosSchemaAttributes,
	},
}

var CertificateAuthserviceResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"auto_populate_login": schema.StringAttribute{
		Default: stringdefault.StaticString("S_DN_CN"),
		Validators: []validator.String{
			stringvalidator.OneOf("SERIAL_NUMBER", "S_DN_CN", "S_DN_EMAIL", "SAN_UPN", "SAN_EMAIL", "AD_SUBJECT_ISSUER"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Specifies the value of the client certificate for automatically populating the NIOS login name.",
	},
	"ca_certificates": schema.ListAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of CA certificates.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The descriptive comment for the certificate authentication service.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this certificate authentication service is enabled or disabled.",
	},
	"enable_password_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if username/password authentication together with client certificate authentication is enabled or disabled.",
	},
	"enable_remote_lookup": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the lookup for user group membership information on remote services is enabled or disabled.",
	},
	"max_retries": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(0, 5),
		},
		MarkdownDescription: "The number of validation attempts before the appliance contacts the next responder.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the certificate authentication service.",
	},
	"ocsp_check": schema.StringAttribute{
		Default: stringdefault.StaticString("MANUAL"),
		Validators: []validator.String{
			stringvalidator.OneOf("MANUAL", "AIA_ONLY", "AIA_AND_MANUAL", "DISABLED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Specifies the source of OCSP settings.",
	},
	"ocsp_responders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: CertificateAuthserviceOcspRespondersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An ordered list of OCSP responders that are part of the certificate authentication service.",
	},
	"recovery_interval": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(30),
		Validators: []validator.Int64{
			int64validator.Between(1, 600),
		},
		MarkdownDescription: "The period of time the appliance waits before it attempts to contact a responder that is out of service again. The value must be between 1 and 600 seconds.",
	},
	"remote_lookup_password": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The password for the service account.",
	},
	"remote_lookup_service": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "",
	},
	"remote_lookup_username": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The username for the service account.",
	},
	"response_timeout": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(1000),
		Validators: []validator.Int64{
			int64validator.Between(1000, 60000),
		},
		MarkdownDescription: "The validation timeout period in milliseconds.",
	},
	"trust_model": schema.StringAttribute{
		Default: stringdefault.StaticString("DIRECT"),
		Validators: []validator.String{
			stringvalidator.OneOf("DIRECT", "DELEGATED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The OCSP trust model.",
	},
	"user_match_type": schema.StringAttribute{
		Default: stringdefault.StaticString("AUTO_MATCH"),
		Validators: []validator.String{
			stringvalidator.OneOf("DIRECT_MATCH", "AUTO_MATCH"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Specifies how to search for a user.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *CertificateAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.CertificateAuthservice {
	if m == nil {
		return nil
	}

	obj := &coremodel.CertificateAuthservice{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSCertificateAuthserviceModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSCertificateAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSCertificateAuthserviceExt {
	return &coremodel.NIOSCertificateAuthserviceExt{
		AutoPopulateLogin:     flex.ExpandStringPointerNullAsEmpty(m.AutoPopulateLogin),
		CaCertificates:        flex.ExpandFrameworkListString(ctx, m.CaCertificates, diags),
		Comment:               flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disabled:              flex.ExpandBoolPointer(m.Disabled),
		EnablePasswordRequest: flex.ExpandBoolPointer(m.EnablePasswordRequest),
		EnableRemoteLookup:    flex.ExpandBoolPointer(m.EnableRemoteLookup),
		MaxRetries:            flex.ExpandInt64Pointer(m.MaxRetries),
		Name:                  flex.ExpandStringPointerNullAsEmpty(m.Name),
		OcspCheck:             flex.ExpandStringPointerNullAsEmpty(m.OcspCheck),
		OcspResponders:        flex.ExpandFrameworkListNestedBlock(ctx, m.OcspResponders, diags, ExpandCertificateAuthserviceOcspResponders),
		RecoveryInterval:      flex.ExpandInt64Pointer(m.RecoveryInterval),
		RemoteLookupPassword:  flex.ExpandStringPointerNullAsEmpty(m.RemoteLookupPassword),
		RemoteLookupUsername:  flex.ExpandStringPointerNullAsEmpty(m.RemoteLookupUsername),
		ResponseTimeout:       flex.ExpandInt64Pointer(m.ResponseTimeout),
		TrustModel:            flex.ExpandStringPointerNullAsEmpty(m.TrustModel),
		UserMatchType:         flex.ExpandStringPointerNullAsEmpty(m.UserMatchType),
	}
}

// Flatten populates the TF model from a core response.
func (m *CertificateAuthserviceModel) Flatten(ctx context.Context, resp *coremodel.CertificateAuthservice, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSCertificateAuthserviceModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSCertificateAuthserviceModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSCertificateAuthserviceModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenCertificateAuthserviceNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSCertificateAuthserviceAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSCertificateAuthserviceAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSCertificateAuthserviceModel) Flatten(ctx context.Context, from *coremodel.NIOSCertificateAuthserviceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AutoPopulateLogin = flex.FlattenStringPointerEmptyAsNull(from.AutoPopulateLogin)
	m.CaCertificates = flex.FlattenFrameworkListString(ctx, from.CaCertificates, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.EnablePasswordRequest = flex.FlattenBoolPointer(from.EnablePasswordRequest)
	m.EnableRemoteLookup = flex.FlattenBoolPointer(from.EnableRemoteLookup)
	m.MaxRetries = flex.FlattenInt64Pointer(from.MaxRetries)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.OcspCheck = flex.FlattenStringPointerEmptyAsNull(from.OcspCheck)
	m.OcspResponders = flex.FlattenFrameworkListNestedBlock(ctx, from.OcspResponders, CertificateAuthserviceOcspRespondersAttrTypes, diags, FlattenCertificateAuthserviceOcspResponders)
	m.RecoveryInterval = flex.FlattenInt64Pointer(from.RecoveryInterval)
	m.RemoteLookupPassword = flex.FlattenStringPointerEmptyAsNull(from.RemoteLookupPassword)
	m.RemoteLookupUsername = flex.FlattenStringPointerEmptyAsNull(from.RemoteLookupUsername)
	m.ResponseTimeout = flex.FlattenInt64Pointer(from.ResponseTimeout)
	m.TrustModel = flex.FlattenStringPointerEmptyAsNull(from.TrustModel)
	m.UserMatchType = flex.FlattenStringPointerEmptyAsNull(from.UserMatchType)
}
