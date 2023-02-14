package jossoappi

import (
	"fmt"
)

func (p *InternalSaml2ServiceProviderDTO) GetSamlR2SPConfig() (*SamlR2SPConfigDTO, error) {
	return p.GetConfig().ToSamlR2SPConfig()
}

func (p *InternalSaml2ServiceProviderDTO) SetSamlR2SPConfig(spCfg *SamlR2SPConfigDTO) error {
	cfg, err := spCfg.ToProviderConfig()
	if err != nil {
		return err
	}
	p.SetConfig(*cfg)
	return nil
}
func (p *InternalSaml2ServiceProviderDTO) GetIdentityLookup(name string) *IdentityLookupDTO {
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

func (p *InternalSaml2ServiceProviderDTO) AddIdentityLookup(name string) (IdentityLookupDTO, error) {

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
func (p *InternalSaml2ServiceProviderDTO) RemoveIdentityLookup(name string) bool {
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
