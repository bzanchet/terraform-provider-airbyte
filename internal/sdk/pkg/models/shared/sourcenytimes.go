// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

// SourceNytimesPeriodUsedForMostPopularStreamsEnum - Period of time (in days)
type SourceNytimesPeriodUsedForMostPopularStreamsEnum int64

const (
	SourceNytimesPeriodUsedForMostPopularStreamsEnumOne    SourceNytimesPeriodUsedForMostPopularStreamsEnum = 1
	SourceNytimesPeriodUsedForMostPopularStreamsEnumSeven  SourceNytimesPeriodUsedForMostPopularStreamsEnum = 7
	SourceNytimesPeriodUsedForMostPopularStreamsEnumThirty SourceNytimesPeriodUsedForMostPopularStreamsEnum = 30
)

func (e SourceNytimesPeriodUsedForMostPopularStreamsEnum) ToPointer() *SourceNytimesPeriodUsedForMostPopularStreamsEnum {
	return &e
}

func (e *SourceNytimesPeriodUsedForMostPopularStreamsEnum) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 30:
		*e = SourceNytimesPeriodUsedForMostPopularStreamsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNytimesPeriodUsedForMostPopularStreamsEnum: %v", v)
	}
}

// SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum - Share Type
type SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum string

const (
	SourceNytimesShareTypeUsedForMostPopularSharedStreamEnumFacebook SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum = "facebook"
)

func (e SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum) ToPointer() *SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum {
	return &e
}

func (e *SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "facebook":
		*e = SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum: %v", v)
	}
}

type SourceNytimesNytimesEnum string

const (
	SourceNytimesNytimesEnumNytimes SourceNytimesNytimesEnum = "nytimes"
)

func (e SourceNytimesNytimesEnum) ToPointer() *SourceNytimesNytimesEnum {
	return &e
}

func (e *SourceNytimesNytimesEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "nytimes":
		*e = SourceNytimesNytimesEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNytimesNytimesEnum: %v", v)
	}
}

type SourceNytimes struct {
	// API Key
	APIKey string `json:"api_key"`
	// End date to stop the article retrieval (format YYYY-MM)
	EndDate *types.Date `json:"end_date,omitempty"`
	// Period of time (in days)
	Period SourceNytimesPeriodUsedForMostPopularStreamsEnum `json:"period"`
	// Share Type
	ShareType  *SourceNytimesShareTypeUsedForMostPopularSharedStreamEnum `json:"share_type,omitempty"`
	SourceType SourceNytimesNytimesEnum                                  `json:"sourceType"`
	// Start date to begin the article retrieval (format YYYY-MM)
	StartDate types.Date `json:"start_date"`
}
