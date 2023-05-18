// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationGcsResponse struct {
	ContentType string
	// Successful operation
	DestinationResponse *shared.DestinationResponse
	StatusCode          int
	RawResponse         *http.Response
}
