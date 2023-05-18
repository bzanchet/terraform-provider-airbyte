// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceTodoistResourceModel) ToCreateSDKType() *shared.SourceTodoistCreateRequest {
	sourceType := shared.SourceTodoistTodoistEnum(r.Configuration.SourceType.ValueString())
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceTodoist{
		SourceType: sourceType,
		Token:      token,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTodoistCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTodoistResourceModel) ToDeleteSDKType() *shared.SourceTodoistCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
