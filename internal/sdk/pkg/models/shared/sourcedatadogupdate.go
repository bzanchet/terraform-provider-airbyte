// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceDatadogUpdateQueriesDataSource - A data source that is powered by the platform.
type SourceDatadogUpdateQueriesDataSource string

const (
	SourceDatadogUpdateQueriesDataSourceMetrics   SourceDatadogUpdateQueriesDataSource = "metrics"
	SourceDatadogUpdateQueriesDataSourceCloudCost SourceDatadogUpdateQueriesDataSource = "cloud_cost"
	SourceDatadogUpdateQueriesDataSourceLogs      SourceDatadogUpdateQueriesDataSource = "logs"
	SourceDatadogUpdateQueriesDataSourceRum       SourceDatadogUpdateQueriesDataSource = "rum"
)

func (e SourceDatadogUpdateQueriesDataSource) ToPointer() *SourceDatadogUpdateQueriesDataSource {
	return &e
}

func (e *SourceDatadogUpdateQueriesDataSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "metrics":
		fallthrough
	case "cloud_cost":
		fallthrough
	case "logs":
		fallthrough
	case "rum":
		*e = SourceDatadogUpdateQueriesDataSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDatadogUpdateQueriesDataSource: %v", v)
	}
}

type SourceDatadogUpdateQueries struct {
	// A data source that is powered by the platform.
	DataSource SourceDatadogUpdateQueriesDataSource `json:"data_source"`
	// The variable name for use in queries.
	Name string `json:"name"`
	// A classic query string.
	Query string `json:"query"`
}

// SourceDatadogUpdateSite - The site where Datadog data resides in.
type SourceDatadogUpdateSite string

const (
	SourceDatadogUpdateSiteDatadoghqCom    SourceDatadogUpdateSite = "datadoghq.com"
	SourceDatadogUpdateSiteUs3DatadoghqCom SourceDatadogUpdateSite = "us3.datadoghq.com"
	SourceDatadogUpdateSiteUs5DatadoghqCom SourceDatadogUpdateSite = "us5.datadoghq.com"
	SourceDatadogUpdateSiteDatadoghqEu     SourceDatadogUpdateSite = "datadoghq.eu"
	SourceDatadogUpdateSiteDdogGovCom      SourceDatadogUpdateSite = "ddog-gov.com"
)

func (e SourceDatadogUpdateSite) ToPointer() *SourceDatadogUpdateSite {
	return &e
}

func (e *SourceDatadogUpdateSite) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "datadoghq.com":
		fallthrough
	case "us3.datadoghq.com":
		fallthrough
	case "us5.datadoghq.com":
		fallthrough
	case "datadoghq.eu":
		fallthrough
	case "ddog-gov.com":
		*e = SourceDatadogUpdateSite(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDatadogUpdateSite: %v", v)
	}
}

type SourceDatadogUpdate struct {
	// Datadog API key
	APIKey string `json:"api_key"`
	// Datadog application key
	ApplicationKey string `json:"application_key"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Data after this date will  not be replicated. An empty value will represent the current datetime for each  execution. This just applies to Incremental syncs.
	EndDate *string `json:"end_date,omitempty"`
	// Maximum number of records to collect per request.
	MaxRecordsPerRequest *int64 `json:"max_records_per_request,omitempty"`
	// List of queries to be run and used as inputs.
	Queries []SourceDatadogUpdateQueries `json:"queries,omitempty"`
	// The search query. This just applies to Incremental syncs. If empty, it'll collect all logs.
	Query *string `json:"query,omitempty"`
	// The site where Datadog data resides in.
	Site *SourceDatadogUpdateSite `json:"site,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. This just applies to Incremental syncs.
	StartDate *string `json:"start_date,omitempty"`
}