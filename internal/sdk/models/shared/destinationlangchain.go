// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Langchain string

const (
	LangchainLangchain Langchain = "langchain"
)

func (e Langchain) ToPointer() *Langchain {
	return &e
}

func (e *Langchain) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "langchain":
		*e = Langchain(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Langchain: %v", v)
	}
}

type DestinationLangchainSchemasMode string

const (
	DestinationLangchainSchemasModeFake DestinationLangchainSchemasMode = "fake"
)

func (e DestinationLangchainSchemasMode) ToPointer() *DestinationLangchainSchemasMode {
	return &e
}

func (e *DestinationLangchainSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationLangchainSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainSchemasMode: %v", v)
	}
}

// DestinationLangchainFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationLangchainFake struct {
	mode *DestinationLangchainSchemasMode `const:"fake" json:"mode"`
}

func (d DestinationLangchainFake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainFake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainFake) GetMode() *DestinationLangchainSchemasMode {
	return DestinationLangchainSchemasModeFake.ToPointer()
}

type DestinationLangchainMode string

const (
	DestinationLangchainModeOpenai DestinationLangchainMode = "openai"
)

func (e DestinationLangchainMode) ToPointer() *DestinationLangchainMode {
	return &e
}

func (e *DestinationLangchainMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationLangchainMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainMode: %v", v)
	}
}

// DestinationLangchainOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationLangchainOpenAI struct {
	mode      *DestinationLangchainMode `const:"openai" json:"mode"`
	OpenaiKey string                    `json:"openai_key"`
}

func (d DestinationLangchainOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainOpenAI) GetMode() *DestinationLangchainMode {
	return DestinationLangchainModeOpenai.ToPointer()
}

