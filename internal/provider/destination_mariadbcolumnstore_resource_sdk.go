// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationMariadbColumnstoreResourceModel) ToCreateSDKType() *shared.DestinationMariadbColumnstoreCreateRequest {
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationMariadbColumnstoreMariadbColumnstore(r.Configuration.DestinationType.ValueString())
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := r.Configuration.Port.ValueInt64()
	var tunnelMethod *shared.DestinationMariadbColumnstoreSSHTunnelMethod
	var destinationMariadbColumnstoreSSHTunnelMethodNoTunnel *shared.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationMariadbColumnstoreSSHTunnelMethodNoTunnel = &shared.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationMariadbColumnstoreSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationMariadbColumnstoreSSHTunnelMethod{
			DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel: destinationMariadbColumnstoreSSHTunnelMethodNoTunnel,
		}
	}
	var destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication *shared.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication != nil {
		sshKey := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
		tunnelHost := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication{
			SSHKey:       sshKey,
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationMariadbColumnstoreSSHTunnelMethod{
			DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication: destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication *shared.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		tunnelUserPassword := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
		destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication = &shared.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication{
			TunnelHost:         tunnelHost1,
			TunnelMethod:       tunnelMethod3,
			TunnelPort:         tunnelPort1,
			TunnelUser:         tunnelUser1,
			TunnelUserPassword: tunnelUserPassword,
		}
	}
	if destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationMariadbColumnstoreSSHTunnelMethod{
			DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication: destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMariadbColumnstore{
		Database:        database,
		DestinationType: destinationType,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		Password:        password,
		Port:            port,
		TunnelMethod:    tunnelMethod,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMariadbColumnstoreCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMariadbColumnstoreResourceModel) ToUpdateSDKType() *shared.DestinationMariadbColumnstorePutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := r.Configuration.Port.ValueInt64()
	var tunnelMethod *shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethod
	var destinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnel *shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnel = &shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethod{
			DestinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnel: destinationMariadbColumnstoreUpdateSSHTunnelMethodNoTunnel,
		}
	}
	var destinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthentication *shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication != nil {
		sshKey := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
		tunnelHost := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthentication{
			SSHKey:       sshKey,
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethod{
			DestinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthentication: destinationMariadbColumnstoreUpdateSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthentication *shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		tunnelUserPassword := r.Configuration.TunnelMethod.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
		destinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthentication = &shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthentication{
			TunnelHost:         tunnelHost1,
			TunnelMethod:       tunnelMethod3,
			TunnelPort:         tunnelPort1,
			TunnelUser:         tunnelUser1,
			TunnelUserPassword: tunnelUserPassword,
		}
	}
	if destinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationMariadbColumnstoreUpdateSSHTunnelMethod{
			DestinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthentication: destinationMariadbColumnstoreUpdateSSHTunnelMethodPasswordAuthentication,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMariadbColumnstoreUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMariadbColumnstorePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMariadbColumnstoreResourceModel) ToDeleteSDKType() *shared.DestinationMariadbColumnstoreCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMariadbColumnstoreResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
