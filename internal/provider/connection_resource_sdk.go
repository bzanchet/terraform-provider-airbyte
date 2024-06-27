// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ConnectionResourceModel) ToSharedConnectionCreateRequest() *shared.ConnectionCreateRequest {
	var configurations *shared.StreamConfigurations
	if r.Configurations != nil {
		var streams []shared.StreamConfiguration = []shared.StreamConfiguration{}
		for _, streamsItem := range r.Configurations.Streams {
			var cursorField []string = []string{}
			for _, cursorFieldItem := range streamsItem.CursorField {
				cursorField = append(cursorField, cursorFieldItem.ValueString())
			}
			name := streamsItem.Name.ValueString()
			var primaryKey [][]string = [][]string{}
			for _, primaryKeyItem := range streamsItem.PrimaryKey {
				var primaryKeyTmp []string = []string{}
				for _, item := range primaryKeyItem {
					primaryKeyTmp = append(primaryKeyTmp, item.ValueString())
				}
				primaryKey = append(primaryKey, primaryKeyTmp)
			}
			var selectedFields []shared.SelectedFieldInfo = []shared.SelectedFieldInfo{}
			for _, selectedFieldsItem := range streamsItem.SelectedFields {
				var fieldPath []string = []string{}
				for _, fieldPathItem := range selectedFieldsItem.FieldPath {
					fieldPath = append(fieldPath, fieldPathItem.ValueString())
				}
				selectedFields = append(selectedFields, shared.SelectedFieldInfo{
					FieldPath: fieldPath,
				})
			}
			syncMode := new(shared.ConnectionSyncModeEnum)
			if !streamsItem.SyncMode.IsUnknown() && !streamsItem.SyncMode.IsNull() {
				*syncMode = shared.ConnectionSyncModeEnum(streamsItem.SyncMode.ValueString())
			} else {
				syncMode = nil
			}
			streams = append(streams, shared.StreamConfiguration{
				CursorField:    cursorField,
				Name:           name,
				PrimaryKey:     primaryKey,
				SelectedFields: selectedFields,
				SyncMode:       syncMode,
			})
		}
		configurations = &shared.StreamConfigurations{
			Streams: streams,
		}
	}
	dataResidency := new(shared.GeographyEnum)
	if !r.DataResidency.IsUnknown() && !r.DataResidency.IsNull() {
		*dataResidency = shared.GeographyEnum(r.DataResidency.ValueString())
	} else {
		dataResidency = nil
	}
	destinationID := r.DestinationID.ValueString()
	name1 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name1 = r.Name.ValueString()
	} else {
		name1 = nil
	}
	namespaceDefinition := new(shared.NamespaceDefinitionEnum)
	if !r.NamespaceDefinition.IsUnknown() && !r.NamespaceDefinition.IsNull() {
		*namespaceDefinition = shared.NamespaceDefinitionEnum(r.NamespaceDefinition.ValueString())
	} else {
		namespaceDefinition = nil
	}
	namespaceFormat := new(string)
	if !r.NamespaceFormat.IsUnknown() && !r.NamespaceFormat.IsNull() {
		*namespaceFormat = r.NamespaceFormat.ValueString()
	} else {
		namespaceFormat = nil
	}
	nonBreakingSchemaUpdatesBehavior := new(shared.NonBreakingSchemaUpdatesBehaviorEnum)
	if !r.NonBreakingSchemaUpdatesBehavior.IsUnknown() && !r.NonBreakingSchemaUpdatesBehavior.IsNull() {
		*nonBreakingSchemaUpdatesBehavior = shared.NonBreakingSchemaUpdatesBehaviorEnum(r.NonBreakingSchemaUpdatesBehavior.ValueString())
	} else {
		nonBreakingSchemaUpdatesBehavior = nil
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = nil
	}
	var schedule *shared.AirbyteAPIConnectionSchedule
	if r.Schedule != nil {
		cronExpression := new(string)
		if !r.Schedule.CronExpression.IsUnknown() && !r.Schedule.CronExpression.IsNull() {
			*cronExpression = r.Schedule.CronExpression.ValueString()
		} else {
			cronExpression = nil
		}
		scheduleType := shared.ScheduleTypeEnum(r.Schedule.ScheduleType.ValueString())
		schedule = &shared.AirbyteAPIConnectionSchedule{
			CronExpression: cronExpression,
			ScheduleType:   scheduleType,
		}
	}
	sourceID := r.SourceID.ValueString()
	status := new(shared.ConnectionStatusEnum)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.ConnectionStatusEnum(r.Status.ValueString())
	} else {
		status = nil
	}
	out := shared.ConnectionCreateRequest{
		Configurations:                   configurations,
		DataResidency:                    dataResidency,
		DestinationID:                    destinationID,
		Name:                             name1,
		NamespaceDefinition:              namespaceDefinition,
		NamespaceFormat:                  namespaceFormat,
		NonBreakingSchemaUpdatesBehavior: nonBreakingSchemaUpdatesBehavior,
		Prefix:                           prefix,
		Schedule:                         schedule,
		SourceID:                         sourceID,
		Status:                           status,
	}
	return &out
}

