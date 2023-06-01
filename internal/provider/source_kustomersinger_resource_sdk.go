// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceKustomerSingerResourceModel) ToCreateSDKType() *shared.SourceKustomerSingerCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	sourceType := shared.SourceKustomerSingerKustomerSinger(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceKustomerSinger{
		APIToken:   apiToken,
		SourceType: sourceType,
		StartDate:  startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceKustomerSingerCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKustomerSingerResourceModel) ToUpdateSDKType() *shared.SourceKustomerSingerPutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceKustomerSingerUpdate{
		APIToken:  apiToken,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceKustomerSingerPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKustomerSingerResourceModel) ToDeleteSDKType() *shared.SourceKustomerSingerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceKustomerSingerResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