func (o *DestinationLangchainOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationLangchainEmbeddingType string

const (
	DestinationLangchainEmbeddingTypeDestinationLangchainOpenAI DestinationLangchainEmbeddingType = "destination-langchain_OpenAI"
	DestinationLangchainEmbeddingTypeDestinationLangchainFake   DestinationLangchainEmbeddingType = "destination-langchain_Fake"
)

// DestinationLangchainEmbedding - Embedding configuration
type DestinationLangchainEmbedding struct {
	DestinationLangchainOpenAI *DestinationLangchainOpenAI
	DestinationLangchainFake   *DestinationLangchainFake

	Type DestinationLangchainEmbeddingType
}

func CreateDestinationLangchainEmbeddingDestinationLangchainOpenAI(destinationLangchainOpenAI DestinationLangchainOpenAI) DestinationLangchainEmbedding {
	typ := DestinationLangchainEmbeddingTypeDestinationLangchainOpenAI

	return DestinationLangchainEmbedding{
		DestinationLangchainOpenAI: &destinationLangchainOpenAI,
		Type:                       typ,
	}
}

func CreateDestinationLangchainEmbeddingDestinationLangchainFake(destinationLangchainFake DestinationLangchainFake) DestinationLangchainEmbedding {
	typ := DestinationLangchainEmbeddingTypeDestinationLangchainFake

	return DestinationLangchainEmbedding{
		DestinationLangchainFake: &destinationLangchainFake,
		Type:                     typ,
	}
}

func (u *DestinationLangchainEmbedding) UnmarshalJSON(data []byte) error {

	var destinationLangchainFake DestinationLangchainFake = DestinationLangchainFake{}
	if err := utils.UnmarshalJSON(data, &destinationLangchainFake, "", true, true); err == nil {
		u.DestinationLangchainFake = &destinationLangchainFake
		u.Type = DestinationLangchainEmbeddingTypeDestinationLangchainFake
		return nil
	}

	var destinationLangchainOpenAI DestinationLangchainOpenAI = DestinationLangchainOpenAI{}
	if err := utils.UnmarshalJSON(data, &destinationLangchainOpenAI, "", true, true); err == nil {
		u.DestinationLangchainOpenAI = &destinationLangchainOpenAI
		u.Type = DestinationLangchainEmbeddingTypeDestinationLangchainOpenAI
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationLangchainEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationLangchainOpenAI != nil {
		return utils.MarshalJSON(u.DestinationLangchainOpenAI, "", true)
	}

	if u.DestinationLangchainFake != nil {
		return utils.MarshalJSON(u.DestinationLangchainFake, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationLangchainSchemasIndexingIndexing3Mode string

const (
	DestinationLangchainSchemasIndexingIndexing3ModeChromaLocal DestinationLangchainSchemasIndexingIndexing3Mode = "chroma_local"
)

func (e DestinationLangchainSchemasIndexingIndexing3Mode) ToPointer() *DestinationLangchainSchemasIndexingIndexing3Mode {
	return &e
}

func (e *DestinationLangchainSchemasIndexingIndexing3Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "chroma_local":
		*e = DestinationLangchainSchemasIndexingIndexing3Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainSchemasIndexingIndexing3Mode: %v", v)
	}
}

// DestinationLangchainChromaLocalPersistance - Chroma is a popular vector store that can be used to store and retrieve embeddings. It will build its index in memory and persist it to disk by the end of the sync.
type DestinationLangchainChromaLocalPersistance struct {
	// Name of the collection to use.
	CollectionName *string `default:"langchain" json:"collection_name"`
	// Path to the directory where chroma files will be written. The files will be placed inside that local mount.
	DestinationPath string                                            `json:"destination_path"`
	mode            *DestinationLangchainSchemasIndexingIndexing3Mode `const:"chroma_local" json:"mode"`
}

func (d DestinationLangchainChromaLocalPersistance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainChromaLocalPersistance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainChromaLocalPersistance) GetCollectionName() *string {
	if o == nil {
		return nil
	}
	return o.CollectionName
}

func (o *DestinationLangchainChromaLocalPersistance) GetDestinationPath() string {
	if o == nil {
		return ""
	}
	return o.DestinationPath
}

func (o *DestinationLangchainChromaLocalPersistance) GetMode() *DestinationLangchainSchemasIndexingIndexing3Mode {
	return DestinationLangchainSchemasIndexingIndexing3ModeChromaLocal.ToPointer()
}

type DestinationLangchainSchemasIndexingIndexingMode string

const (
	DestinationLangchainSchemasIndexingIndexingModeDocArrayHnswSearch DestinationLangchainSchemasIndexingIndexingMode = "DocArrayHnswSearch"
)

func (e DestinationLangchainSchemasIndexingIndexingMode) ToPointer() *DestinationLangchainSchemasIndexingIndexingMode {
	return &e
}

func (e *DestinationLangchainSchemasIndexingIndexingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DocArrayHnswSearch":
		*e = DestinationLangchainSchemasIndexingIndexingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainSchemasIndexingIndexingMode: %v", v)
	}
}

// DestinationLangchainDocArrayHnswSearch - DocArrayHnswSearch is a lightweight Document Index implementation provided by Docarray that runs fully locally and is best suited for small- to medium-sized datasets. It stores vectors on disk in hnswlib, and stores all other data in SQLite.
type DestinationLangchainDocArrayHnswSearch struct {
	// Path to the directory where hnswlib and meta data files will be written. The files will be placed inside that local mount. All files in the specified destination directory will be deleted on each run.
	DestinationPath string                                           `json:"destination_path"`
	mode            *DestinationLangchainSchemasIndexingIndexingMode `const:"DocArrayHnswSearch" json:"mode"`
}

func (d DestinationLangchainDocArrayHnswSearch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainDocArrayHnswSearch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainDocArrayHnswSearch) GetDestinationPath() string {
	if o == nil {
		return ""
	}
	return o.DestinationPath
}

func (o *DestinationLangchainDocArrayHnswSearch) GetMode() *DestinationLangchainSchemasIndexingIndexingMode {
	return DestinationLangchainSchemasIndexingIndexingModeDocArrayHnswSearch.ToPointer()
}

type DestinationLangchainSchemasIndexingMode string

const (
	DestinationLangchainSchemasIndexingModePinecone DestinationLangchainSchemasIndexingMode = "pinecone"
)

func (e DestinationLangchainSchemasIndexingMode) ToPointer() *DestinationLangchainSchemasIndexingMode {
	return &e
}

func (e *DestinationLangchainSchemasIndexingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinecone":
		*e = DestinationLangchainSchemasIndexingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationLangchainSchemasIndexingMode: %v", v)
	}
}

// DestinationLangchainPinecone - Pinecone is a popular vector store that can be used to store and retrieve embeddings. It is a managed service and can also be queried from outside of langchain.
type DestinationLangchainPinecone struct {
	// Pinecone index to use
	Index string                                   `json:"index"`
	mode  *DestinationLangchainSchemasIndexingMode `const:"pinecone" json:"mode"`
	// Pinecone environment to use
	PineconeEnvironment string `json:"pinecone_environment"`
	PineconeKey         string `json:"pinecone_key"`
}

func (d DestinationLangchainPinecone) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainPinecone) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainPinecone) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *DestinationLangchainPinecone) GetMode() *DestinationLangchainSchemasIndexingMode {
	return DestinationLangchainSchemasIndexingModePinecone.ToPointer()
}

