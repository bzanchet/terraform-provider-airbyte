// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationAmazonSqsUpdateAWSRegion - AWS Region of the SQS Queue
type DestinationAmazonSqsUpdateAWSRegion string

const (
	DestinationAmazonSqsUpdateAWSRegionUsEast1      DestinationAmazonSqsUpdateAWSRegion = "us-east-1"
	DestinationAmazonSqsUpdateAWSRegionUsEast2      DestinationAmazonSqsUpdateAWSRegion = "us-east-2"
	DestinationAmazonSqsUpdateAWSRegionUsWest1      DestinationAmazonSqsUpdateAWSRegion = "us-west-1"
	DestinationAmazonSqsUpdateAWSRegionUsWest2      DestinationAmazonSqsUpdateAWSRegion = "us-west-2"
	DestinationAmazonSqsUpdateAWSRegionAfSouth1     DestinationAmazonSqsUpdateAWSRegion = "af-south-1"
	DestinationAmazonSqsUpdateAWSRegionApEast1      DestinationAmazonSqsUpdateAWSRegion = "ap-east-1"
	DestinationAmazonSqsUpdateAWSRegionApSouth1     DestinationAmazonSqsUpdateAWSRegion = "ap-south-1"
	DestinationAmazonSqsUpdateAWSRegionApNortheast1 DestinationAmazonSqsUpdateAWSRegion = "ap-northeast-1"
	DestinationAmazonSqsUpdateAWSRegionApNortheast2 DestinationAmazonSqsUpdateAWSRegion = "ap-northeast-2"
	DestinationAmazonSqsUpdateAWSRegionApNortheast3 DestinationAmazonSqsUpdateAWSRegion = "ap-northeast-3"
	DestinationAmazonSqsUpdateAWSRegionApSoutheast1 DestinationAmazonSqsUpdateAWSRegion = "ap-southeast-1"
	DestinationAmazonSqsUpdateAWSRegionApSoutheast2 DestinationAmazonSqsUpdateAWSRegion = "ap-southeast-2"
	DestinationAmazonSqsUpdateAWSRegionCaCentral1   DestinationAmazonSqsUpdateAWSRegion = "ca-central-1"
	DestinationAmazonSqsUpdateAWSRegionCnNorth1     DestinationAmazonSqsUpdateAWSRegion = "cn-north-1"
	DestinationAmazonSqsUpdateAWSRegionCnNorthwest1 DestinationAmazonSqsUpdateAWSRegion = "cn-northwest-1"
	DestinationAmazonSqsUpdateAWSRegionEuCentral1   DestinationAmazonSqsUpdateAWSRegion = "eu-central-1"
	DestinationAmazonSqsUpdateAWSRegionEuNorth1     DestinationAmazonSqsUpdateAWSRegion = "eu-north-1"
	DestinationAmazonSqsUpdateAWSRegionEuSouth1     DestinationAmazonSqsUpdateAWSRegion = "eu-south-1"
	DestinationAmazonSqsUpdateAWSRegionEuWest1      DestinationAmazonSqsUpdateAWSRegion = "eu-west-1"
	DestinationAmazonSqsUpdateAWSRegionEuWest2      DestinationAmazonSqsUpdateAWSRegion = "eu-west-2"
	DestinationAmazonSqsUpdateAWSRegionEuWest3      DestinationAmazonSqsUpdateAWSRegion = "eu-west-3"
	DestinationAmazonSqsUpdateAWSRegionSaEast1      DestinationAmazonSqsUpdateAWSRegion = "sa-east-1"
	DestinationAmazonSqsUpdateAWSRegionMeSouth1     DestinationAmazonSqsUpdateAWSRegion = "me-south-1"
	DestinationAmazonSqsUpdateAWSRegionUsGovEast1   DestinationAmazonSqsUpdateAWSRegion = "us-gov-east-1"
	DestinationAmazonSqsUpdateAWSRegionUsGovWest1   DestinationAmazonSqsUpdateAWSRegion = "us-gov-west-1"
)

func (e DestinationAmazonSqsUpdateAWSRegion) ToPointer() *DestinationAmazonSqsUpdateAWSRegion {
	return &e
}

func (e *DestinationAmazonSqsUpdateAWSRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
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
		*e = DestinationAmazonSqsUpdateAWSRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAmazonSqsUpdateAWSRegion: %v", v)
	}
}

type DestinationAmazonSqsUpdate struct {
	// The Access Key ID of the AWS IAM Role to use for sending  messages
	AccessKey *string `json:"access_key,omitempty"`
	// Use this property to extract the contents of the named key in the input record to use as the SQS message body. If not set, the entire content of the input record data is used as the message body.
	MessageBodyKey *string `json:"message_body_key,omitempty"`
	// Modify the Message Delay of the individual message from the Queue's default (seconds).
	MessageDelay *int64 `json:"message_delay,omitempty"`
	// The tag that specifies that a message belongs to a specific message group. This parameter applies only to, and is REQUIRED by, FIFO queues.
	MessageGroupID *string `json:"message_group_id,omitempty"`
	// URL of the SQS Queue
	QueueURL string `json:"queue_url"`
	// AWS Region of the SQS Queue
	Region DestinationAmazonSqsUpdateAWSRegion `json:"region"`
	// The Secret Key of the AWS IAM Role to use for sending messages
	SecretKey *string `json:"secret_key,omitempty"`
}
