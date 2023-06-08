// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DestinationGoogleSheetsUpdateAuthenticationViaGoogleOAuth - Google API Credentials for connecting to Google Sheets and Google Drive APIs
type DestinationGoogleSheetsUpdateAuthenticationViaGoogleOAuth struct {
	// The Client ID of your Google Sheets developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Google Sheets developer application.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining new access token.
	RefreshToken string `json:"refresh_token"`
}

type DestinationGoogleSheetsUpdate struct {
	// Google API Credentials for connecting to Google Sheets and Google Drive APIs
	Credentials DestinationGoogleSheetsUpdateAuthenticationViaGoogleOAuth `json:"credentials"`
	// The link to your spreadsheet. See <a href='https://docs.airbyte.com/integrations/destinations/google-sheets#sheetlink'>this guide</a> for more details.
	SpreadsheetID string `json:"spreadsheet_id"`
}