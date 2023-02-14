package jossoappi

func (p *VirtualSaml2ServiceProviderDTO) GetSamlR2IDPConfig() (*SamlR2IDPConfigDTO, error) {
	return p.GetConfig().ToSamlR2IDPConfig()
}

func (p *VirtualSaml2ServiceProviderDTO) SetSamlR2IDPConfig(idpCfg *SamlR2IDPConfigDTO) error {
	cfg, err := idpCfg.ToProviderConfig()

	if err != nil {
		return err
	}
	p.SetConfig(*cfg)
	return nil

}

//Add federated connection
func (p *VirtualSaml2ServiceProviderDTO) AddFederatedConnection(target string,
	spChannel *InternalSaml2ServiceProviderChannelDTO,
	idpChannel *IdentityProviderChannelDTO) error {

	fcs, err := addFederatedConnection(p.GetFederatedConnectionsB(), target, spChannel, idpChannel)
	if err != nil {
		return err
	}
	p.SetFederatedConnectionsB(fcs)
	return nil
}

//Remove federated connection
func (p *VirtualSaml2ServiceProviderDTO) RemoveFederatedConnection(target string) error {
	fcs, err := removeFederatedConnection(p.GetFederatedConnectionsB(), target)
	if err != nil {
		return err
	}
	p.SetFederatedConnectionsB(fcs)
	return nil
}
