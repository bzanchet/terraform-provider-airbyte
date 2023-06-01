// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceOutreachUpdate struct {
	// The Client ID of your Outreach developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Outreach developer application.
	ClientSecret string `json:"client_secret"`
	// A Redirect URI is the location where the authorization server sends the user once the app has been successfully authorized and granted an authorization code or access token.
	RedirectURI string `json:"redirect_uri"`
	// The token for obtaining the new access token.
	RefreshToken string `json:"refresh_token"`
	// The date from which you'd like to replicate data for Outreach API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate string `json:"start_date"`
}
