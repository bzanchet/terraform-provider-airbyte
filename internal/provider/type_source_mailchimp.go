// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMailchimp struct {
	CampaignID  types.String                   `tfsdk:"campaign_id"`
	Credentials *SourceMailchimpAuthentication `tfsdk:"credentials"`
	StartDate   types.String                   `tfsdk:"start_date"`
}
