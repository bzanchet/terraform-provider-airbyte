// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationTypesenseTypesenseEnum string

const (
	DestinationTypesenseTypesenseEnumTypesense DestinationTypesenseTypesenseEnum = "typesense"
)

func (e DestinationTypesenseTypesenseEnum) ToPointer() *DestinationTypesenseTypesenseEnum {
	return &e
}

func (e *DestinationTypesenseTypesenseEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "typesense":
		*e = DestinationTypesenseTypesenseEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationTypesenseTypesenseEnum: %v", v)
	}
}

type DestinationTypesense struct {
	// Typesense API Key
	APIKey string `json:"api_key"`
	// How many documents should be imported together. Default 1000
	BatchSize       *string                           `json:"batch_size,omitempty"`
	DestinationType DestinationTypesenseTypesenseEnum `json:"destinationType"`
	// Hostname of the Typesense instance without protocol.
	Host string `json:"host"`
	// Port of the Typesense instance. Ex: 8108, 80, 443. Default is 443
	Port *string `json:"port,omitempty"`
	// Protocol of the Typesense instance. Ex: http or https. Default is https
	Protocol *string `json:"protocol,omitempty"`
}
