// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceRecruiteeResourceModel) ToSharedSourceRecruiteeCreateRequest() *shared.SourceRecruiteeCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	companyID := r.Configuration.CompanyID.ValueInt64()
	configuration := shared.SourceRecruitee{
		APIKey:    apiKey,
		CompanyID: companyID,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecruiteeCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecruiteeResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceRecruiteeResourceModel) ToSharedSourceRecruiteePutRequest() *shared.SourceRecruiteePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	companyID := r.Configuration.CompanyID.ValueInt64()
	configuration := shared.SourceRecruiteeUpdate{
		APIKey:    apiKey,
		CompanyID: companyID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecruiteePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
