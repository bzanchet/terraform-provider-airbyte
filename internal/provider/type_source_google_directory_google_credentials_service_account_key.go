// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleDirectoryGoogleCredentialsServiceAccountKey struct {
	CredentialsJSON  types.String `tfsdk:"credentials_json"`
	CredentialsTitle types.String `tfsdk:"credentials_title"`
	Email            types.String `tfsdk:"email"`
}