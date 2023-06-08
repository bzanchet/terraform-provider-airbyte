// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationRedshiftResourceModel) ToCreateSDKType() *shared.DestinationRedshiftCreateRequest {
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationRedshiftRedshift(r.Configuration.DestinationType.ValueString())
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := r.Configuration.Password.ValueString()
	port := r.Configuration.Port.ValueInt64()
	schema := r.Configuration.Schema.ValueString()
	var tunnelMethod *shared.DestinationRedshiftSSHTunnelMethod
	var destinationRedshiftSSHTunnelMethodNoTunnel *shared.DestinationRedshiftSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationRedshiftSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationRedshiftSSHTunnelMethodNoTunnel = &shared.DestinationRedshiftSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationRedshiftSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationRedshiftSSHTunnelMethod{
			DestinationRedshiftSSHTunnelMethodNoTunnel: destinationRedshiftSSHTunnelMethodNoTunnel,
		}
	}
	var destinationRedshiftSSHTunnelMethodSSHKeyAuthentication *shared.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationRedshiftSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		destinationRedshiftSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication{
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationRedshiftSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftSSHTunnelMethod{
			DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication: destinationRedshiftSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationRedshiftSSHTunnelMethodPasswordAuthentication *shared.DestinationRedshiftSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationRedshiftSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationRedshiftSSHTunnelMethodPasswordAuthentication = &shared.DestinationRedshiftSSHTunnelMethodPasswordAuthentication{
			TunnelHost:   tunnelHost1,
			TunnelMethod: tunnelMethod3,
			TunnelPort:   tunnelPort1,
			TunnelUser:   tunnelUser1,
		}
	}
	if destinationRedshiftSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftSSHTunnelMethod{
			DestinationRedshiftSSHTunnelMethodPasswordAuthentication: destinationRedshiftSSHTunnelMethodPasswordAuthentication,
		}
	}
	var uploadingMethod *shared.DestinationRedshiftUploadingMethod
	var destinationRedshiftUploadingMethodStandard *shared.DestinationRedshiftUploadingMethodStandard
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging != nil {
		method := shared.DestinationRedshiftUploadingMethodStandardMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Method.ValueString())
		destinationRedshiftUploadingMethodStandard = &shared.DestinationRedshiftUploadingMethodStandard{
			Method: method,
		}
	}
	if destinationRedshiftUploadingMethodStandard != nil {
		uploadingMethod = &shared.DestinationRedshiftUploadingMethod{
			DestinationRedshiftUploadingMethodStandard: destinationRedshiftUploadingMethodStandard,
		}
	}
	var destinationRedshiftUploadingMethodS3Staging *shared.DestinationRedshiftUploadingMethodS3Staging
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard != nil {
		method1 := shared.DestinationRedshiftUploadingMethodS3StagingMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard.Method.ValueString())
		destinationRedshiftUploadingMethodS3Staging = &shared.DestinationRedshiftUploadingMethodS3Staging{
			Method: method1,
		}
	}
	if destinationRedshiftUploadingMethodS3Staging != nil {
		uploadingMethod = &shared.DestinationRedshiftUploadingMethod{
			DestinationRedshiftUploadingMethodS3Staging: destinationRedshiftUploadingMethodS3Staging,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationRedshift{
		Database:        database,
		DestinationType: destinationType,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		Password:        password,
		Port:            port,
		Schema:          schema,
		TunnelMethod:    tunnelMethod,
		UploadingMethod: uploadingMethod,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationRedshiftCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationRedshiftResourceModel) ToUpdateSDKType() *shared.DestinationRedshiftPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := r.Configuration.Password.ValueString()
	port := r.Configuration.Port.ValueInt64()
	schema := r.Configuration.Schema.ValueString()
	var tunnelMethod *shared.DestinationRedshiftUpdateSSHTunnelMethod
	var destinationRedshiftUpdateSSHTunnelMethodNoTunnel *shared.DestinationRedshiftUpdateSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationRedshiftUpdateSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationRedshiftUpdateSSHTunnelMethodNoTunnel = &shared.DestinationRedshiftUpdateSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationRedshiftUpdateSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationRedshiftUpdateSSHTunnelMethod{
			DestinationRedshiftUpdateSSHTunnelMethodNoTunnel: destinationRedshiftUpdateSSHTunnelMethodNoTunnel,
		}
	}
	var destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication *shared.DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication{
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftUpdateSSHTunnelMethod{
			DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication: destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication *shared.DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication = &shared.DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication{
			TunnelHost:   tunnelHost1,
			TunnelMethod: tunnelMethod3,
			TunnelPort:   tunnelPort1,
			TunnelUser:   tunnelUser1,
		}
	}
	if destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftUpdateSSHTunnelMethod{
			DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication: destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication,
		}
	}
	var uploadingMethod *shared.DestinationRedshiftUpdateUploadingMethod
	var destinationRedshiftUpdateUploadingMethodStandard *shared.DestinationRedshiftUpdateUploadingMethodStandard
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging != nil {
		method := shared.DestinationRedshiftUpdateUploadingMethodStandardMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Method.ValueString())
		destinationRedshiftUpdateUploadingMethodStandard = &shared.DestinationRedshiftUpdateUploadingMethodStandard{
			Method: method,
		}
	}
	if destinationRedshiftUpdateUploadingMethodStandard != nil {
		uploadingMethod = &shared.DestinationRedshiftUpdateUploadingMethod{
			DestinationRedshiftUpdateUploadingMethodStandard: destinationRedshiftUpdateUploadingMethodStandard,
		}
	}
	var destinationRedshiftUpdateUploadingMethodS3Staging *shared.DestinationRedshiftUpdateUploadingMethodS3Staging
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard != nil {
		method1 := shared.DestinationRedshiftUpdateUploadingMethodS3StagingMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard.Method.ValueString())
		destinationRedshiftUpdateUploadingMethodS3Staging = &shared.DestinationRedshiftUpdateUploadingMethodS3Staging{
			Method: method1,
		}
	}
	if destinationRedshiftUpdateUploadingMethodS3Staging != nil {
		uploadingMethod = &shared.DestinationRedshiftUpdateUploadingMethod{
			DestinationRedshiftUpdateUploadingMethodS3Staging: destinationRedshiftUpdateUploadingMethodS3Staging,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationRedshiftUpdate{
		Database:        database,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		Password:        password,
		Port:            port,
		Schema:          schema,
		TunnelMethod:    tunnelMethod,
		UploadingMethod: uploadingMethod,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationRedshiftPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationRedshiftResourceModel) ToDeleteSDKType() *shared.DestinationRedshiftCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationRedshiftResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}