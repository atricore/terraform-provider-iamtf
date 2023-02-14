package jossoappi

import (
	"fmt"
)

/*

This builds a specific type for the IdP configuration.

Since the Java model uses polymorphism, we need to transform the abstarct type api.ProviderConfigDTO into api.SamlR2IDPConfigDTO

SAMPLE DUMP OF PROVIDERDTO TYPE

{
	Description:(*string)(nil),
	DisplayName:(*string)(nil),
	ElementId:(*string)(0xc0001a9820),
	Id:(*int64)(0xc0001abb00),
	Name:(*string)(nil),
	AdditionalProperties:map[string]interface {}
	{
		"@c": ".SamlR2IDPConfigDTO",
		"@id": 4,
		"encrypter": 5,
		"signer":map[string]interface {}{
			"@id": 5,
			"certificateAlias": "jcog-saml",
			"displayName":interface {}(nil),
			"elementId": "id8A9B3A5314E128",
			"id": 0, : IGNORED!!
			"keystorePassOnly": false,
			"name": "keystore-547427557",
			"password": "@WSX3edc",
			"privateKeyName": "jcog-saml",
			"privateKeyPassword": "@WSX3edc",
			"store":map[string]interface {}{
				"@id": 6,
				"displayName": "keystore-547427557.jks",
				"elementId":interface {}(nil),
				"id": 0,
				"name": "keystore-547427557",
				"uri": "keystore-547427557.jks",
				"value": "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"
			},
		}
	}
}
*/

func (p *IdentityProviderDTO) GetSamlR2IDPConfig() (*SamlR2IDPConfigDTO, error) {
	return p.GetConfig().ToSamlR2IDPConfig()
}

func (p *IdentityProviderDTO) SetSamlR2IDPConfig(idpCfg *SamlR2IDPConfigDTO) error {
	cfg, err := idpCfg.ToProviderConfig()
	if err != nil {
		return err
	}
	p.SetConfig(*cfg)
	return nil
}

func (p *IdentityProviderDTO) GetIdentityLookup(name string) *IdentityLookupDTO {

	if p.IdentityLookups == nil {
		return nil
	}

	for _, l := range p.IdentityLookups {
		if l.GetName() == name {
			return &l
		}
	}

	return nil

}

func (p *IdentityProviderDTO) AddIdentityLookup(name string) (IdentityLookupDTO, error) {

	// Initialize id lookup dto
	l := NewIdentityLookupDTO()
	l.SetName(name)
	l.AdditionalProperties = make(map[string]interface{})
	l.AdditionalProperties["@c"] = ".IdentityLookupDTO"

	var ls []IdentityLookupDTO

	if p.IdentityLookups == nil {
		ls = make([]IdentityLookupDTO, 0)
	} else {
		if p.GetIdentityLookup(name) != nil {
			return *l, fmt.Errorf("name already in use for identity lookup %s", name)
		}
		ls = p.IdentityLookups
	}

	// Add a new element to : p.IdentityLookups
	ls = append(ls, *l)

	p.IdentityLookups = ls

	return *l, nil

}

// Return true if the element was deleted, false otherwise
func (p *IdentityProviderDTO) RemoveIdentityLookup(name string) bool {
	// Remove an element from : p.IdentityLookups

	if p.IdentityLookups == nil {
		return false
	}

	ls := p.IdentityLookups
	var newLs []IdentityLookupDTO
	deleted := false

	for _, l := range ls {

		if l.GetName() == name {
			deleted = true
			continue
		}

		newLs = append(newLs, l)

	}

	p.IdentityLookups = newLs

	return deleted
}

func (p *IdentityProviderDTO) GetBasicAuthns() ([]*BasicAuthenticationDTO, error) {

	bas := make([]*BasicAuthenticationDTO, 0)

	for _, m := range p.GetAuthenticationMechanisms() {

		if m.IsBasicAuthn() {
			ba, err := m.ToBasicAuthn()
			if err != nil {
				return bas, err
			}
			bas = append(bas, ba)
		}
	}

	return bas, nil
}

func (p *IdentityProviderDTO) AddBasicAuthns(ms []*BasicAuthenticationDTO) error {
	for _, ba := range ms {
		p.AddBasicAuthn(ba)
	}

	return nil
}

