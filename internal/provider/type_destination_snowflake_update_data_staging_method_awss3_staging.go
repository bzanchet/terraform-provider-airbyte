// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationSnowflakeUpdateDataStagingMethodAWSS3Staging struct {
	AccessKeyID      types.String                                                       `tfsdk:"access_key_id"`
	Encryption       *DestinationSnowflakeUpdateDataStagingMethodAWSS3StagingEncryption `tfsdk:"encryption"`
	FileNamePattern  types.String                                                       `tfsdk:"file_name_pattern"`
	Method           types.String                                                       `tfsdk:"method"`
	PurgeStagingData types.Bool                                                         `tfsdk:"purge_staging_data"`
	S3BucketName     types.String                                                       `tfsdk:"s3_bucket_name"`
	S3BucketRegion   types.String                                                       `tfsdk:"s3_bucket_region"`
	SecretAccessKey  types.String                                                       `tfsdk:"secret_access_key"`
}
