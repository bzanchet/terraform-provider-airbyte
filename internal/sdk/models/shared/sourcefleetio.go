// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Fleetio string

const (
	FleetioFleetio Fleetio = "fleetio"
)

func (e Fleetio) ToPointer() *Fleetio {
	return &e
}
func (e *Fleetio) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fleetio":
		*e = Fleetio(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Fleetio: %v", v)
	}
}

type SourceFleetio struct {
	AccountToken string  `json:"account_token"`
	APIKey       string  `json:"api_key"`
	sourceType   Fleetio `const:"fleetio" json:"sourceType"`
}

func (s SourceFleetio) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFleetio) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFleetio) GetAccountToken() string {
	if o == nil {
		return ""
	}
	return o.AccountToken
}

func (o *SourceFleetio) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceFleetio) GetSourceType() Fleetio {
	return FleetioFleetio
}
