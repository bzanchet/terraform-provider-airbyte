// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *DestinationSnowflakeResourceModel) ToCreateSDKType() *shared.DestinationSnowflakeCreateRequest {
	var credentials *shared.DestinationSnowflakeAuthorizationMethod
	var destinationSnowflakeAuthorizationMethodOAuth20 *shared.DestinationSnowflakeAuthorizationMethodOAuth20
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.DestinationSnowflakeAuthorizationMethodOAuth20AuthTypeEnum)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.DestinationSnowflakeAuthorizationMethodOAuth20AuthTypeEnum(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := new(string)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientID.IsNull() {
			*clientID = r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientID.ValueString()
		} else {
			clientID = nil
		}
		clientSecret := new(string)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientSecret.IsNull() {
			*clientSecret = r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientSecret.ValueString()
		} else {
			clientSecret = nil
		}
		refreshToken := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.RefreshToken.ValueString()
		destinationSnowflakeAuthorizationMethodOAuth20 = &shared.DestinationSnowflakeAuthorizationMethodOAuth20{
			AccessToken:  accessToken,
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if destinationSnowflakeAuthorizationMethodOAuth20 != nil {
		credentials = &shared.DestinationSnowflakeAuthorizationMethod{
			DestinationSnowflakeAuthorizationMethodOAuth20: destinationSnowflakeAuthorizationMethodOAuth20,
		}
	}
	var destinationSnowflakeAuthorizationMethodKeyPairAuthentication *shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication != nil {
		authType1 := new(shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthenticationAuthTypeEnum)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsNull() {
			*authType1 = shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthenticationAuthTypeEnum(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		privateKey := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKey.ValueString()
		privateKeyPassword := new(string)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.IsNull() {
			*privateKeyPassword = r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.ValueString()
		} else {
			privateKeyPassword = nil
		}
		destinationSnowflakeAuthorizationMethodKeyPairAuthentication = &shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication{
			AuthType:           authType1,
			PrivateKey:         privateKey,
			PrivateKeyPassword: privateKeyPassword,
		}
	}
	if destinationSnowflakeAuthorizationMethodKeyPairAuthentication != nil {
		credentials = &shared.DestinationSnowflakeAuthorizationMethod{
			DestinationSnowflakeAuthorizationMethodKeyPairAuthentication: destinationSnowflakeAuthorizationMethodKeyPairAuthentication,
		}
	}
	var destinationSnowflakeAuthorizationMethodUsernameAndPassword *shared.DestinationSnowflakeAuthorizationMethodUsernameAndPassword
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		authType2 := new(shared.DestinationSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsNull() {
			*authType2 = shared.DestinationSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.ValueString())
		} else {
			authType2 = nil
		}
		password := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.Password.ValueString()
		destinationSnowflakeAuthorizationMethodUsernameAndPassword = &shared.DestinationSnowflakeAuthorizationMethodUsernameAndPassword{
			AuthType: authType2,
			Password: password,
		}
	}
	if destinationSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		credentials = &shared.DestinationSnowflakeAuthorizationMethod{
			DestinationSnowflakeAuthorizationMethodUsernameAndPassword: destinationSnowflakeAuthorizationMethodUsernameAndPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationSnowflakeSnowflakeEnum(r.Configuration.DestinationType.ValueString())
	fileBufferCount := new(int64)
	if !r.Configuration.FileBufferCount.IsUnknown() && !r.Configuration.FileBufferCount.IsNull() {
		*fileBufferCount = r.Configuration.FileBufferCount.ValueInt64()
	} else {
		fileBufferCount = nil
	}
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	var loadingMethod *shared.DestinationSnowflakeDataStagingMethod
	var destinationSnowflakeDataStagingMethodSelectAnotherOption *shared.DestinationSnowflakeDataStagingMethodSelectAnotherOption
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodSelectAnotherOption != nil {
		method := shared.DestinationSnowflakeDataStagingMethodSelectAnotherOptionMethodEnum(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodSelectAnotherOption.Method.ValueString())
		destinationSnowflakeDataStagingMethodSelectAnotherOption = &shared.DestinationSnowflakeDataStagingMethodSelectAnotherOption{
			Method: method,
		}
	}
	if destinationSnowflakeDataStagingMethodSelectAnotherOption != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodSelectAnotherOption: destinationSnowflakeDataStagingMethodSelectAnotherOption,
		}
	}
	var destinationSnowflakeDataStagingMethodRecommendedInternalStaging *shared.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging != nil {
		method1 := shared.DestinationSnowflakeDataStagingMethodRecommendedInternalStagingMethodEnum(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging.Method.ValueString())
		destinationSnowflakeDataStagingMethodRecommendedInternalStaging = &shared.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging{
			Method: method1,
		}
	}
	if destinationSnowflakeDataStagingMethodRecommendedInternalStaging != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodRecommendedInternalStaging: destinationSnowflakeDataStagingMethodRecommendedInternalStaging,
		}
	}
	var destinationSnowflakeDataStagingMethodAWSS3Staging *shared.DestinationSnowflakeDataStagingMethodAWSS3Staging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging != nil {
		accessKeyID := r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.AccessKeyID.ValueString()
		var encryption *shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryption
		var destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption *shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption
		if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Encryption.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption != nil {
			encryptionType := shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryptionEncryptionTypeEnum(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Encryption.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption.EncryptionType.ValueString())
			destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption = &shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption{
				EncryptionType: encryptionType,
			}
		}
		if destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption != nil {
			encryption = &shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryption{
				DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption: destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionNoEncryption,
			}
		}
		var destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption *shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption
		if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Encryption.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption != nil {
			encryptionType1 := shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryptionEncryptionTypeEnum(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Encryption.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption.EncryptionType.ValueString())
			keyEncryptingKey := new(string)
			if !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Encryption.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.IsUnknown() && !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Encryption.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.IsNull() {
				*keyEncryptingKey = r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Encryption.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.ValueString()
			} else {
				keyEncryptingKey = nil
			}
			destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption = &shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption{
				EncryptionType:   encryptionType1,
				KeyEncryptingKey: keyEncryptingKey,
			}
		}
		if destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption != nil {
			encryption = &shared.DestinationSnowflakeDataStagingMethodAWSS3StagingEncryption{
				DestinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption: destinationSnowflakeDataStagingMethodAWSS3StagingEncryptionAESCBCEnvelopeEncryption,
			}
		}
		fileNamePattern := new(string)
		if !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.FileNamePattern.IsUnknown() && !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.FileNamePattern.IsNull() {
			*fileNamePattern = r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.FileNamePattern.ValueString()
		} else {
			fileNamePattern = nil
		}
		method2 := shared.DestinationSnowflakeDataStagingMethodAWSS3StagingMethodEnum(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Method.ValueString())
		purgeStagingData := new(bool)
		if !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.PurgeStagingData.IsUnknown() && !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.PurgeStagingData.IsNull() {
			*purgeStagingData = r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.PurgeStagingData.ValueBool()
		} else {
			purgeStagingData = nil
		}
		s3BucketName := r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.S3BucketName.ValueString()
		s3BucketRegion := new(shared.DestinationSnowflakeDataStagingMethodAWSS3StagingS3BucketRegionEnum)
		if !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.S3BucketRegion.IsUnknown() && !r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.S3BucketRegion.IsNull() {
			*s3BucketRegion = shared.DestinationSnowflakeDataStagingMethodAWSS3StagingS3BucketRegionEnum(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.S3BucketRegion.ValueString())
		} else {
			s3BucketRegion = nil
		}
		secretAccessKey := r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.SecretAccessKey.ValueString()
		destinationSnowflakeDataStagingMethodAWSS3Staging = &shared.DestinationSnowflakeDataStagingMethodAWSS3Staging{
			AccessKeyID:      accessKeyID,
			Encryption:       encryption,
			FileNamePattern:  fileNamePattern,
			Method:           method2,
			PurgeStagingData: purgeStagingData,
			S3BucketName:     s3BucketName,
			S3BucketRegion:   s3BucketRegion,
			SecretAccessKey:  secretAccessKey,
		}
	}
	if destinationSnowflakeDataStagingMethodAWSS3Staging != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodAWSS3Staging: destinationSnowflakeDataStagingMethodAWSS3Staging,
		}
	}
	var destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging *shared.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging != nil {
		bucketName := r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging.BucketName.ValueString()
		credentialsJSON := r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging.CredentialsJSON.ValueString()
		method3 := shared.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStagingMethodEnum(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging.Method.ValueString())
		projectID := r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging.ProjectID.ValueString()
		destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging = &shared.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging{
			BucketName:      bucketName,
			CredentialsJSON: credentialsJSON,
			Method:          method3,
			ProjectID:       projectID,
		}
	}
	if destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging: destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging,
		}
	}
	role := r.Configuration.Role.ValueString()
	schema := r.Configuration.Schema.ValueString()
	username := r.Configuration.Username.ValueString()
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.DestinationSnowflake{
		Credentials:     credentials,
		Database:        database,
		DestinationType: destinationType,
		FileBufferCount: fileBufferCount,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		LoadingMethod:   loadingMethod,
		Role:            role,
		Schema:          schema,
		Username:        username,
		Warehouse:       warehouse,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSnowflakeCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSnowflakeResourceModel) ToDeleteSDKType() *shared.DestinationSnowflakeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}
