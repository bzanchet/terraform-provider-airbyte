// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceConvexResourceModel) ToSharedSourceConvexCreateRequest() *shared.SourceConvexCreateRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	deploymentURL := r.Configuration.DeploymentURL.ValueString()
	configuration := shared.SourceConvex{
		AccessKey:     accessKey,
		DeploymentURL: deploymentURL,
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
	out := shared.SourceConvexCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceConvexResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceConvexResourceModel) ToSharedSourceConvexPutRequest() *shared.SourceConvexPutRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	deploymentURL := r.Configuration.DeploymentURL.ValueString()
	configuration := shared.SourceConvexUpdate{
		AccessKey:     accessKey,
		DeploymentURL: deploymentURL,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceConvexPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
