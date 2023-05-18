// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceEmailoctopusResourceModel) ToCreateSDKType() *shared.SourceEmailoctopusCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceEmailoctopusEmailoctopusEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceEmailoctopus{
		APIKey:     apiKey,
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
	out := shared.SourceEmailoctopusCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceEmailoctopusResourceModel) ToDeleteSDKType() *shared.SourceEmailoctopusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
