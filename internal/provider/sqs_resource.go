package provider

import (
	"context"
	stream "github.com/GetStream/stream-chat-go/v6"
	"github.com/dcarbone/terraform-plugin-framework-utils/validation"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.ResourceType = sqsResourceType{}
var _ tfsdk.Resource = sqsResource{}
var _ tfsdk.ResourceWithImportState = sqsResource{}

type sqsResourceType struct{}

func (t sqsResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Add an sqsResource connection to the GetStream.io application",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Computed:            true,
				MarkdownDescription: "Identifier",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"sqs_url": {
				MarkdownDescription: "URL to send messages on the SQS queue",
				Required:            true,
				Type:                types.StringType,
				Validators: []tfsdk.AttributeValidator{
					validation.IsURL(),
				},
			},
			"sqs_access_key": {
				MarkdownDescription: "Access key with privileges to send message on the SQS queue",
				Required:            true,
				Type:                types.StringType,
			},
			"sqs_secret_key": {
				MarkdownDescription: "Secret key with privileges to send message on the SQS queue",
				Required:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func (t sqsResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return sqsResource{
		provider: provider,
	}, diags
}

type sqsResourceData struct {
	Id           types.String `tfsdk:"id"`
	SqsUrl       types.String `tfsdk:"sqs_url"`
	SqsAccessKey types.String `tfsdk:"sqs_access_key"`
	SqsSecretKey types.String `tfsdk:"sqs_secret_key"`
}

type sqsResource struct {
	provider provider
}

func (r sqsResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data sqsResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating the sqs link on the GetStream.io...")
	tflog.Trace(ctx, "URL: "+data.SqsUrl.Value)
	tflog.Trace(ctx, "AccessKey: "+data.SqsAccessKey.Value)
	tflog.Trace(ctx, "SecretKey: "+data.SqsSecretKey.Value)
	// set your sqsResource queue details
	settings := &stream.AppSettings{
		SqsURL:    &data.SqsUrl.Value,
		SqsKey:    &data.SqsAccessKey.Value,
		SqsSecret: &data.SqsSecretKey.Value,
	}
	_, err := r.provider.client.UpdateAppSettings(ctx, settings)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Error during the SQS creation.", err.Error()))
		tflog.Error(ctx, err.Error())
		return
	}
	tflog.Debug(ctx, "SQS link on the GetStream.io created.")

	data.Id = types.String{Value: "getstream-sqs-1"}
	diags.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(diags...)
}

func (r sqsResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data sqsResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r sqsResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data sqsResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// set your sqsResource queue details
	tflog.Debug(ctx, "Updating the sqs link on the GetStream.io...")
	tflog.Trace(ctx, "URL: "+data.SqsUrl.Value)
	tflog.Trace(ctx, "AccessKey: "+data.SqsAccessKey.Value)
	tflog.Trace(ctx, "SecretKey: "+data.SqsSecretKey.Value)
	settings := &stream.AppSettings{
		SqsURL:    &data.SqsUrl.Value,
		SqsKey:    &data.SqsAccessKey.Value,
		SqsSecret: &data.SqsSecretKey.Value,
	}
	_, err := r.provider.client.UpdateAppSettings(ctx, settings)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Error during the SQS update.", err.Error()))
		tflog.Error(ctx, err.Error())
		return
	}

	tflog.Debug(ctx, "SQS link on the GetStream.io updated.")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r sqsResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data sqsResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.SqsUrl.Value = ""
	data.SqsAccessKey.Value = ""
	data.SqsSecretKey.Value = ""

	// set your sqsResource queue details
	tflog.Debug(ctx, "Deleting the sqs link on the GetStream.io...")
	settings := &stream.AppSettings{
		SqsURL:    &data.SqsUrl.Value,
		SqsKey:    &data.SqsAccessKey.Value,
		SqsSecret: &data.SqsSecretKey.Value,
	}
	_, err := r.provider.client.UpdateAppSettings(ctx, settings)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Error during the SQS deletion.", err.Error()))
		tflog.Error(ctx, err.Error())
		return
	}
	tflog.Debug(ctx, "SQS link on the GetStream.io deleted.")
}

func (r sqsResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
