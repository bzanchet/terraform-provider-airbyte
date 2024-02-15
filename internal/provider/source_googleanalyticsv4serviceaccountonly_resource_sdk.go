// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel) ToSharedSourceGoogleAnalyticsV4ServiceAccountOnlyCreateRequest() *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyCreateRequest {
	var credentials *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyCredentials
	if r.Configuration.Credentials != nil {
		var sourceGoogleAnalyticsV4ServiceAccountOnlyServiceAccountKeyAuthentication *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyServiceAccountKeyAuthentication
		if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
			credentialsJSON := r.Configuration.Credentials.ServiceAccountKeyAuthentication.CredentialsJSON.ValueString()
			sourceGoogleAnalyticsV4ServiceAccountOnlyServiceAccountKeyAuthentication = &shared.SourceGoogleAnalyticsV4ServiceAccountOnlyServiceAccountKeyAuthentication{
				CredentialsJSON: credentialsJSON,
			}
		}
		if sourceGoogleAnalyticsV4ServiceAccountOnlyServiceAccountKeyAuthentication != nil {
			credentials = &shared.SourceGoogleAnalyticsV4ServiceAccountOnlyCredentials{
				SourceGoogleAnalyticsV4ServiceAccountOnlyServiceAccountKeyAuthentication: sourceGoogleAnalyticsV4ServiceAccountOnlyServiceAccountKeyAuthentication,
			}
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	viewID := r.Configuration.ViewID.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceGoogleAnalyticsV4ServiceAccountOnly{
		Credentials:   credentials,
		CustomReports: customReports,
		EndDate:       endDate,
		StartDate:     startDate,
		ViewID:        viewID,
		WindowInDays:  windowInDays,
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
	out := shared.SourceGoogleAnalyticsV4ServiceAccountOnlyCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel) ToSharedSourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest() *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest {
	var credentials *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyUpdateCredentials
	if r.Configuration.Credentials != nil {
		var sourceGoogleAnalyticsV4ServiceAccountOnlyUpdateServiceAccountKeyAuthentication *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyUpdateServiceAccountKeyAuthentication
		if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
			credentialsJSON := r.Configuration.Credentials.ServiceAccountKeyAuthentication.CredentialsJSON.ValueString()
			sourceGoogleAnalyticsV4ServiceAccountOnlyUpdateServiceAccountKeyAuthentication = &shared.SourceGoogleAnalyticsV4ServiceAccountOnlyUpdateServiceAccountKeyAuthentication{
				CredentialsJSON: credentialsJSON,
			}
		}
		if sourceGoogleAnalyticsV4ServiceAccountOnlyUpdateServiceAccountKeyAuthentication != nil {
			credentials = &shared.SourceGoogleAnalyticsV4ServiceAccountOnlyUpdateCredentials{
				SourceGoogleAnalyticsV4ServiceAccountOnlyUpdateServiceAccountKeyAuthentication: sourceGoogleAnalyticsV4ServiceAccountOnlyUpdateServiceAccountKeyAuthentication,
			}
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	viewID := r.Configuration.ViewID.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceGoogleAnalyticsV4ServiceAccountOnlyUpdate{
		Credentials:   credentials,
		CustomReports: customReports,
		EndDate:       endDate,
		StartDate:     startDate,
		ViewID:        viewID,
		WindowInDays:  windowInDays,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
