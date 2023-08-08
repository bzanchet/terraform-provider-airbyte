---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_google_sheets Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationGoogleSheets DataSource
---

# airbyte_destination_google_sheets (Data Source)

DestinationGoogleSheets DataSource

## Example Usage

```terraform
data "airbyte_destination_google_sheets" "my_destination_googlesheets" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `credentials` (Attributes) Google API Credentials for connecting to Google Sheets and Google Drive APIs (see [below for nested schema](#nestedatt--configuration--credentials))
- `destination_type` (String) must be one of ["google-sheets"]
- `spreadsheet_id` (String) The link to your spreadsheet. See <a href='https://docs.airbyte.com/integrations/destinations/google-sheets#sheetlink'>this guide</a> for more details.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `client_id` (String) The Client ID of your Google Sheets developer application.
- `client_secret` (String) The Client Secret of your Google Sheets developer application.
- `refresh_token` (String) The token for obtaining new access token.

