---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_datadog Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceDatadog DataSource
---

# airbyte_source_datadog (Data Source)

SourceDatadog DataSource

## Example Usage

```terraform
data "airbyte_source_datadog" "my_source_datadog" {
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

- `api_key` (String) Datadog API key
- `application_key` (String) Datadog application key
- `end_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Data after this date will  not be replicated. An empty value will represent the current datetime for each  execution. This just applies to Incremental syncs.
- `max_records_per_request` (Number) Maximum number of records to collect per request.
- `queries` (Attributes List) List of queries to be run and used as inputs. (see [below for nested schema](#nestedatt--configuration--queries))
- `query` (String) The search query. This just applies to Incremental syncs. If empty, it'll collect all logs.
- `site` (String) must be one of ["datadoghq.com", "us3.datadoghq.com", "us5.datadoghq.com", "datadoghq.eu", "ddog-gov.com"]
The site where Datadog data resides in.
- `source_type` (String) must be one of ["datadog"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. This just applies to Incremental syncs.

<a id="nestedatt--configuration--queries"></a>
### Nested Schema for `configuration.queries`

Read-Only:

- `data_source` (String) must be one of ["metrics", "cloud_cost", "logs", "rum"]
A data source that is powered by the platform.
- `name` (String) The variable name for use in queries.
- `query` (String) A classic query string.

