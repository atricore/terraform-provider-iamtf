package jossoappi

import "fmt"

// AuthenticationMechanismDTO -> BindAuthenticationDTO+DelegatedAuthenticationDTO+AuthenticationMechanism

func (m AuthenticationMechanismDTO) ToBasicAuthn() (*BasicAuthenticationDTO, error) {
	ba := NewBasicAuthenticationDTOInit()

	if m.AdditionalProperties["@c"] != ba.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", m.AdditionalProperties["@c"])
	}

	ba.SetName(m.GetName())
	ba.SetDisplayName(m.GetDisplayName())
	ba.SetPriority(m.GetPriority())
	ba.SetEnabled(AsBool(m.AdditionalProperties["enabled"], false))

	ba.SetHashAlgorithm(AsString(m.AdditionalProperties["hashAlgorithm"], ""))
	ba.SetHashEncoding(AsString(m.AdditionalProperties["hashEncoding"], ""))
	ba.SetIgnoreUsernameCase(AsBool(m.AdditionalProperties["ignoreUsernamecase"], false))

	//ba.SetIgnorePasswordCase(m.AdditionalProperties["ignorePassowordCase"].(bool))

	ba.SetSaltLength(AsInt32(m.AdditionalProperties["saltLength"], 0))
	ba.SetSaltPrefix(AsString(m.AdditionalProperties["saltPrefix"], ""))
	ba.SetSaltSuffix(AsString(m.AdditionalProperties["saltSuffix"], ""))
	//authn.AdditionalProperties["impersonateUserPolicy"]
	ba.SetSimpleAuthnSaml2AuthnCtxClass(AsString(m.AdditionalProperties["simpleAuthnSaml2AuthnCtxClass"], ""))

	return ba, nil
}

func (m AuthenticationMechanismDTO) IsBasicAuthn() bool {
	return m.AdditionalProperties["@c"] == ".BasicAuthenticationDTO"
}
