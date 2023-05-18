// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceDremioResourceModel) ToCreateSDKType() *shared.SourceDremioCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	baseURL := r.Configuration.BaseURL.ValueString()
	sourceType := shared.SourceDremioDremioEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceDremio{
		APIKey:     apiKey,
		BaseURL:    baseURL,
		SourceType: sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDremioCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDremioResourceModel) ToDeleteSDKType() *shared.SourceDremioCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
