package jossoappi

// Transforms the federated connection's channelA (FederatedChannelDTO) into a IdentityProviderChannelDTO
func (f *FederatedConnectionDTO) GetIDPChannel() (*IdentityProviderChannelDTO, error) {
	c := f.GetChannelB()
	var idpc IdentityProviderChannelDTO

	idpc.SetActiveBindings(c.GetActiveBindings())
	idpc.SetActiveProfiles(c.GetActiveProfiles())
	idpc.SetDescription(c.GetDescription())
	idpc.SetDisplayName(c.GetDisplayName())
	idpc.SetElementId(c.GetElementId())
	idpc.SetId(c.GetId())
	idpc.SetLocation(c.GetLocation())
	idpc.SetName(c.GetName())
	idpc.SetOverrideProviderSetup(c.GetOverrideProviderSetup())

	prefered := false
	if c.AdditionalProperties["preferred"] != nil {
		prefered = (c.AdditionalProperties["preferred"].(bool))
	}
	idpc.SetPreferred(prefered)

	if idpc.GetOverrideProviderSetup() {
		idpc.SetSignatureHash(AsString(c.AdditionalProperties["signatureHash"], ""))
		idpc.SetMessageTtl(AsInt32(c.AdditionalProperties["messageTtl"], 0))
		idpc.SetMessageTtlTolerance(AsInt32(c.AdditionalProperties["messageTtlTolerance"], 0))
		accountLinkage := toAccountLinkagePolicyDTO(c.AdditionalProperties["accountLinkagePolicy"].(map[string]interface{}))
		idpc.SetAccountLinkagePolicy(accountLinkage)
		idpc.SetEnableProxyExtension(AsBool(c.AdditionalProperties["enableProxyExtension"], false))
		idMapping := toIdentityMappingPolicyDTO(c.AdditionalProperties["identityMappingPolicy"].(map[string]interface{}))
		idpc.SetIdentityMappingPolicy(idMapping)
		idpc.SetSignAuthenticationRequests(AsBool(c.AdditionalProperties["signAuthenticationRequests"], false))
		idpc.SetWantAssertionSigned(AsBool(c.AdditionalProperties["wantAssertionSigned"], false))
	}

	idpc.AdditionalProperties = map[string]interface{}{
		"@c": ".IdentityProviderChannelDTO",
	}

	return &idpc, nil
}

// Transforms the IdentityProviderChannelDTO into a FederatedChannel and sets it into channelA
func (f *FederatedConnectionDTO) SetIDPChannel(idpc *IdentityProviderChannelDTO) error {

	var c FederatedChannelDTO

	c.SetActiveBindings(idpc.GetActiveBindings())
	c.SetActiveProfiles(idpc.GetActiveProfiles())
	c.SetDescription(idpc.GetDescription())
	c.SetDisplayName(idpc.GetDisplayName())
	c.SetElementId(idpc.GetElementId())
	c.SetId(idpc.GetId())
	c.SetLocation(idpc.GetLocation())
	c.SetName(idpc.GetName())
	c.SetOverrideProviderSetup(idpc.GetOverrideProviderSetup())

	c.AdditionalProperties = make(map[string]interface{})
	c.AdditionalProperties["preferred"] = idpc.GetPreferred()
	c.AdditionalProperties["@c"] = ".IdentityProviderChannelDTO"

	if idpc.GetOverrideProviderSetup() {
		c.AdditionalProperties["signatureHash"] = idpc.GetSignatureHash()
		c.AdditionalProperties["messageTtl"] = idpc.GetMessageTtl()
		c.AdditionalProperties["messageTtlTolerance"] = idpc.GetMessageTtlTolerance()
		c.AdditionalProperties["accountLinkagePolicy"] = toAccountLinkagePolicyMap(idpc.GetAccountLinkagePolicy())
		c.AdditionalProperties["enableProxyExtension"] = idpc.GetEnableProxyExtension()
		c.AdditionalProperties["identityMappingPolicy"] = toIdentityMappingPolicyMap(idpc.GetIdentityMappingPolicy())
		c.AdditionalProperties["signAuthenticationRequests"] = idpc.GetSignAuthenticationRequests()
		c.AdditionalProperties["wantAssertionSigned"] = idpc.GetWantAssertionSigned()
	}

	f.SetChannelB(c)
	return nil
}

