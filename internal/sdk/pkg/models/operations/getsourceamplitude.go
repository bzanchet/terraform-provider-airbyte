// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceAmplitudeRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceAmplitudeResponse struct {
	ContentType string
	// Get a Source by the id in the path.
	SourceResponse *shared.SourceResponse
	StatusCode     int
	RawResponse    *http.Response
}
