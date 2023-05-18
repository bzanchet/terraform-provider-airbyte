// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceZohoCrmResourceModel) ToCreateSDKType() *shared.SourceZohoCrmCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dcRegion := shared.SourceZohoCrmDataCenterLocationEnum(r.Configuration.DcRegion.ValueString())
	edition := shared.SourceZohoCRMZohoCRMEditionEnum(r.Configuration.Edition.ValueString())
	environment := shared.SourceZohoCrmEnvironmentEnum(r.Configuration.Environment.ValueString())
	refreshToken := r.Configuration.RefreshToken.ValueString()
	sourceType := shared.SourceZohoCrmZohoCrmEnum(r.Configuration.SourceType.ValueString())
	startDatetime := new(time.Time)
	if !r.Configuration.StartDatetime.IsUnknown() && !r.Configuration.StartDatetime.IsNull() {
		*startDatetime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDatetime.ValueString())
	} else {
		startDatetime = nil
	}
	configuration := shared.SourceZohoCrm{
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		DcRegion:      dcRegion,
		Edition:       edition,
		Environment:   environment,
		RefreshToken:  refreshToken,
		SourceType:    sourceType,
		StartDatetime: startDatetime,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZohoCrmCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZohoCrmResourceModel) ToDeleteSDKType() *shared.SourceZohoCrmCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
