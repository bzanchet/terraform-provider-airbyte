// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// NonBreakingSchemaUpdatesBehaviorEnumEnum - Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
type NonBreakingSchemaUpdatesBehaviorEnumEnum string

const (
	NonBreakingSchemaUpdatesBehaviorEnumEnumIgnore            NonBreakingSchemaUpdatesBehaviorEnumEnum = "ignore"
	NonBreakingSchemaUpdatesBehaviorEnumEnumDisableConnection NonBreakingSchemaUpdatesBehaviorEnumEnum = "disable_connection"
)

func (e NonBreakingSchemaUpdatesBehaviorEnumEnum) ToPointer() *NonBreakingSchemaUpdatesBehaviorEnumEnum {
	return &e
}

func (e *NonBreakingSchemaUpdatesBehaviorEnumEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ignore":
		fallthrough
	case "disable_connection":
		*e = NonBreakingSchemaUpdatesBehaviorEnumEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NonBreakingSchemaUpdatesBehaviorEnumEnum: %v", v)
	}
}