func (p *IdentityProviderDTO) AddBasicAuthn(ba *BasicAuthenticationDTO) (*AuthenticationMechanismDTO, error) {
	m, err := ba.ToAuthnMechansim()
	if err != nil {
		return m, err
	}
	p.SetAuthenticationMechanisms(append(p.GetAuthenticationMechanisms(), *m))

	return m, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *IdentityProviderDTO) GetDirectoryAuthnSvc() ([]*DirectoryAuthenticationServiceDTO, error) {

	r := make([]*DirectoryAuthenticationServiceDTO, 0)

	for _, m := range p.GetAuthenticationMechanisms() {
		da := m.GetDelegatedAuthentication()
		as := da.GetAuthnService()
		if as.IsDirectoryAuthnSvs() {
			a, err := as.ToDirectoryAuthnSvc()
			if err != nil {
				return r, err
			}
			r = append(r, a)
		}
	}

	return r, nil
}

func (p *IdentityProviderDTO) AddDirectoryAuthnsSvc(ms []*DirectoryAuthenticationServiceDTO, pr int32) error {
	for _, das := range ms {
		p.AddDirectoryAuthnSvc(das, pr)
	}

	return nil
}

func (p *IdentityProviderDTO) AddDirectoryAuthnSvc(das *DirectoryAuthenticationServiceDTO, pr int32) (*AuthenticationServiceDTO, error) {
	m, err := das.toAuthnSvc()
	if err != nil {
		return m, err
	}
	dauthDTO := NewDelegatedAuthenticationDTO()
	dauthDTO.AdditionalProperties = make(map[string]interface{})
	dauthDTO.AdditionalProperties["@c"] = ".DelegatedAuthenticationDTO"
	dauthDTO.SetAuthnService(*m)

	authMechDTO := NewAuthenticationMechanismDTO()
	authMechDTO.SetPriority(pr)
	authMechDTO.SetDelegatedAuthentication(*dauthDTO)
	authMechDTO.AdditionalProperties = make(map[string]interface{})
	authMechDTO.AdditionalProperties["@c"] = ".BindAuthenticationDTO"

	p.GetAuthenticationMechanisms()
	authMech := append(p.GetAuthenticationMechanisms(), *authMechDTO)
	p.SetAuthenticationMechanisms(authMech)
	return m, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *IdentityProviderDTO) GetClientCertAuthnSvs() ([]*ClientCertAuthnServiceDTO, error) {

	r := make([]*ClientCertAuthnServiceDTO, 0)

	for _, m := range p.GetAuthenticationMechanisms() {
		da := m.GetDelegatedAuthentication()
		as := da.GetAuthnService()
		if as.IsClientCertAuthnSvs() {
			a, err := as.ToClientCertAuthnSvc()
			if err != nil {
				return r, err
			}
			r = append(r, a)
		}
	}

	return r, nil
}

func (p *IdentityProviderDTO) AddClientCertAuthnsSvs(ms []*ClientCertAuthnServiceDTO, pr int32) error {
	for _, cas := range ms {
		p.AddClientCertAuthnSvs(cas, pr)
	}

	return nil
}

func (p *IdentityProviderDTO) AddClientCertAuthnSvs(cas *ClientCertAuthnServiceDTO, pr int32) (*AuthenticationServiceDTO, error) {
	m, err := cas.toClientCertAuthSvc()
	if err != nil {
		return m, err
	}
	dauthDTO := NewDelegatedAuthenticationDTO()
	dauthDTO.AdditionalProperties = make(map[string]interface{})
	dauthDTO.AdditionalProperties["@c"] = ".DelegatedAuthenticationDTO"
	dauthDTO.SetAuthnService(*m)

	authMechDTO := NewAuthenticationMechanismDTO()
	authMechDTO.SetPriority(pr)
	authMechDTO.SetDelegatedAuthentication(*dauthDTO)
	authMechDTO.AdditionalProperties = make(map[string]interface{})
	authMechDTO.AdditionalProperties["@c"] = ".ClientCertAuthenticationDTO"

	p.GetAuthenticationMechanisms()
	authMech := append(p.GetAuthenticationMechanisms(), *authMechDTO)
	p.SetAuthenticationMechanisms(authMech)
	return m, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *IdentityProviderDTO) GetWindowsIntegratedAuthn() ([]*WindowsIntegratedAuthenticationDTO, error) {

	r := make([]*WindowsIntegratedAuthenticationDTO, 0)

	for _, m := range p.GetAuthenticationMechanisms() {
		da := m.GetDelegatedAuthentication()
		as := da.GetAuthnService()
		if as.IsWindowsIntegratedAuthn() {
			a, err := as.ToWindowsIntegratedAuthn()
			if err != nil {
				return r, err
			}
			r = append(r, a)
		}
	}

	return r, nil
}

