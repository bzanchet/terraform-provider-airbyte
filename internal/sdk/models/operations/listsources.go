// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type ListSourcesRequest struct {
	// Include deleted sources in the returned results.
	IncludeDeleted *bool `default:"false" queryParam:"style=form,explode=true,name=includeDeleted"`
	// Set the limit on the number of sources returned. The default is 20.
	Limit *int `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Set the offset to start at when returning sources. The default is 0
	Offset *int `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// The UUIDs of the workspaces you wish to list sources for. Empty list will retrieve all allowed workspaces.
	WorkspaceIds []string `queryParam:"style=form,explode=true,name=workspaceIds"`
}

func (l ListSourcesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSourcesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListSourcesRequest) GetIncludeDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeleted
}

func (o *ListSourcesRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListSourcesRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListSourcesRequest) GetWorkspaceIds() []string {
	if o == nil {
		return nil
	}
	return o.WorkspaceIds
}

type ListSourcesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	SourcesResponse *shared.SourcesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListSourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSourcesResponse) GetSourcesResponse() *shared.SourcesResponse {
	if o == nil {
		return nil
	}
	return o.SourcesResponse
}

func (o *ListSourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}