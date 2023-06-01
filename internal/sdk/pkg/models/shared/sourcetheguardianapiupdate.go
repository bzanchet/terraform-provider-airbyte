// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceTheGuardianAPIUpdate struct {
	// Your API Key. See <a href="https://open-platform.theguardian.com/access/">here</a>. The key is case sensitive.
	APIKey string `json:"api_key"`
	// (Optional) Use this to set the maximum date (YYYY-MM-DD) of the results. Results newer than the end_date will not be shown. Default is set to the current date (today) for incremental syncs.
	EndDate *string `json:"end_date,omitempty"`
	// (Optional) The query (q) parameter filters the results to only those that include that search term. The q parameter supports AND, OR and NOT operators.
	Query *string `json:"query,omitempty"`
	// (Optional) Use this to filter the results by a particular section. See <a href="https://content.guardianapis.com/sections?api-key=test">here</a> for a list of all sections, and <a href="https://open-platform.theguardian.com/documentation/section">here</a> for the sections endpoint documentation.
	Section *string `json:"section,omitempty"`
	// Use this to set the minimum date (YYYY-MM-DD) of the results. Results older than the start_date will not be shown.
	StartDate string `json:"start_date"`
	// (Optional) A tag is a piece of data that is used by The Guardian to categorise content. Use this parameter to filter results by showing only the ones matching the entered tag. See <a href="https://content.guardianapis.com/tags?api-key=test">here</a> for a list of all tags, and <a href="https://open-platform.theguardian.com/documentation/tag">here</a> for the tags endpoint documentation.
	Tag *string `json:"tag,omitempty"`
}
