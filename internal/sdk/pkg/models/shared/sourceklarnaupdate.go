// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceKlarnaUpdateRegion - Base url region (For playground eu https://docs.klarna.com/klarna-payments/api/payments-api/#tag/API-URLs). Supported 'eu', 'us', 'oc'
type SourceKlarnaUpdateRegion string

const (
	SourceKlarnaUpdateRegionEu SourceKlarnaUpdateRegion = "eu"
	SourceKlarnaUpdateRegionUs SourceKlarnaUpdateRegion = "us"
	SourceKlarnaUpdateRegionOc SourceKlarnaUpdateRegion = "oc"
)

func (e SourceKlarnaUpdateRegion) ToPointer() *SourceKlarnaUpdateRegion {
	return &e
}

func (e *SourceKlarnaUpdateRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "eu":
		fallthrough
	case "us":
		fallthrough
	case "oc":
		*e = SourceKlarnaUpdateRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceKlarnaUpdateRegion: %v", v)
	}
}

type SourceKlarnaUpdate struct {
	// A string which is associated with your Merchant ID and is used to authorize use of Klarna's APIs (https://developers.klarna.com/api/#authentication)
	Password string `json:"password"`
	// Propertie defining if connector is used against playground or production environment
	Playground bool `json:"playground"`
	// Base url region (For playground eu https://docs.klarna.com/klarna-payments/api/payments-api/#tag/API-URLs). Supported 'eu', 'us', 'oc'
	Region SourceKlarnaUpdateRegion `json:"region"`
	// Consists of your Merchant ID (eid) - a unique number that identifies your e-store, combined with a random string (https://developers.klarna.com/api/#authentication)
	Username string `json:"username"`
}
