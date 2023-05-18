// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceOktaResourceModel) ToCreateSDKType() *shared.SourceOktaCreateRequest {
	var credentials *shared.SourceOktaAuthorizationMethod
	var sourceOktaAuthorizationMethodOAuth20 *shared.SourceOktaAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20 != nil {
		authType := shared.SourceOktaAuthorizationMethodOAuth20AuthTypeEnum(r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.AuthType.ValueString())
		clientID := r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.RefreshToken.ValueString()
		sourceOktaAuthorizationMethodOAuth20 = &shared.SourceOktaAuthorizationMethodOAuth20{
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceOktaAuthorizationMethodOAuth20 != nil {
		credentials = &shared.SourceOktaAuthorizationMethod{
			SourceOktaAuthorizationMethodOAuth20: sourceOktaAuthorizationMethodOAuth20,
		}
	}
	var sourceOktaAuthorizationMethodAPIToken *shared.SourceOktaAuthorizationMethodAPIToken
	if r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken != nil {
		apiToken := r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken.APIToken.ValueString()
		authType1 := shared.SourceOktaAuthorizationMethodAPITokenAuthTypeEnum(r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken.AuthType.ValueString())
		sourceOktaAuthorizationMethodAPIToken = &shared.SourceOktaAuthorizationMethodAPIToken{
			APIToken: apiToken,
			AuthType: authType1,
		}
	}
	if sourceOktaAuthorizationMethodAPIToken != nil {
		credentials = &shared.SourceOktaAuthorizationMethod{
			SourceOktaAuthorizationMethodAPIToken: sourceOktaAuthorizationMethodAPIToken,
		}
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	sourceType := shared.SourceOktaOktaEnum(r.Configuration.SourceType.ValueString())
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	configuration := shared.SourceOkta{
		Credentials: credentials,
		Domain:      domain,
		SourceType:  sourceType,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOktaCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOktaResourceModel) ToDeleteSDKType() *shared.SourceOktaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
