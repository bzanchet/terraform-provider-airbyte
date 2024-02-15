// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationClickhouseResourceModel) ToSharedDestinationClickhouseCreateRequest() *shared.DestinationClickhouseCreateRequest {
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
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	var tunnelMethod *shared.DestinationClickhouseSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationClickhouseNoTunnel *shared.DestinationClickhouseNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationClickhouseNoTunnel = &shared.DestinationClickhouseNoTunnel{}
		}
		if destinationClickhouseNoTunnel != nil {
			tunnelMethod = &shared.DestinationClickhouseSSHTunnelMethod{
				DestinationClickhouseNoTunnel: destinationClickhouseNoTunnel,
			}
		}
		var destinationClickhouseSSHKeyAuthentication *shared.DestinationClickhouseSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationClickhouseSSHKeyAuthentication = &shared.DestinationClickhouseSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationClickhouseSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationClickhouseSSHTunnelMethod{
				DestinationClickhouseSSHKeyAuthentication: destinationClickhouseSSHKeyAuthentication,
			}
		}
		var destinationClickhousePasswordAuthentication *shared.DestinationClickhousePasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationClickhousePasswordAuthentication = &shared.DestinationClickhousePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationClickhousePasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationClickhouseSSHTunnelMethod{
				DestinationClickhousePasswordAuthentication: destinationClickhousePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationClickhouse{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationClickhouseCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationClickhouseResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationClickhouseResourceModel) ToSharedDestinationClickhousePutRequest() *shared.DestinationClickhousePutRequest {
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
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	var tunnelMethod *shared.SSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var noTunnel *shared.NoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			noTunnel = &shared.NoTunnel{}
		}
		if noTunnel != nil {
			tunnelMethod = &shared.SSHTunnelMethod{
				NoTunnel: noTunnel,
			}
		}
		var sshKeyAuthentication *shared.SSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			sshKeyAuthentication = &shared.SSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if sshKeyAuthentication != nil {
			tunnelMethod = &shared.SSHTunnelMethod{
				SSHKeyAuthentication: sshKeyAuthentication,
			}
		}
		var passwordAuthentication *shared.PasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			passwordAuthentication = &shared.PasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if passwordAuthentication != nil {
			tunnelMethod = &shared.SSHTunnelMethod{
				PasswordAuthentication: passwordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationClickhouseUpdate{
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
	out := shared.DestinationClickhousePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
