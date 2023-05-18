// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *DestinationConvexResourceModel) ToCreateSDKType() *shared.DestinationConvexCreateRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	deploymentURL := r.Configuration.DeploymentURL.ValueString()
	destinationType := shared.DestinationConvexConvexEnum(r.Configuration.DestinationType.ValueString())
	configuration := shared.DestinationConvex{
		AccessKey:       accessKey,
		DeploymentURL:   deploymentURL,
		DestinationType: destinationType,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationConvexCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationConvexResourceModel) ToDeleteSDKType() *shared.DestinationConvexCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
