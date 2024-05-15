// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceS3Request struct {
	SourceS3PutRequest *shared.SourceS3PutRequest `request:"mediaType=application/json"`
	SourceID           string                     `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceS3Request) GetSourceS3PutRequest() *shared.SourceS3PutRequest {
	if o == nil {
		return nil
	}
	return o.SourceS3PutRequest
}

func (o *PutSourceS3Request) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceS3Response struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceS3Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceS3Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceS3Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}