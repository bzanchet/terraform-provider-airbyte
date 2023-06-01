// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage struct {
	SasToken       types.String `tfsdk:"sas_token"`
	SharedKey      types.String `tfsdk:"shared_key"`
	Storage        types.String `tfsdk:"storage"`
	StorageAccount types.String `tfsdk:"storage_account"`
}