// Transforms IdentityMappingPolicyDTO a map
func toIdentityMappingPolicyMap(dto IdentityMappingPolicyDTO) map[string]interface{} {

	props := make(map[string]interface{})

	props["customMapper"] = dto.GetCustomMapper()
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["mappingType"] = dto.GetMappingType()
	props["name"] = dto.GetName()
	props["useLocalId"] = dto.GetUseLocalId()
	props["@c"] = ".IdentityMappingPolicyDTO"
	return props
}

// Transforms a map into an IdentityMappingPolicyDTO
func toIdentityMappingPolicyDTO(props map[string]interface{}) IdentityMappingPolicyDTO {
	dto := NewIdentityMappingPolicyDTO()
	dto.SetCustomMapper(AsString(props["customMapper"], ""))
	dto.SetElementId(AsString(props["elementId"], ""))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetMappingType(AsString(props["mappingType"], ""))
	dto.SetName(AsString(props["name"], ""))
	dto.SetUseLocalId(AsBool(props["useLocalId"], false))
	dto.AdditionalProperties = map[string]interface{}{
		"@c": ".IdentityMappingPolicyDTO",
	}
	return *dto
}

// Transforms AccountLinkagePolicy a map
func toAccountLinkagePolicyMap(dto AccountLinkagePolicyDTO) map[string]interface{} {
	props := make(map[string]interface{})

	props["customLinkEmitter"] = dto.GetCustomLinkEmitter()
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["linkEmitterType"] = dto.GetLinkEmitterType()
	props["name"] = dto.GetName()
	props["@c"] = ".AccountLinkagePolicyDTO"
	return props
}

// Transforms a map into an AccountLinkagePolicyDTO
func toAccountLinkagePolicyDTO(props map[string]interface{}) AccountLinkagePolicyDTO {

	dto := NewAccountLinkagePolicyDTO()
	dto.SetCustomLinkEmitter(AsString(props["customLinkEmitter"], ""))
	dto.SetElementId(AsString(props["elementId"], ""))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetLinkEmitterType(AsString(props["linkEmitterType"], ""))
	dto.SetName(AsString(props["name"], ""))
	dto.AdditionalProperties = map[string]interface{}{
		"@c": ".AccountLinkagePolicyDTO",
	}
	return *dto
}

// Transforms a map into an EmissionPolicyDTO
func toEmissionPolicyDTO(props map[string]interface{}) AuthenticationAssertionEmissionPolicyDTO {
	dto := NewAuthenticationAssertionEmissionPolicyDTO()
	dto.SetElementId(AsString(props["elementId"], ""))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetName(AsString(props["name"], ""))
	dto.AdditionalProperties = map[string]interface{}{
		"@c": ".AuthenticationAssertionEmissionPolicyDTO",
	}
	return *dto
}

// Transforms a map into an AuthenticationContractDTO
func toAuthenticationContractDTO(props map[string]interface{}) AuthenticationContractDTO {
	dto := NewAuthenticationContractDTO()
	dto.SetElementId(AsString(props["elementId"], ""))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetName(AsString(props["name"], ""))
	dto.AdditionalProperties = map[string]interface{}{
		"@c": ".AuthenticationContractDTO",
	}
	return *dto
}

// Transforms a map into an AttributeProfileDTO
func toAttributeProfileDTO(props map[string]interface{}) AttributeProfileDTO {
	dto := NewAttributeProfileDTO()
	dto.SetElementId(AsString(props["elementId"], ""))
	dto.SetId(AsInt64(props["id"], 0))
	dto.SetName(AsString(props["name"], ""))
	dto.SetProfileType(AsString(props["profileType"], ""))
	// TODO : Support custom attribute profile
	dto.AdditionalProperties = map[string]interface{}{
		"@c": ".BuiltInAttributeProfileDTO",
	}
	return *dto
}

