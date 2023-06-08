// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// SourceChargebeeUpdateProductCatalog - Product Catalog version of your Chargebee site. Instructions on how to find your version you may find <a href="https://apidocs.chargebee.com/docs/api?prod_cat_ver=2">here</a> under `API Version` section.
type SourceChargebeeUpdateProductCatalog string

const (
	SourceChargebeeUpdateProductCatalogOne0 SourceChargebeeUpdateProductCatalog = "1.0"
	SourceChargebeeUpdateProductCatalogTwo0 SourceChargebeeUpdateProductCatalog = "2.0"
)

func (e SourceChargebeeUpdateProductCatalog) ToPointer() *SourceChargebeeUpdateProductCatalog {
	return &e
}

func (e *SourceChargebeeUpdateProductCatalog) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1.0":
		fallthrough
	case "2.0":
		*e = SourceChargebeeUpdateProductCatalog(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceChargebeeUpdateProductCatalog: %v", v)
	}
}

type SourceChargebeeUpdate struct {
	// Product Catalog version of your Chargebee site. Instructions on how to find your version you may find <a href="https://apidocs.chargebee.com/docs/api?prod_cat_ver=2">here</a> under `API Version` section.
	ProductCatalog SourceChargebeeUpdateProductCatalog `json:"product_catalog"`
	// The site prefix for your Chargebee instance.
	Site string `json:"site"`
	// Chargebee API Key. See the <a href="https://docs.airbyte.com/integrations/sources/chargebee">docs</a> for more information on how to obtain this key.
	SiteAPIKey string `json:"site_api_key"`
	// UTC date and time in the format 2021-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}