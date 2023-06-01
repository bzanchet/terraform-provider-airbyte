// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSmailyResourceModel) ToCreateSDKType() *shared.SourceSmailyCreateRequest {
	apiPassword := r.Configuration.APIPassword.ValueString()
	apiSubdomain := r.Configuration.APISubdomain.ValueString()
	apiUsername := r.Configuration.APIUsername.ValueString()
	sourceType := shared.SourceSmailySmaily(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceSmaily{
		APIPassword:  apiPassword,
		APISubdomain: apiSubdomain,
		APIUsername:  apiUsername,
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
	out := shared.SourceSmailyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmailyResourceModel) ToUpdateSDKType() *shared.SourceSmailyPutRequest {
	apiPassword := r.Configuration.APIPassword.ValueString()
	apiSubdomain := r.Configuration.APISubdomain.ValueString()
	apiUsername := r.Configuration.APIUsername.ValueString()
	configuration := shared.SourceSmailyUpdate{
		APIPassword:  apiPassword,
		APISubdomain: apiSubdomain,
		APIUsername:  apiUsername,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmailyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmailyResourceModel) ToDeleteSDKType() *shared.SourceSmailyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSmailyResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
