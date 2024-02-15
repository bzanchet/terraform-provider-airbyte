// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationElasticsearchSchemasMethod string

const (
	DestinationElasticsearchSchemasMethodBasic DestinationElasticsearchSchemasMethod = "basic"
)

func (e DestinationElasticsearchSchemasMethod) ToPointer() *DestinationElasticsearchSchemasMethod {
	return &e
}

func (e *DestinationElasticsearchSchemasMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = DestinationElasticsearchSchemasMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchSchemasMethod: %v", v)
	}
}

// DestinationElasticsearchUsernamePassword - Basic auth header with a username and password
type DestinationElasticsearchUsernamePassword struct {
	method DestinationElasticsearchSchemasMethod `const:"basic" json:"method"`
	// Basic auth password to access a secure Elasticsearch server
	Password string `json:"password"`
	// Basic auth username to access a secure Elasticsearch server
	Username string `json:"username"`
}

func (d DestinationElasticsearchUsernamePassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchUsernamePassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchUsernamePassword) GetMethod() DestinationElasticsearchSchemasMethod {
	return DestinationElasticsearchSchemasMethodBasic
}

func (o *DestinationElasticsearchUsernamePassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *DestinationElasticsearchUsernamePassword) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type DestinationElasticsearchMethod string

const (
	DestinationElasticsearchMethodSecret DestinationElasticsearchMethod = "secret"
)

func (e DestinationElasticsearchMethod) ToPointer() *DestinationElasticsearchMethod {
	return &e
}

func (e *DestinationElasticsearchMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		*e = DestinationElasticsearchMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchMethod: %v", v)
	}
}

// DestinationElasticsearchAPIKeySecret - Use a api key and secret combination to authenticate
type DestinationElasticsearchAPIKeySecret struct {
	// The Key ID to used when accessing an enterprise Elasticsearch instance.
	APIKeyID string `json:"apiKeyId"`
	// The secret associated with the API Key ID.
	APIKeySecret string                         `json:"apiKeySecret"`
	method       DestinationElasticsearchMethod `const:"secret" json:"method"`
}

func (d DestinationElasticsearchAPIKeySecret) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearchAPIKeySecret) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearchAPIKeySecret) GetAPIKeyID() string {
	if o == nil {
		return ""
	}
	return o.APIKeyID
}

func (o *DestinationElasticsearchAPIKeySecret) GetAPIKeySecret() string {
	if o == nil {
		return ""
	}
	return o.APIKeySecret
}

func (o *DestinationElasticsearchAPIKeySecret) GetMethod() DestinationElasticsearchMethod {
	return DestinationElasticsearchMethodSecret
}

type DestinationElasticsearchAuthenticationMethodType string

const (
	DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAPIKeySecret     DestinationElasticsearchAuthenticationMethodType = "destination-elasticsearch_Api Key/Secret"
	DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchUsernamePassword DestinationElasticsearchAuthenticationMethodType = "destination-elasticsearch_Username/Password"
)

// DestinationElasticsearchAuthenticationMethod - The type of authentication to be used
type DestinationElasticsearchAuthenticationMethod struct {
	DestinationElasticsearchAPIKeySecret     *DestinationElasticsearchAPIKeySecret
	DestinationElasticsearchUsernamePassword *DestinationElasticsearchUsernamePassword

	Type DestinationElasticsearchAuthenticationMethodType
}

func CreateDestinationElasticsearchAuthenticationMethodDestinationElasticsearchAPIKeySecret(destinationElasticsearchAPIKeySecret DestinationElasticsearchAPIKeySecret) DestinationElasticsearchAuthenticationMethod {
	typ := DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAPIKeySecret

	return DestinationElasticsearchAuthenticationMethod{
		DestinationElasticsearchAPIKeySecret: &destinationElasticsearchAPIKeySecret,
		Type:                                 typ,
	}
}

func CreateDestinationElasticsearchAuthenticationMethodDestinationElasticsearchUsernamePassword(destinationElasticsearchUsernamePassword DestinationElasticsearchUsernamePassword) DestinationElasticsearchAuthenticationMethod {
	typ := DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchUsernamePassword

	return DestinationElasticsearchAuthenticationMethod{
		DestinationElasticsearchUsernamePassword: &destinationElasticsearchUsernamePassword,
		Type:                                     typ,
	}
}

func (u *DestinationElasticsearchAuthenticationMethod) UnmarshalJSON(data []byte) error {

	destinationElasticsearchAPIKeySecret := new(DestinationElasticsearchAPIKeySecret)
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchAPIKeySecret, "", true, true); err == nil {
		u.DestinationElasticsearchAPIKeySecret = destinationElasticsearchAPIKeySecret
		u.Type = DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAPIKeySecret
		return nil
	}

	destinationElasticsearchUsernamePassword := new(DestinationElasticsearchUsernamePassword)
	if err := utils.UnmarshalJSON(data, &destinationElasticsearchUsernamePassword, "", true, true); err == nil {
		u.DestinationElasticsearchUsernamePassword = destinationElasticsearchUsernamePassword
		u.Type = DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchUsernamePassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationElasticsearchAuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationElasticsearchAPIKeySecret != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchAPIKeySecret, "", true)
	}

	if u.DestinationElasticsearchUsernamePassword != nil {
		return utils.MarshalJSON(u.DestinationElasticsearchUsernamePassword, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Elasticsearch string

const (
	ElasticsearchElasticsearch Elasticsearch = "elasticsearch"
)

func (e Elasticsearch) ToPointer() *Elasticsearch {
	return &e
}

func (e *Elasticsearch) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "elasticsearch":
		*e = Elasticsearch(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Elasticsearch: %v", v)
	}
}

type DestinationElasticsearch struct {
	// The type of authentication to be used
	AuthenticationMethod *DestinationElasticsearchAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// CA certificate
	CaCertificate   *string       `json:"ca_certificate,omitempty"`
	destinationType Elasticsearch `const:"elasticsearch" json:"destinationType"`
	// The full url of the Elasticsearch server
	Endpoint string `json:"endpoint"`
	// If a primary key identifier is defined in the source, an upsert will be performed using the primary key value as the elasticsearch doc id. Does not support composite primary keys.
	Upsert *bool `default:"true" json:"upsert"`
}

func (d DestinationElasticsearch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationElasticsearch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationElasticsearch) GetAuthenticationMethod() *DestinationElasticsearchAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthenticationMethod
}

func (o *DestinationElasticsearch) GetCaCertificate() *string {
	if o == nil {
		return nil
	}
	return o.CaCertificate
}

func (o *DestinationElasticsearch) GetDestinationType() Elasticsearch {
	return ElasticsearchElasticsearch
}

func (o *DestinationElasticsearch) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *DestinationElasticsearch) GetUpsert() *bool {
	if o == nil {
		return nil
	}
	return o.Upsert
}
