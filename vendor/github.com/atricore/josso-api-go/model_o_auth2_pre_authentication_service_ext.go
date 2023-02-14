package jossoappi

func (dsvc OAuth2PreAuthenticationServiceDTO) toOauth2PreAuthnSvc() (*AuthenticationServiceDTO, error) {

	m := NewAuthenticationServiceDTO()

	m.SetId(dsvc.GetId())
	m.SetElementId(dsvc.GetElementId())
	m.SetName(dsvc.GetName())
	m.SetDisplayName(dsvc.GetDisplayName())
	m.SetDescription(dsvc.GetDescription())

	m.AdditionalProperties = make(map[string]interface{})
	m.AdditionalProperties["@c"] = ".OAuth2PreAuthenticationServiceDTO"
	m.AdditionalProperties["authnService"] = dsvc.GetAuthnService()
	m.AdditionalProperties["externalAuth"] = dsvc.GetExternalAuth()
	m.AdditionalProperties["rememberMe"] = dsvc.GetRememberMe()

	return m, nil
}

func NewOauth2PreAuthnSvcDTOInit() *OAuth2PreAuthenticationServiceDTO {
	oaut2 := NewOAuth2PreAuthenticationServiceDTO()
	oaut2.AdditionalProperties = make(map[string]interface{})
	oaut2.AdditionalProperties["@c"] = ".OAuth2PreAuthenticationServiceDTO"

	return oaut2
}
