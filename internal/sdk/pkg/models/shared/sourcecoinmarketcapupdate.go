// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceCoinmarketcapUpdateDataType - /latest: Latest market ticker quotes and averages for cryptocurrencies and exchanges. /historical: Intervals of historic market data like OHLCV data or data for use in charting libraries. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Endpoint-Overview">here</a>.
type SourceCoinmarketcapUpdateDataType string

const (
	SourceCoinmarketcapUpdateDataTypeLatest     SourceCoinmarketcapUpdateDataType = "latest"
	SourceCoinmarketcapUpdateDataTypeHistorical SourceCoinmarketcapUpdateDataType = "historical"
)

func (e SourceCoinmarketcapUpdateDataType) ToPointer() *SourceCoinmarketcapUpdateDataType {
	return &e
}

func (e *SourceCoinmarketcapUpdateDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "latest":
		fallthrough
	case "historical":
		*e = SourceCoinmarketcapUpdateDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceCoinmarketcapUpdateDataType: %v", v)
	}
}

type SourceCoinmarketcapUpdate struct {
	// Your API Key. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Authentication">here</a>. The token is case sensitive.
	APIKey string `json:"api_key"`
	// /latest: Latest market ticker quotes and averages for cryptocurrencies and exchanges. /historical: Intervals of historic market data like OHLCV data or data for use in charting libraries. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Endpoint-Overview">here</a>.
	DataType SourceCoinmarketcapUpdateDataType `json:"data_type"`
	// Cryptocurrency symbols. (only used for quotes stream)
	Symbols []string `json:"symbols,omitempty"`
}
