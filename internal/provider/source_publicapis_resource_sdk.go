// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePublicApisResourceModel) ToCreateSDKType() *shared.SourcePublicApisCreateRequest {
	sourceType := shared.SourcePublicApisPublicApis(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePublicApis{
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
	out := shared.SourcePublicApisCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePublicApisResourceModel) ToUpdateSDKType() *shared.SourcePublicApisPutRequest {
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePublicApisPutRequest{
		Name:        name,
		WorkspaceID: workspaceID,
	}
	return &out
}

func (r *SourcePublicApisResourceModel) ToDeleteSDKType() *shared.SourcePublicApisCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePublicApisResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
