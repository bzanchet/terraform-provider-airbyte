// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceSlackSchemasOptionTitle string

const (
	SourceSlackSchemasOptionTitleAPITokenCredentials SourceSlackSchemasOptionTitle = "API Token Credentials"
)

func (e SourceSlackSchemasOptionTitle) ToPointer() *SourceSlackSchemasOptionTitle {
	return &e
}

func (e *SourceSlackSchemasOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "API Token Credentials":
		*e = SourceSlackSchemasOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSlackSchemasOptionTitle: %v", v)
	}
}

type SourceSlackAPIToken struct {
	// A Slack bot token. See the <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> for instructions on how to generate it.
	APIToken    string                        `json:"api_token"`
	optionTitle SourceSlackSchemasOptionTitle `const:"API Token Credentials" json:"option_title"`
}

func (s SourceSlackAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSlackAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSlackAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceSlackAPIToken) GetOptionTitle() SourceSlackSchemasOptionTitle {
	return SourceSlackSchemasOptionTitleAPITokenCredentials
}

type SourceSlackOptionTitle string

const (
	SourceSlackOptionTitleDefaultOAuth20Authorization SourceSlackOptionTitle = "Default OAuth2.0 authorization"
)

func (e SourceSlackOptionTitle) ToPointer() *SourceSlackOptionTitle {
	return &e
}

func (e *SourceSlackOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Default OAuth2.0 authorization":
		*e = SourceSlackOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSlackOptionTitle: %v", v)
	}
}

type SourceSlackSignInViaSlackOAuth struct {
	// Slack access_token. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help generating the token.
	AccessToken string `json:"access_token"`
	// Slack client_id. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this id.
	ClientID string `json:"client_id"`
	// Slack client_secret. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this secret.
	ClientSecret string                 `json:"client_secret"`
	optionTitle  SourceSlackOptionTitle `const:"Default OAuth2.0 authorization" json:"option_title"`
}

func (s SourceSlackSignInViaSlackOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSlackSignInViaSlackOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSlackSignInViaSlackOAuth) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceSlackSignInViaSlackOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceSlackSignInViaSlackOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceSlackSignInViaSlackOAuth) GetOptionTitle() SourceSlackOptionTitle {
	return SourceSlackOptionTitleDefaultOAuth20Authorization
}

type SourceSlackAuthenticationMechanismType string

const (
	SourceSlackAuthenticationMechanismTypeSourceSlackSignInViaSlackOAuth SourceSlackAuthenticationMechanismType = "source-slack_Sign in via Slack (OAuth)"
	SourceSlackAuthenticationMechanismTypeSourceSlackAPIToken            SourceSlackAuthenticationMechanismType = "source-slack_API Token"
)

// SourceSlackAuthenticationMechanism - Choose how to authenticate into Slack
type SourceSlackAuthenticationMechanism struct {
	SourceSlackSignInViaSlackOAuth *SourceSlackSignInViaSlackOAuth
	SourceSlackAPIToken            *SourceSlackAPIToken

	Type SourceSlackAuthenticationMechanismType
}

func CreateSourceSlackAuthenticationMechanismSourceSlackSignInViaSlackOAuth(sourceSlackSignInViaSlackOAuth SourceSlackSignInViaSlackOAuth) SourceSlackAuthenticationMechanism {
	typ := SourceSlackAuthenticationMechanismTypeSourceSlackSignInViaSlackOAuth

	return SourceSlackAuthenticationMechanism{
		SourceSlackSignInViaSlackOAuth: &sourceSlackSignInViaSlackOAuth,
		Type:                           typ,
	}
}

func CreateSourceSlackAuthenticationMechanismSourceSlackAPIToken(sourceSlackAPIToken SourceSlackAPIToken) SourceSlackAuthenticationMechanism {
	typ := SourceSlackAuthenticationMechanismTypeSourceSlackAPIToken

	return SourceSlackAuthenticationMechanism{
		SourceSlackAPIToken: &sourceSlackAPIToken,
		Type:                typ,
	}
}

func (u *SourceSlackAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	sourceSlackAPIToken := new(SourceSlackAPIToken)
	if err := utils.UnmarshalJSON(data, &sourceSlackAPIToken, "", true, true); err == nil {
		u.SourceSlackAPIToken = sourceSlackAPIToken
		u.Type = SourceSlackAuthenticationMechanismTypeSourceSlackAPIToken
		return nil
	}

	sourceSlackSignInViaSlackOAuth := new(SourceSlackSignInViaSlackOAuth)
	if err := utils.UnmarshalJSON(data, &sourceSlackSignInViaSlackOAuth, "", true, true); err == nil {
		u.SourceSlackSignInViaSlackOAuth = sourceSlackSignInViaSlackOAuth
		u.Type = SourceSlackAuthenticationMechanismTypeSourceSlackSignInViaSlackOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSlackAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceSlackSignInViaSlackOAuth != nil {
		return utils.MarshalJSON(u.SourceSlackSignInViaSlackOAuth, "", true)
	}

	if u.SourceSlackAPIToken != nil {
		return utils.MarshalJSON(u.SourceSlackAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Slack string

const (
	SlackSlack Slack = "slack"
)

func (e Slack) ToPointer() *Slack {
	return &e
}

func (e *Slack) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "slack":
		*e = Slack(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Slack: %v", v)
	}
}

type SourceSlack struct {
	// A channel name list (without leading '#' char) which limit the channels from which you'd like to sync. Empty list means no filter.
	ChannelFilter []string `json:"channel_filter,omitempty"`
	// Choose how to authenticate into Slack
	Credentials *SourceSlackAuthenticationMechanism `json:"credentials,omitempty"`
	// Whether to join all channels or to sync data only from channels the bot is already in.  If false, you'll need to manually add the bot to all the channels from which you'd like to sync messages.
	JoinChannels *bool `default:"true" json:"join_channels"`
	// How far into the past to look for messages in threads, default is 0 days
	LookbackWindow *int64 `default:"0" json:"lookback_window"`
	sourceType     Slack  `const:"slack" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceSlack) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSlack) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSlack) GetChannelFilter() []string {
	if o == nil {
		return nil
	}
	return o.ChannelFilter
}

func (o *SourceSlack) GetCredentials() *SourceSlackAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceSlack) GetJoinChannels() *bool {
	if o == nil {
		return nil
	}
	return o.JoinChannels
}

func (o *SourceSlack) GetLookbackWindow() *int64 {
	if o == nil {
		return nil
	}
	return o.LookbackWindow
}

func (o *SourceSlack) GetSourceType() Slack {
	return SlackSlack
}

func (o *SourceSlack) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
