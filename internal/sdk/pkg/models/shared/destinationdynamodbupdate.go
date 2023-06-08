// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationDynamodbUpdateDynamoDBRegion - The region of the DynamoDB.
type DestinationDynamodbUpdateDynamoDBRegion string

const (
	DestinationDynamodbUpdateDynamoDBRegionUnknown      DestinationDynamodbUpdateDynamoDBRegion = ""
	DestinationDynamodbUpdateDynamoDBRegionUsEast1      DestinationDynamodbUpdateDynamoDBRegion = "us-east-1"
	DestinationDynamodbUpdateDynamoDBRegionUsEast2      DestinationDynamodbUpdateDynamoDBRegion = "us-east-2"
	DestinationDynamodbUpdateDynamoDBRegionUsWest1      DestinationDynamodbUpdateDynamoDBRegion = "us-west-1"
	DestinationDynamodbUpdateDynamoDBRegionUsWest2      DestinationDynamodbUpdateDynamoDBRegion = "us-west-2"
	DestinationDynamodbUpdateDynamoDBRegionAfSouth1     DestinationDynamodbUpdateDynamoDBRegion = "af-south-1"
	DestinationDynamodbUpdateDynamoDBRegionApEast1      DestinationDynamodbUpdateDynamoDBRegion = "ap-east-1"
	DestinationDynamodbUpdateDynamoDBRegionApSouth1     DestinationDynamodbUpdateDynamoDBRegion = "ap-south-1"
	DestinationDynamodbUpdateDynamoDBRegionApNortheast1 DestinationDynamodbUpdateDynamoDBRegion = "ap-northeast-1"
	DestinationDynamodbUpdateDynamoDBRegionApNortheast2 DestinationDynamodbUpdateDynamoDBRegion = "ap-northeast-2"
	DestinationDynamodbUpdateDynamoDBRegionApNortheast3 DestinationDynamodbUpdateDynamoDBRegion = "ap-northeast-3"
	DestinationDynamodbUpdateDynamoDBRegionApSoutheast1 DestinationDynamodbUpdateDynamoDBRegion = "ap-southeast-1"
	DestinationDynamodbUpdateDynamoDBRegionApSoutheast2 DestinationDynamodbUpdateDynamoDBRegion = "ap-southeast-2"
	DestinationDynamodbUpdateDynamoDBRegionCaCentral1   DestinationDynamodbUpdateDynamoDBRegion = "ca-central-1"
	DestinationDynamodbUpdateDynamoDBRegionCnNorth1     DestinationDynamodbUpdateDynamoDBRegion = "cn-north-1"
	DestinationDynamodbUpdateDynamoDBRegionCnNorthwest1 DestinationDynamodbUpdateDynamoDBRegion = "cn-northwest-1"
	DestinationDynamodbUpdateDynamoDBRegionEuCentral1   DestinationDynamodbUpdateDynamoDBRegion = "eu-central-1"
	DestinationDynamodbUpdateDynamoDBRegionEuNorth1     DestinationDynamodbUpdateDynamoDBRegion = "eu-north-1"
	DestinationDynamodbUpdateDynamoDBRegionEuSouth1     DestinationDynamodbUpdateDynamoDBRegion = "eu-south-1"
	DestinationDynamodbUpdateDynamoDBRegionEuWest1      DestinationDynamodbUpdateDynamoDBRegion = "eu-west-1"
	DestinationDynamodbUpdateDynamoDBRegionEuWest2      DestinationDynamodbUpdateDynamoDBRegion = "eu-west-2"
	DestinationDynamodbUpdateDynamoDBRegionEuWest3      DestinationDynamodbUpdateDynamoDBRegion = "eu-west-3"
	DestinationDynamodbUpdateDynamoDBRegionSaEast1      DestinationDynamodbUpdateDynamoDBRegion = "sa-east-1"
	DestinationDynamodbUpdateDynamoDBRegionMeSouth1     DestinationDynamodbUpdateDynamoDBRegion = "me-south-1"
	DestinationDynamodbUpdateDynamoDBRegionUsGovEast1   DestinationDynamodbUpdateDynamoDBRegion = "us-gov-east-1"
	DestinationDynamodbUpdateDynamoDBRegionUsGovWest1   DestinationDynamodbUpdateDynamoDBRegion = "us-gov-west-1"
)

func (e DestinationDynamodbUpdateDynamoDBRegion) ToPointer() *DestinationDynamodbUpdateDynamoDBRegion {
	return &e
}

func (e *DestinationDynamodbUpdateDynamoDBRegion) UnmarshalJSON(data []byte) error {
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
		*e = DestinationDynamodbUpdateDynamoDBRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDynamodbUpdateDynamoDBRegion: %v", v)
	}
}

type DestinationDynamodbUpdate struct {
	// The access key id to access the DynamoDB. Airbyte requires Read and Write permissions to the DynamoDB.
	AccessKeyID string `json:"access_key_id"`
	// This is your DynamoDB endpoint url.(if you are working with AWS DynamoDB, just leave empty).
	DynamodbEndpoint *string `json:"dynamodb_endpoint,omitempty"`
	// The region of the DynamoDB.
	DynamodbRegion DestinationDynamodbUpdateDynamoDBRegion `json:"dynamodb_region"`
	// The prefix to use when naming DynamoDB tables.
	DynamodbTableNamePrefix string `json:"dynamodb_table_name_prefix"`
	// The corresponding secret to the access key id.
	SecretAccessKey string `json:"secret_access_key"`
}