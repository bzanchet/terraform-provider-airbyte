// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourcePostmarkappRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourcePostmarkappResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
