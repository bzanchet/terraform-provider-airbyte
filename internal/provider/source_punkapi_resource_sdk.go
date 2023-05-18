// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourcePunkAPIResourceModel) ToCreateSDKType() *shared.SourcePunkAPICreateRequest {
	brewedAfter := r.Configuration.BrewedAfter.ValueString()
	brewedBefore := r.Configuration.BrewedBefore.ValueString()
	id := new(string)
	if !r.Configuration.ID.IsUnknown() && !r.Configuration.ID.IsNull() {
		*id = r.Configuration.ID.ValueString()
	} else {
		id = nil
	}
	sourceType := shared.SourcePunkAPIPunkAPIEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePunkAPI{
		BrewedAfter:  brewedAfter,
		BrewedBefore: brewedBefore,
		ID:           id,
		SourceType:   sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePunkAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePunkAPIResourceModel) ToDeleteSDKType() *shared.SourcePunkAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}
