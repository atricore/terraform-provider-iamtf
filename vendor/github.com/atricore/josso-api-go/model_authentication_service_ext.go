package jossoappi

import "fmt"

// AuthenticationServiceDTO -> DirectoryAuthenticationServiceDTO
func (svc AuthenticationServiceDTO) ToDirectoryAuthnSvc() (*DirectoryAuthenticationServiceDTO, error) {
	das := NewDirectoryAuthnSvcDTOInit()

	if svc.AdditionalProperties["@c"] != das.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", das.AdditionalProperties["@c"])
	}

	das.SetId(svc.GetId())
	das.SetElementId(svc.GetElementId())
	das.SetName(svc.GetName())
	das.SetDisplayName(svc.GetDisplayName())
	das.SetDescription(svc.GetDescription())
	das.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	das.SetInitialContextFactory(AsString(svc.AdditionalProperties["initialContextFactory"], ""))
	das.SetProviderUrl(AsString(svc.AdditionalProperties["providerUrl"], ""))
	das.SetPerformDnSearch(AsBool(svc.AdditionalProperties["performDnSearch"], false))
	das.SetPasswordPolicy(AsString(svc.AdditionalProperties["passwordPolicy"], ""))
	das.SetSecurityAuthentication(AsString(svc.AdditionalProperties["securityAuthentication"], ""))
	das.SetUsersCtxDN(AsString(svc.AdditionalProperties["usersCtxDN"], ""))
	das.SetPrincipalUidAttributeID(AsString(svc.AdditionalProperties["principalUidAttributeID"], ""))
	das.SetSecurityPrincipal(AsString(svc.AdditionalProperties["securityPrincipal"], ""))
	das.SetSecurityCredential(AsString(svc.AdditionalProperties["securityCredential"], ""))
	das.SetLdapSearchScope(AsString(svc.AdditionalProperties["ldapSearchScope"], ""))
	das.SetSimpleAuthnSaml2AuthnCtxClass(AsString(svc.AdditionalProperties["simpleAuthnSaml2AuthnCtxClass"], ""))
	das.SetReferrals(AsString(svc.AdditionalProperties["referrals"], ""))
	das.SetIncludeOperationalAttributes(AsBool(svc.AdditionalProperties["includeOperationalAttributes"], false))

	return das, nil
}

// AuthenticationServiceDTO -> ClientCertAuthnServiceDTO
func (svc AuthenticationServiceDTO) ToClientCertAuthnSvc() (*ClientCertAuthnServiceDTO, error) {
	cas := NewClientCertAuthnSvcDTOInit()

	if svc.AdditionalProperties["@c"] != cas.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", cas.AdditionalProperties["@c"])
	}

	cas.SetId(svc.GetId())
	cas.SetElementId(svc.GetElementId())
	cas.SetName(svc.GetName())
	cas.SetDisplayName(svc.GetDisplayName())
	cas.SetDescription(svc.GetDescription())
	cas.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	cas.SetClrEnabled(AsBool(svc.AdditionalProperties["clrEnabled"], false))
	cas.SetCrlRefreshSeconds(AsInt32(svc.AdditionalProperties["crlRefreshSeconds"], 0))
	cas.SetCrlUrl(AsString(svc.AdditionalProperties["crlUrl"], ""))
	cas.SetOcspEnabled(AsBool(svc.AdditionalProperties["ocspEnabled"], false))
	cas.SetOcspServer(AsString(svc.AdditionalProperties["ocspServer"], ""))
	cas.SetOcspserver(AsString(svc.AdditionalProperties["ocspserver"], ""))
	cas.SetUid(AsString(svc.AdditionalProperties["uid"], ""))

	return cas, nil
}

// AuthenticationServiceDTO -> WindowsIntegratedAuthenticationDTO
func (svc AuthenticationServiceDTO) ToWindowsIntegratedAuthn() (*WindowsIntegratedAuthenticationDTO, error) {
	wia := NewWindowsintegratedAuthnDTOInit()

	if svc.AdditionalProperties["@c"] != wia.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", wia.AdditionalProperties["@c"])
	}
	wia.SetId(svc.GetId())
	wia.SetElementId(svc.GetElementId())
	wia.SetName(svc.GetName())
	wia.SetDisplayName(svc.GetDisplayName())
	wia.SetDescription(svc.GetDescription())
	wia.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	// TODO : wia.SetKeyTab ?
	// ktMap := svc.AdditionalProperties["keytab"].(map[string]interface{})
	// ktValue := ktMap["value"]

	wia.SetDomain(AsString(svc.AdditionalProperties["domain"], ""))
	wia.SetDomainController(AsString(svc.AdditionalProperties["domainController"], ""))
	wia.SetHost(AsString(svc.AdditionalProperties["host"], ""))
	wia.SetOverwriteKerberosSetup(AsBool(svc.AdditionalProperties["overwriteKerberosSetup"], false))
	wia.SetPort(AsInt32(svc.AdditionalProperties["port"], 0))
	wia.SetProtocol(AsString(svc.AdditionalProperties["protocol"], ""))
	wia.SetServiceClass(AsString(svc.AdditionalProperties["serviceClass"], ""))
	wia.SetServiceName(AsString(svc.AdditionalProperties["serviceName"], ""))

	if _, ok := svc.AdditionalProperties["keyTab"].(map[string]interface{}); ok {

		ktMap := svc.AdditionalProperties["keyTab"].(map[string]interface{})
		kt := NewResourceDTO()
		kt.SetValue(AsString(ktMap["value"], ""))
		wia.SetKeyTab(*kt)
	}

	return wia, nil
}

// AuthenticationServiceDTO -> OAuth2PreAuthenticationServiceDTO
func (svc AuthenticationServiceDTO) ToOauth2PreAuthnSvs() (*OAuth2PreAuthenticationServiceDTO, error) {
	oaut2 := NewOauth2PreAuthnSvcDTOInit()

	if svc.AdditionalProperties["@c"] != oaut2.AdditionalProperties["@c"] {
		return nil, fmt.Errorf("invalid authentication mechanism java class %s", oaut2.AdditionalProperties["@c"])
	}

	oaut2.SetId(svc.GetId())
	oaut2.SetElementId(svc.GetElementId())
	oaut2.SetName(svc.GetName())
	oaut2.SetDisplayName(svc.GetDisplayName())
	oaut2.SetDescription(svc.GetDescription())
	oaut2.SetDelegatedAuthentications(svc.GetDelegatedAuthentications())

	oaut2.SetAuthnService(AsString(svc.AdditionalProperties["authnService"], ""))
	oaut2.SetExternalAuth((AsBool(svc.AdditionalProperties["externalAuth"], false)))
	oaut2.SetRememberMe(AsBool(svc.AdditionalProperties["rememberMe"], false))

	return oaut2, nil
}

func (m AuthenticationServiceDTO) IsDirectoryAuthnSvs() bool {
	return m.AdditionalProperties["@c"] == ".DirectoryAuthenticationServiceDTO"
}

func (m AuthenticationServiceDTO) IsClientCertAuthnSvs() bool {
	return m.AdditionalProperties["@c"] == ".ClientCertAuthnServiceDTO"
}

func (m AuthenticationServiceDTO) IsWindowsIntegratedAuthn() bool {
	return m.AdditionalProperties["@c"] == ".WindowsIntegratedAuthenticationDTO"
}

func (m AuthenticationServiceDTO) IsOauth2PreAuthnSvc() bool {
	return m.AdditionalProperties["@c"] == ".OAuth2PreAuthenticationServiceDTO"
}
