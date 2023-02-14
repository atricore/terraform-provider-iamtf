package jossoappi

func (idpCfg SamlR2IDPConfigDTO) ToProviderConfig() (*ProviderConfigDTO, error) {
	cfg := NewProviderConfigDTO()
	cfg.AdditionalProperties = make(map[string]interface{})

	// Build specific type
	//cfg.AdditionalProperties["@id"] = idpCfg.AdditionalProperties["@id"]
	cfg.AdditionalProperties["@c"] = ".SamlR2IDPConfigDTO"

	cfg.Description = idpCfg.Description
	cfg.DisplayName = idpCfg.DisplayName
	cfg.ElementId = idpCfg.ElementId
	cfg.Name = idpCfg.Name
	cfg.AdditionalProperties["useSampleStore"] = AsBool(idpCfg.UseSampleStore, false)
	cfg.AdditionalProperties["useSystemStore"] = AsBool(idpCfg.UseSystemStore, false)

	if !*idpCfg.UseSampleStore && !*idpCfg.UseSystemStore {
		storeProps := toKeyStoreMap(idpCfg.GetSigner())
		cfg.AdditionalProperties["signer"] = storeProps
		cfg.AdditionalProperties["encrypter"] = storeProps
	}

	return cfg, nil

}

func NewSamlR2IDPConfigDTOInit() *SamlR2IDPConfigDTO {
	idpCfg := NewSamlR2IDPConfigDTOWithDefaults()
	idpCfg.SetUseSampleStore(false)
	idpCfg.SetUseSystemStore(false)
	idpCfg.AdditionalProperties = make(map[string]interface{})
	idpCfg.AdditionalProperties["@c"] = ".SamlR2IDPConfigDTO"

	return idpCfg

}
