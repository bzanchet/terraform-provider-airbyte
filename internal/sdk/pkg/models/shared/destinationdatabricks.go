// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum string

const (
	DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnumAzureBlobStorage DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum = "AZURE_BLOB_STORAGE"
)

func (e DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum) ToPointer() *DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum {
	return &e
}

func (e *DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AZURE_BLOB_STORAGE":
		*e = DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum: %v", v)
	}
}

// DestinationDatabricksDataSourceAzureBlobStorage - Storage on which the delta lake is built.
type DestinationDatabricksDataSourceAzureBlobStorage struct {
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The name of the Azure blob storage container.
	AzureBlobStorageContainerName string `json:"azure_blob_storage_container_name"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpointDomainName *string `json:"azure_blob_storage_endpoint_domain_name,omitempty"`
	// Shared access signature (SAS) token to grant limited access to objects in your storage account.
	AzureBlobStorageSasToken string                                                            `json:"azure_blob_storage_sas_token"`
	DataSourceType           DestinationDatabricksDataSourceAzureBlobStorageDataSourceTypeEnum `json:"data_source_type"`
}

type DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum string

const (
	DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnumS3Storage DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum = "S3_STORAGE"
)

func (e DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum) ToPointer() *DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum {
	return &e
}

func (e *DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3_STORAGE":
		*e = DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum: %v", v)
	}
}

// DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum - The region of the S3 staging bucket to use if utilising a copy strategy.
type DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum string

const (
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumUnknown      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = ""
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumUsEast1      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "us-east-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumUsEast2      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "us-east-2"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumUsWest1      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "us-west-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumUsWest2      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "us-west-2"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumAfSouth1     DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "af-south-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumApEast1      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ap-east-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumApSouth1     DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ap-south-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumApNortheast1 DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ap-northeast-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumApNortheast2 DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ap-northeast-2"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumApNortheast3 DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ap-northeast-3"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumApSoutheast1 DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ap-southeast-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumApSoutheast2 DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ap-southeast-2"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumCaCentral1   DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "ca-central-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumCnNorth1     DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "cn-north-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumCnNorthwest1 DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "cn-northwest-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumEuCentral1   DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "eu-central-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumEuNorth1     DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "eu-north-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumEuSouth1     DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "eu-south-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumEuWest1      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "eu-west-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumEuWest2      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "eu-west-2"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumEuWest3      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "eu-west-3"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumSaEast1      DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "sa-east-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumMeSouth1     DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "me-south-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumUsGovEast1   DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "us-gov-east-1"
	DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnumUsGovWest1   DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum = "us-gov-west-1"
)

func (e DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum) ToPointer() *DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum {
	return &e
}

func (e *DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum) UnmarshalJSON(data []byte) error {
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
		*e = DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum: %v", v)
	}
}

// DestinationDatabricksDataSourceAmazonS3 - Storage on which the delta lake is built.
type DestinationDatabricksDataSourceAmazonS3 struct {
	DataSourceType DestinationDatabricksDataSourceAmazonS3DataSourceTypeEnum `json:"data_source_type"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string `json:"file_name_pattern,omitempty"`
	// The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.
	S3AccessKeyID string `json:"s3_access_key_id"`
	// The name of the S3 bucket to use for intermittent staging of the data.
	S3BucketName string `json:"s3_bucket_name"`
	// The directory under the S3 bucket where data will be written.
	S3BucketPath string `json:"s3_bucket_path"`
	// The region of the S3 staging bucket to use if utilising a copy strategy.
	S3BucketRegion DestinationDatabricksDataSourceAmazonS3S3BucketRegionEnum `json:"s3_bucket_region"`
	// The corresponding secret to the above access key id.
	S3SecretAccessKey string `json:"s3_secret_access_key"`
}

type DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum string

const (
	DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnumManagedTablesStorage DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum = "MANAGED_TABLES_STORAGE"
)

func (e DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum) ToPointer() *DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum {
	return &e
}

func (e *DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANAGED_TABLES_STORAGE":
		*e = DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum: %v", v)
	}
}

// DestinationDatabricksDataSourceRecommendedManagedTables - Storage on which the delta lake is built.
type DestinationDatabricksDataSourceRecommendedManagedTables struct {
	DataSourceType DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceTypeEnum `json:"data_source_type"`
}

type DestinationDatabricksDataSourceType string

const (
	DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceRecommendedManagedTables DestinationDatabricksDataSourceType = "destination-databricks_Data Source_[Recommended] Managed tables"
	DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceAmazonS3                 DestinationDatabricksDataSourceType = "destination-databricks_Data Source_Amazon S3"
	DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceAzureBlobStorage         DestinationDatabricksDataSourceType = "destination-databricks_Data Source_Azure Blob Storage"
)

type DestinationDatabricksDataSource struct {
	DestinationDatabricksDataSourceRecommendedManagedTables *DestinationDatabricksDataSourceRecommendedManagedTables
	DestinationDatabricksDataSourceAmazonS3                 *DestinationDatabricksDataSourceAmazonS3
	DestinationDatabricksDataSourceAzureBlobStorage         *DestinationDatabricksDataSourceAzureBlobStorage

	Type DestinationDatabricksDataSourceType
}

