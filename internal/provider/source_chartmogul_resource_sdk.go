// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceChartmogulResourceModel) ToCreateSDKType() *shared.SourceChartmogulCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	interval := shared.SourceChartmogulInterval(r.Configuration.Interval.ValueString())
	sourceType := shared.SourceChartmogulChartmogul(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceChartmogul{
		APIKey:     apiKey,
		Interval:   interval,
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
	out := shared.SourceChartmogulCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceChartmogulResourceModel) ToUpdateSDKType() *shared.SourceChartmogulPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	interval := shared.SourceChartmogulUpdateInterval(r.Configuration.Interval.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceChartmogulUpdate{
		APIKey:    apiKey,
		Interval:  interval,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceChartmogulPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceChartmogulResourceModel) ToDeleteSDKType() *shared.SourceChartmogulCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceChartmogulResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
