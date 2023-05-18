// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationFireboltFireboltEnum string

const (
	DestinationFireboltFireboltEnumFirebolt DestinationFireboltFireboltEnum = "firebolt"
)

func (e DestinationFireboltFireboltEnum) ToPointer() *DestinationFireboltFireboltEnum {
	return &e
}

func (e *DestinationFireboltFireboltEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "firebolt":
		*e = DestinationFireboltFireboltEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationFireboltFireboltEnum: %v", v)
	}
}

type DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum string

const (
	DestinationFireboltLoadingMethodExternalTableViaS3MethodEnumS3 DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum = "S3"
)

func (e DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum) ToPointer() *DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum {
	return &e
}

func (e *DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3":
		*e = DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum: %v", v)
	}
}

// DestinationFireboltLoadingMethodExternalTableViaS3 - Loading method used to select the way data will be uploaded to Firebolt
type DestinationFireboltLoadingMethodExternalTableViaS3 struct {
	// AWS access key granting read and write access to S3.
	AwsKeyID string `json:"aws_key_id"`
	// Corresponding secret part of the AWS Key
	AwsKeySecret string                                                       `json:"aws_key_secret"`
	Method       DestinationFireboltLoadingMethodExternalTableViaS3MethodEnum `json:"method"`
	// The name of the S3 bucket.
	S3Bucket string `json:"s3_bucket"`
	// Region name of the S3 bucket.
	S3Region string `json:"s3_region"`
}

type DestinationFireboltLoadingMethodSQLInsertsMethodEnum string

const (
	DestinationFireboltLoadingMethodSQLInsertsMethodEnumSQL DestinationFireboltLoadingMethodSQLInsertsMethodEnum = "SQL"
)

func (e DestinationFireboltLoadingMethodSQLInsertsMethodEnum) ToPointer() *DestinationFireboltLoadingMethodSQLInsertsMethodEnum {
	return &e
}

func (e *DestinationFireboltLoadingMethodSQLInsertsMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SQL":
		*e = DestinationFireboltLoadingMethodSQLInsertsMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationFireboltLoadingMethodSQLInsertsMethodEnum: %v", v)
	}
}

// DestinationFireboltLoadingMethodSQLInserts - Loading method used to select the way data will be uploaded to Firebolt
type DestinationFireboltLoadingMethodSQLInserts struct {
	Method DestinationFireboltLoadingMethodSQLInsertsMethodEnum `json:"method"`
}

type DestinationFireboltLoadingMethodType string

const (
	DestinationFireboltLoadingMethodTypeDestinationFireboltLoadingMethodSQLInserts         DestinationFireboltLoadingMethodType = "destination-firebolt_Loading Method_SQL Inserts"
	DestinationFireboltLoadingMethodTypeDestinationFireboltLoadingMethodExternalTableViaS3 DestinationFireboltLoadingMethodType = "destination-firebolt_Loading Method_External Table via S3"
)

type DestinationFireboltLoadingMethod struct {
	DestinationFireboltLoadingMethodSQLInserts         *DestinationFireboltLoadingMethodSQLInserts
	DestinationFireboltLoadingMethodExternalTableViaS3 *DestinationFireboltLoadingMethodExternalTableViaS3

	Type DestinationFireboltLoadingMethodType
}

func CreateDestinationFireboltLoadingMethodDestinationFireboltLoadingMethodSQLInserts(destinationFireboltLoadingMethodSQLInserts DestinationFireboltLoadingMethodSQLInserts) DestinationFireboltLoadingMethod {
	typ := DestinationFireboltLoadingMethodTypeDestinationFireboltLoadingMethodSQLInserts

	return DestinationFireboltLoadingMethod{
		DestinationFireboltLoadingMethodSQLInserts: &destinationFireboltLoadingMethodSQLInserts,
		Type: typ,
	}
}

func CreateDestinationFireboltLoadingMethodDestinationFireboltLoadingMethodExternalTableViaS3(destinationFireboltLoadingMethodExternalTableViaS3 DestinationFireboltLoadingMethodExternalTableViaS3) DestinationFireboltLoadingMethod {
	typ := DestinationFireboltLoadingMethodTypeDestinationFireboltLoadingMethodExternalTableViaS3

	return DestinationFireboltLoadingMethod{
		DestinationFireboltLoadingMethodExternalTableViaS3: &destinationFireboltLoadingMethodExternalTableViaS3,
		Type: typ,
	}
}

func (u *DestinationFireboltLoadingMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationFireboltLoadingMethodSQLInserts := new(DestinationFireboltLoadingMethodSQLInserts)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationFireboltLoadingMethodSQLInserts); err == nil {
		u.DestinationFireboltLoadingMethodSQLInserts = destinationFireboltLoadingMethodSQLInserts
		u.Type = DestinationFireboltLoadingMethodTypeDestinationFireboltLoadingMethodSQLInserts
		return nil
	}

	destinationFireboltLoadingMethodExternalTableViaS3 := new(DestinationFireboltLoadingMethodExternalTableViaS3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationFireboltLoadingMethodExternalTableViaS3); err == nil {
		u.DestinationFireboltLoadingMethodExternalTableViaS3 = destinationFireboltLoadingMethodExternalTableViaS3
		u.Type = DestinationFireboltLoadingMethodTypeDestinationFireboltLoadingMethodExternalTableViaS3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationFireboltLoadingMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationFireboltLoadingMethodSQLInserts != nil {
		return json.Marshal(u.DestinationFireboltLoadingMethodSQLInserts)
	}

	if u.DestinationFireboltLoadingMethodExternalTableViaS3 != nil {
		return json.Marshal(u.DestinationFireboltLoadingMethodExternalTableViaS3)
	}

	return nil, nil
}

type DestinationFirebolt struct {
	// Firebolt account to login.
	Account *string `json:"account,omitempty"`
	// The database to connect to.
	Database        string                          `json:"database"`
	DestinationType DestinationFireboltFireboltEnum `json:"destinationType"`
	// Engine name or url to connect to.
	Engine *string `json:"engine,omitempty"`
	// The host name of your Firebolt database.
	Host *string `json:"host,omitempty"`
	// Loading method used to select the way data will be uploaded to Firebolt
	LoadingMethod *DestinationFireboltLoadingMethod `json:"loading_method,omitempty"`
	// Firebolt password.
	Password string `json:"password"`
	// Firebolt email address you use to login.
	Username string `json:"username"`
}
