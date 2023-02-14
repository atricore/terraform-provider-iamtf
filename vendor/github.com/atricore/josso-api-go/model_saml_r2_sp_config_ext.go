package jossoappi

func (spCfg SamlR2SPConfigDTO) ToProviderConfig() (*ProviderConfigDTO, error) {
	cfg := NewProviderConfigDTO()
	cfg.AdditionalProperties = make(map[string]interface{})

	// Build specific type
	//cfg.AdditionalProperties["@id"] = spCfg.AdditionalProperties["@id"]
	cfg.AdditionalProperties["@c"] = ".SamlR2SPConfigDTO"

	cfg.Description = spCfg.Description
	cfg.DisplayName = spCfg.DisplayName
	cfg.ElementId = spCfg.ElementId
	cfg.Name = spCfg.Name
	cfg.AdditionalProperties["useSampleStore"] = AsBool(spCfg.UseSampleStore, false)
	cfg.AdditionalProperties["useSystemStore"] = AsBool(spCfg.UseSystemStore, false)

	if !*spCfg.UseSampleStore && !*spCfg.UseSystemStore {
		storeProps := toKeyStoreMap(spCfg.GetSigner())
		cfg.AdditionalProperties["signer"] = storeProps
		cfg.AdditionalProperties["encrypter"] = storeProps
	}

	return cfg, nil

}

func NewSamlR2SPConfigDTOInit() *SamlR2SPConfigDTO {
	spCfg := NewSamlR2SPConfigDTOWithDefaults()
	spCfg.SetUseSampleStore(false)
	spCfg.SetUseSystemStore(false)
	spCfg.AdditionalProperties = make(map[string]interface{})
	spCfg.AdditionalProperties["@c"] = ".SamlR2SPConfigDTO"

	return spCfg

}
