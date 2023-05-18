// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceOpenweatherResourceModel) ToCreateSDKType() *shared.SourceOpenweatherCreateRequest {
	appid := r.Configuration.Appid.ValueString()
	lang := new(shared.SourceOpenweatherLanguageEnum)
	if !r.Configuration.Lang.IsUnknown() && !r.Configuration.Lang.IsNull() {
		*lang = shared.SourceOpenweatherLanguageEnum(r.Configuration.Lang.ValueString())
	} else {
		lang = nil
	}
	lat := r.Configuration.Lat.ValueString()
	lon := r.Configuration.Lon.ValueString()
	sourceType := shared.SourceOpenweatherOpenweatherEnum(r.Configuration.SourceType.ValueString())
	units := new(shared.SourceOpenweatherUnitsEnum)
	if !r.Configuration.Units.IsUnknown() && !r.Configuration.Units.IsNull() {
		*units = shared.SourceOpenweatherUnitsEnum(r.Configuration.Units.ValueString())
	} else {
		units = nil
	}
	configuration := shared.SourceOpenweather{
		Appid:      appid,
		Lang:       lang,
		Lat:        lat,
		Lon:        lon,
		SourceType: sourceType,
		Units:      units,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOpenweatherCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOpenweatherResourceModel) ToDeleteSDKType() *shared.SourceOpenweatherCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
