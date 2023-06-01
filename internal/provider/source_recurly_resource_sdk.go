// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceRecurlyResourceModel) ToCreateSDKType() *shared.SourceRecurlyCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	beginTime := new(string)
	if !r.Configuration.BeginTime.IsUnknown() && !r.Configuration.BeginTime.IsNull() {
		*beginTime = r.Configuration.BeginTime.ValueString()
	} else {
		beginTime = nil
	}
	endTime := new(string)
	if !r.Configuration.EndTime.IsUnknown() && !r.Configuration.EndTime.IsNull() {
		*endTime = r.Configuration.EndTime.ValueString()
	} else {
		endTime = nil
	}
	sourceType := shared.SourceRecurlyRecurly(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceRecurly{
		APIKey:     apiKey,
		BeginTime:  beginTime,
		EndTime:    endTime,
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
	out := shared.SourceRecurlyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecurlyResourceModel) ToUpdateSDKType() *shared.SourceRecurlyPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	beginTime := new(string)
	if !r.Configuration.BeginTime.IsUnknown() && !r.Configuration.BeginTime.IsNull() {
		*beginTime = r.Configuration.BeginTime.ValueString()
	} else {
		beginTime = nil
	}
	endTime := new(string)
	if !r.Configuration.EndTime.IsUnknown() && !r.Configuration.EndTime.IsNull() {
		*endTime = r.Configuration.EndTime.ValueString()
	} else {
		endTime = nil
	}
	configuration := shared.SourceRecurlyUpdate{
		APIKey:    apiKey,
		BeginTime: beginTime,
		EndTime:   endTime,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecurlyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecurlyResourceModel) ToDeleteSDKType() *shared.SourceRecurlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRecurlyResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
