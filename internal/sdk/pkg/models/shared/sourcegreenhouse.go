// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGreenhouseGreenhouseEnum string

const (
	SourceGreenhouseGreenhouseEnumGreenhouse SourceGreenhouseGreenhouseEnum = "greenhouse"
)

func (e SourceGreenhouseGreenhouseEnum) ToPointer() *SourceGreenhouseGreenhouseEnum {
	return &e
}

func (e *SourceGreenhouseGreenhouseEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "greenhouse":
		*e = SourceGreenhouseGreenhouseEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGreenhouseGreenhouseEnum: %v", v)
	}
}

type SourceGreenhouse struct {
	// Greenhouse API Key. See the <a href="https://docs.airbyte.com/integrations/sources/greenhouse">docs</a> for more information on how to generate this key.
	APIKey     string                         `json:"api_key"`
	SourceType SourceGreenhouseGreenhouseEnum `json:"sourceType"`
}
