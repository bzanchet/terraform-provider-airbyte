// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceIp2whoisRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceIp2whoisResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
