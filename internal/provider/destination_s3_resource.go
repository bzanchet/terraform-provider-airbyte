// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_objectplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationS3Resource{}
var _ resource.ResourceWithImportState = &DestinationS3Resource{}

func NewDestinationS3Resource() resource.Resource {
	return &DestinationS3Resource{}
}

// DestinationS3Resource defines the resource implementation.
type DestinationS3Resource struct {
	client *sdk.SDK
}

// DestinationS3ResourceModel describes the resource data model.
type DestinationS3ResourceModel struct {
	Configuration   tfTypes.DestinationS3 `tfsdk:"configuration"`
	DefinitionID    types.String          `tfsdk:"definition_id"`
	DestinationID   types.String          `tfsdk:"destination_id"`
	DestinationType types.String          `tfsdk:"destination_type"`
	Name            types.String          `tfsdk:"name"`
	WorkspaceID     types.String          `tfsdk:"workspace_id"`
}

func (r *DestinationS3Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_s3"
}

func (r *DestinationS3Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationS3 Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"access_key_id": schema.StringAttribute{
						Optional:    true,
						Sensitive:   true,
						Description: `The access key ID to access the S3 bucket. Airbyte requires Read and Write permissions to the given bucket. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>.`,
					},
					"file_name_pattern": schema.StringAttribute{
						Optional:    true,
						Description: `The pattern allows you to set the file-name format for the S3 staging file(s)`,
					},
					"format": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"avro_apache_avro": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"compression_codec": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"bzip2": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("bzip2"),
														Description: `must be one of ["bzip2"]; Default: "bzip2"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"bzip2",
															),
														},
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("deflate"),
														path.MatchRelative().AtParent().AtName("no_compression"),
														path.MatchRelative().AtParent().AtName("snappy"),
														path.MatchRelative().AtParent().AtName("xz"),
														path.MatchRelative().AtParent().AtName("zstandard"),
													}...),
												},
											},
											"deflate": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("Deflate"),
														Description: `must be one of ["Deflate"]; Default: "Deflate"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"Deflate",
															),
														},
													},
													"compression_level": schema.Int64Attribute{
														Computed:    true,
														Optional:    true,
														Default:     int64default.StaticInt64(0),
														Description: `0: no compression & fastest, 9: best compression & slowest. Default: 0`,
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("bzip2"),
														path.MatchRelative().AtParent().AtName("no_compression"),
														path.MatchRelative().AtParent().AtName("snappy"),
														path.MatchRelative().AtParent().AtName("xz"),
														path.MatchRelative().AtParent().AtName("zstandard"),
													}...),
												},
											},
											"no_compression": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("no compression"),
														Description: `must be one of ["no compression"]; Default: "no compression"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"no compression",
															),
														},
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("bzip2"),
														path.MatchRelative().AtParent().AtName("deflate"),
														path.MatchRelative().AtParent().AtName("snappy"),
														path.MatchRelative().AtParent().AtName("xz"),
														path.MatchRelative().AtParent().AtName("zstandard"),
													}...),
												},
											},
											"snappy": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("snappy"),
														Description: `must be one of ["snappy"]; Default: "snappy"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"snappy",
															),
														},
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("bzip2"),
														path.MatchRelative().AtParent().AtName("deflate"),
														path.MatchRelative().AtParent().AtName("no_compression"),
														path.MatchRelative().AtParent().AtName("xz"),
														path.MatchRelative().AtParent().AtName("zstandard"),
													}...),
												},
											},
											"xz": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("xz"),
														Description: `must be one of ["xz"]; Default: "xz"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"xz",
															),
														},
													},
													"compression_level": schema.Int64Attribute{
														Computed:    true,
														Optional:    true,
														Default:     int64default.StaticInt64(6),
														Description: `See <a href="https://commons.apache.org/proper/commons-compress/apidocs/org/apache/commons/compress/compressors/xz/XZCompressorOutputStream.html#XZCompressorOutputStream-java.io.OutputStream-int-">here</a> for details. Default: 6`,
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("bzip2"),
														path.MatchRelative().AtParent().AtName("deflate"),
														path.MatchRelative().AtParent().AtName("no_compression"),
														path.MatchRelative().AtParent().AtName("snappy"),
														path.MatchRelative().AtParent().AtName("zstandard"),
													}...),
												},
											},
											"zstandard": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("zstandard"),
														Description: `must be one of ["zstandard"]; Default: "zstandard"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"zstandard",
															),
														},
													},
													"compression_level": schema.Int64Attribute{
														Computed:    true,
														Optional:    true,
														Default:     int64default.StaticInt64(3),
														Description: `Negative levels are 'fast' modes akin to lz4 or snappy, levels above 9 are generally for archival purposes, and levels above 18 use a lot of memory. Default: 3`,
													},
													"include_checksum": schema.BoolAttribute{
														Computed:    true,
														Optional:    true,
														Default:     booldefault.StaticBool(false),
														Description: `If true, include a checksum with each data block. Default: false`,
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("bzip2"),
														path.MatchRelative().AtParent().AtName("deflate"),
														path.MatchRelative().AtParent().AtName("no_compression"),
														path.MatchRelative().AtParent().AtName("snappy"),
														path.MatchRelative().AtParent().AtName("xz"),
													}...),
												},
											},
										},
										Description: `The compression algorithm used to compress data. Default to no compression.`,
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"format_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Avro"),
										Description: `must be one of ["Avro"]; Default: "Avro"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Avro",
											),
										},
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("csv_comma_separated_values"),
										path.MatchRelative().AtParent().AtName("json_lines_newline_delimited_json"),
										path.MatchRelative().AtParent().AtName("parquet_columnar_storage"),
									}...),
								},
							},
							"csv_comma_separated_values": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"compression": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"gzip": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("GZIP"),
														Description: `must be one of ["GZIP"]; Default: "GZIP"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"GZIP",
															),
														},
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("no_compression"),
													}...),
												},
											},
											"no_compression": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("No Compression"),
														Description: `must be one of ["No Compression"]; Default: "No Compression"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"No Compression",
															),
														},
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("gzip"),
													}...),
												},
											},
										},
										Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz").`,
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"flattening": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("No flattening"),
										Description: `Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details. must be one of ["No flattening", "Root level flattening"]; Default: "No flattening"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"No flattening",
												"Root level flattening",
											),
										},
									},
									"format_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("CSV"),
										Description: `must be one of ["CSV"]; Default: "CSV"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CSV",
											),
										},
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("avro_apache_avro"),
										path.MatchRelative().AtParent().AtName("json_lines_newline_delimited_json"),
										path.MatchRelative().AtParent().AtName("parquet_columnar_storage"),
									}...),
								},
							},
							"json_lines_newline_delimited_json": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"compression": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"gzip": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("GZIP"),
														Description: `must be one of ["GZIP"]; Default: "GZIP"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"GZIP",
															),
														},
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("no_compression"),
													}...),
												},
											},
											"no_compression": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Computed:    true,
														Optional:    true,
														Default:     stringdefault.StaticString("No Compression"),
														Description: `must be one of ["No Compression"]; Default: "No Compression"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"No Compression",
															),
														},
													},
												},
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("gzip"),
													}...),
												},
											},
										},
										Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"flattening": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("No flattening"),
										Description: `Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details. must be one of ["No flattening", "Root level flattening"]; Default: "No flattening"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"No flattening",
												"Root level flattening",
											),
										},
									},
									"format_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("JSONL"),
										Description: `must be one of ["JSONL"]; Default: "JSONL"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"JSONL",
											),
										},
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("avro_apache_avro"),
										path.MatchRelative().AtParent().AtName("csv_comma_separated_values"),
										path.MatchRelative().AtParent().AtName("parquet_columnar_storage"),
									}...),
								},
							},
							"parquet_columnar_storage": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"block_size_mb": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(128),
										Description: `This is the size of a row group being buffered in memory. It limits the memory usage when writing. Larger values will improve the IO when reading, but consume more memory when writing. Default: 128 MB. Default: 128`,
									},
									"compression_codec": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("UNCOMPRESSED"),
										Description: `The compression algorithm used to compress data pages. must be one of ["UNCOMPRESSED", "SNAPPY", "GZIP", "LZO", "BROTLI", "LZ4", "ZSTD"]; Default: "UNCOMPRESSED"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"UNCOMPRESSED",
												"SNAPPY",
												"GZIP",
												"LZO",
												"BROTLI",
												"LZ4",
												"ZSTD",
											),
										},
									},
									"dictionary_encoding": schema.BoolAttribute{
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(true),
										Description: `Default: true. Default: true`,
									},
									"dictionary_page_size_kb": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(1024),
										Description: `There is one dictionary page per column per row group when dictionary encoding is used. The dictionary page size works like the page size but for dictionary. Default: 1024 KB. Default: 1024`,
									},
									"format_type": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Parquet"),
										Description: `must be one of ["Parquet"]; Default: "Parquet"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Parquet",
											),
										},
									},
									"max_padding_size_mb": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(8),
										Description: `Maximum size allowed as padding to align row groups. This is also the minimum size of a row group. Default: 8 MB. Default: 8`,
									},
									"page_size_kb": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(1024),
										Description: `The page size is for compression. A block is composed of pages. A page is the smallest unit that must be read fully to access a single record. If this value is too small, the compression will deteriorate. Default: 1024 KB. Default: 1024`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("avro_apache_avro"),
										path.MatchRelative().AtParent().AtName("csv_comma_separated_values"),
										path.MatchRelative().AtParent().AtName("json_lines_newline_delimited_json"),
									}...),
								},
							},
						},
						Description: `Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"role_arn": schema.StringAttribute{
						Optional:    true,
						Description: `The Role ARN`,
					},
					"s3_bucket_name": schema.StringAttribute{
						Required:    true,
						Description: `The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.`,
					},
					"s3_bucket_path": schema.StringAttribute{
						Required:    true,
						Description: `Directory under the S3 bucket where data will be written. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=to%20format%20the-,bucket%20path,-%3A">here</a>`,
					},
					"s3_bucket_region": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(""),
						Description: `The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes. must be one of ["", "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-4", "ca-central-1", "ca-west-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-central-2", "eu-north-1", "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3", "il-central-1", "me-central-1", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"]; Default: ""`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"af-south-1",
								"ap-east-1",
								"ap-northeast-1",
								"ap-northeast-2",
								"ap-northeast-3",
								"ap-south-1",
								"ap-south-2",
								"ap-southeast-1",
								"ap-southeast-2",
								"ap-southeast-3",
								"ap-southeast-4",
								"ca-central-1",
								"ca-west-1",
								"cn-north-1",
								"cn-northwest-1",
								"eu-central-1",
								"eu-central-2",
								"eu-north-1",
								"eu-south-1",
								"eu-south-2",
								"eu-west-1",
								"eu-west-2",
								"eu-west-3",
								"il-central-1",
								"me-central-1",
								"me-south-1",
								"sa-east-1",
								"us-east-1",
								"us-east-2",
								"us-gov-east-1",
								"us-gov-west-1",
								"us-west-1",
								"us-west-2",
							),
						},
					},
					"s3_endpoint": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(""),
						Description: `Your S3 endpoint url. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/s3.html#:~:text=Service%20endpoints-,Amazon%20S3%20endpoints,-When%20you%20use">here</a>. Default: ""`,
					},
					"s3_path_format": schema.StringAttribute{
						Optional:    true,
						Description: `Format string on how data will be organized inside the S3 bucket directory. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=The%20full%20path%20of%20the%20output%20data%20with%20the%20default%20S3%20path%20format">here</a>`,
					},
					"secret_access_key": schema.StringAttribute{
						Optional:    true,
						Sensitive:   true,
						Description: `The corresponding secret to the access key ID. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>`,
					},
				},
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed. `,
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Name of the destination e.g. dev-mysql-instance.`,
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
		},
	}
}

func (r *DestinationS3Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationS3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationS3ResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedDestinationS3CreateRequest()
	res, err := r.client.Destinations.CreateDestinationS3(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
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
	if !(res.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	destinationID := data.DestinationID.ValueString()
	request1 := operations.GetDestinationS3Request{
		DestinationID: destinationID,
	}
	res1, err := r.client.Destinations.GetDestinationS3(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationS3Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationS3ResourceModel
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
	request := operations.GetDestinationS3Request{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationS3(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationS3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationS3ResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationS3PutRequest := data.ToSharedDestinationS3PutRequest()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationS3Request{
		DestinationS3PutRequest: destinationS3PutRequest,
		DestinationID:           destinationID,
	}
	res, err := r.client.Destinations.PutDestinationS3(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
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
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	destinationId1 := data.DestinationID.ValueString()
	request1 := operations.GetDestinationS3Request{
		DestinationID: destinationId1,
	}
	res1, err := r.client.Destinations.GetDestinationS3(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.DestinationResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationS3Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationS3ResourceModel
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
	request := operations.DeleteDestinationS3Request{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationS3(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
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

}

func (r *DestinationS3Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
