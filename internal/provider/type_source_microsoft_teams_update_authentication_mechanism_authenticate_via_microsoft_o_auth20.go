// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 struct {
	AuthType     types.String `tfsdk:"auth_type"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	RefreshToken types.String `tfsdk:"refresh_token"`
	TenantID     types.String `tfsdk:"tenant_id"`
}