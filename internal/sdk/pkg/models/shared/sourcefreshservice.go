// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceFreshserviceFreshservice string

const (
	SourceFreshserviceFreshserviceFreshservice SourceFreshserviceFreshservice = "freshservice"
)

func (e SourceFreshserviceFreshservice) ToPointer() *SourceFreshserviceFreshservice {
	return &e
}

func (e *SourceFreshserviceFreshservice) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "freshservice":
		*e = SourceFreshserviceFreshservice(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFreshserviceFreshservice: %v", v)
	}
}

type SourceFreshservice struct {
	// Freshservice API Key. See <a href="https://api.freshservice.com/#authentication">here</a>. The key is case sensitive.
	APIKey string `json:"api_key"`
	// The name of your Freshservice domain
	DomainName string                         `json:"domain_name"`
	SourceType SourceFreshserviceFreshservice `json:"sourceType"`
	// UTC date and time in the format 2020-10-01T00:00:00Z. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}