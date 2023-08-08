---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_smartsheets Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSmartsheets DataSource
---

# airbyte_source_smartsheets (Data Source)

SourceSmartsheets DataSource

## Example Usage

```terraform
data "airbyte_source_smartsheets" "my_source_smartsheets" {
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

- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `source_type` (String) must be one of ["smartsheets"]
- `spreadsheet_id` (String) The spreadsheet ID. Find it by opening the spreadsheet then navigating to File > Properties
- `start_datetime` (String) Only rows modified after this date/time will be replicated. This should be an ISO 8601 string, for instance: `2000-01-01T13:00:00`

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_smartsheets_authorization_method_api_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_smartsheets_authorization_method_api_access_token))
- `source_smartsheets_authorization_method_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_smartsheets_authorization_method_o_auth2_0))
- `source_smartsheets_update_authorization_method_api_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_smartsheets_update_authorization_method_api_access_token))
- `source_smartsheets_update_authorization_method_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_smartsheets_update_authorization_method_o_auth2_0))

<a id="nestedatt--configuration--credentials--source_smartsheets_authorization_method_api_access_token"></a>
### Nested Schema for `configuration.credentials.source_smartsheets_authorization_method_api_access_token`

Read-Only:

- `access_token` (String) The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account > Apps & Integrations > API Access. See the <a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide">setup guide</a> for information on how to obtain this token.
- `auth_type` (String) must be one of ["access_token"]


<a id="nestedatt--configuration--credentials--source_smartsheets_authorization_method_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_smartsheets_authorization_method_o_auth2_0`

Read-Only:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_type` (String) must be one of ["oauth2.0"]
- `client_id` (String) The API ID of the SmartSheets developer application.
- `client_secret` (String) The API Secret the SmartSheets developer application.
- `refresh_token` (String) The key to refresh the expired access_token.
- `token_expiry_date` (String) The date-time when the access token should be refreshed.


<a id="nestedatt--configuration--credentials--source_smartsheets_update_authorization_method_api_access_token"></a>
### Nested Schema for `configuration.credentials.source_smartsheets_update_authorization_method_api_access_token`

Read-Only:

- `access_token` (String) The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account > Apps & Integrations > API Access. See the <a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide">setup guide</a> for information on how to obtain this token.
- `auth_type` (String) must be one of ["access_token"]


<a id="nestedatt--configuration--credentials--source_smartsheets_update_authorization_method_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_smartsheets_update_authorization_method_o_auth2_0`

Read-Only:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_type` (String) must be one of ["oauth2.0"]
- `client_id` (String) The API ID of the SmartSheets developer application.
- `client_secret` (String) The API Secret the SmartSheets developer application.
- `refresh_token` (String) The key to refresh the expired access_token.
- `token_expiry_date` (String) The date-time when the access token should be refreshed.

