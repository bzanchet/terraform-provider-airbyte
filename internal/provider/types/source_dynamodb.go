// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceDynamodb struct {
	Credentials                        *SourceDynamodbCredentials `tfsdk:"credentials"`
	Endpoint                           types.String               `tfsdk:"endpoint"`
	IgnoreMissingReadPermissionsTables types.Bool                 `tfsdk:"ignore_missing_read_permissions_tables"`
	Region                             types.String               `tfsdk:"region"`
	ReservedAttributeNames             types.String               `tfsdk:"reserved_attribute_names"`
}