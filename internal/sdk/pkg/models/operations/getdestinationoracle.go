// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationOracleRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type GetDestinationOracleResponse struct {
	ContentType string
	// Get a Destination by the id in the path.
	DestinationResponse *shared.DestinationResponse
	StatusCode          int
	RawResponse         *http.Response
}
