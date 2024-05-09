// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// OAuth20Credentials - OAuth2.0 credentials used to authenticate admin actions (creating/deleting corpora)
type OAuth20Credentials struct {
	// OAuth2.0 client id
	ClientID string `json:"client_id"`
	// OAuth2.0 client secret
	ClientSecret string `json:"client_secret"`
}

func (o *OAuth20Credentials) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *OAuth20Credentials) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

// DestinationVectaraUpdate - Configuration to connect to the Vectara instance
type DestinationVectaraUpdate struct {
	// The Name of Corpus to load data into
	CorpusName string `json:"corpus_name"`
	// Your customer id as it is in the authenticaion url
	CustomerID string `json:"customer_id"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// OAuth2.0 credentials used to authenticate admin actions (creating/deleting corpora)
	Oauth2 OAuth20Credentials `json:"oauth2"`
	// Parallelize indexing into Vectara with multiple threads
	Parallelize *bool `default:"false" json:"parallelize"`
	// List of fields in the record that should be in the section of the document. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
	// A field that will be used to populate the `title` of each document. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TitleField *string `default:"" json:"title_field"`
}

func (d DestinationVectaraUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVectaraUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVectaraUpdate) GetCorpusName() string {
	if o == nil {
		return ""
	}
	return o.CorpusName
}

func (o *DestinationVectaraUpdate) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *DestinationVectaraUpdate) GetMetadataFields() []string {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *DestinationVectaraUpdate) GetOauth2() OAuth20Credentials {
	if o == nil {
		return OAuth20Credentials{}
	}
	return o.Oauth2
}

func (o *DestinationVectaraUpdate) GetParallelize() *bool {
	if o == nil {
		return nil
	}
	return o.Parallelize
}

func (o *DestinationVectaraUpdate) GetTextFields() []string {
	if o == nil {
		return nil
	}
	return o.TextFields
}

func (o *DestinationVectaraUpdate) GetTitleField() *string {
	if o == nil {
		return nil
	}
	return o.TitleField
}