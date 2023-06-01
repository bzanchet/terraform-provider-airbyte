// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type SourceCloseComUpdate struct {
	// Close.com API key (usually starts with 'api_'; find yours <a href="https://app.close.com/settings/api/">here</a>).
	APIKey string `json:"api_key"`
	// The start date to sync data. Leave blank for full sync. Format: YYYY-MM-DD.
	StartDate *time.Time `json:"start_date,omitempty"`
}
