// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType string

const (
	DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceTypeAzureBlobStorage DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType = "AZURE_BLOB_STORAGE"
)

func (e DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType) ToPointer() *DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType {
	return &e
}

func (e *DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AZURE_BLOB_STORAGE":
		*e = DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType: %v", v)
	}
}

// DestinationDatabricksUpdateDataSourceAzureBlobStorage - Storage on which the delta lake is built.
type DestinationDatabricksUpdateDataSourceAzureBlobStorage struct {
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The name of the Azure blob storage container.
	AzureBlobStorageContainerName string `json:"azure_blob_storage_container_name"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpointDomainName *string `json:"azure_blob_storage_endpoint_domain_name,omitempty"`
	// Shared access signature (SAS) token to grant limited access to objects in your storage account.
	AzureBlobStorageSasToken string                                                              `json:"azure_blob_storage_sas_token"`
	DataSourceType           DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType `json:"data_source_type"`
}

type DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType string

const (
	DestinationDatabricksUpdateDataSourceAmazonS3DataSourceTypeS3Storage DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType = "S3_STORAGE"
)

func (e DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType) ToPointer() *DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType {
	return &e
}

func (e *DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3_STORAGE":
		*e = DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType: %v", v)
	}
}

// DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion - The region of the S3 staging bucket to use if utilising a copy strategy.
type DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion string

const (
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionUnknown      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = ""
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionUsEast1      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "us-east-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionUsEast2      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "us-east-2"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionUsWest1      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "us-west-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionUsWest2      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "us-west-2"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionAfSouth1     DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "af-south-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionApEast1      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ap-east-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionApSouth1     DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ap-south-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionApNortheast1 DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ap-northeast-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionApNortheast2 DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ap-northeast-2"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionApNortheast3 DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ap-northeast-3"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionApSoutheast1 DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ap-southeast-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionApSoutheast2 DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ap-southeast-2"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionCaCentral1   DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "ca-central-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionCnNorth1     DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "cn-north-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionCnNorthwest1 DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "cn-northwest-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionEuCentral1   DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "eu-central-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionEuNorth1     DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "eu-north-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionEuSouth1     DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "eu-south-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionEuWest1      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "eu-west-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionEuWest2      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "eu-west-2"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionEuWest3      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "eu-west-3"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionSaEast1      DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "sa-east-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionMeSouth1     DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "me-south-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionUsGovEast1   DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "us-gov-east-1"
	DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegionUsGovWest1   DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion = "us-gov-west-1"
)

func (e DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion) ToPointer() *DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion {
	return &e
}

func (e *DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion) UnmarshalJSON(data []byte) error {
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
		*e = DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion: %v", v)
	}
}

// DestinationDatabricksUpdateDataSourceAmazonS3 - Storage on which the delta lake is built.
type DestinationDatabricksUpdateDataSourceAmazonS3 struct {
	DataSourceType DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType `json:"data_source_type"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string `json:"file_name_pattern,omitempty"`
	// The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.
	S3AccessKeyID string `json:"s3_access_key_id"`
	// The name of the S3 bucket to use for intermittent staging of the data.
	S3BucketName string `json:"s3_bucket_name"`
	// The directory under the S3 bucket where data will be written.
	S3BucketPath string `json:"s3_bucket_path"`
	// The region of the S3 staging bucket to use if utilising a copy strategy.
	S3BucketRegion DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion `json:"s3_bucket_region"`
	// The corresponding secret to the above access key id.
	S3SecretAccessKey string `json:"s3_secret_access_key"`
}

type DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType string

const (
	DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceTypeManagedTablesStorage DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType = "MANAGED_TABLES_STORAGE"
)

func (e DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType) ToPointer() *DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType {
	return &e
}

func (e *DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANAGED_TABLES_STORAGE":
		*e = DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType: %v", v)
	}
}

// DestinationDatabricksUpdateDataSourceRecommendedManagedTables - Storage on which the delta lake is built.
type DestinationDatabricksUpdateDataSourceRecommendedManagedTables struct {
	DataSourceType DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType `json:"data_source_type"`
}

type DestinationDatabricksUpdateDataSourceType string

const (
	DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceRecommendedManagedTables DestinationDatabricksUpdateDataSourceType = "destination-databricks-update_Data Source_[Recommended] Managed tables"
	DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceAmazonS3                 DestinationDatabricksUpdateDataSourceType = "destination-databricks-update_Data Source_Amazon S3"
	DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceAzureBlobStorage         DestinationDatabricksUpdateDataSourceType = "destination-databricks-update_Data Source_Azure Blob Storage"
)

type DestinationDatabricksUpdateDataSource struct {
	DestinationDatabricksUpdateDataSourceRecommendedManagedTables *DestinationDatabricksUpdateDataSourceRecommendedManagedTables
	DestinationDatabricksUpdateDataSourceAmazonS3                 *DestinationDatabricksUpdateDataSourceAmazonS3
	DestinationDatabricksUpdateDataSourceAzureBlobStorage         *DestinationDatabricksUpdateDataSourceAzureBlobStorage

	Type DestinationDatabricksUpdateDataSourceType
}

