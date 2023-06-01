// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSalesforceSingerResourceModel) ToCreateSDKType() *shared.SourceSalesforceSingerCreateRequest {
	apiType := shared.SourceSalesforceSingerAPIType(r.Configuration.APIType.ValueString())
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	isSandbox := new(bool)
	if !r.Configuration.IsSandbox.IsUnknown() && !r.Configuration.IsSandbox.IsNull() {
		*isSandbox = r.Configuration.IsSandbox.ValueBool()
	} else {
		isSandbox = nil
	}
	quotaPercentPerRun := new(float64)
	if !r.Configuration.QuotaPercentPerRun.IsUnknown() && !r.Configuration.QuotaPercentPerRun.IsNull() {
		*quotaPercentPerRun, _ = r.Configuration.QuotaPercentPerRun.ValueBigFloat().Float64()
	} else {
		quotaPercentPerRun = nil
	}
	quotaPercentTotal := new(float64)
	if !r.Configuration.QuotaPercentTotal.IsUnknown() && !r.Configuration.QuotaPercentTotal.IsNull() {
		*quotaPercentTotal, _ = r.Configuration.QuotaPercentTotal.ValueBigFloat().Float64()
	} else {
		quotaPercentTotal = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	sourceType := shared.SourceSalesforceSingerSalesforceSinger(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceSalesforceSinger{
		APIType:            apiType,
		ClientID:           clientID,
		ClientSecret:       clientSecret,
		IsSandbox:          isSandbox,
		QuotaPercentPerRun: quotaPercentPerRun,
		QuotaPercentTotal:  quotaPercentTotal,
		RefreshToken:       refreshToken,
		SourceType:         sourceType,
		StartDate:          startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSalesforceSingerCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSalesforceSingerResourceModel) ToUpdateSDKType() *shared.SourceSalesforceSingerPutRequest {
	apiType := shared.SourceSalesforceSingerUpdateAPIType(r.Configuration.APIType.ValueString())
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	isSandbox := new(bool)
	if !r.Configuration.IsSandbox.IsUnknown() && !r.Configuration.IsSandbox.IsNull() {
		*isSandbox = r.Configuration.IsSandbox.ValueBool()
	} else {
		isSandbox = nil
	}
	quotaPercentPerRun := new(float64)
	if !r.Configuration.QuotaPercentPerRun.IsUnknown() && !r.Configuration.QuotaPercentPerRun.IsNull() {
		*quotaPercentPerRun, _ = r.Configuration.QuotaPercentPerRun.ValueBigFloat().Float64()
	} else {
		quotaPercentPerRun = nil
	}
	quotaPercentTotal := new(float64)
	if !r.Configuration.QuotaPercentTotal.IsUnknown() && !r.Configuration.QuotaPercentTotal.IsNull() {
		*quotaPercentTotal, _ = r.Configuration.QuotaPercentTotal.ValueBigFloat().Float64()
	} else {
		quotaPercentTotal = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceSalesforceSingerUpdate{
		APIType:            apiType,
		ClientID:           clientID,
		ClientSecret:       clientSecret,
		IsSandbox:          isSandbox,
		QuotaPercentPerRun: quotaPercentPerRun,
		QuotaPercentTotal:  quotaPercentTotal,
		RefreshToken:       refreshToken,
		StartDate:          startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSalesforceSingerPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSalesforceSingerResourceModel) ToDeleteSDKType() *shared.SourceSalesforceSingerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSalesforceSingerResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
