// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourcePokeapiRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourcePokeapiResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
