// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceConfluenceResourceModel) ToCreateSDKType() *shared.SourceConfluenceCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	domainName := r.Configuration.DomainName.ValueString()
	email := r.Configuration.Email.ValueString()
	sourceType := shared.SourceConfluenceConfluence(r.Configuration.SourceType.ValueString())
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

func (r *SourceConfluenceResourceModel) ToUpdateSDKType() *shared.SourceConfluencePutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	domainName := r.Configuration.DomainName.ValueString()
	email := r.Configuration.Email.ValueString()
	configuration := shared.SourceConfluenceUpdate{
		APIToken:   apiToken,
		DomainName: domainName,
		Email:      email,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceConfluencePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceConfluenceResourceModel) ToDeleteSDKType() *shared.SourceConfluenceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceConfluenceResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
