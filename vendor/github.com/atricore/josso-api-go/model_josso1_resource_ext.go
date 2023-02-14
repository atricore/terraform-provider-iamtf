package jossoappi

func (rs *JOSSO1ResourceDTO) NewActivation(name string) (ActivationDTO, error) {
	var ac ActivationDTO
	ac.AdditionalProperties = make(map[string]interface{})
	ac.AdditionalProperties["@c"] = ".ActivationDTO"
	ac.SetName(name)
	rs.SetActivation(ac)

	return ac, nil
}

func (rs *JOSSO1ResourceDTO) RemoveActivation() bool {
	// Remove an element from : p.IdentityLookups

	if rs.Activation == nil {
		return false
	}

	rs.Activation = nil

	return true
}

func (rs *JOSSO1ResourceDTO) NewServiceConnection(name string) (ServiceConnectionDTO, error) {
	var sc ServiceConnectionDTO
	sc.AdditionalProperties = make(map[string]interface{})
	sc.AdditionalProperties["@c"] = ".ServiceConnectionDTO"
	sc.SetName(name)
	rs.SetServiceConnection(sc)

	return sc, nil
}

func (rs *JOSSO1ResourceDTO) RemoveServiceConnection() bool {
	// Remove an element from : p.IdentityLookups

	if rs.ServiceConnection == nil {
		return false
	}

	rs.ServiceConnection = nil

	return true
}
