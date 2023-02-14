package jossoappi

func (rs *SharepointResourceDTO) NewServiceConnection(name string) (ServiceConnectionDTO, error) {
	var sc ServiceConnectionDTO
	sc.AdditionalProperties = make(map[string]interface{})
	sc.AdditionalProperties["@c"] = ".ServiceConnectionDTO"
	sc.SetName(name)
	rs.SetServiceConnection(sc)

	return sc, nil
}
