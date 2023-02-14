package jossoappi

func (dsvc WindowsIntegratedAuthenticationDTO) toWindowsIntegratedAuth() (*AuthenticationServiceDTO, error) {

	m := NewAuthenticationServiceDTO()

	m.SetId(dsvc.GetId())
	m.SetElementId(dsvc.GetElementId())
	m.SetName(dsvc.GetName())
	m.SetDisplayName(dsvc.GetDisplayName())
	m.SetDescription(dsvc.GetDescription())

	m.AdditionalProperties = make(map[string]interface{})
	m.AdditionalProperties["@c"] = ".WindowsIntegratedAuthenticationDTO"
	m.AdditionalProperties["domain"] = dsvc.GetDomain()
	m.AdditionalProperties["domainController"] = dsvc.GetDomainController()
	m.AdditionalProperties["host"] = dsvc.GetHost()
	m.AdditionalProperties["overwriteKerberosSetup"] = dsvc.GetOverwriteKerberosSetup()
	m.AdditionalProperties["port"] = dsvc.GetPort()
	m.AdditionalProperties["protocol"] = dsvc.GetProtocol()
	m.AdditionalProperties["serviceClass"] = dsvc.GetServiceClass()
	m.AdditionalProperties["serviceName"] = dsvc.GetServiceName()
	
	m.AdditionalProperties["keyTab"] = dsvc.GetKeyTab()

	return m, nil
}

func NewWindowsintegratedAuthnDTOInit() *WindowsIntegratedAuthenticationDTO {
	wia := NewWindowsIntegratedAuthenticationDTO()
	wia.AdditionalProperties = make(map[string]interface{})
	wia.AdditionalProperties["@c"] = ".WindowsIntegratedAuthenticationDTO"

	return wia
}
