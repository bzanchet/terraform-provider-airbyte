// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceMailjetSmsUpdate struct {
	// Retrieve SMS messages created before the specified timestamp. Required format - Unix timestamp.
	EndDate *int64 `json:"end_date,omitempty"`
	// Retrieve SMS messages created after the specified timestamp. Required format - Unix timestamp.
	StartDate *int64 `json:"start_date,omitempty"`
	// Your access token. See <a href="https://dev.mailjet.com/sms/reference/overview/authentication">here</a>.
	Token string `json:"token"`
}
