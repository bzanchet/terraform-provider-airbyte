// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceZoomResourceModel) ToCreateSDKType() *shared.SourceZoomCreateRequest {
	jwtToken := r.Configuration.JwtToken.ValueString()
	sourceType := shared.SourceZoomZoomEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceZoom{
		JwtToken:   jwtToken,
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
	out := shared.SourceZoomCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZoomResourceModel) ToDeleteSDKType() *shared.SourceZoomCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
