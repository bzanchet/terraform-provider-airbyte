// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceGoogleAnalyticsDataAPIResource{}
var _ resource.ResourceWithImportState = &SourceGoogleAnalyticsDataAPIResource{}

func NewSourceGoogleAnalyticsDataAPIResource() resource.Resource {
	return &SourceGoogleAnalyticsDataAPIResource{}
}

// SourceGoogleAnalyticsDataAPIResource defines the resource implementation.
type SourceGoogleAnalyticsDataAPIResource struct {
	client *sdk.SDK
}

// SourceGoogleAnalyticsDataAPIResourceModel describes the resource data model.
type SourceGoogleAnalyticsDataAPIResourceModel struct {
	Configuration SourceGoogleAnalyticsDataAPI `tfsdk:"configuration"`
	Name          types.String                 `tfsdk:"name"`
	SecretID      types.String                 `tfsdk:"secret_id"`
	SourceID      types.String                 `tfsdk:"source_id"`
	SourceType    types.String                 `tfsdk:"source_type"`
	WorkspaceID   types.String                 `tfsdk:"workspace_id"`
}

func (r *SourceGoogleAnalyticsDataAPIResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_analytics_data_api"
}

func (r *SourceGoogleAnalyticsDataAPIResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleAnalyticsDataAPI Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_google_analytics_data_api_credentials_authenticate_via_google_oauth_": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Optional: true,
									},
									"auth_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
									},
									"client_id": schema.StringAttribute{
										Required: true,
									},
									"client_secret": schema.StringAttribute{
										Required: true,
									},
									"refresh_token": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Credentials for the service`,
							},
							"source_google_analytics_data_api_credentials_service_account_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Service",
											),
										},
									},
									"credentials_json": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Credentials for the service`,
							},
							"source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth_": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed: true,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
									},
									"client_id": schema.StringAttribute{
										Computed: true,
									},
									"client_secret": schema.StringAttribute{
										Computed: true,
									},
									"refresh_token": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Credentials for the service`,
							},
							"source_google_analytics_data_api_update_credentials_service_account_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Service",
											),
										},
									},
									"credentials_json": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Credentials for the service`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"custom_reports": schema.StringAttribute{
						Optional: true,
					},
					"date_ranges_start_date": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
					},
					"property_id": schema.StringAttribute{
						Required: true,
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"google-analytics-data-api",
							),
						},
					},
					"window_in_days": schema.Int64Attribute{
						Optional: true,
					},
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				Optional: true,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
			},
			"source_type": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *SourceGoogleAnalyticsDataAPIResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceGoogleAnalyticsDataAPIResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceGoogleAnalyticsDataAPIResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceGoogleAnalyticsDataAPI(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAnalyticsDataAPIResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceGoogleAnalyticsDataAPIResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAnalyticsDataAPIResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceGoogleAnalyticsDataAPIResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceGoogleAnalyticsDataAPIPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceGoogleAnalyticsDataAPIRequest{
		SourceGoogleAnalyticsDataAPIPutRequest: sourceGoogleAnalyticsDataAPIPutRequest,
		SourceID:                               sourceID,
	}
	res, err := r.client.Sources.PutSourceGoogleAnalyticsDataAPI(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAnalyticsDataAPIResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceGoogleAnalyticsDataAPIResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.DeleteSourceGoogleAnalyticsDataAPIRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceGoogleAnalyticsDataAPI(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceGoogleAnalyticsDataAPIResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}