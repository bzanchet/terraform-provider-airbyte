// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type DestinationPostgresUpdateSSHTunnelMethod struct {
	DestinationPostgresUpdateSSHTunnelMethodNoTunnel               *DestinationClickhouseUpdateSSHTunnelMethodNoTunnel               `tfsdk:"destination_postgres_update_ssh_tunnel_method_no_tunnel"`
	DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationClickhouseUpdateSSHTunnelMethodSSHKeyAuthentication   `tfsdk:"destination_postgres_update_ssh_tunnel_method_ssh_key_authentication"`
	DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication *DestinationClickhouseUpdateSSHTunnelMethodPasswordAuthentication `tfsdk:"destination_postgres_update_ssh_tunnel_method_password_authentication"`
	DestinationPostgresSSHTunnelMethodNoTunnel                     *DestinationClickhouseSSHTunnelMethodNoTunnel                     `tfsdk:"destination_postgres_ssh_tunnel_method_no_tunnel"`
	DestinationPostgresSSHTunnelMethodSSHKeyAuthentication         *DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication         `tfsdk:"destination_postgres_ssh_tunnel_method_ssh_key_authentication"`
	DestinationPostgresSSHTunnelMethodPasswordAuthentication       *DestinationClickhouseSSHTunnelMethodPasswordAuthentication       `tfsdk:"destination_postgres_ssh_tunnel_method_password_authentication"`
}
