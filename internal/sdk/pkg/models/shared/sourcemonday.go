// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceMondayAuthorizationMethodAPITokenAuthTypeEnum string

const (
	SourceMondayAuthorizationMethodAPITokenAuthTypeEnumAPIToken SourceMondayAuthorizationMethodAPITokenAuthTypeEnum = "api_token"
)

func (e SourceMondayAuthorizationMethodAPITokenAuthTypeEnum) ToPointer() *SourceMondayAuthorizationMethodAPITokenAuthTypeEnum {
	return &e
}

func (e *SourceMondayAuthorizationMethodAPITokenAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceMondayAuthorizationMethodAPITokenAuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMondayAuthorizationMethodAPITokenAuthTypeEnum: %v", v)
	}
}

type SourceMondayAuthorizationMethodAPIToken struct {
	// API Token for making authenticated requests.
	APIToken string                                              `json:"api_token"`
	AuthType SourceMondayAuthorizationMethodAPITokenAuthTypeEnum `json:"auth_type"`
}

type SourceMondayAuthorizationMethodOAuth20AuthTypeEnum string

const (
	SourceMondayAuthorizationMethodOAuth20AuthTypeEnumOauth20 SourceMondayAuthorizationMethodOAuth20AuthTypeEnum = "oauth2.0"
)

func (e SourceMondayAuthorizationMethodOAuth20AuthTypeEnum) ToPointer() *SourceMondayAuthorizationMethodOAuth20AuthTypeEnum {
	return &e
}

func (e *SourceMondayAuthorizationMethodOAuth20AuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceMondayAuthorizationMethodOAuth20AuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMondayAuthorizationMethodOAuth20AuthTypeEnum: %v", v)
	}
}

type SourceMondayAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                                             `json:"access_token"`
	AuthType    SourceMondayAuthorizationMethodOAuth20AuthTypeEnum `json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
	// Slug/subdomain of the account, or the first part of the URL that comes before .monday.com
	Subdomain *string `json:"subdomain,omitempty"`
}

type SourceMondayAuthorizationMethodType string

const (
	SourceMondayAuthorizationMethodTypeSourceMondayAuthorizationMethodOAuth20  SourceMondayAuthorizationMethodType = "source-monday_Authorization Method_OAuth2.0"
	SourceMondayAuthorizationMethodTypeSourceMondayAuthorizationMethodAPIToken SourceMondayAuthorizationMethodType = "source-monday_Authorization Method_API Token"
)

type SourceMondayAuthorizationMethod struct {
	SourceMondayAuthorizationMethodOAuth20  *SourceMondayAuthorizationMethodOAuth20
	SourceMondayAuthorizationMethodAPIToken *SourceMondayAuthorizationMethodAPIToken

	Type SourceMondayAuthorizationMethodType
}

func CreateSourceMondayAuthorizationMethodSourceMondayAuthorizationMethodOAuth20(sourceMondayAuthorizationMethodOAuth20 SourceMondayAuthorizationMethodOAuth20) SourceMondayAuthorizationMethod {
	typ := SourceMondayAuthorizationMethodTypeSourceMondayAuthorizationMethodOAuth20

	return SourceMondayAuthorizationMethod{
		SourceMondayAuthorizationMethodOAuth20: &sourceMondayAuthorizationMethodOAuth20,
		Type:                                   typ,
	}
}

func CreateSourceMondayAuthorizationMethodSourceMondayAuthorizationMethodAPIToken(sourceMondayAuthorizationMethodAPIToken SourceMondayAuthorizationMethodAPIToken) SourceMondayAuthorizationMethod {
	typ := SourceMondayAuthorizationMethodTypeSourceMondayAuthorizationMethodAPIToken

	return SourceMondayAuthorizationMethod{
		SourceMondayAuthorizationMethodAPIToken: &sourceMondayAuthorizationMethodAPIToken,
		Type:                                    typ,
	}
}

func (u *SourceMondayAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMondayAuthorizationMethodOAuth20 := new(SourceMondayAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMondayAuthorizationMethodOAuth20); err == nil {
		u.SourceMondayAuthorizationMethodOAuth20 = sourceMondayAuthorizationMethodOAuth20
		u.Type = SourceMondayAuthorizationMethodTypeSourceMondayAuthorizationMethodOAuth20
		return nil
	}

	sourceMondayAuthorizationMethodAPIToken := new(SourceMondayAuthorizationMethodAPIToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMondayAuthorizationMethodAPIToken); err == nil {
		u.SourceMondayAuthorizationMethodAPIToken = sourceMondayAuthorizationMethodAPIToken
		u.Type = SourceMondayAuthorizationMethodTypeSourceMondayAuthorizationMethodAPIToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMondayAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMondayAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceMondayAuthorizationMethodOAuth20)
	}

	if u.SourceMondayAuthorizationMethodAPIToken != nil {
		return json.Marshal(u.SourceMondayAuthorizationMethodAPIToken)
	}

	return nil, nil
}

type SourceMondayMondayEnum string

const (
	SourceMondayMondayEnumMonday SourceMondayMondayEnum = "monday"
)

func (e SourceMondayMondayEnum) ToPointer() *SourceMondayMondayEnum {
	return &e
}

func (e *SourceMondayMondayEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "monday":
		*e = SourceMondayMondayEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMondayMondayEnum: %v", v)
	}
}

type SourceMonday struct {
	Credentials *SourceMondayAuthorizationMethod `json:"credentials,omitempty"`
	SourceType  SourceMondayMondayEnum           `json:"sourceType"`
}