// Transforms a map into an SubjectNameIDPolicyDTO
func toSubjectNameIDPolicyDTO(props map[string]interface{}) SubjectNameIdentifierPolicyDTO {
	dto := NewSubjectNameIdentifierPolicyDTO()
	dto.SetDescriptionKey(AsString(props["descriptionKey"], ""))
	dto.SetId(AsString(props["id"], ""))
	dto.SetName(AsString(props["name"], ""))
	dto.SetSubjectAttribute(AsString(props["subjectAttribute"], ""))
	dto.SetType(AsString(props["type"], ""))
	dto.AdditionalProperties = map[string]interface{}{
		"@c": ".SubjectNameIdentifierPolicyDTO",
	}
	return *dto
}

// Transforms the InternalSaml2ServiceProviderChannelDTO into a FederatedChannel and sets it into channelA
func (f *FederatedConnectionDTO) SetSPChannel(spc *InternalSaml2ServiceProviderChannelDTO) error {

	var c FederatedChannelDTO

	c.SetId(spc.GetId())
	c.SetActiveBindings(spc.GetActiveBindings())
	c.SetActiveProfiles(spc.GetActiveProfiles())
	c.SetDescription(spc.GetDescription())
	c.SetDisplayName(spc.GetDisplayName())
	c.SetElementId(spc.GetElementId())
	c.SetLocation(spc.GetLocation())
	c.SetName(spc.GetName())
	c.SetOverrideProviderSetup(spc.GetOverrideProviderSetup())

	c.AdditionalProperties = make(map[string]interface{})
	c.AdditionalProperties["@c"] = ".InternalSaml2ServiceProviderChannelDTO"

	if spc.GetOverrideProviderSetup() {
		c.AdditionalProperties["emissionPolicy"] = toEmissionPolicyMap(spc.GetEmissionPolicy())
		c.AdditionalProperties["restrictedRoles"] = spc.GetRestrictedRoles()
		c.AdditionalProperties["requiredRoles"] = spc.GetRequiredRoles()
		c.AdditionalProperties["messageTtl"] = spc.GetMessageTtl()
		c.AdditionalProperties["requiredRolesMatchMode"] = spc.GetRequiredRolesMatchMode()
		c.AdditionalProperties["restrictedRolesMatchMode"] = spc.GetRestrictedRolesMatchMode()
		c.AdditionalProperties["encryptAssertionAlgorithm"] = spc.GetEncryptAssertionAlgorithm()
		c.AdditionalProperties["ignoreRequestedNameIDPolicy"] = spc.GetIgnoreRequestedNameIDPolicy()

		if p, ok := spc.GetSubjectNameIDPolicyOk(); ok {
			c.AdditionalProperties["subjectNameIDPolicy"] = toSubjectNameIDPolicyMap(p)
		}
		c.AdditionalProperties["encryptAssertion"] = spc.GetEncryptAssertion()
		c.AdditionalProperties["signatureHash"] = spc.GetSignatureHash()
		c.AdditionalProperties["attributeProfile"] = toAttributeProfilemap(spc.GetAttributeProfile())
		c.AdditionalProperties["authenticationContract"] = toAuthenticationContractmap(spc.GetAuthenticationContract())
		c.AdditionalProperties["wantAuthnRequestsSigned"] = spc.GetWantAuthnRequestsSigned()
	}

	f.SetChannelA(c)
	return nil

}

