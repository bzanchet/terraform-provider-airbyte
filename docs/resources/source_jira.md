---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_jira Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceJira Resource
---

# airbyte_source_jira (Resource)

SourceJira Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Optional

- `secret_id` (String)

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `api_token` (String)
- `domain` (String)
- `email` (String)
- `source_type` (String)

Optional:

- `enable_experimental_streams` (Boolean)
- `expand_issue_changelog` (Boolean)
- `projects` (List of String)
- `render_fields` (Boolean)
- `start_date` (String)

