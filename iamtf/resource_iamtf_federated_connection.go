package iamtf

import (
	"fmt"

	api "github.com/atricore/josso-api-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func spSamlSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Computed:    true,
		Description: "SP SAML 2 settings",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"account_linkage": {
					Type:             schema.TypeString,
					Description:      "account linkage: which attribute to use as UID from the IdP.",
					ValidateDiagFunc: stringInSlice([]string{"ONE_TO_ONE", "EMAIL", "UID", "CUSTOM"}),
					Computed:         true,
					Optional:         true,
				},
				"message_ttl": {
					Type:        schema.TypeInt,
					Optional:    true,
					Computed:    true,
					Description: "SAML message time to live",
				},
				"message_ttl_tolerance": {
					Type:        schema.TypeInt,
					Optional:    true,
					Computed:    true,
					Description: "SAML message time to live tolerance",
				},
				"identity_mapping": {
					Type:             schema.TypeString,
					Description:      "AppAgent_idp Resource identity_mapping",
					ValidateDiagFunc: stringInSlice([]string{"LOCAL", "REMOTE", "MERGED", "CUSTOM"}),
					Computed:         true,
					Optional:         true,
				},
				"sign_requests": {
					Type:        schema.TypeBool,
					Optional:    true,
					Computed:    true,
					Description: "sign requests issued to IdPs",
				},
				"sign_authentication_requests": {
					Type:        schema.TypeBool,
					Optional:    true,
					Computed:    true,
					Description: "sign authentication requests issued to IdPs",
				},
				"signature_hash": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
					Description: "saml signature hash algorithm",
				},
				"want_assertion_signed": {
					Type:        schema.TypeBool,
					Optional:    true,
					Computed:    true,
					Description: "require signed assertions from IdPs",
				},
			},
		},
	}

}

// This adds federated connection specific attributes to SP saml information
func idpConnectionSchema() *schema.Schema {

	idpConn := &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Computed:    true,
		Description: "SP to IDP SAML 2 settings",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "name of the trusted IdP",
				},
				"is_preferred": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "identifies this IdP as the preferred one (only one IdP must be set to preferred)",
				},
				"saml2": spSamlSchema(),
			},
		},
	}

	return idpConn
}

