// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
)

type SourceSnapchatMarketingUpdate struct {
	// The Client ID of your Snapchat developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Snapchat developer application.
	ClientSecret string `json:"client_secret"`
	// Date in the format 2017-01-25. Any data after this date will not be replicated.
	EndDate *types.Date `json:"end_date,omitempty"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// Date in the format 2022-01-01. Any data before this date will not be replicated.
	StartDate *types.Date `json:"start_date,omitempty"`
}
