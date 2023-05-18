// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
)

func (r *SourceBrazeResourceModel) ToCreateSDKType() *shared.SourceBrazeCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceBrazeBrazeEnum(r.Configuration.SourceType.ValueString())
	startDate, _ := customTypes.NewDate(r.Configuration.StartDate.ValueString())
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceBraze{
		APIKey:     apiKey,
		SourceType: sourceType,
		StartDate:  startDate,
		URL:        url,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBrazeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBrazeResourceModel) ToDeleteSDKType() *shared.SourceBrazeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
