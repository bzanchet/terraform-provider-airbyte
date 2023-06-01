// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationSftpJSONRequest struct {
	DestinationSftpJSONPutRequest *shared.DestinationSftpJSONPutRequest `request:"mediaType=application/json"`
	DestinationID                 string                                `pathParam:"style=simple,explode=false,name=destinationId"`
}

type PutDestinationSftpJSONResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