func (r *ConnectionResourceModel) RefreshFromSharedConnectionResponse(resp *shared.ConnectionResponse) {
	if resp != nil {
		if r.Configurations == nil {
			r.Configurations = &tfTypes.StreamConfigurations{}
		}
		r.Configurations.Streams = []tfTypes.StreamConfiguration{}
		if len(r.Configurations.Streams) > len(resp.Configurations.Streams) {
			r.Configurations.Streams = r.Configurations.Streams[:len(resp.Configurations.Streams)]
		}
		for streamsCount, streamsItem := range resp.Configurations.Streams {
			var streams1 tfTypes.StreamConfiguration
			streams1.CursorField = []types.String{}
			for _, v := range streamsItem.CursorField {
				streams1.CursorField = append(streams1.CursorField, types.StringValue(v))
			}
			streams1.Name = types.StringValue(streamsItem.Name)
			streams1.PrimaryKey = nil
			for _, primaryKeyItem := range streamsItem.PrimaryKey {
				var primaryKey1 []types.String
				primaryKey1 = []types.String{}
				for _, v := range primaryKeyItem {
					primaryKey1 = append(primaryKey1, types.StringValue(v))
				}
				streams1.PrimaryKey = append(streams1.PrimaryKey, primaryKey1)
			}
			streams1.SelectedFields = []tfTypes.SelectedFieldInfo{}
			for selectedFieldsCount, selectedFieldsItem := range streamsItem.SelectedFields {
				var selectedFields1 tfTypes.SelectedFieldInfo
				selectedFields1.FieldPath = []types.String{}
				for _, v := range selectedFieldsItem.FieldPath {
					selectedFields1.FieldPath = append(selectedFields1.FieldPath, types.StringValue(v))
				}
				if selectedFieldsCount+1 > len(streams1.SelectedFields) {
					streams1.SelectedFields = append(streams1.SelectedFields, selectedFields1)
				} else {
					streams1.SelectedFields[selectedFieldsCount].FieldPath = selectedFields1.FieldPath
				}
			}
			if streamsItem.SyncMode != nil {
				streams1.SyncMode = types.StringValue(string(*streamsItem.SyncMode))
			} else {
				streams1.SyncMode = types.StringNull()
			}
			if streamsCount+1 > len(r.Configurations.Streams) {
				r.Configurations.Streams = append(r.Configurations.Streams, streams1)
			} else {
				r.Configurations.Streams[streamsCount].CursorField = streams1.CursorField
				r.Configurations.Streams[streamsCount].Name = streams1.Name
				r.Configurations.Streams[streamsCount].PrimaryKey = streams1.PrimaryKey
				r.Configurations.Streams[streamsCount].SelectedFields = streams1.SelectedFields
				r.Configurations.Streams[streamsCount].SyncMode = streams1.SyncMode
			}
		}
		r.ConnectionID = types.StringValue(resp.ConnectionID)
		if resp.DataResidency != nil {
			r.DataResidency = types.StringValue(string(*resp.DataResidency))
		} else {
			r.DataResidency = types.StringNull()
		}
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.Name = types.StringValue(resp.Name)
		if resp.NamespaceDefinition != nil {
			r.NamespaceDefinition = types.StringValue(string(*resp.NamespaceDefinition))
		} else {
			r.NamespaceDefinition = types.StringNull()
		}
		r.NamespaceFormat = types.StringPointerValue(resp.NamespaceFormat)
		if resp.NonBreakingSchemaUpdatesBehavior != nil {
			r.NonBreakingSchemaUpdatesBehavior = types.StringValue(string(*resp.NonBreakingSchemaUpdatesBehavior))
		} else {
			r.NonBreakingSchemaUpdatesBehavior = types.StringNull()
		}
		r.Prefix = types.StringPointerValue(resp.Prefix)
		if r.Schedule == nil {
			r.Schedule = &tfTypes.AirbyteAPIConnectionSchedule{}
		}
		r.Schedule.BasicTiming = types.StringPointerValue(resp.Schedule.BasicTiming)
		r.Schedule.CronExpression = types.StringPointerValue(resp.Schedule.CronExpression)
		r.Schedule.ScheduleType = types.StringValue(string(resp.Schedule.ScheduleType))
		r.SourceID = types.StringValue(resp.SourceID)
		r.Status = types.StringValue(string(resp.Status))
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *ConnectionResourceModel) ToSharedConnectionPatchRequest() *shared.ConnectionPatchRequest {
	var configurations *shared.StreamConfigurations
	if r.Configurations != nil {
		var streams []shared.StreamConfiguration = []shared.StreamConfiguration{}
		for _, streamsItem := range r.Configurations.Streams {
			var cursorField []string = []string{}
			for _, cursorFieldItem := range streamsItem.CursorField {
				cursorField = append(cursorField, cursorFieldItem.ValueString())
			}
			name := streamsItem.Name.ValueString()
			var primaryKey [][]string = [][]string{}
			for _, primaryKeyItem := range streamsItem.PrimaryKey {
				var primaryKeyTmp []string = []string{}
				for _, item := range primaryKeyItem {
					primaryKeyTmp = append(primaryKeyTmp, item.ValueString())
				}
				primaryKey = append(primaryKey, primaryKeyTmp)
			}
			var selectedFields []shared.SelectedFieldInfo = []shared.SelectedFieldInfo{}
			for _, selectedFieldsItem := range streamsItem.SelectedFields {
				var fieldPath []string = []string{}
				for _, fieldPathItem := range selectedFieldsItem.FieldPath {
					fieldPath = append(fieldPath, fieldPathItem.ValueString())
				}
				selectedFields = append(selectedFields, shared.SelectedFieldInfo{
					FieldPath: fieldPath,
				})
			}
			syncMode := new(shared.ConnectionSyncModeEnum)
			if !streamsItem.SyncMode.IsUnknown() && !streamsItem.SyncMode.IsNull() {
				*syncMode = shared.ConnectionSyncModeEnum(streamsItem.SyncMode.ValueString())
			} else {
				syncMode = nil
			}
			streams = append(streams, shared.StreamConfiguration{
				CursorField:    cursorField,
				Name:           name,
				PrimaryKey:     primaryKey,
				SelectedFields: selectedFields,
				SyncMode:       syncMode,
			})
		}
		configurations = &shared.StreamConfigurations{
			Streams: streams,
		}
	}
	dataResidency := new(shared.GeographyEnumNoDefault)
	if !r.DataResidency.IsUnknown() && !r.DataResidency.IsNull() {
		*dataResidency = shared.GeographyEnumNoDefault(r.DataResidency.ValueString())
	} else {
		dataResidency = nil
	}
	name1 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name1 = r.Name.ValueString()
	} else {
		name1 = nil
	}
	namespaceDefinition := new(shared.NamespaceDefinitionEnumNoDefault)
	if !r.NamespaceDefinition.IsUnknown() && !r.NamespaceDefinition.IsNull() {
		*namespaceDefinition = shared.NamespaceDefinitionEnumNoDefault(r.NamespaceDefinition.ValueString())
	} else {
		namespaceDefinition = nil
	}
	namespaceFormat := new(string)
	if !r.NamespaceFormat.IsUnknown() && !r.NamespaceFormat.IsNull() {
		*namespaceFormat = r.NamespaceFormat.ValueString()
	} else {
		namespaceFormat = nil
	}
	nonBreakingSchemaUpdatesBehavior := new(shared.NonBreakingSchemaUpdatesBehaviorEnumNoDefault)
	if !r.NonBreakingSchemaUpdatesBehavior.IsUnknown() && !r.NonBreakingSchemaUpdatesBehavior.IsNull() {
		*nonBreakingSchemaUpdatesBehavior = shared.NonBreakingSchemaUpdatesBehaviorEnumNoDefault(r.NonBreakingSchemaUpdatesBehavior.ValueString())
	} else {
		nonBreakingSchemaUpdatesBehavior = nil
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = nil
	}
	var schedule *shared.AirbyteAPIConnectionSchedule
	if r.Schedule != nil {
		cronExpression := new(string)
		if !r.Schedule.CronExpression.IsUnknown() && !r.Schedule.CronExpression.IsNull() {
			*cronExpression = r.Schedule.CronExpression.ValueString()
		} else {
			cronExpression = nil
		}
		scheduleType := shared.ScheduleTypeEnum(r.Schedule.ScheduleType.ValueString())
		schedule = &shared.AirbyteAPIConnectionSchedule{
			CronExpression: cronExpression,
			ScheduleType:   scheduleType,
		}
	}
	status := new(shared.ConnectionStatusEnum)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.ConnectionStatusEnum(r.Status.ValueString())
	} else {
		status = nil
	}
	out := shared.ConnectionPatchRequest{
		Configurations:                   configurations,
		DataResidency:                    dataResidency,
		Name:                             name1,
		NamespaceDefinition:              namespaceDefinition,
		NamespaceFormat:                  namespaceFormat,
		NonBreakingSchemaUpdatesBehavior: nonBreakingSchemaUpdatesBehavior,
		Prefix:                           prefix,
		Schedule:                         schedule,
		Status:                           status,
	}
	return &out
}
