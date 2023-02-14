package jossoappi

func (dsvc ClientCertAuthnServiceDTO) toClientCertAuthSvc() (*AuthenticationServiceDTO, error) {

	m := NewAuthenticationServiceDTO()

	m.SetId(dsvc.GetId())
	m.SetElementId(dsvc.GetElementId())
	m.SetName(dsvc.GetName())
	m.SetDisplayName(dsvc.GetDisplayName())
	m.SetDescription(dsvc.GetDescription())

	m.AdditionalProperties = make(map[string]interface{})
	m.AdditionalProperties["@c"] = ".ClientCertAuthnServiceDTO"
	m.AdditionalProperties["clrEnabled"] = dsvc.GetClrEnabled()
	m.AdditionalProperties["crlRefreshSeconds"] = dsvc.GetCrlRefreshSeconds()
	m.AdditionalProperties["crlUrl"] = dsvc.GetCrlUrl()
	m.AdditionalProperties["ocspEnabled"] = dsvc.GetOcspEnabled()
	m.AdditionalProperties["ocspServer"] = dsvc.GetOcspServer()
	m.AdditionalProperties["ocspserver"] = dsvc.GetOcspserver()
	m.AdditionalProperties["uid"] = dsvc.GetUid()

	return m, nil
}

func NewClientCertAuthnSvcDTOInit() *ClientCertAuthnServiceDTO {
	cas := NewClientCertAuthnServiceDTO()
	cas.AdditionalProperties = make(map[string]interface{})
	cas.AdditionalProperties["@c"] = ".ClientCertAuthnServiceDTO"

	return cas
}
