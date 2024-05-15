// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceCoinmarketcapRequest struct {
	SourceCoinmarketcapPutRequest *shared.SourceCoinmarketcapPutRequest `request:"mediaType=application/json"`
	SourceID                      string                                `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceCoinmarketcapRequest) GetSourceCoinmarketcapPutRequest() *shared.SourceCoinmarketcapPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceCoinmarketcapPutRequest
}

func (o *PutSourceCoinmarketcapRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceCoinmarketcapResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceCoinmarketcapResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceCoinmarketcapResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceCoinmarketcapResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}