// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePatchRequest struct {
	// The values required to configure the source.
	Configuration interface{} `json:"configuration,omitempty"`
	Name          *string     `json:"name,omitempty"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID *string `json:"workspaceId,omitempty"`
}