package jossoappi

func (dsvc DirectoryAuthenticationServiceDTO) toAuthnSvc() (*AuthenticationServiceDTO, error) {

	m := NewAuthenticationServiceDTO()

	m.SetId(dsvc.GetId())
	m.SetElementId(dsvc.GetElementId())
	m.SetName(dsvc.GetName())
	m.SetDisplayName(dsvc.GetDisplayName())
	m.SetDescription(dsvc.GetDescription())

	m.AdditionalProperties = make(map[string]interface{})
	m.AdditionalProperties["@c"] = ".DirectoryAuthenticationServiceDTO"
	m.AdditionalProperties["initialContextFactory"] = dsvc.GetInitialContextFactory()
	m.AdditionalProperties["providerUrl"] = dsvc.GetProviderUrl()
	m.AdditionalProperties["performDnSearch"] = dsvc.GetPerformDnSearch()
	m.AdditionalProperties["passwordPolicy"] = dsvc.GetPasswordPolicy()
	m.AdditionalProperties["securityAuthentication"] = dsvc.GetSecurityAuthentication()
	m.AdditionalProperties["usersCtxDN"] = dsvc.GetUsersCtxDN()
	m.AdditionalProperties["principalUidAttributeID"] = dsvc.GetPrincipalUidAttributeID()
	m.AdditionalProperties["securityPrincipal"] = dsvc.GetSecurityPrincipal()
	m.AdditionalProperties["securityCredential"] = dsvc.GetSecurityCredential()
	m.AdditionalProperties["ldapSearchScope"] = dsvc.GetLdapSearchScope()
	m.AdditionalProperties["simpleAuthnSaml2AuthnCtxClass"] = dsvc.GetSimpleAuthnSaml2AuthnCtxClass()
	m.AdditionalProperties["referrals"] = dsvc.GetReferrals()
	m.AdditionalProperties["includeOperationalAttributes"] = dsvc.GetIncludeOperationalAttributes()

	return m, nil
}

func NewDirectoryAuthnSvcDTOInit() *DirectoryAuthenticationServiceDTO {
	das := NewDirectoryAuthenticationServiceDTO()
	das.AdditionalProperties = make(map[string]interface{})
	das.AdditionalProperties["@c"] = ".DirectoryAuthenticationServiceDTO"

	return das
}
