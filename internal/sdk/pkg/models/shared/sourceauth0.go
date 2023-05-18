// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum string

const (
	SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnumOauth2AccessToken SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum = "oauth2_access_token"
)

func (e SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum) ToPointer() *SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum {
	return &e
}

func (e *SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2_access_token":
		*e = SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum: %v", v)
	}
}

type SourceAuth0AuthenticationMethodOAuth2AccessToken struct {
	// Also called <a href="https://auth0.com/docs/secure/tokens/access-tokens/get-management-api-access-tokens-for-testing">API Access Token </a> The access token used to call the Auth0 Management API Token. It's a JWT that contains specific grant permissions knowns as scopes.
	AccessToken string                                                                   `json:"access_token"`
	AuthType    SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodEnum `json:"auth_type"`
}

type SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum string

const (
	SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnumOauth2ConfidentialApplication SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum = "oauth2_confidential_application"
)

func (e SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum) ToPointer() *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum {
	return &e
}

func (e *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2_confidential_application":
		*e = SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum: %v", v)
	}
}

type SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication struct {
	// The audience for the token, which is your API. You can find this in the Identifier field on your  <a href="https://manage.auth0.com/#/apis">API's settings tab</a>
	Audience string                                                                               `json:"audience"`
	AuthType SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodEnum `json:"auth_type"`
	// Your application's Client ID. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
	ClientID string `json:"client_id"`
	// Your application's Client Secret. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
	ClientSecret string `json:"client_secret"`
}

type SourceAuth0AuthenticationMethodType string

const (
	SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication SourceAuth0AuthenticationMethodType = "source-auth0_Authentication Method_OAuth2 Confidential Application"
	SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2AccessToken             SourceAuth0AuthenticationMethodType = "source-auth0_Authentication Method_OAuth2 Access Token"
)

type SourceAuth0AuthenticationMethod struct {
	SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication
	SourceAuth0AuthenticationMethodOAuth2AccessToken             *SourceAuth0AuthenticationMethodOAuth2AccessToken

	Type SourceAuth0AuthenticationMethodType
}

func CreateSourceAuth0AuthenticationMethodSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication(sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) SourceAuth0AuthenticationMethod {
	typ := SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication

	return SourceAuth0AuthenticationMethod{
		SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication: &sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication,
		Type: typ,
	}
}

func CreateSourceAuth0AuthenticationMethodSourceAuth0AuthenticationMethodOAuth2AccessToken(sourceAuth0AuthenticationMethodOAuth2AccessToken SourceAuth0AuthenticationMethodOAuth2AccessToken) SourceAuth0AuthenticationMethod {
	typ := SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2AccessToken

	return SourceAuth0AuthenticationMethod{
		SourceAuth0AuthenticationMethodOAuth2AccessToken: &sourceAuth0AuthenticationMethodOAuth2AccessToken,
		Type: typ,
	}
}

func (u *SourceAuth0AuthenticationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication := new(SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication); err == nil {
		u.SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication = sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication
		u.Type = SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication
		return nil
	}

	sourceAuth0AuthenticationMethodOAuth2AccessToken := new(SourceAuth0AuthenticationMethodOAuth2AccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceAuth0AuthenticationMethodOAuth2AccessToken); err == nil {
		u.SourceAuth0AuthenticationMethodOAuth2AccessToken = sourceAuth0AuthenticationMethodOAuth2AccessToken
		u.Type = SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2AccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAuth0AuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication != nil {
		return json.Marshal(u.SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication)
	}

	if u.SourceAuth0AuthenticationMethodOAuth2AccessToken != nil {
		return json.Marshal(u.SourceAuth0AuthenticationMethodOAuth2AccessToken)
	}

	return nil, nil
}

type SourceAuth0Auth0Enum string

const (
	SourceAuth0Auth0EnumAuth0 SourceAuth0Auth0Enum = "auth0"
)

func (e SourceAuth0Auth0Enum) ToPointer() *SourceAuth0Auth0Enum {
	return &e
}

func (e *SourceAuth0Auth0Enum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auth0":
		*e = SourceAuth0Auth0Enum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0Auth0Enum: %v", v)
	}
}

type SourceAuth0 struct {
	// The Authentication API is served over HTTPS. All URLs referenced in the documentation have the following base `https://YOUR_DOMAIN`
	BaseURL     string                          `json:"base_url"`
	Credentials SourceAuth0AuthenticationMethod `json:"credentials"`
	SourceType  SourceAuth0Auth0Enum            `json:"sourceType"`
}
