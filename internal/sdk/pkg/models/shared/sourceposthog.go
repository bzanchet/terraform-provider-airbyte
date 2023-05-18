// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourcePosthogPosthogEnum string

const (
	SourcePosthogPosthogEnumPosthog SourcePosthogPosthogEnum = "posthog"
)

func (e SourcePosthogPosthogEnum) ToPointer() *SourcePosthogPosthogEnum {
	return &e
}

func (e *SourcePosthogPosthogEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "posthog":
		*e = SourcePosthogPosthogEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePosthogPosthogEnum: %v", v)
	}
}

type SourcePosthog struct {
	// API Key. See the <a href="https://docs.airbyte.com/integrations/sources/posthog">docs</a> for information on how to generate this key.
	APIKey string `json:"api_key"`
	// Base PostHog url. Defaults to PostHog Cloud (https://app.posthog.com).
	BaseURL    *string                  `json:"base_url,omitempty"`
	SourceType SourcePosthogPosthogEnum `json:"sourceType"`
	// The date from which you'd like to replicate the data. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
