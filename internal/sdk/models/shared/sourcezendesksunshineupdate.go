// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceZendeskSunshineUpdateSchemasAuthMethod string

const (
	SourceZendeskSunshineUpdateSchemasAuthMethodAPIToken SourceZendeskSunshineUpdateSchemasAuthMethod = "api_token"
)

func (e SourceZendeskSunshineUpdateSchemasAuthMethod) ToPointer() *SourceZendeskSunshineUpdateSchemasAuthMethod {
	return &e
}
func (e *SourceZendeskSunshineUpdateSchemasAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceZendeskSunshineUpdateSchemasAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSunshineUpdateSchemasAuthMethod: %v", v)
	}
}

type SourceZendeskSunshineUpdateAPIToken struct {
	// API Token. See the <a href="https://docs.airbyte.com/integrations/sources/zendesk_sunshine">docs</a> for information on how to generate this key.
	APIToken   string                                        `json:"api_token"`
	authMethod *SourceZendeskSunshineUpdateSchemasAuthMethod `const:"api_token" json:"auth_method"`
	// The user email for your Zendesk account
	Email string `json:"email"`
}

func (s SourceZendeskSunshineUpdateAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSunshineUpdateAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSunshineUpdateAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceZendeskSunshineUpdateAPIToken) GetAuthMethod() *SourceZendeskSunshineUpdateSchemasAuthMethod {
	return SourceZendeskSunshineUpdateSchemasAuthMethodAPIToken.ToPointer()
}

func (o *SourceZendeskSunshineUpdateAPIToken) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

type SourceZendeskSunshineUpdateAuthMethod string

const (
	SourceZendeskSunshineUpdateAuthMethodOauth20 SourceZendeskSunshineUpdateAuthMethod = "oauth2.0"
)

func (e SourceZendeskSunshineUpdateAuthMethod) ToPointer() *SourceZendeskSunshineUpdateAuthMethod {
	return &e
}
func (e *SourceZendeskSunshineUpdateAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskSunshineUpdateAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSunshineUpdateAuthMethod: %v", v)
	}
}

type SourceZendeskSunshineUpdateOAuth20 struct {
	// Long-term access Token for making authenticated requests.
	AccessToken string                                 `json:"access_token"`
	authMethod  *SourceZendeskSunshineUpdateAuthMethod `const:"oauth2.0" json:"auth_method"`
	// The Client ID of your OAuth application.
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
}

func (s SourceZendeskSunshineUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSunshineUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSunshineUpdateOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceZendeskSunshineUpdateOAuth20) GetAuthMethod() *SourceZendeskSunshineUpdateAuthMethod {
	return SourceZendeskSunshineUpdateAuthMethodOauth20.ToPointer()
}

func (o *SourceZendeskSunshineUpdateOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceZendeskSunshineUpdateOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

type SourceZendeskSunshineUpdateAuthorizationMethodType string

const (
	SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateOAuth20  SourceZendeskSunshineUpdateAuthorizationMethodType = "source-zendesk-sunshine-update_OAuth2.0"
	SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAPIToken SourceZendeskSunshineUpdateAuthorizationMethodType = "source-zendesk-sunshine-update_API Token"
)

type SourceZendeskSunshineUpdateAuthorizationMethod struct {
	SourceZendeskSunshineUpdateOAuth20  *SourceZendeskSunshineUpdateOAuth20
	SourceZendeskSunshineUpdateAPIToken *SourceZendeskSunshineUpdateAPIToken

	Type SourceZendeskSunshineUpdateAuthorizationMethodType
}

func CreateSourceZendeskSunshineUpdateAuthorizationMethodSourceZendeskSunshineUpdateOAuth20(sourceZendeskSunshineUpdateOAuth20 SourceZendeskSunshineUpdateOAuth20) SourceZendeskSunshineUpdateAuthorizationMethod {
	typ := SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateOAuth20

	return SourceZendeskSunshineUpdateAuthorizationMethod{
		SourceZendeskSunshineUpdateOAuth20: &sourceZendeskSunshineUpdateOAuth20,
		Type:                               typ,
	}
}

func CreateSourceZendeskSunshineUpdateAuthorizationMethodSourceZendeskSunshineUpdateAPIToken(sourceZendeskSunshineUpdateAPIToken SourceZendeskSunshineUpdateAPIToken) SourceZendeskSunshineUpdateAuthorizationMethod {
	typ := SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAPIToken

	return SourceZendeskSunshineUpdateAuthorizationMethod{
		SourceZendeskSunshineUpdateAPIToken: &sourceZendeskSunshineUpdateAPIToken,
		Type:                                typ,
	}
}

func (u *SourceZendeskSunshineUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {

	var sourceZendeskSunshineUpdateAPIToken SourceZendeskSunshineUpdateAPIToken = SourceZendeskSunshineUpdateAPIToken{}
	if err := utils.UnmarshalJSON(data, &sourceZendeskSunshineUpdateAPIToken, "", true, true); err == nil {
		u.SourceZendeskSunshineUpdateAPIToken = &sourceZendeskSunshineUpdateAPIToken
		u.Type = SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAPIToken
		return nil
	}

	var sourceZendeskSunshineUpdateOAuth20 SourceZendeskSunshineUpdateOAuth20 = SourceZendeskSunshineUpdateOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceZendeskSunshineUpdateOAuth20, "", true, true); err == nil {
		u.SourceZendeskSunshineUpdateOAuth20 = &sourceZendeskSunshineUpdateOAuth20
		u.Type = SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceZendeskSunshineUpdateAuthorizationMethod", string(data))
}

func (u SourceZendeskSunshineUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskSunshineUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceZendeskSunshineUpdateOAuth20, "", true)
	}

	if u.SourceZendeskSunshineUpdateAPIToken != nil {
		return utils.MarshalJSON(u.SourceZendeskSunshineUpdateAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type SourceZendeskSunshineUpdateAuthorizationMethod: all fields are null")
}

type SourceZendeskSunshineUpdate struct {
	Credentials *SourceZendeskSunshineUpdateAuthorizationMethod `json:"credentials,omitempty"`
	// The date from which you'd like to replicate data for Zendesk Sunshine API, in the format YYYY-MM-DDT00:00:00Z.
	StartDate time.Time `json:"start_date"`
	// The subdomain for your Zendesk Account.
	Subdomain string `json:"subdomain"`
}

func (s SourceZendeskSunshineUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSunshineUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSunshineUpdate) GetCredentials() *SourceZendeskSunshineUpdateAuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceZendeskSunshineUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceZendeskSunshineUpdate) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
