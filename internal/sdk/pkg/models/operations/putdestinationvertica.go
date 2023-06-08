// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationVerticaRequest struct {
	DestinationVerticaPutRequest *shared.DestinationVerticaPutRequest `request:"mediaType=application/json"`
	DestinationID                string                               `pathParam:"style=simple,explode=false,name=destinationId"`
}

type PutDestinationVerticaResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}