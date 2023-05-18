// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourcePipedriveResourceModel) ToCreateSDKType() *shared.SourcePipedriveCreateRequest {
	var authorization *shared.SourcePipedriveAPIKeyAuthentication
	if r.Configuration.Authorization != nil {
		apiToken := r.Configuration.Authorization.APIToken.ValueString()
		authType := shared.SourcePipedriveAPIKeyAuthenticationAuthTypeEnum(r.Configuration.Authorization.AuthType.ValueString())
		authorization = &shared.SourcePipedriveAPIKeyAuthentication{
			APIToken: apiToken,
			AuthType: authType,
		}
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	sourceType := shared.SourcePipedrivePipedriveEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePipedrive{
		Authorization:        authorization,
		ReplicationStartDate: replicationStartDate,
		SourceType:           sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePipedriveCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePipedriveResourceModel) ToDeleteSDKType() *shared.SourcePipedriveCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
