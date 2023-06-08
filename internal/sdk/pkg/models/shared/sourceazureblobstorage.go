// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType string

const (
	SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatTypeJsonl SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType = "JSONL"
)

func (e SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType) ToPointer() *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType {
	return &e
}

func (e *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType: %v", v)
	}
}

// SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON - Input data format
type SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON struct {
	FormatType SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType `json:"format_type"`
}

type SourceAzureBlobStorageInputFormatType string

const (
	SourceAzureBlobStorageInputFormatTypeSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON SourceAzureBlobStorageInputFormatType = "source-azure-blob-storage_Input Format_JSON Lines: newline-delimited JSON"
)

type SourceAzureBlobStorageInputFormat struct {
	SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON

	Type SourceAzureBlobStorageInputFormatType
}

func CreateSourceAzureBlobStorageInputFormatSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON(sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON) SourceAzureBlobStorageInputFormat {
	typ := SourceAzureBlobStorageInputFormatTypeSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON

	return SourceAzureBlobStorageInputFormat{
		SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON: &sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func (u *SourceAzureBlobStorageInputFormat) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON := new(SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON); err == nil {
		u.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON = sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON
		u.Type = SourceAzureBlobStorageInputFormatTypeSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAzureBlobStorageInputFormat) MarshalJSON() ([]byte, error) {
	if u.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON != nil {
		return json.Marshal(u.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON)
	}

	return nil, nil
}

type SourceAzureBlobStorageAzureBlobStorage string

const (
	SourceAzureBlobStorageAzureBlobStorageAzureBlobStorage SourceAzureBlobStorageAzureBlobStorage = "azure-blob-storage"
)

func (e SourceAzureBlobStorageAzureBlobStorage) ToPointer() *SourceAzureBlobStorageAzureBlobStorage {
	return &e
}

func (e *SourceAzureBlobStorageAzureBlobStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure-blob-storage":
		*e = SourceAzureBlobStorageAzureBlobStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAzureBlobStorageAzureBlobStorage: %v", v)
	}
}

type SourceAzureBlobStorage struct {
	// The Azure blob storage account key.
	AzureBlobStorageAccountKey string `json:"azure_blob_storage_account_key"`
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The Azure blob storage prefix to be applied
	AzureBlobStorageBlobsPrefix *string `json:"azure_blob_storage_blobs_prefix,omitempty"`
	// The name of the Azure blob storage container.
	AzureBlobStorageContainerName string `json:"azure_blob_storage_container_name"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpoint *string `json:"azure_blob_storage_endpoint,omitempty"`
	// The Azure blob storage blobs to scan for inferring the schema, useful on large amounts of data with consistent structure
	AzureBlobStorageSchemaInferenceLimit *int64 `json:"azure_blob_storage_schema_inference_limit,omitempty"`
	// Input data format
	Format     SourceAzureBlobStorageInputFormat      `json:"format"`
	SourceType SourceAzureBlobStorageAzureBlobStorage `json:"sourceType"`
}