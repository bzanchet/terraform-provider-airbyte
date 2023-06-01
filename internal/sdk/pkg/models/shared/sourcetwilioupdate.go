// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type SourceTwilioUpdate struct {
	// Twilio account SID
	AccountSid string `json:"account_sid"`
	// Twilio Auth Token.
	AuthToken string `json:"auth_token"`
	// How far into the past to look for records. (in minutes)
	LookbackWindow *int64 `json:"lookback_window,omitempty"`
	// UTC date and time in the format 2020-10-01T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