func (o *DestinationLangchainPinecone) GetPineconeEnvironment() string {
	if o == nil {
		return ""
	}
	return o.PineconeEnvironment
}

func (o *DestinationLangchainPinecone) GetPineconeKey() string {
	if o == nil {
		return ""
	}
	return o.PineconeKey
}

type DestinationLangchainIndexingType string

const (
	DestinationLangchainIndexingTypeDestinationLangchainPinecone               DestinationLangchainIndexingType = "destination-langchain_Pinecone"
	DestinationLangchainIndexingTypeDestinationLangchainDocArrayHnswSearch     DestinationLangchainIndexingType = "destination-langchain_DocArrayHnswSearch"
	DestinationLangchainIndexingTypeDestinationLangchainChromaLocalPersistance DestinationLangchainIndexingType = "destination-langchain_Chroma (local persistance)"
)

// DestinationLangchainIndexing - Indexing configuration
type DestinationLangchainIndexing struct {
	DestinationLangchainPinecone               *DestinationLangchainPinecone
	DestinationLangchainDocArrayHnswSearch     *DestinationLangchainDocArrayHnswSearch
	DestinationLangchainChromaLocalPersistance *DestinationLangchainChromaLocalPersistance

	Type DestinationLangchainIndexingType
}

func CreateDestinationLangchainIndexingDestinationLangchainPinecone(destinationLangchainPinecone DestinationLangchainPinecone) DestinationLangchainIndexing {
	typ := DestinationLangchainIndexingTypeDestinationLangchainPinecone

	return DestinationLangchainIndexing{
		DestinationLangchainPinecone: &destinationLangchainPinecone,
		Type:                         typ,
	}
}

func CreateDestinationLangchainIndexingDestinationLangchainDocArrayHnswSearch(destinationLangchainDocArrayHnswSearch DestinationLangchainDocArrayHnswSearch) DestinationLangchainIndexing {
	typ := DestinationLangchainIndexingTypeDestinationLangchainDocArrayHnswSearch

	return DestinationLangchainIndexing{
		DestinationLangchainDocArrayHnswSearch: &destinationLangchainDocArrayHnswSearch,
		Type:                                   typ,
	}
}

func CreateDestinationLangchainIndexingDestinationLangchainChromaLocalPersistance(destinationLangchainChromaLocalPersistance DestinationLangchainChromaLocalPersistance) DestinationLangchainIndexing {
	typ := DestinationLangchainIndexingTypeDestinationLangchainChromaLocalPersistance

	return DestinationLangchainIndexing{
		DestinationLangchainChromaLocalPersistance: &destinationLangchainChromaLocalPersistance,
		Type: typ,
	}
}