func CreateDestinationDatabricksDataSourceDestinationDatabricksDataSourceRecommendedManagedTables(destinationDatabricksDataSourceRecommendedManagedTables DestinationDatabricksDataSourceRecommendedManagedTables) DestinationDatabricksDataSource {
	typ := DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceRecommendedManagedTables

	return DestinationDatabricksDataSource{
		DestinationDatabricksDataSourceRecommendedManagedTables: &destinationDatabricksDataSourceRecommendedManagedTables,
		Type: typ,
	}
}

func CreateDestinationDatabricksDataSourceDestinationDatabricksDataSourceAmazonS3(destinationDatabricksDataSourceAmazonS3 DestinationDatabricksDataSourceAmazonS3) DestinationDatabricksDataSource {
	typ := DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceAmazonS3

	return DestinationDatabricksDataSource{
		DestinationDatabricksDataSourceAmazonS3: &destinationDatabricksDataSourceAmazonS3,
		Type:                                    typ,
	}
}

func CreateDestinationDatabricksDataSourceDestinationDatabricksDataSourceAzureBlobStorage(destinationDatabricksDataSourceAzureBlobStorage DestinationDatabricksDataSourceAzureBlobStorage) DestinationDatabricksDataSource {
	typ := DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceAzureBlobStorage

	return DestinationDatabricksDataSource{
		DestinationDatabricksDataSourceAzureBlobStorage: &destinationDatabricksDataSourceAzureBlobStorage,
		Type: typ,
	}
}

func (u *DestinationDatabricksDataSource) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationDatabricksDataSourceRecommendedManagedTables := new(DestinationDatabricksDataSourceRecommendedManagedTables)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationDatabricksDataSourceRecommendedManagedTables); err == nil {
		u.DestinationDatabricksDataSourceRecommendedManagedTables = destinationDatabricksDataSourceRecommendedManagedTables
		u.Type = DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceRecommendedManagedTables
		return nil
	}

	destinationDatabricksDataSourceAmazonS3 := new(DestinationDatabricksDataSourceAmazonS3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationDatabricksDataSourceAmazonS3); err == nil {
		u.DestinationDatabricksDataSourceAmazonS3 = destinationDatabricksDataSourceAmazonS3
		u.Type = DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceAmazonS3
		return nil
	}

	destinationDatabricksDataSourceAzureBlobStorage := new(DestinationDatabricksDataSourceAzureBlobStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationDatabricksDataSourceAzureBlobStorage); err == nil {
		u.DestinationDatabricksDataSourceAzureBlobStorage = destinationDatabricksDataSourceAzureBlobStorage
		u.Type = DestinationDatabricksDataSourceTypeDestinationDatabricksDataSourceAzureBlobStorage
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationDatabricksDataSource) MarshalJSON() ([]byte, error) {
	if u.DestinationDatabricksDataSourceRecommendedManagedTables != nil {
		return json.Marshal(u.DestinationDatabricksDataSourceRecommendedManagedTables)
	}

	if u.DestinationDatabricksDataSourceAmazonS3 != nil {
		return json.Marshal(u.DestinationDatabricksDataSourceAmazonS3)
	}

	if u.DestinationDatabricksDataSourceAzureBlobStorage != nil {
		return json.Marshal(u.DestinationDatabricksDataSourceAzureBlobStorage)
	}

	return nil, nil
}

type DestinationDatabricksDatabricksEnum string

const (
	DestinationDatabricksDatabricksEnumDatabricks DestinationDatabricksDatabricksEnum = "databricks"
)

func (e DestinationDatabricksDatabricksEnum) ToPointer() *DestinationDatabricksDatabricksEnum {
	return &e
}

func (e *DestinationDatabricksDatabricksEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "databricks":
		*e = DestinationDatabricksDatabricksEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksDatabricksEnum: %v", v)
	}
}

type DestinationDatabricks struct {
	// You must agree to the Databricks JDBC Driver <a href="https://databricks.com/jdbc-odbc-driver-license">Terms & Conditions</a> to use this connector.
	AcceptTerms bool `json:"accept_terms"`
	// Storage on which the delta lake is built.
	DataSource DestinationDatabricksDataSource `json:"data_source"`
	// The name of the catalog. If not specified otherwise, the "hive_metastore" will be used.
	Database *string `json:"database,omitempty"`
	// Databricks Cluster HTTP Path.
	DatabricksHTTPPath string `json:"databricks_http_path"`
	// Databricks Personal Access Token for making authenticated requests.
	DatabricksPersonalAccessToken string `json:"databricks_personal_access_token"`
	// Databricks Cluster Port.
	DatabricksPort *string `json:"databricks_port,omitempty"`
	// Databricks Cluster Server Hostname.
	DatabricksServerHostname string                              `json:"databricks_server_hostname"`
	DestinationType          DestinationDatabricksDatabricksEnum `json:"destinationType"`
	// Default to 'true'. Switch it to 'false' for debugging purpose.
	PurgeStagingData *bool `json:"purge_staging_data,omitempty"`
	// The default schema tables are written. If not specified otherwise, the "default" will be used.
	Schema *string `json:"schema,omitempty"`
}
