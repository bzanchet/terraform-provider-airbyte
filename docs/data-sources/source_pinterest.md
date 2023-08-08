---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_pinterest Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePinterest DataSource
---

# airbyte_source_pinterest (Data Source)

SourcePinterest DataSource

## Example Usage

```terraform
data "airbyte_source_pinterest" "my_source_pinterest" {
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
- `source_type` (String) must be one of ["pinterest"]
- `start_date` (String) A date in the format YYYY-MM-DD. If you have not set a date, it would be defaulted to latest allowed date by api (89 days from today).
- `status` (List of String) Entity statuses based off of campaigns, ad_groups, and ads. If you do not have a status set, it will be ignored completely.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_pinterest_authorization_method_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_pinterest_authorization_method_access_token))
- `source_pinterest_authorization_method_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_pinterest_authorization_method_o_auth2_0))
- `source_pinterest_update_authorization_method_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_pinterest_update_authorization_method_access_token))
- `source_pinterest_update_authorization_method_o_auth2_0` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_pinterest_update_authorization_method_o_auth2_0))

<a id="nestedatt--configuration--credentials--source_pinterest_authorization_method_access_token"></a>
### Nested Schema for `configuration.credentials.source_pinterest_authorization_method_access_token`

Read-Only:

- `access_token` (String) The Access Token to make authenticated requests.
- `auth_method` (String) must be one of ["access_token"]


<a id="nestedatt--configuration--credentials--source_pinterest_authorization_method_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_pinterest_authorization_method_o_auth2_0`

Read-Only:

- `auth_method` (String) must be one of ["oauth2.0"]
- `client_id` (String) The Client ID of your OAuth application
- `client_secret` (String) The Client Secret of your OAuth application.
- `refresh_token` (String) Refresh Token to obtain new Access Token, when it's expired.


<a id="nestedatt--configuration--credentials--source_pinterest_update_authorization_method_access_token"></a>
### Nested Schema for `configuration.credentials.source_pinterest_update_authorization_method_access_token`

Read-Only:

- `access_token` (String) The Access Token to make authenticated requests.
- `auth_method` (String) must be one of ["access_token"]


<a id="nestedatt--configuration--credentials--source_pinterest_update_authorization_method_o_auth2_0"></a>
### Nested Schema for `configuration.credentials.source_pinterest_update_authorization_method_o_auth2_0`

Read-Only:

- `auth_method` (String) must be one of ["oauth2.0"]
- `client_id` (String) The Client ID of your OAuth application
- `client_secret` (String) The Client Secret of your OAuth application.
- `refresh_token` (String) Refresh Token to obtain new Access Token, when it's expired.

