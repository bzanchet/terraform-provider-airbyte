// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceSmartengageResourceModel) ToCreateSDKType() *shared.SourceSmartengageCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceSmartengageSmartengageEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceSmartengage{
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
	out := shared.SourceSmartengageCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmartengageResourceModel) ToDeleteSDKType() *shared.SourceSmartengageCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
