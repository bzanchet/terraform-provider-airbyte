// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationPineconeCreateRequest struct {
	Configuration DestinationPinecone `json:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationPineconeCreateRequest) GetConfiguration() DestinationPinecone {
	if o == nil {
		return DestinationPinecone{}
	}
	return o.Configuration
}

func (o *DestinationPineconeCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationPineconeCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationPineconeCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}