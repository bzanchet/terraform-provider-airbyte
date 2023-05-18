// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceInsightlyInsightlyEnum string

const (
	SourceInsightlyInsightlyEnumInsightly SourceInsightlyInsightlyEnum = "insightly"
)

func (e SourceInsightlyInsightlyEnum) ToPointer() *SourceInsightlyInsightlyEnum {
	return &e
}

func (e *SourceInsightlyInsightlyEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "insightly":
		*e = SourceInsightlyInsightlyEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceInsightlyInsightlyEnum: %v", v)
	}
}

type SourceInsightly struct {
	SourceType SourceInsightlyInsightlyEnum `json:"sourceType"`
	// The date from which you'd like to replicate data for Insightly in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated. Note that it will be used only for incremental streams.
	StartDate string `json:"start_date"`
	// Your Insightly API token.
	Token string `json:"token"`
}
