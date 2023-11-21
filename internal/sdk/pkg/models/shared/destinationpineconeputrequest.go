// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationPineconePutRequest struct {
	Configuration DestinationPineconeUpdate `json:"configuration"`
	Name          string                    `json:"name"`
	WorkspaceID   string                    `json:"workspaceId"`
}

func (o *DestinationPineconePutRequest) GetConfiguration() DestinationPineconeUpdate {
	if o == nil {
		return DestinationPineconeUpdate{}
	}
	return o.Configuration
}

func (o *DestinationPineconePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationPineconePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}