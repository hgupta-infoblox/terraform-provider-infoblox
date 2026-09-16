package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type TacacsplusAuthserviceModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var TacacsplusAuthserviceAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSTacacsplusAuthserviceAttrTypes},
}

type NIOSTacacsplusAuthserviceModel struct {
	AcctRetries types.Int64  `tfsdk:"acct_retries"`
	AcctTimeout types.Int64  `tfsdk:"acct_timeout"`
	AuthRetries types.Int64  `tfsdk:"auth_retries"`
	AuthTimeout types.Int64  `tfsdk:"auth_timeout"`
	Comment     types.String `tfsdk:"comment"`
	Disable     types.Bool   `tfsdk:"disable"`
	Name        types.String `tfsdk:"name"`
	Servers     types.List   `tfsdk:"servers"`
}

var NIOSTacacsplusAuthserviceAttrTypes = map[string]attr.Type{
	"acct_retries": types.Int64Type,
	"acct_timeout": types.Int64Type,
	"auth_retries": types.Int64Type,
	"auth_timeout": types.Int64Type,
	"comment":      types.StringType,
	"disable":      types.BoolType,
	"name":         types.StringType,
	"servers":      types.ListType{ElemType: types.ObjectType{AttrTypes: TacacsplusAuthserviceServersAttrTypes}},
}

const (
	TacacsplusAuthserviceReturnFields = "acct_retries,acct_timeout,auth_retries,auth_timeout,comment,disable,name,servers"
)

var TacacsplusAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          TacacsplusAuthserviceResourceNiosSchemaAttributes,
	},
}

var TacacsplusAuthserviceResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"acct_retries": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(0, 5),
		},
		MarkdownDescription: "The number of the accounting retries before giving up and moving on to the next server.",
	},
	"acct_timeout": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(1000),
		Validators: []validator.Int64{
			int64validator.Between(1, 4294967295),
		},
		MarkdownDescription: "The accounting retry period in milliseconds.",
	},
	"auth_retries": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(0, 5),
		},
		MarkdownDescription: "The number of the authentication/authorization retries before giving up and moving on to the next server.",
	},
	"auth_timeout": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(5000),
		Validators: []validator.Int64{
			int64validator.Between(5000, 60000),
		},
		MarkdownDescription: "The authentication/authorization timeout period in milliseconds.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The TACACS+ authentication service descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the TACACS+ authentication service object is disabled.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The TACACS+ authentication service name.",
	},
	"servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TacacsplusAuthserviceServersResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of the TACACS+ servers used for authentication.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *TacacsplusAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.TacacsplusAuthservice {
	if m == nil {
		return nil
	}

	obj := &coremodel.TacacsplusAuthservice{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSTacacsplusAuthserviceModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSTacacsplusAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSTacacsplusAuthserviceExt {
	return &coremodel.NIOSTacacsplusAuthserviceExt{
		AcctRetries: flex.ExpandInt64Pointer(m.AcctRetries),
		AcctTimeout: flex.ExpandInt64Pointer(m.AcctTimeout),
		AuthRetries: flex.ExpandInt64Pointer(m.AuthRetries),
		AuthTimeout: flex.ExpandInt64Pointer(m.AuthTimeout),
		Comment:     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:     flex.ExpandBoolPointer(m.Disable),
		Name:        flex.ExpandStringPointerNullAsEmpty(m.Name),
		Servers:     flex.ExpandFrameworkListNestedBlock(ctx, m.Servers, diags, ExpandTacacsplusAuthserviceServers),
	}
}

// Flatten populates the TF model from a core response.
func (m *TacacsplusAuthserviceModel) Flatten(ctx context.Context, resp *coremodel.TacacsplusAuthservice, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSTacacsplusAuthserviceModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSTacacsplusAuthserviceModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSTacacsplusAuthserviceModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenTacacsplusAuthserviceNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSTacacsplusAuthserviceAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSTacacsplusAuthserviceAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSTacacsplusAuthserviceModel) Flatten(ctx context.Context, from *coremodel.NIOSTacacsplusAuthserviceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AcctRetries = flex.FlattenInt64Pointer(from.AcctRetries)
	m.AcctTimeout = flex.FlattenInt64Pointer(from.AcctTimeout)
	m.AuthRetries = flex.FlattenInt64Pointer(from.AuthRetries)
	m.AuthTimeout = flex.FlattenInt64Pointer(from.AuthTimeout)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Servers = flex.FlattenFrameworkListNestedBlock(ctx, from.Servers, TacacsplusAuthserviceServersAttrTypes, diags, FlattenTacacsplusAuthserviceServers)
}
