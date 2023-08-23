// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceClockifyClockify string

const (
	SourceClockifyClockifyClockify SourceClockifyClockify = "clockify"
)

func (e SourceClockifyClockify) ToPointer() *SourceClockifyClockify {
	return &e
}

func (e *SourceClockifyClockify) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clockify":
		*e = SourceClockifyClockify(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceClockifyClockify: %v", v)
	}
}

type SourceClockify struct {
	// You can get your api access_key <a href="https://app.clockify.me/user/settings">here</a> This API is Case Sensitive.
	APIKey string `json:"api_key"`
	// The URL for the Clockify API. This should only need to be modified if connecting to an enterprise version of Clockify.
	APIURL     *string                `json:"api_url,omitempty"`
	SourceType SourceClockifyClockify `json:"sourceType"`
	// WorkSpace Id
	WorkspaceID string `json:"workspace_id"`
}