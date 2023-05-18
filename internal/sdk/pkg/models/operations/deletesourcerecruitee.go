// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceRecruiteeRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceRecruiteeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
