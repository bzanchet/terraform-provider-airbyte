// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAlloydbResourceModel) ToCreateSDKType() *shared.SourceAlloydbCreateRequest {
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
	var replicationMethod *shared.SourceAlloydbReplicationMethod
	var sourceAlloydbReplicationMethodStandard *shared.SourceAlloydbReplicationMethodStandard
	if r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodStandard != nil {
		method := shared.SourceAlloydbReplicationMethodStandardMethod(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodStandard.Method.ValueString())
		sourceAlloydbReplicationMethodStandard = &shared.SourceAlloydbReplicationMethodStandard{
			Method: method,
		}
	}
	if sourceAlloydbReplicationMethodStandard != nil {
		replicationMethod = &shared.SourceAlloydbReplicationMethod{
			SourceAlloydbReplicationMethodStandard: sourceAlloydbReplicationMethodStandard,
		}
	}
	var sourceAlloydbReplicationMethodLogicalReplicationCDC *shared.SourceAlloydbReplicationMethodLogicalReplicationCDC
	if r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC != nil {
		initialWaitingSeconds := new(int64)
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.InitialWaitingSeconds.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.InitialWaitingSeconds.IsNull() {
			*initialWaitingSeconds = r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.InitialWaitingSeconds.ValueInt64()
		} else {
			initialWaitingSeconds = nil
		}
		lsnCommitBehaviour := new(shared.SourceAlloydbReplicationMethodLogicalReplicationCDCLSNCommitBehaviour)
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.LsnCommitBehaviour.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.LsnCommitBehaviour.IsNull() {
			*lsnCommitBehaviour = shared.SourceAlloydbReplicationMethodLogicalReplicationCDCLSNCommitBehaviour(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.LsnCommitBehaviour.ValueString())
		} else {
			lsnCommitBehaviour = nil
		}
		method1 := shared.SourceAlloydbReplicationMethodLogicalReplicationCDCMethod(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Method.ValueString())
		plugin := new(shared.SourceAlloydbReplicationMethodLogicalReplicationCDCPlugin)
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Plugin.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Plugin.IsNull() {
			*plugin = shared.SourceAlloydbReplicationMethodLogicalReplicationCDCPlugin(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Plugin.ValueString())
		} else {
			plugin = nil
		}
		publication := r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Publication.ValueString()
		replicationSlot := r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.ReplicationSlot.ValueString()
		var additionalProperties interface{}
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.AdditionalProperties.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.AdditionalProperties.ValueString()), &additionalProperties)
		}
		sourceAlloydbReplicationMethodLogicalReplicationCDC = &shared.SourceAlloydbReplicationMethodLogicalReplicationCDC{
			InitialWaitingSeconds: initialWaitingSeconds,
			LsnCommitBehaviour:    lsnCommitBehaviour,
			Method:                method1,
			Plugin:                plugin,
			Publication:           publication,
			ReplicationSlot:       replicationSlot,
			AdditionalProperties:  additionalProperties,
		}
	}
	if sourceAlloydbReplicationMethodLogicalReplicationCDC != nil {
		replicationMethod = &shared.SourceAlloydbReplicationMethod{
			SourceAlloydbReplicationMethodLogicalReplicationCDC: sourceAlloydbReplicationMethodLogicalReplicationCDC,
		}
	}
	schemas := make([]string, 0)
	for _, schemasItem := range r.Configuration.Schemas {
		schemas = append(schemas, schemasItem.ValueString())
	}
	sourceType := shared.SourceAlloydbAlloydb(r.Configuration.SourceType.ValueString())
	var sslMode *shared.SourceAlloydbSSLModes
	var sourceAlloydbSSLModesDisable *shared.SourceAlloydbSSLModesDisable
	if r.Configuration.SslMode.SourceAlloydbSSLModesDisable != nil {
		mode := shared.SourceAlloydbSSLModesDisableMode(r.Configuration.SslMode.SourceAlloydbSSLModesDisable.Mode.ValueString())
		var additionalProperties1 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesDisable.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesDisable.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesDisable.AdditionalProperties.ValueString()), &additionalProperties1)
		}
		sourceAlloydbSSLModesDisable = &shared.SourceAlloydbSSLModesDisable{
			Mode:                 mode,
			AdditionalProperties: additionalProperties1,
		}
	}
	if sourceAlloydbSSLModesDisable != nil {
		sslMode = &shared.SourceAlloydbSSLModes{
			SourceAlloydbSSLModesDisable: sourceAlloydbSSLModesDisable,
		}
	}
	var sourceAlloydbSSLModesAllow *shared.SourceAlloydbSSLModesAllow
	if r.Configuration.SslMode.SourceAlloydbSSLModesAllow != nil {
		mode1 := shared.SourceAlloydbSSLModesAllowMode(r.Configuration.SslMode.SourceAlloydbSSLModesAllow.Mode.ValueString())
		var additionalProperties2 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesAllow.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesAllow.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesAllow.AdditionalProperties.ValueString()), &additionalProperties2)
		}
		sourceAlloydbSSLModesAllow = &shared.SourceAlloydbSSLModesAllow{
			Mode:                 mode1,
			AdditionalProperties: additionalProperties2,
		}
	}
	if sourceAlloydbSSLModesAllow != nil {
		sslMode = &shared.SourceAlloydbSSLModes{
			SourceAlloydbSSLModesAllow: sourceAlloydbSSLModesAllow,
		}
	}
	var sourceAlloydbSSLModesPrefer *shared.SourceAlloydbSSLModesPrefer
	if r.Configuration.SslMode.SourceAlloydbSSLModesPrefer != nil {
		mode2 := shared.SourceAlloydbSSLModesPreferMode(r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.Mode.ValueString())
		var additionalProperties3 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.AdditionalProperties.ValueString()), &additionalProperties3)
		}
		sourceAlloydbSSLModesPrefer = &shared.SourceAlloydbSSLModesPrefer{
			Mode:                 mode2,
			AdditionalProperties: additionalProperties3,
		}
	}
	if sourceAlloydbSSLModesPrefer != nil {
		sslMode = &shared.SourceAlloydbSSLModes{
			SourceAlloydbSSLModesPrefer: sourceAlloydbSSLModesPrefer,
		}
	}
	var sourceAlloydbSSLModesRequire *shared.SourceAlloydbSSLModesRequire
	if r.Configuration.SslMode.SourceAlloydbSSLModesRequire != nil {
		mode3 := shared.SourceAlloydbSSLModesRequireMode(r.Configuration.SslMode.SourceAlloydbSSLModesRequire.Mode.ValueString())
		var additionalProperties4 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesRequire.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesRequire.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesRequire.AdditionalProperties.ValueString()), &additionalProperties4)
		}
		sourceAlloydbSSLModesRequire = &shared.SourceAlloydbSSLModesRequire{
			Mode:                 mode3,
			AdditionalProperties: additionalProperties4,
		}
	}
	if sourceAlloydbSSLModesRequire != nil {
		sslMode = &shared.SourceAlloydbSSLModes{
			SourceAlloydbSSLModesRequire: sourceAlloydbSSLModesRequire,
		}
	}
	var sourceAlloydbSSLModesVerifyCa *shared.SourceAlloydbSSLModesVerifyCa
	if r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa != nil {
		caCertificate := r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.CaCertificate.ValueString()
		clientCertificate := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientCertificate.IsNull() {
			*clientCertificate = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientCertificate.ValueString()
		} else {
			clientCertificate = nil
		}
		clientKey := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKey.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKey.IsNull() {
			*clientKey = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKey.ValueString()
		} else {
			clientKey = nil
		}
		clientKeyPassword := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKeyPassword.IsNull() {
			*clientKeyPassword = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword = nil
		}
		mode4 := shared.SourceAlloydbSSLModesVerifyCaMode(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.Mode.ValueString())
		var additionalProperties5 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.AdditionalProperties.ValueString()), &additionalProperties5)
		}
		sourceAlloydbSSLModesVerifyCa = &shared.SourceAlloydbSSLModesVerifyCa{
			CaCertificate:        caCertificate,
			ClientCertificate:    clientCertificate,
			ClientKey:            clientKey,
			ClientKeyPassword:    clientKeyPassword,
			Mode:                 mode4,
			AdditionalProperties: additionalProperties5,
		}
	}
	if sourceAlloydbSSLModesVerifyCa != nil {
		sslMode = &shared.SourceAlloydbSSLModes{
			SourceAlloydbSSLModesVerifyCa: sourceAlloydbSSLModesVerifyCa,
		}
	}
	var sourceAlloydbSSLModesVerifyFull *shared.SourceAlloydbSSLModesVerifyFull
	if r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull != nil {
		caCertificate1 := r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.CaCertificate.ValueString()
		clientCertificate1 := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientCertificate.IsNull() {
			*clientCertificate1 = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientCertificate.ValueString()
		} else {
			clientCertificate1 = nil
		}
		clientKey1 := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKey.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKey.IsNull() {
			*clientKey1 = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKey.ValueString()
		} else {
			clientKey1 = nil
		}
		clientKeyPassword1 := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKeyPassword.IsNull() {
			*clientKeyPassword1 = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword1 = nil
		}
		mode5 := shared.SourceAlloydbSSLModesVerifyFullMode(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.Mode.ValueString())
		var additionalProperties6 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.AdditionalProperties.ValueString()), &additionalProperties6)
		}
		sourceAlloydbSSLModesVerifyFull = &shared.SourceAlloydbSSLModesVerifyFull{
			CaCertificate:        caCertificate1,
			ClientCertificate:    clientCertificate1,
			ClientKey:            clientKey1,
			ClientKeyPassword:    clientKeyPassword1,
			Mode:                 mode5,
			AdditionalProperties: additionalProperties6,
		}
	}
	if sourceAlloydbSSLModesVerifyFull != nil {
		sslMode = &shared.SourceAlloydbSSLModes{
			SourceAlloydbSSLModesVerifyFull: sourceAlloydbSSLModesVerifyFull,
		}
	}
	var tunnelMethod *shared.SourceAlloydbSSHTunnelMethod
	var sourceAlloydbSSHTunnelMethodNoTunnel *shared.SourceAlloydbSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.SourceAlloydbSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		sourceAlloydbSSHTunnelMethodNoTunnel = &shared.SourceAlloydbSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if sourceAlloydbSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.SourceAlloydbSSHTunnelMethod{
			SourceAlloydbSSHTunnelMethodNoTunnel: sourceAlloydbSSHTunnelMethodNoTunnel,
		}
	}
	var sourceAlloydbSSHTunnelMethodSSHKeyAuthentication *shared.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication != nil {
		sshKey := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
		tunnelHost := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.SourceAlloydbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		sourceAlloydbSSHTunnelMethodSSHKeyAuthentication = &shared.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication{
			SSHKey:       sshKey,
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if sourceAlloydbSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.SourceAlloydbSSHTunnelMethod{
			SourceAlloydbSSHTunnelMethodSSHKeyAuthentication: sourceAlloydbSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var sourceAlloydbSSHTunnelMethodPasswordAuthentication *shared.SourceAlloydbSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.SourceAlloydbSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		tunnelUserPassword := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
		sourceAlloydbSSHTunnelMethodPasswordAuthentication = &shared.SourceAlloydbSSHTunnelMethodPasswordAuthentication{
			TunnelHost:         tunnelHost1,
			TunnelMethod:       tunnelMethod3,
			TunnelPort:         tunnelPort1,
			TunnelUser:         tunnelUser1,
			TunnelUserPassword: tunnelUserPassword,
		}
	}
	if sourceAlloydbSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.SourceAlloydbSSHTunnelMethod{
			SourceAlloydbSSHTunnelMethodPasswordAuthentication: sourceAlloydbSSHTunnelMethodPasswordAuthentication,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceAlloydb{
		Database:          database,
		Host:              host,
		JdbcURLParams:     jdbcURLParams,
		Password:          password,
		Port:              port,
		ReplicationMethod: replicationMethod,
		Schemas:           schemas,
		SourceType:        sourceType,
		SslMode:           sslMode,
		TunnelMethod:      tunnelMethod,
		Username:          username,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAlloydbCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAlloydbResourceModel) ToUpdateSDKType() *shared.SourceAlloydbPutRequest {
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
	var replicationMethod *shared.SourceAlloydbUpdateReplicationMethod
	var sourceAlloydbUpdateReplicationMethodStandard *shared.SourceAlloydbUpdateReplicationMethodStandard
	if r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodStandard != nil {
		method := shared.SourceAlloydbUpdateReplicationMethodStandardMethod(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodStandard.Method.ValueString())
		sourceAlloydbUpdateReplicationMethodStandard = &shared.SourceAlloydbUpdateReplicationMethodStandard{
			Method: method,
		}
	}
	if sourceAlloydbUpdateReplicationMethodStandard != nil {
		replicationMethod = &shared.SourceAlloydbUpdateReplicationMethod{
			SourceAlloydbUpdateReplicationMethodStandard: sourceAlloydbUpdateReplicationMethodStandard,
		}
	}
	var sourceAlloydbUpdateReplicationMethodLogicalReplicationCDC *shared.SourceAlloydbUpdateReplicationMethodLogicalReplicationCDC
	if r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC != nil {
		initialWaitingSeconds := new(int64)
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.InitialWaitingSeconds.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.InitialWaitingSeconds.IsNull() {
			*initialWaitingSeconds = r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.InitialWaitingSeconds.ValueInt64()
		} else {
			initialWaitingSeconds = nil
		}
		lsnCommitBehaviour := new(shared.SourceAlloydbUpdateReplicationMethodLogicalReplicationCDCLSNCommitBehaviour)
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.LsnCommitBehaviour.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.LsnCommitBehaviour.IsNull() {
			*lsnCommitBehaviour = shared.SourceAlloydbUpdateReplicationMethodLogicalReplicationCDCLSNCommitBehaviour(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.LsnCommitBehaviour.ValueString())
		} else {
			lsnCommitBehaviour = nil
		}
		method1 := shared.SourceAlloydbUpdateReplicationMethodLogicalReplicationCDCMethod(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Method.ValueString())
		plugin := new(shared.SourceAlloydbUpdateReplicationMethodLogicalReplicationCDCPlugin)
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Plugin.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Plugin.IsNull() {
			*plugin = shared.SourceAlloydbUpdateReplicationMethodLogicalReplicationCDCPlugin(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Plugin.ValueString())
		} else {
			plugin = nil
		}
		publication := r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.Publication.ValueString()
		replicationSlot := r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.ReplicationSlot.ValueString()
		var additionalProperties interface{}
		if !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.AdditionalProperties.IsUnknown() && !r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.ReplicationMethod.SourceAlloydbReplicationMethodLogicalReplicationCDC.AdditionalProperties.ValueString()), &additionalProperties)
		}
		sourceAlloydbUpdateReplicationMethodLogicalReplicationCDC = &shared.SourceAlloydbUpdateReplicationMethodLogicalReplicationCDC{
			InitialWaitingSeconds: initialWaitingSeconds,
			LsnCommitBehaviour:    lsnCommitBehaviour,
			Method:                method1,
			Plugin:                plugin,
			Publication:           publication,
			ReplicationSlot:       replicationSlot,
			AdditionalProperties:  additionalProperties,
		}
	}
	if sourceAlloydbUpdateReplicationMethodLogicalReplicationCDC != nil {
		replicationMethod = &shared.SourceAlloydbUpdateReplicationMethod{
			SourceAlloydbUpdateReplicationMethodLogicalReplicationCDC: sourceAlloydbUpdateReplicationMethodLogicalReplicationCDC,
		}
	}
	schemas := make([]string, 0)
	for _, schemasItem := range r.Configuration.Schemas {
		schemas = append(schemas, schemasItem.ValueString())
	}
	var sslMode *shared.SourceAlloydbUpdateSSLModes
	var sourceAlloydbUpdateSSLModesDisable *shared.SourceAlloydbUpdateSSLModesDisable
	if r.Configuration.SslMode.SourceAlloydbSSLModesDisable != nil {
		mode := shared.SourceAlloydbUpdateSSLModesDisableMode(r.Configuration.SslMode.SourceAlloydbSSLModesDisable.Mode.ValueString())
		var additionalProperties1 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesDisable.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesDisable.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesDisable.AdditionalProperties.ValueString()), &additionalProperties1)
		}
		sourceAlloydbUpdateSSLModesDisable = &shared.SourceAlloydbUpdateSSLModesDisable{
			Mode:                 mode,
			AdditionalProperties: additionalProperties1,
		}
	}
	if sourceAlloydbUpdateSSLModesDisable != nil {
		sslMode = &shared.SourceAlloydbUpdateSSLModes{
			SourceAlloydbUpdateSSLModesDisable: sourceAlloydbUpdateSSLModesDisable,
		}
	}
	var sourceAlloydbUpdateSSLModesAllow *shared.SourceAlloydbUpdateSSLModesAllow
	if r.Configuration.SslMode.SourceAlloydbSSLModesAllow != nil {
		mode1 := shared.SourceAlloydbUpdateSSLModesAllowMode(r.Configuration.SslMode.SourceAlloydbSSLModesAllow.Mode.ValueString())
		var additionalProperties2 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesAllow.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesAllow.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesAllow.AdditionalProperties.ValueString()), &additionalProperties2)
		}
		sourceAlloydbUpdateSSLModesAllow = &shared.SourceAlloydbUpdateSSLModesAllow{
			Mode:                 mode1,
			AdditionalProperties: additionalProperties2,
		}
	}
	if sourceAlloydbUpdateSSLModesAllow != nil {
		sslMode = &shared.SourceAlloydbUpdateSSLModes{
			SourceAlloydbUpdateSSLModesAllow: sourceAlloydbUpdateSSLModesAllow,
		}
	}
	var sourceAlloydbUpdateSSLModesPrefer *shared.SourceAlloydbUpdateSSLModesPrefer
	if r.Configuration.SslMode.SourceAlloydbSSLModesPrefer != nil {
		mode2 := shared.SourceAlloydbUpdateSSLModesPreferMode(r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.Mode.ValueString())
		var additionalProperties3 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesPrefer.AdditionalProperties.ValueString()), &additionalProperties3)
		}
		sourceAlloydbUpdateSSLModesPrefer = &shared.SourceAlloydbUpdateSSLModesPrefer{
			Mode:                 mode2,
			AdditionalProperties: additionalProperties3,
		}
	}
	if sourceAlloydbUpdateSSLModesPrefer != nil {
		sslMode = &shared.SourceAlloydbUpdateSSLModes{
			SourceAlloydbUpdateSSLModesPrefer: sourceAlloydbUpdateSSLModesPrefer,
		}
	}
	var sourceAlloydbUpdateSSLModesRequire *shared.SourceAlloydbUpdateSSLModesRequire
	if r.Configuration.SslMode.SourceAlloydbSSLModesRequire != nil {
		mode3 := shared.SourceAlloydbUpdateSSLModesRequireMode(r.Configuration.SslMode.SourceAlloydbSSLModesRequire.Mode.ValueString())
		var additionalProperties4 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesRequire.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesRequire.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesRequire.AdditionalProperties.ValueString()), &additionalProperties4)
		}
		sourceAlloydbUpdateSSLModesRequire = &shared.SourceAlloydbUpdateSSLModesRequire{
			Mode:                 mode3,
			AdditionalProperties: additionalProperties4,
		}
	}
	if sourceAlloydbUpdateSSLModesRequire != nil {
		sslMode = &shared.SourceAlloydbUpdateSSLModes{
			SourceAlloydbUpdateSSLModesRequire: sourceAlloydbUpdateSSLModesRequire,
		}
	}
	var sourceAlloydbUpdateSSLModesVerifyCa *shared.SourceAlloydbUpdateSSLModesVerifyCa
	if r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa != nil {
		caCertificate := r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.CaCertificate.ValueString()
		clientCertificate := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientCertificate.IsNull() {
			*clientCertificate = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientCertificate.ValueString()
		} else {
			clientCertificate = nil
		}
		clientKey := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKey.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKey.IsNull() {
			*clientKey = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKey.ValueString()
		} else {
			clientKey = nil
		}
		clientKeyPassword := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKeyPassword.IsNull() {
			*clientKeyPassword = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword = nil
		}
		mode4 := shared.SourceAlloydbUpdateSSLModesVerifyCaMode(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.Mode.ValueString())
		var additionalProperties5 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyCa.AdditionalProperties.ValueString()), &additionalProperties5)
		}
		sourceAlloydbUpdateSSLModesVerifyCa = &shared.SourceAlloydbUpdateSSLModesVerifyCa{
			CaCertificate:        caCertificate,
			ClientCertificate:    clientCertificate,
			ClientKey:            clientKey,
			ClientKeyPassword:    clientKeyPassword,
			Mode:                 mode4,
			AdditionalProperties: additionalProperties5,
		}
	}
	if sourceAlloydbUpdateSSLModesVerifyCa != nil {
		sslMode = &shared.SourceAlloydbUpdateSSLModes{
			SourceAlloydbUpdateSSLModesVerifyCa: sourceAlloydbUpdateSSLModesVerifyCa,
		}
	}
	var sourceAlloydbUpdateSSLModesVerifyFull *shared.SourceAlloydbUpdateSSLModesVerifyFull
	if r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull != nil {
		caCertificate1 := r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.CaCertificate.ValueString()
		clientCertificate1 := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientCertificate.IsNull() {
			*clientCertificate1 = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientCertificate.ValueString()
		} else {
			clientCertificate1 = nil
		}
		clientKey1 := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKey.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKey.IsNull() {
			*clientKey1 = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKey.ValueString()
		} else {
			clientKey1 = nil
		}
		clientKeyPassword1 := new(string)
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKeyPassword.IsNull() {
			*clientKeyPassword1 = r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword1 = nil
		}
		mode5 := shared.SourceAlloydbUpdateSSLModesVerifyFullMode(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.Mode.ValueString())
		var additionalProperties6 interface{}
		if !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.AdditionalProperties.IsUnknown() && !r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.Configuration.SslMode.SourceAlloydbSSLModesVerifyFull.AdditionalProperties.ValueString()), &additionalProperties6)
		}
		sourceAlloydbUpdateSSLModesVerifyFull = &shared.SourceAlloydbUpdateSSLModesVerifyFull{
			CaCertificate:        caCertificate1,
			ClientCertificate:    clientCertificate1,
			ClientKey:            clientKey1,
			ClientKeyPassword:    clientKeyPassword1,
			Mode:                 mode5,
			AdditionalProperties: additionalProperties6,
		}
	}
	if sourceAlloydbUpdateSSLModesVerifyFull != nil {
		sslMode = &shared.SourceAlloydbUpdateSSLModes{
			SourceAlloydbUpdateSSLModesVerifyFull: sourceAlloydbUpdateSSLModesVerifyFull,
		}
	}
	var tunnelMethod *shared.SourceAlloydbUpdateSSHTunnelMethod
	var sourceAlloydbUpdateSSHTunnelMethodNoTunnel *shared.SourceAlloydbUpdateSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.SourceAlloydbUpdateSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		sourceAlloydbUpdateSSHTunnelMethodNoTunnel = &shared.SourceAlloydbUpdateSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if sourceAlloydbUpdateSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.SourceAlloydbUpdateSSHTunnelMethod{
			SourceAlloydbUpdateSSHTunnelMethodNoTunnel: sourceAlloydbUpdateSSHTunnelMethodNoTunnel,
		}
	}
	var sourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication *shared.SourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication != nil {
		sshKey := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
		tunnelHost := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.SourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		sourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication = &shared.SourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication{
			SSHKey:       sshKey,
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if sourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.SourceAlloydbUpdateSSHTunnelMethod{
			SourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication: sourceAlloydbUpdateSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var sourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication *shared.SourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.SourceAlloydbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		tunnelUserPassword := r.Configuration.TunnelMethod.SourceAlloydbSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
		sourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication = &shared.SourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication{
			TunnelHost:         tunnelHost1,
			TunnelMethod:       tunnelMethod3,
			TunnelPort:         tunnelPort1,
			TunnelUser:         tunnelUser1,
			TunnelUserPassword: tunnelUserPassword,
		}
	}
	if sourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.SourceAlloydbUpdateSSHTunnelMethod{
			SourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication: sourceAlloydbUpdateSSHTunnelMethodPasswordAuthentication,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceAlloydbUpdate{
		Database:          database,
		Host:              host,
		JdbcURLParams:     jdbcURLParams,
		Password:          password,
		Port:              port,
		ReplicationMethod: replicationMethod,
		Schemas:           schemas,
		SslMode:           sslMode,
		TunnelMethod:      tunnelMethod,
		Username:          username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAlloydbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAlloydbResourceModel) ToDeleteSDKType() *shared.SourceAlloydbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAlloydbResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
