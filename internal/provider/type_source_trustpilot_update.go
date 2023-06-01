// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceTrustpilotUpdate struct {
	BusinessUnits []types.String                            `tfsdk:"business_units"`
	Credentials   SourceTrustpilotUpdateAuthorizationMethod `tfsdk:"credentials"`
	StartDate     types.String                              `tfsdk:"start_date"`
	SourceType    types.String                              `tfsdk:"source_type"`
}
