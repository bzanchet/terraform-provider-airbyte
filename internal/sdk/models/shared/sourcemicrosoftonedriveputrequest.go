// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceMicrosoftOnedrivePutRequest struct {
	// SourceMicrosoftOneDriveSpec class for Microsoft OneDrive Source Specification.
	// This class combines the authentication details with additional configuration for the OneDrive API.
	Configuration SourceMicrosoftOnedriveUpdate `json:"configuration"`
	Name          string                        `json:"name"`
	WorkspaceID   string                        `json:"workspaceId"`
}

func (o *SourceMicrosoftOnedrivePutRequest) GetConfiguration() SourceMicrosoftOnedriveUpdate {
	if o == nil {
		return SourceMicrosoftOnedriveUpdate{}
	}
	return o.Configuration
}

func (o *SourceMicrosoftOnedrivePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMicrosoftOnedrivePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}