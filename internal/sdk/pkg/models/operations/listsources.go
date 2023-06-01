// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListSourcesRequest struct {
	// Include deleted sources in the returned results.
	IncludeDeleted *bool `queryParam:"style=form,explode=true,name=includeDeleted"`
	// Set the limit on the number of sources returned. The default is 20.
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	// Set the offset to start at when returning sources. The default is 0
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// The UUIDs of the workspaces you wish to list sources for. Empty list will retrieve all allowed workspaces.
	WorkspaceIds []string `queryParam:"style=form,explode=true,name=workspaceIds"`
}

type ListSourcesResponse struct {
	ContentType string
	// Successful operation
	SourcesResponse *shared.SourcesResponse
	StatusCode      int
	RawResponse     *http.Response
}
