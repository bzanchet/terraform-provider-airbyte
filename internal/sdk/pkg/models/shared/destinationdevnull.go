// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationDevNullDevNull string

const (
	DestinationDevNullDevNullDevNull DestinationDevNullDevNull = "dev-null"
)

func (e DestinationDevNullDevNull) ToPointer() *DestinationDevNullDevNull {
	return &e
}

func (e *DestinationDevNullDevNull) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dev-null":
		*e = DestinationDevNullDevNull(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullDevNull: %v", v)
	}
}

type DestinationDevNullTestDestinationSilentTestDestinationType string

const (
	DestinationDevNullTestDestinationSilentTestDestinationTypeSilent DestinationDevNullTestDestinationSilentTestDestinationType = "SILENT"
)

func (e DestinationDevNullTestDestinationSilentTestDestinationType) ToPointer() *DestinationDevNullTestDestinationSilentTestDestinationType {
	return &e
}

func (e *DestinationDevNullTestDestinationSilentTestDestinationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SILENT":
		*e = DestinationDevNullTestDestinationSilentTestDestinationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullTestDestinationSilentTestDestinationType: %v", v)
	}
}

// DestinationDevNullTestDestinationSilent - The type of destination to be used
type DestinationDevNullTestDestinationSilent struct {
	TestDestinationType DestinationDevNullTestDestinationSilentTestDestinationType `json:"test_destination_type"`
}

type DestinationDevNullTestDestinationType string

const (
	DestinationDevNullTestDestinationTypeDestinationDevNullTestDestinationSilent DestinationDevNullTestDestinationType = "destination-dev-null_Test Destination_Silent"
)

type DestinationDevNullTestDestination struct {
	DestinationDevNullTestDestinationSilent *DestinationDevNullTestDestinationSilent

	Type DestinationDevNullTestDestinationType
}

func CreateDestinationDevNullTestDestinationDestinationDevNullTestDestinationSilent(destinationDevNullTestDestinationSilent DestinationDevNullTestDestinationSilent) DestinationDevNullTestDestination {
	typ := DestinationDevNullTestDestinationTypeDestinationDevNullTestDestinationSilent

	return DestinationDevNullTestDestination{
		DestinationDevNullTestDestinationSilent: &destinationDevNullTestDestinationSilent,
		Type:                                    typ,
	}
}

func (u *DestinationDevNullTestDestination) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationDevNullTestDestinationSilent := new(DestinationDevNullTestDestinationSilent)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationDevNullTestDestinationSilent); err == nil {
		u.DestinationDevNullTestDestinationSilent = destinationDevNullTestDestinationSilent
		u.Type = DestinationDevNullTestDestinationTypeDestinationDevNullTestDestinationSilent
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationDevNullTestDestination) MarshalJSON() ([]byte, error) {
	if u.DestinationDevNullTestDestinationSilent != nil {
		return json.Marshal(u.DestinationDevNullTestDestinationSilent)
	}

	return nil, nil
}

type DestinationDevNull struct {
	DestinationType DestinationDevNullDevNull `json:"destinationType"`
	// The type of destination to be used
	TestDestination DestinationDevNullTestDestination `json:"test_destination"`
}
