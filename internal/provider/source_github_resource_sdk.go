// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceGithubResourceModel) ToCreateSDKType() *shared.SourceGithubCreateRequest {
	branch := new(string)
	if !r.Configuration.Branch.IsUnknown() && !r.Configuration.Branch.IsNull() {
		*branch = r.Configuration.Branch.ValueString()
	} else {
		branch = nil
	}
	var credentials *shared.SourceGithubAuthentication
	var sourceGithubAuthenticationOAuth *shared.SourceGithubAuthenticationOAuth
	if r.Configuration.Credentials.SourceGithubAuthenticationOAuth != nil {
		accessToken := r.Configuration.Credentials.SourceGithubAuthenticationOAuth.AccessToken.ValueString()
		optionTitle := new(shared.SourceGithubAuthenticationOAuthOptionTitleEnum)
		if !r.Configuration.Credentials.SourceGithubAuthenticationOAuth.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceGithubAuthenticationOAuth.OptionTitle.IsNull() {
			*optionTitle = shared.SourceGithubAuthenticationOAuthOptionTitleEnum(r.Configuration.Credentials.SourceGithubAuthenticationOAuth.OptionTitle.ValueString())
		} else {
			optionTitle = nil
		}
		sourceGithubAuthenticationOAuth = &shared.SourceGithubAuthenticationOAuth{
			AccessToken: accessToken,
			OptionTitle: optionTitle,
		}
	}
	if sourceGithubAuthenticationOAuth != nil {
		credentials = &shared.SourceGithubAuthentication{
			SourceGithubAuthenticationOAuth: sourceGithubAuthenticationOAuth,
		}
	}
	var sourceGithubAuthenticationPersonalAccessToken *shared.SourceGithubAuthenticationPersonalAccessToken
	if r.Configuration.Credentials.SourceGithubAuthenticationPersonalAccessToken != nil {
		optionTitle1 := new(shared.SourceGithubAuthenticationPersonalAccessTokenOptionTitleEnum)
		if !r.Configuration.Credentials.SourceGithubAuthenticationPersonalAccessToken.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceGithubAuthenticationPersonalAccessToken.OptionTitle.IsNull() {
			*optionTitle1 = shared.SourceGithubAuthenticationPersonalAccessTokenOptionTitleEnum(r.Configuration.Credentials.SourceGithubAuthenticationPersonalAccessToken.OptionTitle.ValueString())
		} else {
			optionTitle1 = nil
		}
		personalAccessToken := r.Configuration.Credentials.SourceGithubAuthenticationPersonalAccessToken.PersonalAccessToken.ValueString()
		sourceGithubAuthenticationPersonalAccessToken = &shared.SourceGithubAuthenticationPersonalAccessToken{
			OptionTitle:         optionTitle1,
			PersonalAccessToken: personalAccessToken,
		}
	}
	if sourceGithubAuthenticationPersonalAccessToken != nil {
		credentials = &shared.SourceGithubAuthentication{
			SourceGithubAuthenticationPersonalAccessToken: sourceGithubAuthenticationPersonalAccessToken,
		}
	}
	pageSizeForLargeStreams := new(int64)
	if !r.Configuration.PageSizeForLargeStreams.IsUnknown() && !r.Configuration.PageSizeForLargeStreams.IsNull() {
		*pageSizeForLargeStreams = r.Configuration.PageSizeForLargeStreams.ValueInt64()
	} else {
		pageSizeForLargeStreams = nil
	}
	repository := r.Configuration.Repository.ValueString()
	requestsPerHour := new(int64)
	if !r.Configuration.RequestsPerHour.IsUnknown() && !r.Configuration.RequestsPerHour.IsNull() {
		*requestsPerHour = r.Configuration.RequestsPerHour.ValueInt64()
	} else {
		requestsPerHour = nil
	}
	sourceType := shared.SourceGithubGithubEnum(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceGithub{
		Branch:                  branch,
		Credentials:             credentials,
		PageSizeForLargeStreams: pageSizeForLargeStreams,
		Repository:              repository,
		RequestsPerHour:         requestsPerHour,
		SourceType:              sourceType,
		StartDate:               startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGithubCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGithubResourceModel) ToDeleteSDKType() *shared.SourceGithubCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
