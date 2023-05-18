// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceShortioResourceModel) ToCreateSDKType() *shared.SourceShortioCreateRequest {
	domainID := r.Configuration.DomainID.ValueString()
	secretKey := r.Configuration.SecretKey.ValueString()
	sourceType := shared.SourceShortioShortioEnum(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceShortio{
		DomainID:   domainID,
		SecretKey:  secretKey,
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
	out := shared.SourceShortioCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceShortioResourceModel) ToDeleteSDKType() *shared.SourceShortioCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
