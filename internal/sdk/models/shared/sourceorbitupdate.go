// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceOrbitUpdate struct {
	// Authorizes you to work with Orbit workspaces associated with the token.
	APIToken string `json:"api_token"`
	// Date in the format 2022-06-26. Only load members whose last activities are after this date.
	StartDate *string `json:"start_date,omitempty"`
	// The unique name of the workspace that your API token is associated with.
	Workspace string `json:"workspace"`
}

func (o *SourceOrbitUpdate) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceOrbitUpdate) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceOrbitUpdate) GetWorkspace() string {
	if o == nil {
		return ""
	}
	return o.Workspace
}