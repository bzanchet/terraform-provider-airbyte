// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceBraintreeResourceModel) ToCreateSDKType() *shared.SourceBraintreeCreateRequest {
	environment := shared.SourceBraintreeEnvironment(r.Configuration.Environment.ValueString())
	merchantID := r.Configuration.MerchantID.ValueString()
	privateKey := r.Configuration.PrivateKey.ValueString()
	publicKey := r.Configuration.PublicKey.ValueString()
	sourceType := shared.SourceBraintreeBraintree(r.Configuration.SourceType.ValueString())
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceBraintree{
		Environment: environment,
		MerchantID:  merchantID,
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
		SourceType:  sourceType,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBraintreeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBraintreeResourceModel) ToUpdateSDKType() *shared.SourceBraintreePutRequest {
	environment := shared.SourceBraintreeUpdateEnvironment(r.Configuration.Environment.ValueString())
	merchantID := r.Configuration.MerchantID.ValueString()
	privateKey := r.Configuration.PrivateKey.ValueString()
	publicKey := r.Configuration.PublicKey.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceBraintreeUpdate{
		Environment: environment,
		MerchantID:  merchantID,
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBraintreePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBraintreeResourceModel) ToDeleteSDKType() *shared.SourceBraintreeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceBraintreeResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
