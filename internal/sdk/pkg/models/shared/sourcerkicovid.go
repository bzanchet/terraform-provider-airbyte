// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceRkiCovidRkiCovidEnum string

const (
	SourceRkiCovidRkiCovidEnumRkiCovid SourceRkiCovidRkiCovidEnum = "rki-covid"
)

func (e SourceRkiCovidRkiCovidEnum) ToPointer() *SourceRkiCovidRkiCovidEnum {
	return &e
}

func (e *SourceRkiCovidRkiCovidEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rki-covid":
		*e = SourceRkiCovidRkiCovidEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRkiCovidRkiCovidEnum: %v", v)
	}
}

type SourceRkiCovid struct {
	SourceType SourceRkiCovidRkiCovidEnum `json:"sourceType"`
	// UTC date in the format 2017-01-25. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}
