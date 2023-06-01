// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationFireboltResourceModel) ToCreateSDKType() *shared.DestinationFireboltCreateRequest {
	account := new(string)
	if !r.Configuration.Account.IsUnknown() && !r.Configuration.Account.IsNull() {
		*account = r.Configuration.Account.ValueString()
	} else {
		account = nil
	}
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationFireboltFirebolt(r.Configuration.DestinationType.ValueString())
	engine := new(string)
	if !r.Configuration.Engine.IsUnknown() && !r.Configuration.Engine.IsNull() {
		*engine = r.Configuration.Engine.ValueString()
	} else {
		engine = nil
	}
	host := new(string)
	if !r.Configuration.Host.IsUnknown() && !r.Configuration.Host.IsNull() {
		*host = r.Configuration.Host.ValueString()
	} else {
		host = nil
	}
	var loadingMethod *shared.DestinationFireboltLoadingMethod
	var destinationFireboltLoadingMethodSQLInserts *shared.DestinationFireboltLoadingMethodSQLInserts
	if r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodSQLInserts != nil {
		method := shared.DestinationFireboltLoadingMethodSQLInsertsMethod(r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodSQLInserts.Method.ValueString())
		destinationFireboltLoadingMethodSQLInserts = &shared.DestinationFireboltLoadingMethodSQLInserts{
			Method: method,
		}
	}
	if destinationFireboltLoadingMethodSQLInserts != nil {
		loadingMethod = &shared.DestinationFireboltLoadingMethod{
			DestinationFireboltLoadingMethodSQLInserts: destinationFireboltLoadingMethodSQLInserts,
		}
	}
	var destinationFireboltLoadingMethodExternalTableViaS3 *shared.DestinationFireboltLoadingMethodExternalTableViaS3
	if r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3 != nil {
		awsKeyID := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.AwsKeyID.ValueString()
		awsKeySecret := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.AwsKeySecret.ValueString()
		method1 := shared.DestinationFireboltLoadingMethodExternalTableViaS3Method(r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.Method.ValueString())
		s3Bucket := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.S3Bucket.ValueString()
		s3Region := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.S3Region.ValueString()
		destinationFireboltLoadingMethodExternalTableViaS3 = &shared.DestinationFireboltLoadingMethodExternalTableViaS3{
			AwsKeyID:     awsKeyID,
			AwsKeySecret: awsKeySecret,
			Method:       method1,
			S3Bucket:     s3Bucket,
			S3Region:     s3Region,
		}
	}
	if destinationFireboltLoadingMethodExternalTableViaS3 != nil {
		loadingMethod = &shared.DestinationFireboltLoadingMethod{
			DestinationFireboltLoadingMethodExternalTableViaS3: destinationFireboltLoadingMethodExternalTableViaS3,
		}
	}
	password := r.Configuration.Password.ValueString()
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationFirebolt{
		Account:         account,
		Database:        database,
		DestinationType: destinationType,
		Engine:          engine,
		Host:            host,
		LoadingMethod:   loadingMethod,
		Password:        password,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationFireboltCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationFireboltResourceModel) ToUpdateSDKType() *shared.DestinationFireboltPutRequest {
	account := new(string)
	if !r.Configuration.Account.IsUnknown() && !r.Configuration.Account.IsNull() {
		*account = r.Configuration.Account.ValueString()
	} else {
		account = nil
	}
	database := r.Configuration.Database.ValueString()
	engine := new(string)
	if !r.Configuration.Engine.IsUnknown() && !r.Configuration.Engine.IsNull() {
		*engine = r.Configuration.Engine.ValueString()
	} else {
		engine = nil
	}
	host := new(string)
	if !r.Configuration.Host.IsUnknown() && !r.Configuration.Host.IsNull() {
		*host = r.Configuration.Host.ValueString()
	} else {
		host = nil
	}
	var loadingMethod *shared.DestinationFireboltUpdateLoadingMethod
	var destinationFireboltUpdateLoadingMethodSQLInserts *shared.DestinationFireboltUpdateLoadingMethodSQLInserts
	if r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodSQLInserts != nil {
		method := shared.DestinationFireboltUpdateLoadingMethodSQLInsertsMethod(r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodSQLInserts.Method.ValueString())
		destinationFireboltUpdateLoadingMethodSQLInserts = &shared.DestinationFireboltUpdateLoadingMethodSQLInserts{
			Method: method,
		}
	}
	if destinationFireboltUpdateLoadingMethodSQLInserts != nil {
		loadingMethod = &shared.DestinationFireboltUpdateLoadingMethod{
			DestinationFireboltUpdateLoadingMethodSQLInserts: destinationFireboltUpdateLoadingMethodSQLInserts,
		}
	}
	var destinationFireboltUpdateLoadingMethodExternalTableViaS3 *shared.DestinationFireboltUpdateLoadingMethodExternalTableViaS3
	if r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3 != nil {
		awsKeyID := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.AwsKeyID.ValueString()
		awsKeySecret := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.AwsKeySecret.ValueString()
		method1 := shared.DestinationFireboltUpdateLoadingMethodExternalTableViaS3Method(r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.Method.ValueString())
		s3Bucket := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.S3Bucket.ValueString()
		s3Region := r.Configuration.LoadingMethod.DestinationFireboltUpdateLoadingMethodExternalTableViaS3.S3Region.ValueString()
		destinationFireboltUpdateLoadingMethodExternalTableViaS3 = &shared.DestinationFireboltUpdateLoadingMethodExternalTableViaS3{
			AwsKeyID:     awsKeyID,
			AwsKeySecret: awsKeySecret,
			Method:       method1,
			S3Bucket:     s3Bucket,
			S3Region:     s3Region,
		}
	}
	if destinationFireboltUpdateLoadingMethodExternalTableViaS3 != nil {
		loadingMethod = &shared.DestinationFireboltUpdateLoadingMethod{
			DestinationFireboltUpdateLoadingMethodExternalTableViaS3: destinationFireboltUpdateLoadingMethodExternalTableViaS3,
		}
	}
	password := r.Configuration.Password.ValueString()
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationFireboltUpdate{
		Account:       account,
		Database:      database,
		Engine:        engine,
		Host:          host,
		LoadingMethod: loadingMethod,
		Password:      password,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationFireboltPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationFireboltResourceModel) ToDeleteSDKType() *shared.DestinationFireboltCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationFireboltResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
