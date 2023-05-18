// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceKlaviyoKlaviyoEnum string

const (
	SourceKlaviyoKlaviyoEnumKlaviyo SourceKlaviyoKlaviyoEnum = "klaviyo"
)

func (e SourceKlaviyoKlaviyoEnum) ToPointer() *SourceKlaviyoKlaviyoEnum {
	return &e
}

func (e *SourceKlaviyoKlaviyoEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "klaviyo":
		*e = SourceKlaviyoKlaviyoEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceKlaviyoKlaviyoEnum: %v", v)
	}
}

type SourceKlaviyo struct {
	// Klaviyo API Key. See our <a href="https://docs.airbyte.com/integrations/sources/klaviyo">docs</a> if you need help finding this key.
	APIKey     string                   `json:"api_key"`
	SourceType SourceKlaviyoKlaviyoEnum `json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
