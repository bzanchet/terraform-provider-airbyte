// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum string

const (
	SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnumService SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum = "Service"
)

func (e SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum) ToPointer() *SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum {
	return &e
}

func (e *SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service":
		*e = SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum: %v", v)
	}
}

// SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication - Credentials for connecting to the Google Sheets API
type SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication struct {
	AuthType SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthTypeEnum `json:"auth_type"`
	// Enter your Google Cloud <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">service account key</a> in JSON format
	ServiceAccountInfo string `json:"service_account_info"`
}

type SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum string

const (
	SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnumClient SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum = "Client"
)

func (e SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum) ToPointer() *SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum {
	return &e
}

func (e *SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum: %v", v)
	}
}

// SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth - Credentials for connecting to the Google Sheets API
type SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth struct {
	AuthType SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthTypeEnum `json:"auth_type"`
	// Enter your Google application's Client ID
	ClientID string `json:"client_id"`
	// Enter your Google application's Client Secret
	ClientSecret string `json:"client_secret"`
	// Enter your Google application's refresh token
	RefreshToken string `json:"refresh_token"`
}

type SourceGoogleSheetsAuthenticationType string

const (
	SourceGoogleSheetsAuthenticationTypeSourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth      SourceGoogleSheetsAuthenticationType = "source-google-sheets_Authentication_Authenticate via Google (OAuth)"
	SourceGoogleSheetsAuthenticationTypeSourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication SourceGoogleSheetsAuthenticationType = "source-google-sheets_Authentication_Service Account Key Authentication"
)

type SourceGoogleSheetsAuthentication struct {
	SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth      *SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth
	SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication *SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication

	Type SourceGoogleSheetsAuthenticationType
}

func CreateSourceGoogleSheetsAuthenticationSourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth(sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth) SourceGoogleSheetsAuthentication {
	typ := SourceGoogleSheetsAuthenticationTypeSourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth

	return SourceGoogleSheetsAuthentication{
		SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth: &sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth,
		Type: typ,
	}
}

func CreateSourceGoogleSheetsAuthenticationSourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication(sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication) SourceGoogleSheetsAuthentication {
	typ := SourceGoogleSheetsAuthenticationTypeSourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication

	return SourceGoogleSheetsAuthentication{
		SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication: &sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication,
		Type: typ,
	}
}

func (u *SourceGoogleSheetsAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth := new(SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth); err == nil {
		u.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth = sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth
		u.Type = SourceGoogleSheetsAuthenticationTypeSourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth
		return nil
	}

	sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication := new(SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication); err == nil {
		u.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication = sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication
		u.Type = SourceGoogleSheetsAuthenticationTypeSourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGoogleSheetsAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth != nil {
		return json.Marshal(u.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth)
	}

	if u.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication != nil {
		return json.Marshal(u.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication)
	}

	return nil, nil
}

type SourceGoogleSheetsGoogleSheetsEnum string

const (
	SourceGoogleSheetsGoogleSheetsEnumGoogleSheets SourceGoogleSheetsGoogleSheetsEnum = "google-sheets"
)

func (e SourceGoogleSheetsGoogleSheetsEnum) ToPointer() *SourceGoogleSheetsGoogleSheetsEnum {
	return &e
}

func (e *SourceGoogleSheetsGoogleSheetsEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-sheets":
		*e = SourceGoogleSheetsGoogleSheetsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSheetsGoogleSheetsEnum: %v", v)
	}
}

type SourceGoogleSheets struct {
	// Credentials for connecting to the Google Sheets API
	Credentials SourceGoogleSheetsAuthentication `json:"credentials"`
	// Number of rows fetched when making a Google Sheet API call. Defaults to 200.
	RowBatchSize *int64                             `json:"row_batch_size,omitempty"`
	SourceType   SourceGoogleSheetsGoogleSheetsEnum `json:"sourceType"`
	// Enter the link to the Google spreadsheet you want to sync
	SpreadsheetID string `json:"spreadsheet_id"`
}
