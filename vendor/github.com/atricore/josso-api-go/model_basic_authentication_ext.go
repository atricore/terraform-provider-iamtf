package jossoappi

func (ba BasicAuthenticationDTO) ToAuthnMechansim() (*AuthenticationMechanismDTO, error) {

	m := NewAuthenticationMechanismDTO()

	m.SetName(ba.GetName())
	m.SetDisplayName(ba.GetDisplayName())
	m.SetPriority(ba.GetPriority())

	m.AdditionalProperties = make(map[string]interface{})

	m.AdditionalProperties["@c"] = ".BasicAuthenticationDTO"
	m.AdditionalProperties["enabled"] = ba.GetEnabled()
	m.AdditionalProperties["hashAlgorithm"] = ba.GetHashAlgorithm()
	m.AdditionalProperties["hashEncoding"] = ba.GetHashEncoding()
	m.AdditionalProperties["ignoreUsernamecase"] = ba.GetIgnoreUsernameCase()
	m.AdditionalProperties["ignorePassowordCase"] = ba.GetIgnorePasswordCase()
	m.AdditionalProperties["saltLength"] = ba.GetSaltLength()
	m.AdditionalProperties["saltPrefix"] = ba.GetSaltPrefix()
	m.AdditionalProperties["saltSuffix"] = ba.GetSaltSuffix()
	//authn.AdditionalProperties["impersonateUserPolicy"]
	m.AdditionalProperties["simpleAuthnSaml2AuthnCtxClass"] = ba.GetSimpleAuthnSaml2AuthnCtxClass()

	return m, nil
}

func NewBasicAuthenticationDTOInit() *BasicAuthenticationDTO {
	ba := NewBasicAuthenticationDTO()
	ba.AdditionalProperties = make(map[string]interface{})
	ba.AdditionalProperties["@c"] = ".BasicAuthenticationDTO"

	return ba

}
