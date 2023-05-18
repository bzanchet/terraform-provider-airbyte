// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceRedshiftResourceModel) ToCreateSDKType() *shared.SourceRedshiftCreateRequest {
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
	schemas := make([]string, 0)
	for _, schemasItem := range r.Configuration.Schemas {
		schemas = append(schemas, schemasItem.ValueString())
	}
	sourceType := shared.SourceRedshiftRedshiftEnum(r.Configuration.SourceType.ValueString())
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceRedshift{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schemas:       schemas,
		SourceType:    sourceType,
		Username:      username,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRedshiftCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRedshiftResourceModel) ToDeleteSDKType() *shared.SourceRedshiftCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
