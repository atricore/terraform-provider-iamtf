package jossoappi

func toKeyStoreMap(store KeystoreDTO) map[string]interface{} {

	storeProps := make(map[string]interface{})

	//	if store.AdditionalProperties["@id"] != nil {
	//		storeProps["@id"] = store.AdditionalProperties["@id"].(int)
	//	}
	storeProps["certificateAlias"] = store.GetCertificateAlias()
	storeProps["displayName"] = store.GetDisplayName()
	storeProps["elementId"] = store.GetElementId()
	storeProps["id"] = store.GetId()
	storeProps["keystorePassOnly"] = store.GetKeystorePassOnly()
	storeProps["name"] = store.GetName()
	storeProps["password"] = store.GetPassword()
	storeProps["privateKeyName"] = store.GetPrivateKeyName()
	storeProps["privateKeyPassword"] = store.GetPrivateKeyPassword()
	storeProps["type"] = store.GetType()

	resourceProps := make(map[string]interface{})
	storeProps["store"] = resourceProps

	//	resourceProps["@id"] = store.Store.AdditionalProperties["@id"]
	resourceProps["displayName"] = store.Store.GetDisplayName()
	resourceProps["elementId"] = store.Store.GetElementId()
	resourceProps["name"] = store.Store.GetName()
	resourceProps["uri"] = store.Store.GetUri()
	resourceProps["value"] = store.Store.GetValue()

	return storeProps

}

// Transforms a map into an KeyStoreDTO
func toKeyStoreDTO(props map[string]interface{}) *KeystoreDTO {

	store := NewKeystoreDTO()
	store.AdditionalProperties = make(map[string]interface{})
	//store.AdditionalProperties["@id"] = storeId

	if props["displayName"] != nil {
		store.SetDisplayName((props["displayName"].(string)))
	}
	if props["elementId"] != nil {
		store.SetElementId((props["elementId"].(string)))
	}
	if props["id"] != nil {
		store.SetId(AsInt64(props["id"], 0))
	}
	if props["name"] != nil {
		store.SetName((props["name"].(string)))
	}

	store.SetCertificateAlias((props["certificateAlias"].(string)))
	store.SetKeystorePassOnly((props["keystorePassOnly"].(bool)))
	store.SetPassword((props["password"].(string)))
	store.SetPrivateKeyName((props["privateKeyName"].(string)))
	store.SetPrivateKeyPassword((props["privateKeyPassword"].(string)))
	store.SetPrivateKeyPassword((props["privateKeyPassword"].(string)))
	store.SetType((props["type"].(string)))

	resourceProps := props["store"].(map[string]interface{})
	store.Store = NewResourceDTO()
	store.Store.AdditionalProperties = map[string]interface{}{}
	//store.Store.AdditionalProperties["@id"] = resourceProps["@id"]
	store.Store.DisplayName = IPtrString(resourceProps["displayName"])
	store.Store.ElementId = IPtrString(resourceProps["elementId"])
	store.Store.Name = IPtrString(resourceProps["name"])
	store.Store.Uri = IPtrString(resourceProps["uri"])
	store.Store.Value = IPtrString(resourceProps["value"])

	return store
}

func NewKeystoreDTOInit(n string, d string, s *ResourceDTO) *KeystoreDTO {
	ks := NewKeystoreDTO()
	ks.SetName(n)
	ks.SetDisplayName(d)
	ks.AdditionalProperties = make(map[string]interface{})
	ks.AdditionalProperties["@c"] = ".KeystoreDTO"

	ks.SetStore(*s)

	return ks
}
