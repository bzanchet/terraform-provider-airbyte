// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourcePinterestUpdate struct {
	Credentials *SourcePinterestUpdateAuthorizationMethod `tfsdk:"credentials"`
	StartDate   types.String                              `tfsdk:"start_date"`
	Status      []types.String                            `tfsdk:"status"`
	SourceType  types.String                              `tfsdk:"source_type"`
}
