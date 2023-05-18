// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceK6CloudResourceModel) ToCreateSDKType() *shared.SourceK6CloudCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	sourceType := shared.SourceK6CloudK6CloudEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceK6Cloud{
		APIToken:   apiToken,
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
	out := shared.SourceK6CloudCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceK6CloudResourceModel) ToDeleteSDKType() *shared.SourceK6CloudCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
