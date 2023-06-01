// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceAirtableRequest struct {
	SourceAirtablePutRequest *shared.SourceAirtablePutRequest `request:"mediaType=application/json"`
	SourceID                 string                           `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceAirtableResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