// Name of the parent provider
func convertIdPFederatedConnectionsMapArrToDTOs(sp *api.InternalSaml2ServiceProviderDTO, d *schema.ResourceData, idp interface{}) ([]api.FederatedConnectionDTO, error) {
	result := make([]api.FederatedConnectionDTO, 0)
	ls, ok := idp.([]interface{})
	if !ok {
		return result, fmt.Errorf("invalid type: %T", idp)
	}

	for idpIdx, v := range ls {
		// 1. For each IdP(terraform), create a FederatedConnection
		// Store all connections in sp.FederatedConnectionsB array.
		// The idp name is to be used as the federated connection name
		// 2. Each connection will have a FederatedChannel,
		// Create an IDPChannel and use convertion function to get FederatecChannel
		// Store the FederatedChannel in the federatedConnection.federatedChannelB member/element/var
		m, ok := v.(map[string]interface{})
		if !ok {
			return result, fmt.Errorf("invalid element type: %T", v)
		}

		// build new federatedconnectionDTO
		c := api.NewFederatedConnectionDTO()
		c.AdditionalProperties = make(map[string]interface{})

		c.SetName(m["name"].(string))
		c.AdditionalProperties["@c"] = ".FederatedConnectionDTO"

		// from federatedconnectionDTO.Channelb values
		idpChannel := api.NewIdentityProviderChannelDTO()
		// Assing values for preferred option
		idpChannel.SetPreferred(m["is_preferred"].(bool))

		saml2_m, err := asTFMapSingle(m["saml2"])
		if err != nil {
			return result, err
		}

		// build new identityMappingPolicyDTO
		it := sp.GetIdentityMappingPolicy()
		if it.MappingType == nil {
			it.SetMappingType("ONE_TO_ONE")
		}

		im := api.NewIdentityMappingPolicyDTO()
		im.AdditionalProperties = make(map[string]interface{})
		im.AdditionalProperties["@c"] = ".IdentityMappingPolicyDTO"
		im.SetName(fmt.Sprintf("%s-identity-mapping", sp.GetName()))
		im.SetMappingType(GetAsString(d, fmt.Sprintf("idp.%d.saml2.0.identity_mapping", idpIdx), it.GetMappingType()))
		idpChannel.SetIdentityMappingPolicy(*im)

		// build new accountLinkagePolicyDTO
		at := sp.GetAccountLinkagePolicy()
		if at.LinkEmitterType == nil {
			at.SetLinkEmitterType("REMOTE")
		}
		al := api.NewAccountLinkagePolicyDTO()
		al.AdditionalProperties = make(map[string]interface{})
		al.AdditionalProperties["@c"] = ".AccountLinkagePolicyDTO"
		al.SetName(fmt.Sprintf("%s-account-linkage", sp.GetName()))
		al.SetLinkEmitterType(GetAsString(d, fmt.Sprintf("idp.%d.saml2.0.account_linkage", idpIdx), at.GetLinkEmitterType()))
		idpChannel.SetAccountLinkagePolicy(*al)

		if len(saml2_m) > 0 {

			idpChannel.SetOverrideProviderSetup(true)

			idpChannel.SetMessageTtl(GetAsInt32(d, fmt.Sprintf("idp.%d.saml2.0.message_ttl", idpIdx), sp.GetMessageTtl()))
			idpChannel.SetMessageTtlTolerance(GetAsInt32(d, fmt.Sprintf("idp.%d.saml2.0.message_ttl_tolerance", idpIdx), sp.GetMessageTtlTolerance()))
			// NOT SUPPORTED BY SERVER :idpChannel.SetSignRequests(api.AsBool(saml2_m["sign_requests"], true))
			idpChannel.SetSignAuthenticationRequests(
				GetAsBool(d, fmt.Sprintf("idp.%d.saml2.0.sign_authentication_requests", idpIdx), sp.GetSignAuthenticationRequests()))
			idpChannel.SetWantAssertionSigned(
				GetAsBool(d, fmt.Sprintf("idp.%d.saml2.0.want_assertion_signed", idpIdx), sp.GetWantAssertionSigned()))

			idpChannel.SetSignatureHash(GetAsString(d, fmt.Sprintf("idp.%d.saml2.0.signature_hash", idpIdx), sp.GetSignatureHash()))

		} else {
			idpChannel.SetOverrideProviderSetup(false)
		}

		c.SetIDPChannel(idpChannel)

		result = append(result, *c)
	}
	return result, nil
}

func convertIdPFederatedConnectionsToMapArr(fcs []api.FederatedConnectionDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	for _, fc := range fcs {

		idpChannel, err := fc.GetIDPChannel()
		if err != nil {
			return result, err
		}
		idp_map := map[string]interface{}{
			"name":         fc.GetName(),
			"is_preferred": idpChannel.GetPreferred(),
		}

		if idpChannel.GetOverrideProviderSetup() {
			al := idpChannel.GetAccountLinkagePolicy()
			im := idpChannel.GetIdentityMappingPolicy()
			// Array of maps
			saml2_map :=
				[]map[string]interface{}{{
					"account_linkage":              al.GetLinkEmitterType(),
					"message_ttl":                  int(idpChannel.GetMessageTtl()),
					"message_ttl_tolerance":        int(idpChannel.GetMessageTtlTolerance()),
					"identity_mapping":             im.GetMappingType(),
					"sign_authentication_requests": idpChannel.GetSignAuthenticationRequests(),
					"signature_hash":               idpChannel.GetSignatureHash(),
					"want_assertion_signed":        idpChannel.GetWantAssertionSigned(),
					//NOT SUPPORTED BY SERVER "sign_requests":                idpChannel.GetSignRequests(),
				}}
			idp_map["saml2"] = saml2_map
		}
		result = append(result, idp_map)

	}

	return result, nil
}

func idpSamlSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Computed:    true,
		MaxItems:    1,
		MinItems:    0,
		Description: "IDP SAML2 protocol settings",

		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"want_authn_req_signed": {
					Type:        schema.TypeBool,
					Description: "want authn requests signed by SPs",
					Optional:    true,
					Default:     false,
				},
				"want_req_signed": {
					Type:        schema.TypeBool,
					Description: "want requests signed by SPs",
					Optional:    true,
					Default:     false,
				},
				"sign_reqs": {
					Type:        schema.TypeBool,
					Description: "sign all requests to SPs",
					Optional:    true,
					Default:     true,
				},
				"signature_hash": {
					Type:             schema.TypeString,
					Description:      "signature hash algorithm",
					Optional:         true,
					ValidateDiagFunc: stringInSlice([]string{"SHA1", "SHA256", "SHA384", "SHA512"}),
					Default:          "SHA256",
				},
				"encrypt_algorithm": {
					Type:             schema.TypeString,
					Description:      "encrypt assertion algorithm",
					Optional:         true,
					ValidateDiagFunc: stringInSlice([]string{"NONE", "AES-128", "AES-256", "AES-3DES"}),
					Default:          "NONE",
				},
				/*
					"metadata_endpoint": {
						Type:        schema.TypeBool,
						Description: "enable metadata endpoint",
						Optional:    true,
						Default:     true,
					},
				*/
				"message_ttl": {
					Type:        schema.TypeInt,
					Description: "message ttl (sec)",
					Optional:    true,
					Computed:    true,
				},
				"message_ttl_tolerance": {
					Type:        schema.TypeInt,
					Description: "message ttl tolerance (sec)",
					Optional:    true,
					Computed:    true,
				},
			},
		},
	}
}

func spConnectionSchema() *schema.Schema {
	spConn := &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Computed:    true,
		Description: "IDP to SP SAML 2 settings",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "name of the trusted sp",
				},
				"saml2": idpSamlSchema(),
			},
		},
	}

	return spConn
}

func convertIdPSaml2DTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	/** TODO : Add this logic to the server
	encAlg := idp.GetEncryptAssertionAlgorithm()
	if encAlg == "" {
		encAlg = "NONE"
	}
	*/

	saml2_map := map[string]interface{}{
		"want_authn_req_signed": idp.GetWantAuthnRequestsSigned(),
		"want_req_signed":       idp.GetWantSignedRequests(),
		"sign_reqs":             idp.GetSignRequests(),
		"signature_hash":        idp.GetSignatureHash(),
		"encrypt_algorithm":     idp.GetEncryptAssertionAlgorithm(),
		//		"metadata_endpoint":     idp.GetEnableMetadataEndpoint(),
		"message_ttl":           int(idp.GetMessageTtl()),
		"message_ttl_tolerance": int(idp.GetMessageTtlTolerance()),
	}
	result = append(result, saml2_map)

	return result, nil
}

func convertIdPSaml2MapArrToDTO(saml2_arr interface{}, idp *api.IdentityProviderDTO) error {
	saml2_map, err := asTFMapSingle(saml2_arr)
	if err != nil {
		return err
	}

	if saml2_map == nil {
		return nil
	}

	idp.SetWantAuthnRequestsSigned(saml2_map["want_authn_req_signed"].(bool))
	idp.SetWantSignedRequests(saml2_map["want_req_signed"].(bool))
	idp.SetSignRequests(saml2_map["sign_reqs"].(bool))
	idp.SetSignatureHash(saml2_map["signature_hash"].(string))
	idp.SetEncryptAssertionAlgorithm(saml2_map["encrypt_algorithm"].(string))
	//idp.SetEnableMetadataEndpoint(saml2_map["metadata_endpoint"].(bool))
	idp.SetEnableMetadataEndpoint(true)
	idp.SetMessageTtl(int32(saml2_map["message_ttl"].(int)))
	idp.SetMessageTtlTolerance(int32(saml2_map["message_ttl_tolerance"].(int)))

	return nil
}

