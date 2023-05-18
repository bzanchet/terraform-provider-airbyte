// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGithub struct {
	Branch                  types.String                `tfsdk:"branch"`
	Credentials             *SourceGithubAuthentication `tfsdk:"credentials"`
	PageSizeForLargeStreams types.Int64                 `tfsdk:"page_size_for_large_streams"`
	Repository              types.String                `tfsdk:"repository"`
	RequestsPerHour         types.Int64                 `tfsdk:"requests_per_hour"`
	SourceType              types.String                `tfsdk:"source_type"`
	StartDate               types.String                `tfsdk:"start_date"`
}
