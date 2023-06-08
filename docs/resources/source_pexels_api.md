---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_pexels_api Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePexelsAPI Resource
---

# airbyte_source_pexels_api (Resource)

SourcePexelsAPI Resource



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

- `api_key` (String)
- `query` (String)
- `source_type` (String)

Optional:

- `color` (String)
- `locale` (String)
- `orientation` (String)
- `size` (String)

