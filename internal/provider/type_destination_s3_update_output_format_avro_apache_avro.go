// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationS3UpdateOutputFormatAvroApacheAvro struct {
	CompressionCodec DestinationS3UpdateOutputFormatAvroApacheAvroCompressionCodec `tfsdk:"compression_codec"`
	FormatType       types.String                                                  `tfsdk:"format_type"`
}