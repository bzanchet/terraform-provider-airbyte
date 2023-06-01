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
var _ resource.Resource = &DestinationSnowflakeResource{}
var _ resource.ResourceWithImportState = &DestinationSnowflakeResource{}

func NewDestinationSnowflakeResource() resource.Resource {
	return &DestinationSnowflakeResource{}
}

// DestinationSnowflakeResource defines the resource implementation.
type DestinationSnowflakeResource struct {
	client *sdk.SDK
}

// DestinationSnowflakeResourceModel describes the resource data model.
type DestinationSnowflakeResourceModel struct {
	Configuration   DestinationSnowflake `tfsdk:"configuration"`
	DestinationID   types.String         `tfsdk:"destination_id"`
	DestinationType types.String         `tfsdk:"destination_type"`
	Name            types.String         `tfsdk:"name"`
	WorkspaceID     types.String         `tfsdk:"workspace_id"`
}

func (r *DestinationSnowflakeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_snowflake"
}

func (r *DestinationSnowflakeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationSnowflake Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"destination_snowflake_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Required: true,
									},
									"auth_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
											),
										},
									},
									"client_id": schema.StringAttribute{
										Optional: true,
									},
									"client_secret": schema.StringAttribute{
										Optional: true,
									},
									"refresh_token": schema.StringAttribute{
										Required: true,
									},
								},
							},
							"destination_snowflake_authorization_method_key_pair_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Key Pair Authentication",
											),
										},
									},
									"private_key": schema.StringAttribute{
										Required: true,
									},
									"private_key_password": schema.StringAttribute{
										Optional: true,
									},
								},
							},
							"destination_snowflake_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Username and Password",
											),
										},
									},
									"password": schema.StringAttribute{
										Required: true,
									},
								},
							},
							"destination_snowflake_update_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed: true,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
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
							},
							"destination_snowflake_update_authorization_method_key_pair_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Key Pair Authentication",
											),
										},
									},
									"private_key": schema.StringAttribute{
										Computed: true,
									},
									"private_key_password": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"destination_snowflake_update_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Username and Password",
											),
										},
									},
									"password": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"database": schema.StringAttribute{
						Required: true,
					},
					"destination_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"snowflake",
							),
						},
					},
					"file_buffer_count": schema.Int64Attribute{
						Optional: true,
					},
					"host": schema.StringAttribute{
						Required: true,
					},
					"jdbc_url_params": schema.StringAttribute{
						Optional: true,
					},
					"loading_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"destination_snowflake_data_staging_method_select_another_option": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
									},
								},
								Description: `Select another option`,
							},
							"destination_snowflake_data_staging_method_recommended_internal_staging": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Internal Staging",
											),
										},
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_data_staging_method_aws_s3_staging": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Required: true,
									},
									"encryption": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"destination_snowflake_data_staging_method_aws_s3_staging_encryption_no_encryption": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"none",
															),
														},
													},
												},
												Description: `Staging data will be stored in plaintext.`,
											},
											"destination_snowflake_data_staging_method_aws_s3_staging_encryption_aes_cbc_envelope_encryption": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"aes_cbc_envelope",
															),
														},
													},
													"key_encrypting_key": schema.StringAttribute{
														Optional: true,
													},
												},
												Description: `Staging data will be encrypted using AES-CBC envelope encryption.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"file_name_pattern": schema.StringAttribute{
										Optional: true,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3 Staging",
											),
										},
									},
									"purge_staging_data": schema.BoolAttribute{
										Optional: true,
									},
									"s3_bucket_name": schema.StringAttribute{
										Required: true,
									},
									"s3_bucket_region": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
												"us-east-1",
												"us-east-2",
												"us-west-1",
												"us-west-2",
												"af-south-1",
												"ap-east-1",
												"ap-south-1",
												"ap-northeast-1",
												"ap-northeast-2",
												"ap-northeast-3",
												"ap-southeast-1",
												"ap-southeast-2",
												"ca-central-1",
												"cn-north-1",
												"cn-northwest-1",
												"eu-central-1",
												"eu-west-1",
												"eu-west-2",
												"eu-west-3",
												"eu-south-1",
												"eu-north-1",
												"sa-east-1",
												"me-south-1",
											),
										},
										Description: `Enter the region where your S3 bucket resides`,
									},
									"secret_access_key": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_data_staging_method_google_cloud_storage_staging": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"bucket_name": schema.StringAttribute{
										Required: true,
									},
									"credentials_json": schema.StringAttribute{
										Required: true,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"GCS Staging",
											),
										},
									},
									"project_id": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_update_data_staging_method_select_another_option": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
									},
								},
								Description: `Select another option`,
							},
							"destination_snowflake_update_data_staging_method_recommended_internal_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Internal Staging",
											),
										},
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_update_data_staging_method_aws_s3_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Computed: true,
									},
									"encryption": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"destination_snowflake_update_data_staging_method_aws_s3_staging_encryption_no_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"none",
															),
														},
													},
												},
												Description: `Staging data will be stored in plaintext.`,
											},
											"destination_snowflake_update_data_staging_method_aws_s3_staging_encryption_aes_cbc_envelope_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"aes_cbc_envelope",
															),
														},
													},
													"key_encrypting_key": schema.StringAttribute{
														Computed: true,
													},
												},
												Description: `Staging data will be encrypted using AES-CBC envelope encryption.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"file_name_pattern": schema.StringAttribute{
										Computed: true,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3 Staging",
											),
										},
									},
									"purge_staging_data": schema.BoolAttribute{
										Computed: true,
									},
									"s3_bucket_name": schema.StringAttribute{
										Computed: true,
									},
									"s3_bucket_region": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
												"us-east-1",
												"us-east-2",
												"us-west-1",
												"us-west-2",
												"af-south-1",
												"ap-east-1",
												"ap-south-1",
												"ap-northeast-1",
												"ap-northeast-2",
												"ap-northeast-3",
												"ap-southeast-1",
												"ap-southeast-2",
												"ca-central-1",
												"cn-north-1",
												"cn-northwest-1",
												"eu-central-1",
												"eu-west-1",
												"eu-west-2",
												"eu-west-3",
												"eu-south-1",
												"eu-north-1",
												"sa-east-1",
												"me-south-1",
											),
										},
										Description: `Enter the region where your S3 bucket resides`,
									},
									"secret_access_key": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_update_data_staging_method_google_cloud_storage_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"bucket_name": schema.StringAttribute{
										Computed: true,
									},
									"credentials_json": schema.StringAttribute{
										Computed: true,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"GCS Staging",
											),
										},
									},
									"project_id": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"role": schema.StringAttribute{
						Required: true,
					},
					"schema": schema.StringAttribute{
						Required: true,
					},
					"username": schema.StringAttribute{
						Required: true,
					},
					"warehouse": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *DestinationSnowflakeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationSnowflakeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationSnowflakeResourceModel
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
	res, err := r.client.Destinations.CreateDestinationSnowflake(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationSnowflakeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationSnowflakeResourceModel
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

func (r *DestinationSnowflakeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationSnowflakeResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationSnowflakePutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationSnowflakeRequest{
		DestinationSnowflakePutRequest: destinationSnowflakePutRequest,
		DestinationID:                  destinationID,
	}
	res, err := r.client.Destinations.PutDestinationSnowflake(ctx, request)
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationSnowflakeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationSnowflakeResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.DeleteDestinationSnowflakeRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationSnowflake(ctx, request)
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

func (r *DestinationSnowflakeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
