// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceYandexMetricaResourceModel) ToCreateSDKType() *shared.SourceYandexMetricaCreateRequest {
	authToken := r.Configuration.AuthToken.ValueString()
	counterID := r.Configuration.CounterID.ValueString()
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	sourceType := shared.SourceYandexMetricaYandexMetrica(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceYandexMetrica{
		AuthToken:  authToken,
		CounterID:  counterID,
		EndDate:    endDate,
		SourceType: sourceType,
		StartDate:  startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceYandexMetricaCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceYandexMetricaResourceModel) ToUpdateSDKType() *shared.SourceYandexMetricaPutRequest {
	authToken := r.Configuration.AuthToken.ValueString()
	counterID := r.Configuration.CounterID.ValueString()
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceYandexMetricaUpdate{
		AuthToken: authToken,
		CounterID: counterID,
		EndDate:   endDate,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceYandexMetricaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceYandexMetricaResourceModel) ToDeleteSDKType() *shared.SourceYandexMetricaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceYandexMetricaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
