// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

type SourceGoogleAdsGoogleCredentials struct {
	// Access Token for making authenticated requests. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>
	AccessToken *string `json:"access_token,omitempty"`
	// The Client ID of your Google Ads developer application. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>
	ClientID string `json:"client_id"`
	// The Client Secret of your Google Ads developer application. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>
	ClientSecret string `json:"client_secret"`
	// Developer token granted by Google to use their APIs. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>
	DeveloperToken string `json:"developer_token"`
	// The token for obtaining a new access token. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>
	RefreshToken string `json:"refresh_token"`
}

type SourceGoogleAdsCustomQueries struct {
	// A custom defined GAQL query for building the report. Should not contain segments.date expression because it is used by incremental streams. See Google's <a href="https://developers.google.com/google-ads/api/fields/v11/overview_query_builder">query builder</a> for more information.
	Query string `json:"query"`
	// The table name in your destination database for choosen query.
	TableName string `json:"table_name"`
}

type SourceGoogleAdsGoogleAdsEnum string

const (
	SourceGoogleAdsGoogleAdsEnumGoogleAds SourceGoogleAdsGoogleAdsEnum = "google-ads"
)

func (e SourceGoogleAdsGoogleAdsEnum) ToPointer() *SourceGoogleAdsGoogleAdsEnum {
	return &e
}

func (e *SourceGoogleAdsGoogleAdsEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-ads":
		*e = SourceGoogleAdsGoogleAdsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleAdsGoogleAdsEnum: %v", v)
	}
}

type SourceGoogleAds struct {
	// A conversion window is the period of time after an ad interaction (such as an ad click or video view) during which a conversion, such as a purchase, is recorded in Google Ads. For more information, see Google's <a href="https://support.google.com/google-ads/answer/3123169?hl=en">documentation</a>.
	ConversionWindowDays *int64                           `json:"conversion_window_days,omitempty"`
	Credentials          SourceGoogleAdsGoogleCredentials `json:"credentials"`
	CustomQueries        []SourceGoogleAdsCustomQueries   `json:"custom_queries,omitempty"`
	// Comma separated list of (client) customer IDs. Each customer ID must be specified as a 10-digit number without dashes. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>. Metrics streams like AdGroupAdReport cannot be requested for a manager account.
	CustomerID string `json:"customer_id"`
	// UTC date and time in the format 2017-01-25. Any data after this date will not be replicated.
	EndDate *types.Date `json:"end_date,omitempty"`
	// If your access to the customer account is through a manager account, this field is required and must be set to the customer ID of the manager account (10-digit number without dashes). More information about this field you can see <a href="https://developers.google.com/google-ads/api/docs/concepts/call-structure#cid">here</a>
	LoginCustomerID *string                      `json:"login_customer_id,omitempty"`
	SourceType      SourceGoogleAdsGoogleAdsEnum `json:"sourceType"`
	// UTC date and time in the format 2017-01-25. Any data before this date will not be replicated.
	StartDate types.Date `json:"start_date"`
}
