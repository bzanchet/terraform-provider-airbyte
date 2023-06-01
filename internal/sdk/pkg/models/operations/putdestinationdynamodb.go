// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationDynamodbRequest struct {
	DestinationDynamodbPutRequest *shared.DestinationDynamodbPutRequest `request:"mediaType=application/json"`
	DestinationID                 string                                `pathParam:"style=simple,explode=false,name=destinationId"`
}

type PutDestinationDynamodbResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
