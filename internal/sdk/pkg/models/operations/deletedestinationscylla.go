// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDestinationScyllaRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type DeleteDestinationScyllaResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
