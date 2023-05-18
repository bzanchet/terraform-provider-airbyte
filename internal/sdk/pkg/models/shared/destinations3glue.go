// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationS3GlueS3GlueEnum string

const (
	DestinationS3GlueS3GlueEnumS3Glue DestinationS3GlueS3GlueEnum = "s3-glue"
)

func (e DestinationS3GlueS3GlueEnum) ToPointer() *DestinationS3GlueS3GlueEnum {
	return &e
}

func (e *DestinationS3GlueS3GlueEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "s3-glue":
		*e = DestinationS3GlueS3GlueEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueS3GlueEnum: %v", v)
	}
}

type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum string

const (
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnumGzip DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum = "GZIP"
)

func (e DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum) ToPointer() *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum {
	return &e
}

func (e *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GZIP":
		*e = DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum: %v", v)
	}
}

// DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP - Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").
type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP struct {
	CompressionType *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIPCompressionTypeEnum `json:"compression_type,omitempty"`
}

type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum string

const (
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnumNoCompression DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum = "No Compression"
)

func (e DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum) ToPointer() *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum {
	return &e
}

func (e *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No Compression":
		*e = DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum: %v", v)
	}
}

// DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression - Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").
type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression struct {
	CompressionType *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompressionCompressionTypeEnum `json:"compression_type,omitempty"`
}

type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionType string

const (
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionType = "destination-s3-glue_Output Format_JSON Lines: Newline-delimited JSON_Compression_No Compression"
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP          DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionType = "destination-s3-glue_Output Format_JSON Lines: Newline-delimited JSON_Compression_GZIP"
)

type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression struct {
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP          *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP

	Type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionType
}

func CreateDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression(destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression) DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression {
	typ := DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression

	return DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression{
		DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression: &destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression,
		Type: typ,
	}
}

func CreateDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP(destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP) DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression {
	typ := DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP

	return DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression{
		DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP: &destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP,
		Type: typ,
	}
}

func (u *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression := new(DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression); err == nil {
		u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression = destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression
		u.Type = DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression
		return nil
	}

	destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP := new(DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP); err == nil {
		u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP = destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP
		u.Type = DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression) MarshalJSON() ([]byte, error) {
	if u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression != nil {
		return json.Marshal(u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression)
	}

	if u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP != nil {
		return json.Marshal(u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP)
	}

	return nil, nil
}

// DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum - Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.
type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum string

const (
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnumNoFlattening        DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum = "No flattening"
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnumRootLevelFlattening DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum = "Root level flattening"
)

func (e DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum) ToPointer() *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum {
	return &e
}

func (e *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No flattening":
		fallthrough
	case "Root level flattening":
		*e = DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum: %v", v)
	}
}

type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum string

const (
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnumJsonl DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum = "JSONL"
)

func (e DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum) ToPointer() *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum {
	return &e
}

func (e *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum: %v", v)
	}
}

// DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON - Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details
type DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON struct {
	// Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").
	Compression *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONCompression `json:"compression,omitempty"`
	// Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.
	Flattening *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFlatteningEnum `json:"flattening,omitempty"`
	FormatType DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeEnum  `json:"format_type"`
}

type DestinationS3GlueOutputFormatType string

const (
	DestinationS3GlueOutputFormatTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON DestinationS3GlueOutputFormatType = "destination-s3-glue_Output Format_JSON Lines: Newline-delimited JSON"
)

type DestinationS3GlueOutputFormat struct {
	DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON *DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON

	Type DestinationS3GlueOutputFormatType
}

func CreateDestinationS3GlueOutputFormatDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON(destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON) DestinationS3GlueOutputFormat {
	typ := DestinationS3GlueOutputFormatTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON

	return DestinationS3GlueOutputFormat{
		DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON: &destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func (u *DestinationS3GlueOutputFormat) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON := new(DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON); err == nil {
		u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON = destinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON
		u.Type = DestinationS3GlueOutputFormatTypeDestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationS3GlueOutputFormat) MarshalJSON() ([]byte, error) {
	if u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		return json.Marshal(u.DestinationS3GlueOutputFormatJSONLinesNewlineDelimitedJSON)
	}

	return nil, nil
}

// DestinationS3GlueSerializationLibraryEnum - The library that your query engine will use for reading and writing data in your lake.
type DestinationS3GlueSerializationLibraryEnum string

const (
	DestinationS3GlueSerializationLibraryEnumOrgOpenxDataJsonserdeJSONSerDe     DestinationS3GlueSerializationLibraryEnum = "org.openx.data.jsonserde.JsonSerDe"
	DestinationS3GlueSerializationLibraryEnumOrgApacheHiveHcatalogDataJSONSerDe DestinationS3GlueSerializationLibraryEnum = "org.apache.hive.hcatalog.data.JsonSerDe"
)

func (e DestinationS3GlueSerializationLibraryEnum) ToPointer() *DestinationS3GlueSerializationLibraryEnum {
	return &e
}

func (e *DestinationS3GlueSerializationLibraryEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "org.openx.data.jsonserde.JsonSerDe":
		fallthrough
	case "org.apache.hive.hcatalog.data.JsonSerDe":
		*e = DestinationS3GlueSerializationLibraryEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueSerializationLibraryEnum: %v", v)
	}
}

// DestinationS3GlueS3BucketRegionEnum - The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
type DestinationS3GlueS3BucketRegionEnum string

