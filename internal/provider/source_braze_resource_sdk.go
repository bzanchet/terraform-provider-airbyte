// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceBrazeResourceModel) ToCreateSDKType() *shared.SourceBrazeCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceBrazeBraze(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
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

func (r *SourceBrazeResourceModel) ToUpdateSDKType() *shared.SourceBrazePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceBrazeUpdate{
		APIKey:    apiKey,
		StartDate: startDate,
		URL:       url,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBrazePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBrazeResourceModel) ToDeleteSDKType() *shared.SourceBrazeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceBrazeResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
