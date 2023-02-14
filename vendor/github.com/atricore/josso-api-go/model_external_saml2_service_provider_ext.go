package jossoappi

func (p *ExternalSaml2ServiceProviderDTO) AddFederatedConnection(target string,
	spChannel *InternalSaml2ServiceProviderChannelDTO,
	idpChannel *IdentityProviderChannelDTO) error {

	fcs, err := addFederatedConnection(p.GetFederatedConnectionsB(), target, spChannel, idpChannel)
	if err != nil {
		return err
	}
	p.SetFederatedConnectionsB(fcs)
	return nil
}

func (p *ExternalSaml2ServiceProviderDTO) RemoveFederatedConnection(target string) error {
	fcs, err := removeFederatedConnection(p.GetFederatedConnectionsB(), target)
	if err != nil {
		return err
	}
	p.SetFederatedConnectionsB(fcs)
	return nil
}
