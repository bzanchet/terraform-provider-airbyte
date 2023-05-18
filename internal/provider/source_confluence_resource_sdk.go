// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceConfluenceResourceModel) ToCreateSDKType() *shared.SourceConfluenceCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	domainName := r.Configuration.DomainName.ValueString()
	email := r.Configuration.Email.ValueString()
	sourceType := shared.SourceConfluenceConfluenceEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceConfluence{
		APIToken:   apiToken,
		DomainName: domainName,
		Email:      email,
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
	out := shared.SourceConfluenceCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceConfluenceResourceModel) ToDeleteSDKType() *shared.SourceConfluenceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
