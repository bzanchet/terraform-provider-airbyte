// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceInsightlyUpdate struct {
	// The date from which you'd like to replicate data for Insightly in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated. Note that it will be used only for incremental streams.
	StartDate string `json:"start_date"`
	// Your Insightly API token.
	Token string `json:"token"`
}