func convertSPFederatedConnectionsMapArrToDTOs(idp *api.IdentityProviderDTO, d *schema.ResourceData, sp interface{}) ([]api.FederatedConnectionDTO, error) {
	result := make([]api.FederatedConnectionDTO, 0)
	ls, ok := sp.([]interface{})
	if !ok {
		return result, fmt.Errorf("invalid type: %T", sp)
	}

	for spIdx, v := range ls {
		m, ok := v.(map[string]interface{})
		if !ok {
			return result, fmt.Errorf("invalid element type: %T", v)
		}

		// build new federatedconnectionDTO
		c := api.NewFederatedConnectionDTO()
		c.SetName(m["name"].(string))

		c.AdditionalProperties = make(map[string]interface{})
		c.AdditionalProperties["@c"] = ".FederatedConnectionDTO"

		// from federatedconnectionDTO.Channelb values
		spChannel := api.NewInternalSaml2ServiceProviderChannelDTO()

		saml2_m, err := asTFMapSingle(m["saml2"])
		if err != nil {
			return result, err
		}

		if len(saml2_m) > 0 {

			spChannel.SetOverrideProviderSetup(true)

			spChannel.SetWantAuthnRequestsSigned(GetAsBool(d, fmt.Sprintf("%d", spIdx), idp.GetWantAuthnRequestsSigned()))
			// NOT SUPPORETD BY SERVER :spChannel.SetWantSignedRequests(api.AsBool(saml2_m["want_req_signed"], true))
			// NOT SUPPORETD BY SERVER :spChannel.SetSignRequests(api.AsBool(saml2_m["sign_reqs"], true))
			spChannel.SetSignatureHash(GetAsString(d, fmt.Sprintf("sp.%d.saml2.0.signature_hash", spIdx), idp.GetSignatureHash()))
			spChannel.SetEncryptAssertionAlgorithm(GetAsString(d, fmt.Sprintf("sp.%d.saml2.0.encrypt_algorithm", spIdx), idp.GetEncryptAssertionAlgorithm()))
			spChannel.SetMessageTtl(GetAsInt32(d, fmt.Sprintf("sp.%d.saml2.0.message_ttl", spIdx), idp.GetMessageTtl()))
			spChannel.SetMessageTtlTolerance(GetAsInt32(d, fmt.Sprintf("sp.%d.saml2.0.message_ttl_tolerance", spIdx), idp.GetMessageTtlTolerance()))

			// TODO : spChannel.SetAttributeProfile()

		} else {
			spChannel.SetOverrideProviderSetup(false)
		}

		// Attribute profile
		af := api.NewBasicAttributeProfileDTOInit(fmt.Sprintf("%s-attr", idp.GetName()))
		ap, err := af.ToAttrProfile()
		if err != nil {
			return result, err
		}
		spChannel.SetAttributeProfile(*ap)

		// Subject Name ID policy : TODO , not working , error on the server while unmarshalling ?!
		/*
			sid := api.NewSubjectNameIdentifierPolicyDTO()
			sid.AdditionalProperties = make(map[string]interface{})
			sid.AdditionalProperties["@c"] = ".CustomNameIdentifierPolicyDTO"
			sid.SetType("PRINCIPAL")
			sid.SetName(fmt.Sprintf("%s-subject-id", idp.GetName()))

			spChannel.SetSubjectNameIDPolicy(*sid)
		*/
		c.SetSPChannel(spChannel)

		result = append(result, *c)
	}

	return result, nil
}

func convertSPFederatedConnectionsToMapArr(fcs []api.FederatedConnectionDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	for _, fc := range fcs {

		spChannel, err := fc.GetSPChannel()
		if err != nil {
			return result, err
		}
		sp_map := map[string]interface{}{
			"name": fc.GetName(),
		}

		if spChannel.GetOverrideProviderSetup() {

			saml2_map := map[string]interface{}{
				"want_authn_req_signed": spChannel.GetWantAuthnRequestsSigned(),
				// NOT SUPPORETD BY SERVER "want_req_signed":     spChannel.GetWantSignedRequests(),
				// NOT SUPPORETD BY SERVER "sign_reqs":           spChannel.GetSignRequests(),
				"sign_reqs":             true,
				"signature_hash":        spChannel.GetSignatureHash(),
				"encrypt_algorithm":     spChannel.GetEncryptAssertionAlgorithm(),
				"message_ttl":           spChannel.GetMessageTtl(),
				"message_ttl_tolerance": spChannel.GetMessageTtlTolerance(),
			}
			var saml2_ls []interface{}
			saml2_ls = append(saml2_ls, saml2_map)
			sp_map["saml2"] = saml2_ls
		}
		result = append(result, sp_map)

	}

	return result, nil
}
