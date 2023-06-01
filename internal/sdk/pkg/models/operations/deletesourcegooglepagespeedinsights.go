// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceGooglePagespeedInsightsRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceGooglePagespeedInsightsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
