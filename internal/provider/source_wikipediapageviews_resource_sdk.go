// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceWikipediaPageviewsResourceModel) ToCreateSDKType() *shared.SourceWikipediaPageviewsCreateRequest {
	access := r.Configuration.Access.ValueString()
	agent := r.Configuration.Agent.ValueString()
	article := r.Configuration.Article.ValueString()
	country := r.Configuration.Country.ValueString()
	end := r.Configuration.End.ValueString()
	project := r.Configuration.Project.ValueString()
	sourceType := shared.SourceWikipediaPageviewsWikipediaPageviews(r.Configuration.SourceType.ValueString())
	start := r.Configuration.Start.ValueString()
	configuration := shared.SourceWikipediaPageviews{
		Access:     access,
		Agent:      agent,
		Article:    article,
		Country:    country,
		End:        end,
		Project:    project,
		SourceType: sourceType,
		Start:      start,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceWikipediaPageviewsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceWikipediaPageviewsResourceModel) ToUpdateSDKType() *shared.SourceWikipediaPageviewsPutRequest {
	access := r.Configuration.Access.ValueString()
	agent := r.Configuration.Agent.ValueString()
	article := r.Configuration.Article.ValueString()
	country := r.Configuration.Country.ValueString()
	end := r.Configuration.End.ValueString()
	project := r.Configuration.Project.ValueString()
	start := r.Configuration.Start.ValueString()
	configuration := shared.SourceWikipediaPageviewsUpdate{
		Access:  access,
		Agent:   agent,
		Article: article,
		Country: country,
		End:     end,
		Project: project,
		Start:   start,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceWikipediaPageviewsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceWikipediaPageviewsResourceModel) ToDeleteSDKType() *shared.SourceWikipediaPageviewsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceWikipediaPageviewsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
