---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_zoom Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceZoom Resource
---

# airbyte_source_zoom (Resource)

SourceZoom Resource



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

- `jwt_token` (String)
- `source_type` (String)

