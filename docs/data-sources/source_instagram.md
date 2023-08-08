---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_instagram Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceInstagram DataSource
---

# airbyte_source_instagram (Data Source)

SourceInstagram DataSource

## Example Usage

```terraform
data "airbyte_source_instagram" "my_source_instagram" {
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

- `access_token` (String) The value of the access token generated with <b>instagram_basic, instagram_manage_insights, pages_show_list, pages_read_engagement, Instagram Public Content Access</b> permissions. See the <a href="https://docs.airbyte.com/integrations/sources/instagram/#step-1-set-up-instagram">docs</a> for more information
- `source_type` (String) must be one of ["instagram"]
- `start_date` (String) The date from which you'd like to replicate data for User Insights, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.

