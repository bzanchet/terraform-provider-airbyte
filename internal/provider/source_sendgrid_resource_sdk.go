// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceSendgridResourceModel) ToCreateSDKType() *shared.SourceSendgridCreateRequest {
	apikey := r.Configuration.Apikey.ValueString()
	sourceType := shared.SourceSendgridSendgridEnum(r.Configuration.SourceType.ValueString())
	startTime := new(time.Time)
	if !r.Configuration.StartTime.IsUnknown() && !r.Configuration.StartTime.IsNull() {
		*startTime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartTime.ValueString())
	} else {
		startTime = nil
	}
	configuration := shared.SourceSendgrid{
		Apikey:     apikey,
		SourceType: sourceType,
		StartTime:  startTime,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSendgridCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSendgridResourceModel) ToDeleteSDKType() *shared.SourceSendgridCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
