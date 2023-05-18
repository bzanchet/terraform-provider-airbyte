// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationRocksetRocksetEnum string

const (
	DestinationRocksetRocksetEnumRockset DestinationRocksetRocksetEnum = "rockset"
)

func (e DestinationRocksetRocksetEnum) ToPointer() *DestinationRocksetRocksetEnum {
	return &e
}

func (e *DestinationRocksetRocksetEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rockset":
		*e = DestinationRocksetRocksetEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRocksetRocksetEnum: %v", v)
	}
}

type DestinationRockset struct {
	// Rockset api key
	APIKey string `json:"api_key"`
	// Rockset api URL
	APIServer       *string                       `json:"api_server,omitempty"`
	DestinationType DestinationRocksetRocksetEnum `json:"destinationType"`
	// The Rockset workspace in which collections will be created + written to.
	Workspace string `json:"workspace"`
}
