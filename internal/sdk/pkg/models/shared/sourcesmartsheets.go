// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum string

const (
	SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnumAccessToken SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum = "access_token"
)

func (e SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum) ToPointer() *SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum {
	return &e
}

func (e *SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum: %v", v)
	}
}

type SourceSmartsheetsAuthorizationMethodAPIAccessToken struct {
	// The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account > Apps & Integrations > API Access. See the <a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide">setup guide</a> for information on how to obtain this token.
	AccessToken string                                                          `json:"access_token"`
	AuthType    *SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthTypeEnum `json:"auth_type,omitempty"`
}

type SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum string

const (
	SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnumOauth20 SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum = "oauth2.0"
)

func (e SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum) ToPointer() *SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum {
	return &e
}

func (e *SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum: %v", v)
	}
}

type SourceSmartsheetsAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                                                   `json:"access_token"`
	AuthType    *SourceSmartsheetsAuthorizationMethodOAuth20AuthTypeEnum `json:"auth_type,omitempty"`
	// The API ID of the SmartSheets developer application.
	ClientID string `json:"client_id"`
	// The API Secret the SmartSheets developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access_token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

type SourceSmartsheetsAuthorizationMethodType string

const (
	SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAuthorizationMethodOAuth20        SourceSmartsheetsAuthorizationMethodType = "source-smartsheets_Authorization Method_OAuth2.0"
	SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAuthorizationMethodAPIAccessToken SourceSmartsheetsAuthorizationMethodType = "source-smartsheets_Authorization Method_API Access Token"
)

type SourceSmartsheetsAuthorizationMethod struct {
	SourceSmartsheetsAuthorizationMethodOAuth20        *SourceSmartsheetsAuthorizationMethodOAuth20
	SourceSmartsheetsAuthorizationMethodAPIAccessToken *SourceSmartsheetsAuthorizationMethodAPIAccessToken

	Type SourceSmartsheetsAuthorizationMethodType
}

func CreateSourceSmartsheetsAuthorizationMethodSourceSmartsheetsAuthorizationMethodOAuth20(sourceSmartsheetsAuthorizationMethodOAuth20 SourceSmartsheetsAuthorizationMethodOAuth20) SourceSmartsheetsAuthorizationMethod {
	typ := SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAuthorizationMethodOAuth20

	return SourceSmartsheetsAuthorizationMethod{
		SourceSmartsheetsAuthorizationMethodOAuth20: &sourceSmartsheetsAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceSmartsheetsAuthorizationMethodSourceSmartsheetsAuthorizationMethodAPIAccessToken(sourceSmartsheetsAuthorizationMethodAPIAccessToken SourceSmartsheetsAuthorizationMethodAPIAccessToken) SourceSmartsheetsAuthorizationMethod {
	typ := SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAuthorizationMethodAPIAccessToken

	return SourceSmartsheetsAuthorizationMethod{
		SourceSmartsheetsAuthorizationMethodAPIAccessToken: &sourceSmartsheetsAuthorizationMethodAPIAccessToken,
		Type: typ,
	}
}

func (u *SourceSmartsheetsAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSmartsheetsAuthorizationMethodOAuth20 := new(SourceSmartsheetsAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSmartsheetsAuthorizationMethodOAuth20); err == nil {
		u.SourceSmartsheetsAuthorizationMethodOAuth20 = sourceSmartsheetsAuthorizationMethodOAuth20
		u.Type = SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAuthorizationMethodOAuth20
		return nil
	}

	sourceSmartsheetsAuthorizationMethodAPIAccessToken := new(SourceSmartsheetsAuthorizationMethodAPIAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSmartsheetsAuthorizationMethodAPIAccessToken); err == nil {
		u.SourceSmartsheetsAuthorizationMethodAPIAccessToken = sourceSmartsheetsAuthorizationMethodAPIAccessToken
		u.Type = SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAuthorizationMethodAPIAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSmartsheetsAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceSmartsheetsAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceSmartsheetsAuthorizationMethodOAuth20)
	}

	if u.SourceSmartsheetsAuthorizationMethodAPIAccessToken != nil {
		return json.Marshal(u.SourceSmartsheetsAuthorizationMethodAPIAccessToken)
	}

	return nil, nil
}

type SourceSmartsheetsSmartsheetsEnum string

const (
	SourceSmartsheetsSmartsheetsEnumSmartsheets SourceSmartsheetsSmartsheetsEnum = "smartsheets"
)

func (e SourceSmartsheetsSmartsheetsEnum) ToPointer() *SourceSmartsheetsSmartsheetsEnum {
	return &e
}

func (e *SourceSmartsheetsSmartsheetsEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smartsheets":
		*e = SourceSmartsheetsSmartsheetsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsSmartsheetsEnum: %v", v)
	}
}

type SourceSmartsheets struct {
	Credentials SourceSmartsheetsAuthorizationMethod `json:"credentials"`
	SourceType  SourceSmartsheetsSmartsheetsEnum     `json:"sourceType"`
	// The spreadsheet ID. Find it by opening the spreadsheet then navigating to File > Properties
	SpreadsheetID string `json:"spreadsheet_id"`
	// Only rows modified after this date/time will be replicated. This should be an ISO 8601 string, for instance: `2000-01-01T13:00:00`
	StartDatetime *time.Time `json:"start_datetime,omitempty"`
}
