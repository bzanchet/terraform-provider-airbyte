// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceFileSecureResourceModel) ToCreateSDKType() *shared.SourceFileSecureCreateRequest {
	datasetName := r.Configuration.DatasetName.ValueString()
	format := shared.SourceFileSecureFileFormatEnum(r.Configuration.Format.ValueString())
	var provider shared.SourceFileSecureStorageProvider
	var sourceFileSecureStorageProviderHTTPSPublicWeb *shared.SourceFileSecureStorageProviderHTTPSPublicWeb
	if r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb != nil {
		storage := shared.SourceFileSecureStorageProviderHTTPSPublicWebStorageEnum(r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb.Storage.ValueString())
		userAgent := new(bool)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb.UserAgent.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb.UserAgent.IsNull() {
			*userAgent = r.Configuration.Provider.SourceFileSecureStorageProviderHTTPSPublicWeb.UserAgent.ValueBool()
		} else {
			userAgent = nil
		}
		sourceFileSecureStorageProviderHTTPSPublicWeb = &shared.SourceFileSecureStorageProviderHTTPSPublicWeb{
			Storage:   storage,
			UserAgent: userAgent,
		}
	}
	if sourceFileSecureStorageProviderHTTPSPublicWeb != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderHTTPSPublicWeb: sourceFileSecureStorageProviderHTTPSPublicWeb,
		}
	}
	var sourceFileSecureStorageProviderGCSGoogleCloudStorage *shared.SourceFileSecureStorageProviderGCSGoogleCloudStorage
	if r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage != nil {
		serviceAccountJSON := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.IsNull() {
			*serviceAccountJSON = r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage.ServiceAccountJSON.ValueString()
		} else {
			serviceAccountJSON = nil
		}
		storage1 := shared.SourceFileSecureStorageProviderGCSGoogleCloudStorageStorageEnum(r.Configuration.Provider.SourceFileSecureStorageProviderGCSGoogleCloudStorage.Storage.ValueString())
		sourceFileSecureStorageProviderGCSGoogleCloudStorage = &shared.SourceFileSecureStorageProviderGCSGoogleCloudStorage{
			ServiceAccountJSON: serviceAccountJSON,
			Storage:            storage1,
		}
	}
	if sourceFileSecureStorageProviderGCSGoogleCloudStorage != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderGCSGoogleCloudStorage: sourceFileSecureStorageProviderGCSGoogleCloudStorage,
		}
	}
	var sourceFileSecureStorageProviderS3AmazonWebServices *shared.SourceFileSecureStorageProviderS3AmazonWebServices
	if r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices != nil {
		awsAccessKeyID := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsAccessKeyID.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsAccessKeyID.IsNull() {
			*awsAccessKeyID = r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsAccessKeyID.ValueString()
		} else {
			awsAccessKeyID = nil
		}
		awsSecretAccessKey := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsSecretAccessKey.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsSecretAccessKey.IsNull() {
			*awsSecretAccessKey = r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.AwsSecretAccessKey.ValueString()
		} else {
			awsSecretAccessKey = nil
		}
		storage2 := shared.SourceFileSecureStorageProviderS3AmazonWebServicesStorageEnum(r.Configuration.Provider.SourceFileSecureStorageProviderS3AmazonWebServices.Storage.ValueString())
		sourceFileSecureStorageProviderS3AmazonWebServices = &shared.SourceFileSecureStorageProviderS3AmazonWebServices{
			AwsAccessKeyID:     awsAccessKeyID,
			AwsSecretAccessKey: awsSecretAccessKey,
			Storage:            storage2,
		}
	}
	if sourceFileSecureStorageProviderS3AmazonWebServices != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderS3AmazonWebServices: sourceFileSecureStorageProviderS3AmazonWebServices,
		}
	}
	var sourceFileSecureStorageProviderAzBlobAzureBlobStorage *shared.SourceFileSecureStorageProviderAzBlobAzureBlobStorage
	if r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage != nil {
		sasToken := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SasToken.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SasToken.IsNull() {
			*sasToken = r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SasToken.ValueString()
		} else {
			sasToken = nil
		}
		sharedKey := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SharedKey.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SharedKey.IsNull() {
			*sharedKey = r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.SharedKey.ValueString()
		} else {
			sharedKey = nil
		}
		storage3 := shared.SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorageEnum(r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.Storage.ValueString())
		storageAccount := r.Configuration.Provider.SourceFileSecureStorageProviderAzBlobAzureBlobStorage.StorageAccount.ValueString()
		sourceFileSecureStorageProviderAzBlobAzureBlobStorage = &shared.SourceFileSecureStorageProviderAzBlobAzureBlobStorage{
			SasToken:       sasToken,
			SharedKey:      sharedKey,
			Storage:        storage3,
			StorageAccount: storageAccount,
		}
	}
	if sourceFileSecureStorageProviderAzBlobAzureBlobStorage != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderAzBlobAzureBlobStorage: sourceFileSecureStorageProviderAzBlobAzureBlobStorage,
		}
	}
	var sourceFileSecureStorageProviderSSHSecureShell *shared.SourceFileSecureStorageProviderSSHSecureShell
	if r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell != nil {
		host := r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Host.ValueString()
		password := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Password.IsNull() {
			*password = r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Password.ValueString()
		} else {
			password = nil
		}
		port := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Port.IsNull() {
			*port = r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Port.ValueString()
		} else {
			port = nil
		}
		storage4 := shared.SourceFileSecureStorageProviderSSHSecureShellStorageEnum(r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.Storage.ValueString())
		user := r.Configuration.Provider.SourceFileSecureStorageProviderSSHSecureShell.User.ValueString()
		sourceFileSecureStorageProviderSSHSecureShell = &shared.SourceFileSecureStorageProviderSSHSecureShell{
			Host:     host,
			Password: password,
			Port:     port,
			Storage:  storage4,
			User:     user,
		}
	}
	if sourceFileSecureStorageProviderSSHSecureShell != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderSSHSecureShell: sourceFileSecureStorageProviderSSHSecureShell,
		}
	}
	var sourceFileSecureStorageProviderSCPSecureCopyProtocol *shared.SourceFileSecureStorageProviderSCPSecureCopyProtocol
	if r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol != nil {
		host1 := r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Host.ValueString()
		password1 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Password.IsNull() {
			*password1 = r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Password.ValueString()
		} else {
			password1 = nil
		}
		port1 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Port.IsNull() {
			*port1 = r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Port.ValueString()
		} else {
			port1 = nil
		}
		storage5 := shared.SourceFileSecureStorageProviderSCPSecureCopyProtocolStorageEnum(r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.Storage.ValueString())
		user1 := r.Configuration.Provider.SourceFileSecureStorageProviderSCPSecureCopyProtocol.User.ValueString()
		sourceFileSecureStorageProviderSCPSecureCopyProtocol = &shared.SourceFileSecureStorageProviderSCPSecureCopyProtocol{
			Host:     host1,
			Password: password1,
			Port:     port1,
			Storage:  storage5,
			User:     user1,
		}
	}
	if sourceFileSecureStorageProviderSCPSecureCopyProtocol != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderSCPSecureCopyProtocol: sourceFileSecureStorageProviderSCPSecureCopyProtocol,
		}
	}
	var sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol *shared.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol
	if r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol != nil {
		host2 := r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Host.ValueString()
		password2 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Password.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Password.IsNull() {
			*password2 = r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Password.ValueString()
		} else {
			password2 = nil
		}
		port2 := new(string)
		if !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Port.IsUnknown() && !r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Port.IsNull() {
			*port2 = r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Port.ValueString()
		} else {
			port2 = nil
		}
		storage6 := shared.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorageEnum(r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.Storage.ValueString())
		user2 := r.Configuration.Provider.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol.User.ValueString()
		sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol = &shared.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol{
			Host:     host2,
			Password: password2,
			Port:     port2,
			Storage:  storage6,
			User:     user2,
		}
	}
	if sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol: sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol,
		}
	}
	readerOptions := new(string)
	if !r.Configuration.ReaderOptions.IsUnknown() && !r.Configuration.ReaderOptions.IsNull() {
		*readerOptions = r.Configuration.ReaderOptions.ValueString()
	} else {
		readerOptions = nil
	}
	sourceType := shared.SourceFileSecureFileSecureEnum(r.Configuration.SourceType.ValueString())
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceFileSecure{
		DatasetName:   datasetName,
		Format:        format,
		Provider:      provider,
		ReaderOptions: readerOptions,
		SourceType:    sourceType,
		URL:           url,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFileSecureCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFileSecureResourceModel) ToDeleteSDKType() *shared.SourceFileSecureCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
