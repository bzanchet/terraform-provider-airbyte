// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePendoCreateRequest struct {
	Configuration SourcePendo `json:"configuration"`
	Name          string      `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}