// IDP Side, has an SP channel
func (f *FederatedConnectionDTO) GetSPChannel() (*InternalSaml2ServiceProviderChannelDTO, error) {
	c := f.GetChannelA()
	var spc InternalSaml2ServiceProviderChannelDTO

	spc.SetId(c.GetId())
	spc.SetActiveBindings(c.GetActiveBindings())
	spc.SetActiveProfiles(c.GetActiveProfiles())
	spc.SetDescription(c.GetDescription())
	spc.SetDisplayName(c.GetDisplayName())
	spc.SetElementId(c.GetElementId())
	spc.SetLocation(c.GetLocation())
	spc.SetName(c.GetName())
	spc.SetOverrideProviderSetup(c.GetOverrideProviderSetup())

	if c.GetOverrideProviderSetup() {

		if c.AdditionalProperties["emissionPolicy"] != nil {
			emissionPolicy := toEmissionPolicyDTO(c.AdditionalProperties["emissionPolicy"].(map[string]interface{}))
			spc.SetEmissionPolicy(emissionPolicy)
		}

		spc.SetRestrictedRoles(AsStringArr(c.AdditionalProperties["restrictedRoles"]))
		spc.SetRequiredRoles(AsStringArr(c.AdditionalProperties["requiredRoles"]))
		spc.SetMessageTtl(AsInt32(c.AdditionalProperties["messageTtl"], 0))
		spc.SetRequiredRolesMatchMode(AsInt32(c.AdditionalProperties["requiredRolesMatchMode"], 0))
		spc.SetRestrictedRolesMatchMode(AsInt32(c.AdditionalProperties["restrictedRolesMatchMode"], 0))
		spc.SetEncryptAssertionAlgorithm(AsString(c.AdditionalProperties["encryptAssertionAlgorithm"], ""))
		spc.SetIgnoreRequestedNameIDPolicy(AsBool(c.AdditionalProperties["ignoreRequestedNameIDPolicy"], false))

		if c.AdditionalProperties["subjectNameIDPolicy"] != nil {
			subjectNameId := toSubjectNameIDPolicyDTO(c.AdditionalProperties["subjectNameIDPolicy"].(map[string]interface{}))
			spc.SetSubjectNameIDPolicy(subjectNameId)
		}

		spc.SetEncryptAssertion(AsBool(c.AdditionalProperties["encryptAssertion"], false))
		spc.SetSignatureHash(AsString(c.AdditionalProperties["signatureHash"], ""))

		if c.AdditionalProperties["attributeProfile"] != nil {
			attrProfile := toAttributeProfileDTO(c.AdditionalProperties["attributeProfile"].(map[string]interface{}))
			spc.SetAttributeProfile(attrProfile)
		}

		if c.AdditionalProperties["authenticationContract"] != nil {
			authnContract := toAuthenticationContractDTO(c.AdditionalProperties["authenticationContract"].(map[string]interface{}))
			spc.SetAuthenticationContract(authnContract)
		}

		spc.SetWantAuthnRequestsSigned(AsBool(c.AdditionalProperties["wantAuthnRequestsSigned"], false))
	}

	spc.AdditionalProperties = map[string]interface{}{
		"@c": ".InternalSaml2ServiceProviderChannelDTO",
	}

	return &spc, nil
}

// Transforms AuthenticationContract a map
func toAuthenticationContractmap(dto AuthenticationContractDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	props["@c"] = ".AuthenticationContractDTO"
	return &props
}

// Transforms AttributeProfile a map
func toAttributeProfilemap(dto AttributeProfileDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	props["profileType"] = dto.GetProfileType()
	props["@c"] = ".BuiltInAttributeProfileDTO"
	return &props
}

// Transforms EmissionPolicy a map
func toEmissionPolicyMap(dto AuthenticationAssertionEmissionPolicyDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["elementId"] = dto.GetElementId()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	props["@c"] = ".AuthenticationAssertionEmissionPolicyDTO"
	return &props
}

// Transforms SubjectNameIDPolicy a map
func toSubjectNameIDPolicyMap(dto *SubjectNameIdentifierPolicyDTO) *map[string]interface{} {
	props := make(map[string]interface{})
	props["descriptionKey"] = dto.GetDescriptionKey()
	props["id"] = dto.GetId()
	props["name"] = dto.GetName()
	props["subjectAttribute"] = dto.GetSubjectAttribute()
	props["type"] = dto.GetType()
	props["@c"] = ".SubjectNameIdentifierPolicyDTO"
	return &props
}

func addFederatedConnection(fcs []FederatedConnectionDTO,
	target string,
	spChannel *InternalSaml2ServiceProviderChannelDTO,
	idpChannel *IdentityProviderChannelDTO) ([]FederatedConnectionDTO, error) {

	// Create new Federated Connection
	var fc FederatedConnectionDTO
	fc.AdditionalProperties = map[string]interface{}{
		"@c": ".FederatedConnectionDTO",
	}
	fc.SetName(target)
	fc.SetIDPChannel(idpChannel)
	fc.SetSPChannel(spChannel)
	fcs = append(fcs, fc)

	return fcs, nil
}

func removeFederatedConnection(fcs []FederatedConnectionDTO, target string) ([]FederatedConnectionDTO, error) {

	var newFcs []FederatedConnectionDTO
	for _, fc := range fcs {
		if fc.GetName() != target {
			newFcs = append(newFcs, fc)
		}
	}

	return newFcs, nil
}
