// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
)

type SourceBrazeUpdate struct {
	// Braze REST API key
	APIKey string `json:"api_key"`
	// Rows after this date will be synced
	StartDate types.Date `json:"start_date"`
	// Braze REST API endpoint
	URL string `json:"url"`
}