func (p *IdentityProviderDTO) AddWindowsIntegratedAuthns(ms []*WindowsIntegratedAuthenticationDTO, pr int32) error {
	for _, wia := range ms {
		p.AddWindowsIntegratedAuthn(wia, pr)
	}

	return nil
}

func (p *IdentityProviderDTO) AddWindowsIntegratedAuthn(wia *WindowsIntegratedAuthenticationDTO, pr int32) (*AuthenticationServiceDTO, error) {
	m, err := wia.toWindowsIntegratedAuth()
	if err != nil {
		return m, err
	}
	dauthDTO := NewDelegatedAuthenticationDTO()
	dauthDTO.AdditionalProperties = make(map[string]interface{})
	dauthDTO.AdditionalProperties["@c"] = ".DelegatedAuthenticationDTO"
	dauthDTO.SetAuthnService(*m)

	authMechDTO := NewAuthenticationMechanismDTO()
	authMechDTO.SetPriority(pr)
	authMechDTO.SetDelegatedAuthentication(*dauthDTO)
	authMechDTO.AdditionalProperties = make(map[string]interface{})
	authMechDTO.AdditionalProperties["@c"] = ".WindowsAuthenticationDTO"

	p.GetAuthenticationMechanisms()
	authMech := append(p.GetAuthenticationMechanisms(), *authMechDTO)
	p.SetAuthenticationMechanisms(authMech)
	return m, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *IdentityProviderDTO) GetOauth2PreAuthnSvs() ([]*OAuth2PreAuthenticationServiceDTO, error) {

	r := make([]*OAuth2PreAuthenticationServiceDTO, 0)

	for _, m := range p.GetAuthenticationMechanisms() {
		da := m.GetDelegatedAuthentication()
		as := da.GetAuthnService()
		if as.IsOauth2PreAuthnSvc() {
			a, err := as.ToOauth2PreAuthnSvs()
			if err != nil {
				return r, err
			}
			r = append(r, a)
		}
	}

	return r, nil
}

func (p *IdentityProviderDTO) AddOauth2PreAuthnsSvs(ms []*OAuth2PreAuthenticationServiceDTO, pr int32) error {
	for _, oaut2 := range ms {
		p.AddOauth2PreAuthnSvs(oaut2, pr)
	}

	return nil
}

func (p *IdentityProviderDTO) AddOauth2PreAuthnSvs(oaut2 *OAuth2PreAuthenticationServiceDTO, pr int32) (*AuthenticationServiceDTO, error) {
	m, err := oaut2.toOauth2PreAuthnSvc()
	if err != nil {
		return m, err
	}
	dauthDTO := NewDelegatedAuthenticationDTO()
	dauthDTO.AdditionalProperties = make(map[string]interface{})
	dauthDTO.AdditionalProperties["@c"] = ".DelegatedAuthenticationDTO"
	dauthDTO.SetAuthnService(*m)

	authMechDTO := NewAuthenticationMechanismDTO()
	authMechDTO.SetPriority(pr)
	authMechDTO.SetDelegatedAuthentication(*dauthDTO)
	authMechDTO.AdditionalProperties = make(map[string]interface{})
	authMechDTO.AdditionalProperties["@c"] = ".OAuth2PreAuthenticationDTO"

	p.GetAuthenticationMechanisms()
	authMech := append(p.GetAuthenticationMechanisms(), *authMechDTO)
	p.SetAuthenticationMechanisms(authMech)
	return m, nil
}
