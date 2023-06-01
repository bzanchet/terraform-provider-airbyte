// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePendoResourceModel) ToCreateSDKType() *shared.SourcePendoCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourcePendoPendo(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePendo{
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
	out := shared.SourcePendoCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePendoResourceModel) ToUpdateSDKType() *shared.SourcePendoPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourcePendoUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePendoPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePendoResourceModel) ToDeleteSDKType() *shared.SourcePendoCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePendoResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
