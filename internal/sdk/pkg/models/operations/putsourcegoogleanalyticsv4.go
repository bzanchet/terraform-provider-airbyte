// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceGoogleAnalyticsV4Request struct {
	SourceGoogleAnalyticsV4PutRequest *shared.SourceGoogleAnalyticsV4PutRequest `request:"mediaType=application/json"`
	SourceID                          string                                    `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceGoogleAnalyticsV4Response struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}