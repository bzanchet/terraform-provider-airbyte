// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceDatascopeResponse struct {
	ContentType string
	// Successful operation
	SourceResponse *shared.SourceResponse
	StatusCode     int
	RawResponse    *http.Response
}
