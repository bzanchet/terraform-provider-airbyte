---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_amazon_seller_partner Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAmazonSellerPartner DataSource
---

# airbyte_source_amazon_seller_partner (Data Source)

SourceAmazonSellerPartner DataSource

## Example Usage

```terraform
data "airbyte_source_amazon_seller_partner" "my_source_amazonsellerpartner" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `advanced_stream_options` (String) Additional information to configure report options. This varies by report type, not every report implement this kind of feature. Must be a valid json string.
- `auth_type` (String) must be one of ["oauth2.0"]
- `aws_access_key` (String) Specifies the AWS access key used as part of the credentials to authenticate the user.
- `aws_environment` (String) must be one of ["PRODUCTION", "SANDBOX"]
Select the AWS Environment.
- `aws_secret_key` (String) Specifies the AWS secret key used as part of the credentials to authenticate the user.
- `lwa_app_id` (String) Your Login with Amazon Client ID.
- `lwa_client_secret` (String) Your Login with Amazon Client Secret.
- `max_wait_seconds` (Number) Sometimes report can take up to 30 minutes to generate. This will set the limit for how long to wait for a successful report.
- `period_in_days` (Number) Will be used for stream slicing for initial full_refresh sync when no updated state is present for reports that support sliced incremental sync.
- `refresh_token` (String) The Refresh Token obtained via OAuth flow authorization.
- `region` (String) must be one of ["AE", "AU", "BE", "BR", "CA", "DE", "EG", "ES", "FR", "GB", "IN", "IT", "JP", "MX", "NL", "PL", "SA", "SE", "SG", "TR", "UK", "US"]
Select the AWS Region.
- `replication_end_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data after this date will not be replicated.
- `replication_start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
- `report_options` (String) Additional information passed to reports. This varies by report type. Must be a valid json string.
- `role_arn` (String) Specifies the Amazon Resource Name (ARN) of an IAM role that you want to use to perform operations requested using this profile. (Needs permission to 'Assume Role' STS).
- `source_type` (String) must be one of ["amazon-seller-partner"]

