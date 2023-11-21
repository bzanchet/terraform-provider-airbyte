---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_file Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceFile Resource
---

# airbyte_source_file (Resource)

SourceFile Resource

## Example Usage

```terraform
resource "airbyte_source_file" "my_source_file" {
  configuration = {
    dataset_name = "...my_dataset_name..."
    format       = "jsonl"
    provider = {
      az_blob_azure_blob_storage = {
        sas_token       = "...my_sas_token..."
        shared_key      = "...my_shared_key..."
        storage_account = "...my_storage_account..."
      }
    }
    reader_options = "{\"sep\": \"\t\", \"header\": 0, \"names\": [\"column1\", \"column2\"] }"
    url            = "https://storage.googleapis.com/covid19-open-data/v2/latest/epidemiology.csv"
  }
  definition_id = "6c5d5cf5-0fbf-4713-864e-d5bf6d67306c"
  name          = "Floyd Goyette"
  secret_id     = "...my_secret_id..."
  workspace_id  = "68cfaeff-480d-4f14-bee1-0f8279e427b2"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the source e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `dataset_name` (String) The Name of the final table to replicate this file into (should include letters, numbers dash and underscores only).
- `provider` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider))
- `url` (String) The URL path to access the file which should be replicated.

Optional:

- `format` (String) must be one of ["csv", "json", "jsonl", "excel", "excel_binary", "feather", "parquet", "yaml"]; Default: "csv"
The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
- `reader_options` (String) This should be a string in JSON format. It depends on the chosen file format to provide additional options and tune its behavior.

<a id="nestedatt--configuration--provider"></a>
### Nested Schema for `configuration.provider`

Optional:

- `az_blob_azure_blob_storage` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--az_blob_azure_blob_storage))
- `gcs_google_cloud_storage` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--gcs_google_cloud_storage))
- `https_public_web` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--https_public_web))
- `s3_amazon_web_services` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--s3_amazon_web_services))
- `scp_secure_copy_protocol` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--scp_secure_copy_protocol))
- `sftp_secure_file_transfer_protocol` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--sftp_secure_file_transfer_protocol))
- `ssh_secure_shell` (Attributes) The storage Provider or Location of the file(s) which should be replicated. (see [below for nested schema](#nestedatt--configuration--provider--ssh_secure_shell))

<a id="nestedatt--configuration--provider--az_blob_azure_blob_storage"></a>
### Nested Schema for `configuration.provider.az_blob_azure_blob_storage`

Required:

- `storage_account` (String) The globally unique name of the storage account that the desired blob sits within. See <a href="https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview" target="_blank">here</a> for more details.

Optional:

- `sas_token` (String, Sensitive) To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a SAS (Shared Access Signature) token. If accessing publicly available data, this field is not necessary.
- `shared_key` (String, Sensitive) To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a storage account shared key (aka account key or access key). If accessing publicly available data, this field is not necessary.


<a id="nestedatt--configuration--provider--gcs_google_cloud_storage"></a>
### Nested Schema for `configuration.provider.gcs_google_cloud_storage`

Optional:

- `service_account_json` (String) In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described <a href="https://cloud.google.com/iam/docs/service-accounts" target="_blank">here</a>. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.


<a id="nestedatt--configuration--provider--https_public_web"></a>
### Nested Schema for `configuration.provider.https_public_web`

Optional:

- `user_agent` (Boolean) Default: false
Add User-Agent to request


<a id="nestedatt--configuration--provider--s3_amazon_web_services"></a>
### Nested Schema for `configuration.provider.s3_amazon_web_services`

Optional:

- `aws_access_key_id` (String, Sensitive) In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `aws_secret_access_key` (String, Sensitive) In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.


<a id="nestedatt--configuration--provider--scp_secure_copy_protocol"></a>
### Nested Schema for `configuration.provider.scp_secure_copy_protocol`

Required:

- `host` (String)
- `user` (String)

Optional:

- `password` (String, Sensitive)
- `port` (String) Default: "22"


<a id="nestedatt--configuration--provider--sftp_secure_file_transfer_protocol"></a>
### Nested Schema for `configuration.provider.sftp_secure_file_transfer_protocol`

Required:

- `host` (String)
- `user` (String)

Optional:

- `password` (String, Sensitive)
- `port` (String) Default: "22"


<a id="nestedatt--configuration--provider--ssh_secure_shell"></a>
### Nested Schema for `configuration.provider.ssh_secure_shell`

Required:

- `host` (String)
- `user` (String)

Optional:

- `password` (String, Sensitive)
- `port` (String) Default: "22"

