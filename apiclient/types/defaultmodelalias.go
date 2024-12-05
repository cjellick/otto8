package types

type DefaultModelAliasType string

const (
	DefaultModelAliasTypeTextEmbedding   DefaultModelAliasType = "text-embedding-3-large"
	DefaultModelAliasTypeLLM             DefaultModelAliasType = "gpt-4o"
	DefaultModelAliasTypeLLMMini         DefaultModelAliasType = "gpt-4o-mini"
	DefaultModelAliasTypeImageGeneration DefaultModelAliasType = "dall-e-3"
	DefaultModelAliasTypeVision          DefaultModelAliasType = "gpt-4o"
)

type DefaultModelAlias struct {
	DefaultModelAliasManifest
}

type DefaultModelAliasManifest struct {
	Alias string `json:"alias"`
	Model string `json:"model"`
}

type DefaultModelAliasList List[DefaultModelAlias]