const (
	DestinationS3GlueS3BucketRegionEnumUnknown      DestinationS3GlueS3BucketRegionEnum = ""
	DestinationS3GlueS3BucketRegionEnumUsEast1      DestinationS3GlueS3BucketRegionEnum = "us-east-1"
	DestinationS3GlueS3BucketRegionEnumUsEast2      DestinationS3GlueS3BucketRegionEnum = "us-east-2"
	DestinationS3GlueS3BucketRegionEnumUsWest1      DestinationS3GlueS3BucketRegionEnum = "us-west-1"
	DestinationS3GlueS3BucketRegionEnumUsWest2      DestinationS3GlueS3BucketRegionEnum = "us-west-2"
	DestinationS3GlueS3BucketRegionEnumAfSouth1     DestinationS3GlueS3BucketRegionEnum = "af-south-1"
	DestinationS3GlueS3BucketRegionEnumApEast1      DestinationS3GlueS3BucketRegionEnum = "ap-east-1"
	DestinationS3GlueS3BucketRegionEnumApSouth1     DestinationS3GlueS3BucketRegionEnum = "ap-south-1"
	DestinationS3GlueS3BucketRegionEnumApNortheast1 DestinationS3GlueS3BucketRegionEnum = "ap-northeast-1"
	DestinationS3GlueS3BucketRegionEnumApNortheast2 DestinationS3GlueS3BucketRegionEnum = "ap-northeast-2"
	DestinationS3GlueS3BucketRegionEnumApNortheast3 DestinationS3GlueS3BucketRegionEnum = "ap-northeast-3"
	DestinationS3GlueS3BucketRegionEnumApSoutheast1 DestinationS3GlueS3BucketRegionEnum = "ap-southeast-1"
	DestinationS3GlueS3BucketRegionEnumApSoutheast2 DestinationS3GlueS3BucketRegionEnum = "ap-southeast-2"
	DestinationS3GlueS3BucketRegionEnumCaCentral1   DestinationS3GlueS3BucketRegionEnum = "ca-central-1"
	DestinationS3GlueS3BucketRegionEnumCnNorth1     DestinationS3GlueS3BucketRegionEnum = "cn-north-1"
	DestinationS3GlueS3BucketRegionEnumCnNorthwest1 DestinationS3GlueS3BucketRegionEnum = "cn-northwest-1"
	DestinationS3GlueS3BucketRegionEnumEuCentral1   DestinationS3GlueS3BucketRegionEnum = "eu-central-1"
	DestinationS3GlueS3BucketRegionEnumEuNorth1     DestinationS3GlueS3BucketRegionEnum = "eu-north-1"
	DestinationS3GlueS3BucketRegionEnumEuSouth1     DestinationS3GlueS3BucketRegionEnum = "eu-south-1"
	DestinationS3GlueS3BucketRegionEnumEuWest1      DestinationS3GlueS3BucketRegionEnum = "eu-west-1"
	DestinationS3GlueS3BucketRegionEnumEuWest2      DestinationS3GlueS3BucketRegionEnum = "eu-west-2"
	DestinationS3GlueS3BucketRegionEnumEuWest3      DestinationS3GlueS3BucketRegionEnum = "eu-west-3"
	DestinationS3GlueS3BucketRegionEnumSaEast1      DestinationS3GlueS3BucketRegionEnum = "sa-east-1"
	DestinationS3GlueS3BucketRegionEnumMeSouth1     DestinationS3GlueS3BucketRegionEnum = "me-south-1"
	DestinationS3GlueS3BucketRegionEnumUsGovEast1   DestinationS3GlueS3BucketRegionEnum = "us-gov-east-1"
	DestinationS3GlueS3BucketRegionEnumUsGovWest1   DestinationS3GlueS3BucketRegionEnum = "us-gov-west-1"
)

func (e DestinationS3GlueS3BucketRegionEnum) ToPointer() *DestinationS3GlueS3BucketRegionEnum {
	return &e
}

func (e *DestinationS3GlueS3BucketRegionEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = DestinationS3GlueS3BucketRegionEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueS3BucketRegionEnum: %v", v)
	}
}

type DestinationS3Glue struct {
	// The access key ID to access the S3 bucket. Airbyte requires Read and Write permissions to the given bucket. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>.
	AccessKeyID     *string                     `json:"access_key_id,omitempty"`
	DestinationType DestinationS3GlueS3GlueEnum `json:"destinationType"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string `json:"file_name_pattern,omitempty"`
	// Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details
	Format DestinationS3GlueOutputFormat `json:"format"`
	// Name of the glue database for creating the tables, leave blank if no integration
	GlueDatabase string `json:"glue_database"`
	// The library that your query engine will use for reading and writing data in your lake.
	GlueSerializationLibrary DestinationS3GlueSerializationLibraryEnum `json:"glue_serialization_library"`
	// The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.
	S3BucketName string `json:"s3_bucket_name"`
	// Directory under the S3 bucket where data will be written. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=to%20format%20the-,bucket%20path,-%3A">here</a>
	S3BucketPath string `json:"s3_bucket_path"`
	// The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
	S3BucketRegion DestinationS3GlueS3BucketRegionEnum `json:"s3_bucket_region"`
	// Your S3 endpoint url. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/s3.html#:~:text=Service%20endpoints-,Amazon%20S3%20endpoints,-When%20you%20use">here</a>
	S3Endpoint *string `json:"s3_endpoint,omitempty"`
	// Format string on how data will be organized inside the S3 bucket directory. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=The%20full%20path%20of%20the%20output%20data%20with%20the%20default%20S3%20path%20format">here</a>
	S3PathFormat *string `json:"s3_path_format,omitempty"`
	// The corresponding secret to the access key ID. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
}
