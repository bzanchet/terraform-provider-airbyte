// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials string

const (
	SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentialsAccessToken SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials = "access_token"
)

func (e SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials) ToPointer() *SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials {
	return &e
}

func (e *SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials: %v", v)
	}
}

type SourceZendeskChatUpdateAuthorizationMethodAccessToken struct {
	// The Access Token to make authenticated requests.
	AccessToken string                                                           `json:"access_token"`
	Credentials SourceZendeskChatUpdateAuthorizationMethodAccessTokenCredentials `json:"credentials"`
}

type SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials string

const (
	SourceZendeskChatUpdateAuthorizationMethodOAuth20CredentialsOauth20 SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials = "oauth2.0"
)

func (e SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials) ToPointer() *SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials {
	return &e
}

func (e *SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials: %v", v)
	}
}

type SourceZendeskChatUpdateAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string `json:"access_token,omitempty"`
	// The Client ID of your OAuth application
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string                                                      `json:"client_secret,omitempty"`
	Credentials  SourceZendeskChatUpdateAuthorizationMethodOAuth20Credentials `json:"credentials"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

type SourceZendeskChatUpdateAuthorizationMethodType string

const (
	SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAuthorizationMethodOAuth20     SourceZendeskChatUpdateAuthorizationMethodType = "source-zendesk-chat-update_Authorization Method_OAuth2.0"
	SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAuthorizationMethodAccessToken SourceZendeskChatUpdateAuthorizationMethodType = "source-zendesk-chat-update_Authorization Method_Access Token"
)

type SourceZendeskChatUpdateAuthorizationMethod struct {
	SourceZendeskChatUpdateAuthorizationMethodOAuth20     *SourceZendeskChatUpdateAuthorizationMethodOAuth20
	SourceZendeskChatUpdateAuthorizationMethodAccessToken *SourceZendeskChatUpdateAuthorizationMethodAccessToken

	Type SourceZendeskChatUpdateAuthorizationMethodType
}

func CreateSourceZendeskChatUpdateAuthorizationMethodSourceZendeskChatUpdateAuthorizationMethodOAuth20(sourceZendeskChatUpdateAuthorizationMethodOAuth20 SourceZendeskChatUpdateAuthorizationMethodOAuth20) SourceZendeskChatUpdateAuthorizationMethod {
	typ := SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAuthorizationMethodOAuth20

	return SourceZendeskChatUpdateAuthorizationMethod{
		SourceZendeskChatUpdateAuthorizationMethodOAuth20: &sourceZendeskChatUpdateAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceZendeskChatUpdateAuthorizationMethodSourceZendeskChatUpdateAuthorizationMethodAccessToken(sourceZendeskChatUpdateAuthorizationMethodAccessToken SourceZendeskChatUpdateAuthorizationMethodAccessToken) SourceZendeskChatUpdateAuthorizationMethod {
	typ := SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAuthorizationMethodAccessToken

	return SourceZendeskChatUpdateAuthorizationMethod{
		SourceZendeskChatUpdateAuthorizationMethodAccessToken: &sourceZendeskChatUpdateAuthorizationMethodAccessToken,
		Type: typ,
	}
}

func (u *SourceZendeskChatUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceZendeskChatUpdateAuthorizationMethodOAuth20 := new(SourceZendeskChatUpdateAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskChatUpdateAuthorizationMethodOAuth20); err == nil {
		u.SourceZendeskChatUpdateAuthorizationMethodOAuth20 = sourceZendeskChatUpdateAuthorizationMethodOAuth20
		u.Type = SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAuthorizationMethodOAuth20
		return nil
	}

	sourceZendeskChatUpdateAuthorizationMethodAccessToken := new(SourceZendeskChatUpdateAuthorizationMethodAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskChatUpdateAuthorizationMethodAccessToken); err == nil {
		u.SourceZendeskChatUpdateAuthorizationMethodAccessToken = sourceZendeskChatUpdateAuthorizationMethodAccessToken
		u.Type = SourceZendeskChatUpdateAuthorizationMethodTypeSourceZendeskChatUpdateAuthorizationMethodAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskChatUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskChatUpdateAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceZendeskChatUpdateAuthorizationMethodOAuth20)
	}

	if u.SourceZendeskChatUpdateAuthorizationMethodAccessToken != nil {
		return json.Marshal(u.SourceZendeskChatUpdateAuthorizationMethodAccessToken)
	}

	return nil, nil
}

type SourceZendeskChatUpdate struct {
	Credentials *SourceZendeskChatUpdateAuthorizationMethod `json:"credentials,omitempty"`
	// The date from which you'd like to replicate data for Zendesk Chat API, in the format YYYY-MM-DDT00:00:00Z.
	StartDate time.Time `json:"start_date"`
	// Required if you access Zendesk Chat from a Zendesk Support subdomain.
	Subdomain *string `json:"subdomain,omitempty"`
}