func (u *DestinationLangchainIndexing) UnmarshalJSON(data []byte) error {

	var destinationLangchainDocArrayHnswSearch DestinationLangchainDocArrayHnswSearch = DestinationLangchainDocArrayHnswSearch{}
	if err := utils.UnmarshalJSON(data, &destinationLangchainDocArrayHnswSearch, "", true, true); err == nil {
		u.DestinationLangchainDocArrayHnswSearch = &destinationLangchainDocArrayHnswSearch
		u.Type = DestinationLangchainIndexingTypeDestinationLangchainDocArrayHnswSearch
		return nil
	}

	var destinationLangchainChromaLocalPersistance DestinationLangchainChromaLocalPersistance = DestinationLangchainChromaLocalPersistance{}
	if err := utils.UnmarshalJSON(data, &destinationLangchainChromaLocalPersistance, "", true, true); err == nil {
		u.DestinationLangchainChromaLocalPersistance = &destinationLangchainChromaLocalPersistance
		u.Type = DestinationLangchainIndexingTypeDestinationLangchainChromaLocalPersistance
		return nil
	}

	var destinationLangchainPinecone DestinationLangchainPinecone = DestinationLangchainPinecone{}
	if err := utils.UnmarshalJSON(data, &destinationLangchainPinecone, "", true, true); err == nil {
		u.DestinationLangchainPinecone = &destinationLangchainPinecone
		u.Type = DestinationLangchainIndexingTypeDestinationLangchainPinecone
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationLangchainIndexing) MarshalJSON() ([]byte, error) {
	if u.DestinationLangchainPinecone != nil {
		return utils.MarshalJSON(u.DestinationLangchainPinecone, "", true)
	}

	if u.DestinationLangchainDocArrayHnswSearch != nil {
		return utils.MarshalJSON(u.DestinationLangchainDocArrayHnswSearch, "", true)
	}

	if u.DestinationLangchainChromaLocalPersistance != nil {
		return utils.MarshalJSON(u.DestinationLangchainChromaLocalPersistance, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationLangchainProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `default:"0" json:"chunk_overlap"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields in the record that should be used to calculate the embedding. All other fields are passed along as meta fields. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields"`
}

func (d DestinationLangchainProcessingConfigModel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchainProcessingConfigModel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchainProcessingConfigModel) GetChunkOverlap() *int64 {
	if o == nil {
		return nil
	}
	return o.ChunkOverlap
}

func (o *DestinationLangchainProcessingConfigModel) GetChunkSize() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkSize
}

func (o *DestinationLangchainProcessingConfigModel) GetTextFields() []string {
	if o == nil {
		return []string{}
	}
	return o.TextFields
}

type DestinationLangchain struct {
	destinationType Langchain `const:"langchain" json:"destinationType"`
	// Embedding configuration
	Embedding DestinationLangchainEmbedding `json:"embedding"`
	// Indexing configuration
	Indexing   DestinationLangchainIndexing              `json:"indexing"`
	Processing DestinationLangchainProcessingConfigModel `json:"processing"`
}

func (d DestinationLangchain) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationLangchain) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationLangchain) GetDestinationType() Langchain {
	return LangchainLangchain
}

func (o *DestinationLangchain) GetEmbedding() DestinationLangchainEmbedding {
	if o == nil {
		return DestinationLangchainEmbedding{}
	}
	return o.Embedding
}

func (o *DestinationLangchain) GetIndexing() DestinationLangchainIndexing {
	if o == nil {
		return DestinationLangchainIndexing{}
	}
	return o.Indexing
}

func (o *DestinationLangchain) GetProcessing() DestinationLangchainProcessingConfigModel {
	if o == nil {
		return DestinationLangchainProcessingConfigModel{}
	}
	return o.Processing
}