// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceGoogleAnalyticsDataAPIRequest struct {
	SourceGoogleAnalyticsDataAPIPutRequest *shared.SourceGoogleAnalyticsDataAPIPutRequest `request:"mediaType=application/json"`
	SourceID                               string                                         `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceGoogleAnalyticsDataAPIResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
