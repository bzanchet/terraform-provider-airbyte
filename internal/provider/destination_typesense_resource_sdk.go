// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationTypesenseResourceModel) ToSharedDestinationTypesenseCreateRequest() *shared.DestinationTypesenseCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	host := r.Configuration.Host.ValueString()
	port := new(string)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueString()
	} else {
		port = nil
	}
	protocol := new(string)
	if !r.Configuration.Protocol.IsUnknown() && !r.Configuration.Protocol.IsNull() {
		*protocol = r.Configuration.Protocol.ValueString()
	} else {
		protocol = nil
	}
	configuration := shared.DestinationTypesense{
		APIKey:    apiKey,
		BatchSize: batchSize,
		Host:      host,
		Port:      port,
		Protocol:  protocol,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationTypesenseCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationTypesenseResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationTypesenseResourceModel) ToSharedDestinationTypesensePutRequest() *shared.DestinationTypesensePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	host := r.Configuration.Host.ValueString()
	port := new(string)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueString()
	} else {
		port = nil
	}
	protocol := new(string)
	if !r.Configuration.Protocol.IsUnknown() && !r.Configuration.Protocol.IsNull() {
		*protocol = r.Configuration.Protocol.ValueString()
	} else {
		protocol = nil
	}
	configuration := shared.DestinationTypesenseUpdate{
		APIKey:    apiKey,
		BatchSize: batchSize,
		Host:      host,
		Port:      port,
		Protocol:  protocol,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationTypesensePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