func CreateDestinationDatabricksUpdateDataSourceDestinationDatabricksUpdateDataSourceRecommendedManagedTables(destinationDatabricksUpdateDataSourceRecommendedManagedTables DestinationDatabricksUpdateDataSourceRecommendedManagedTables) DestinationDatabricksUpdateDataSource {
	typ := DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceRecommendedManagedTables

	return DestinationDatabricksUpdateDataSource{
		DestinationDatabricksUpdateDataSourceRecommendedManagedTables: &destinationDatabricksUpdateDataSourceRecommendedManagedTables,
		Type: typ,
	}
}

func CreateDestinationDatabricksUpdateDataSourceDestinationDatabricksUpdateDataSourceAmazonS3(destinationDatabricksUpdateDataSourceAmazonS3 DestinationDatabricksUpdateDataSourceAmazonS3) DestinationDatabricksUpdateDataSource {
	typ := DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceAmazonS3

	return DestinationDatabricksUpdateDataSource{
		DestinationDatabricksUpdateDataSourceAmazonS3: &destinationDatabricksUpdateDataSourceAmazonS3,
		Type: typ,
	}
}

func CreateDestinationDatabricksUpdateDataSourceDestinationDatabricksUpdateDataSourceAzureBlobStorage(destinationDatabricksUpdateDataSourceAzureBlobStorage DestinationDatabricksUpdateDataSourceAzureBlobStorage) DestinationDatabricksUpdateDataSource {
	typ := DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceAzureBlobStorage

	return DestinationDatabricksUpdateDataSource{
		DestinationDatabricksUpdateDataSourceAzureBlobStorage: &destinationDatabricksUpdateDataSourceAzureBlobStorage,
		Type: typ,
	}
}

func (u *DestinationDatabricksUpdateDataSource) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationDatabricksUpdateDataSourceRecommendedManagedTables := new(DestinationDatabricksUpdateDataSourceRecommendedManagedTables)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationDatabricksUpdateDataSourceRecommendedManagedTables); err == nil {
		u.DestinationDatabricksUpdateDataSourceRecommendedManagedTables = destinationDatabricksUpdateDataSourceRecommendedManagedTables
		u.Type = DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceRecommendedManagedTables
		return nil
	}

	destinationDatabricksUpdateDataSourceAmazonS3 := new(DestinationDatabricksUpdateDataSourceAmazonS3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationDatabricksUpdateDataSourceAmazonS3); err == nil {
		u.DestinationDatabricksUpdateDataSourceAmazonS3 = destinationDatabricksUpdateDataSourceAmazonS3
		u.Type = DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceAmazonS3
		return nil
	}

	destinationDatabricksUpdateDataSourceAzureBlobStorage := new(DestinationDatabricksUpdateDataSourceAzureBlobStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationDatabricksUpdateDataSourceAzureBlobStorage); err == nil {
		u.DestinationDatabricksUpdateDataSourceAzureBlobStorage = destinationDatabricksUpdateDataSourceAzureBlobStorage
		u.Type = DestinationDatabricksUpdateDataSourceTypeDestinationDatabricksUpdateDataSourceAzureBlobStorage
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationDatabricksUpdateDataSource) MarshalJSON() ([]byte, error) {
	if u.DestinationDatabricksUpdateDataSourceRecommendedManagedTables != nil {
		return json.Marshal(u.DestinationDatabricksUpdateDataSourceRecommendedManagedTables)
	}

	if u.DestinationDatabricksUpdateDataSourceAmazonS3 != nil {
		return json.Marshal(u.DestinationDatabricksUpdateDataSourceAmazonS3)
	}

	if u.DestinationDatabricksUpdateDataSourceAzureBlobStorage != nil {
		return json.Marshal(u.DestinationDatabricksUpdateDataSourceAzureBlobStorage)
	}

	return nil, nil
}

type DestinationDatabricksUpdate struct {
	// You must agree to the Databricks JDBC Driver <a href="https://databricks.com/jdbc-odbc-driver-license">Terms & Conditions</a> to use this connector.
	AcceptTerms bool `json:"accept_terms"`
	// Storage on which the delta lake is built.
	DataSource DestinationDatabricksUpdateDataSource `json:"data_source"`
	// The name of the catalog. If not specified otherwise, the "hive_metastore" will be used.
	Database *string `json:"database,omitempty"`
	// Databricks Cluster HTTP Path.
	DatabricksHTTPPath string `json:"databricks_http_path"`
	// Databricks Personal Access Token for making authenticated requests.
	DatabricksPersonalAccessToken string `json:"databricks_personal_access_token"`
	// Databricks Cluster Port.
	DatabricksPort *string `json:"databricks_port,omitempty"`
	// Databricks Cluster Server Hostname.
	DatabricksServerHostname string `json:"databricks_server_hostname"`
	// Default to 'true'. Switch it to 'false' for debugging purpose.
	PurgeStagingData *bool `json:"purge_staging_data,omitempty"`
	// The default schema tables are written. If not specified otherwise, the "default" will be used.
	Schema *string `json:"schema,omitempty"`
}