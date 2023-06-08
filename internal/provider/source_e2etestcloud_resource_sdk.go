// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceE2eTestCloudResourceModel) ToCreateSDKType() *shared.SourceE2eTestCloudCreateRequest {
	maxMessages := r.Configuration.MaxMessages.ValueInt64()
	messageIntervalMs := new(int64)
	if !r.Configuration.MessageIntervalMs.IsUnknown() && !r.Configuration.MessageIntervalMs.IsNull() {
		*messageIntervalMs = r.Configuration.MessageIntervalMs.ValueInt64()
	} else {
		messageIntervalMs = nil
	}
	var mockCatalog shared.SourceE2eTestCloudMockCatalog
	var sourceE2eTestCloudMockCatalogSingleSchema *shared.SourceE2eTestCloudMockCatalogSingleSchema
	if r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogMultiSchema != nil {
		type1 := shared.SourceE2eTestCloudMockCatalogSingleSchemaType(r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogMultiSchema.Type.ValueString())
		sourceE2eTestCloudMockCatalogSingleSchema = &shared.SourceE2eTestCloudMockCatalogSingleSchema{
			Type: type1,
		}
	}
	if sourceE2eTestCloudMockCatalogSingleSchema != nil {
		mockCatalog = shared.SourceE2eTestCloudMockCatalog{
			SourceE2eTestCloudMockCatalogSingleSchema: sourceE2eTestCloudMockCatalogSingleSchema,
		}
	}
	var sourceE2eTestCloudMockCatalogMultiSchema *shared.SourceE2eTestCloudMockCatalogMultiSchema
	if r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogSingleSchema != nil {
		type2 := shared.SourceE2eTestCloudMockCatalogMultiSchemaType(r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogSingleSchema.Type.ValueString())
		sourceE2eTestCloudMockCatalogMultiSchema = &shared.SourceE2eTestCloudMockCatalogMultiSchema{
			Type: type2,
		}
	}
	if sourceE2eTestCloudMockCatalogMultiSchema != nil {
		mockCatalog = shared.SourceE2eTestCloudMockCatalog{
			SourceE2eTestCloudMockCatalogMultiSchema: sourceE2eTestCloudMockCatalogMultiSchema,
		}
	}
	seed := new(int64)
	if !r.Configuration.Seed.IsUnknown() && !r.Configuration.Seed.IsNull() {
		*seed = r.Configuration.Seed.ValueInt64()
	} else {
		seed = nil
	}
	sourceType := shared.SourceE2eTestCloudE2eTestCloud(r.Configuration.SourceType.ValueString())
	type3 := new(shared.SourceE2eTestCloudType)
	if !r.Configuration.Type.IsUnknown() && !r.Configuration.Type.IsNull() {
		*type3 = shared.SourceE2eTestCloudType(r.Configuration.Type.ValueString())
	} else {
		type3 = nil
	}
	configuration := shared.SourceE2eTestCloud{
		MaxMessages:       maxMessages,
		MessageIntervalMs: messageIntervalMs,
		MockCatalog:       mockCatalog,
		Seed:              seed,
		SourceType:        sourceType,
		Type:              type3,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceE2eTestCloudCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceE2eTestCloudResourceModel) ToUpdateSDKType() *shared.SourceE2eTestCloudPutRequest {
	maxMessages := r.Configuration.MaxMessages.ValueInt64()
	messageIntervalMs := new(int64)
	if !r.Configuration.MessageIntervalMs.IsUnknown() && !r.Configuration.MessageIntervalMs.IsNull() {
		*messageIntervalMs = r.Configuration.MessageIntervalMs.ValueInt64()
	} else {
		messageIntervalMs = nil
	}
	var mockCatalog shared.SourceE2eTestCloudUpdateMockCatalog
	var sourceE2eTestCloudUpdateMockCatalogSingleSchema *shared.SourceE2eTestCloudUpdateMockCatalogSingleSchema
	if r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogMultiSchema != nil {
		type1 := shared.SourceE2eTestCloudUpdateMockCatalogSingleSchemaType(r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogMultiSchema.Type.ValueString())
		sourceE2eTestCloudUpdateMockCatalogSingleSchema = &shared.SourceE2eTestCloudUpdateMockCatalogSingleSchema{
			Type: type1,
		}
	}
	if sourceE2eTestCloudUpdateMockCatalogSingleSchema != nil {
		mockCatalog = shared.SourceE2eTestCloudUpdateMockCatalog{
			SourceE2eTestCloudUpdateMockCatalogSingleSchema: sourceE2eTestCloudUpdateMockCatalogSingleSchema,
		}
	}
	var sourceE2eTestCloudUpdateMockCatalogMultiSchema *shared.SourceE2eTestCloudUpdateMockCatalogMultiSchema
	if r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogSingleSchema != nil {
		type2 := shared.SourceE2eTestCloudUpdateMockCatalogMultiSchemaType(r.Configuration.MockCatalog.SourceE2eTestCloudMockCatalogSingleSchema.Type.ValueString())
		sourceE2eTestCloudUpdateMockCatalogMultiSchema = &shared.SourceE2eTestCloudUpdateMockCatalogMultiSchema{
			Type: type2,
		}
	}
	if sourceE2eTestCloudUpdateMockCatalogMultiSchema != nil {
		mockCatalog = shared.SourceE2eTestCloudUpdateMockCatalog{
			SourceE2eTestCloudUpdateMockCatalogMultiSchema: sourceE2eTestCloudUpdateMockCatalogMultiSchema,
		}
	}
	seed := new(int64)
	if !r.Configuration.Seed.IsUnknown() && !r.Configuration.Seed.IsNull() {
		*seed = r.Configuration.Seed.ValueInt64()
	} else {
		seed = nil
	}
	type3 := new(shared.SourceE2eTestCloudUpdateType)
	if !r.Configuration.Type.IsUnknown() && !r.Configuration.Type.IsNull() {
		*type3 = shared.SourceE2eTestCloudUpdateType(r.Configuration.Type.ValueString())
	} else {
		type3 = nil
	}
	configuration := shared.SourceE2eTestCloudUpdate{
		MaxMessages:       maxMessages,
		MessageIntervalMs: messageIntervalMs,
		MockCatalog:       mockCatalog,
		Seed:              seed,
		Type:              type3,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceE2eTestCloudPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceE2eTestCloudResourceModel) ToDeleteSDKType() *shared.SourceE2eTestCloudCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceE2eTestCloudResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}