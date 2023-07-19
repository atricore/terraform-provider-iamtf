package iamtf

import (
	"fmt"

	api "github.com/atricore/josso-api-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IdPRole interface {
	GetName() string
	GetWantAuthnRequestsSigned() bool
	GetSignatureHash() string
	GetEncryptAssertionAlgorithm() string
	GetMessageTtl() int32
	GetMessageTtlTolerance() int32
}

type SPRole interface {
	GetName() string
	GetSignAuthenticationRequests() bool
	GetIdentityMappingPolicy() api.IdentityMappingPolicyDTO
	GetAccountLinkagePolicy() api.AccountLinkagePolicyDTO
	GetWantAssertionSigned() bool
	GetSignatureHash() string
	GetMessageTtl() int32
	GetMessageTtlTolerance() int32
}

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
				"account_linkage_name": {
					Type:        schema.TypeString,
					Description: "account linkage name, only valid when account_linkage is set to CUSTOM",
					Optional:    true,
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
					Description:      "how the user identity should be mapped for this SP. LOCAL means that the user claims will be retrieved from an identity source connected to the SP.  REMOTE means that claims from the IdP will be used. MERGE is a mix of both claim sets (LOCAL and REMOTE)",
					ValidateDiagFunc: stringInSlice([]string{"LOCAL", "REMOTE", "MERGED", "CUSTOM"}),
					Computed:         true,
					Optional:         true,
				},
				"identiyt_mapping_name": {
					Type:        schema.TypeString,
					Description: "identity mapping name, only valid when identity_mapping is set to CUSTOM",
					Optional:    true,
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
				"bindings": {
					Type:        schema.TypeList,
					Optional:    true,
					Computed:    true,
					MaxItems:    1,
					MinItems:    0,
					Description: "enabled SAML bindings",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"http_post": {
								Type:        schema.TypeBool,
								Description: "use HTTP POST binding",
								Optional:    true,
								Default:     false,
							},
							"http_redirect": {
								Type:        schema.TypeBool,
								Description: "use HTTP REDIRECT binding",
								Optional:    true,
								Default:     true,
							},
							"artifact": {
								Type:        schema.TypeBool,
								Description: "use Artifact binding",
								Optional:    true,
								Default:     true,
							},
							"soap": {
								Type:        schema.TypeBool,
								Description: "use SOAP binding",
								Optional:    true,
								Default:     true,
							},
							"local": {
								Type:        schema.TypeBool,
								Description: "use LOCAL binding",
								Optional:    true,
								Default:     true,
							},
						},
					},
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
func convertIdPFederatedConnectionsMapArrToDTOs(sp SPRole, d *schema.ResourceData, idp interface{}) ([]api.FederatedConnectionDTO, error) {
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
			it.SetMappingType("REMOTE")
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
			at.SetLinkEmitterType("ONE_TO_ONE")
		}
		al := api.NewAccountLinkagePolicyDTO()
		al.AdditionalProperties = make(map[string]interface{})
		al.AdditionalProperties["@c"] = ".AccountLinkagePolicyDTO"
		al.SetName(fmt.Sprintf("%s-account-linkage", sp.GetName()))
		al.SetLinkEmitterType(GetAsString(d, fmt.Sprintf("idp.%d.saml2.0.account_linkage", idpIdx), at.GetLinkEmitterType()))
		idpChannel.SetAccountLinkagePolicy(*al)

		if len(saml2_m) > 0 {

			idpChannel.SetOverrideProviderSetup(true)

			//idpChannel.SetActiveBindings()

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

		// asd, err := convertMapArrToActiveBinding(m["bindings"])
		// idpChannel.SetActiveBindings(asd)
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
			ab, err := convertActiveBindingToMapArr(idpChannel.GetActiveBindings())
			if err != nil {
				return ab, err
			}
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
					"bindings":                     ab,
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
					ValidateDiagFunc: stringInSlice([]string{"NONE", "AES128", "AES256", "AES3DES"}),
					Default:          "NONE",
				},
				"bindings": {
					Type:        schema.TypeList,
					Optional:    true,
					Computed:    true,
					MaxItems:    1,
					MinItems:    0,
					Description: "enabled SAML bindings",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"http_post": {
								Type:        schema.TypeBool,
								Description: "use HTTP POST binding",
								Optional:    true,
								Default:     true,
							},
							"http_redirect": {
								Type:        schema.TypeBool,
								Description: "use HTTP REDIRECT binding",
								Optional:    true,
								Default:     true,
							},
							"artifact": {
								Type:        schema.TypeBool,
								Description: "use Artifact binding",
								Optional:    true,
								Default:     true,
							},
							"soap": {
								Type:        schema.TypeBool,
								Description: "use SOAP binding",
								Optional:    true,
								Default:     true,
							},
							"local": {
								Type:        schema.TypeBool,
								Description: "use LOCAL binding",
								Optional:    true,
								Default:     true,
							},
						},
					},
				},

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
		Description: "IDP to SP SAML 2 settings. Optional, only required is specific SAML IdP settings are required by the SP",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "name of the trusted sp. It normally is the name of the application plus the -sp suffix",
				},
				"saml2": idpSamlSchema(),
			},
		},
	}

	return spConn
}

func convertIdPSaml2DTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	encAlg := idp.GetEncryptAssertionAlgorithm()
	if encAlg == "" {
		encAlg = "NONE"
	}

	ab, err := convertActiveBindingToMapArr(idp.GetActiveBindings())
	if err != nil {
		return ab, err
	}

	enc, err := mapSaml2EncryptionToTF(idp.GetEncryptAssertionAlgorithm())
	if err != nil {
		return result, err
	}

	saml2_map := map[string]interface{}{
		"want_authn_req_signed": idp.GetWantAuthnRequestsSigned(),
		"want_req_signed":       idp.GetWantSignedRequests(),
		"sign_reqs":             idp.GetSignRequests(),
		"signature_hash":        idp.GetSignatureHash(),
		"encrypt_algorithm":     enc,
		//		"metadata_endpoint":     idp.GetEnableMetadataEndpoint(),
		"message_ttl":           int(idp.GetMessageTtl()),
		"message_ttl_tolerance": int(idp.GetMessageTtlTolerance()),
		"bindings":              ab,
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

	enc, err := mapTFEncryptionToSaml2(saml2_map["encrypt_algorithm"].(string))
	if err != nil {
		return err
	}

	idp.SetWantAuthnRequestsSigned(saml2_map["want_authn_req_signed"].(bool))
	idp.SetWantSignedRequests(saml2_map["want_req_signed"].(bool))
	idp.SetSignRequests(saml2_map["sign_reqs"].(bool))
	idp.SetSignatureHash(saml2_map["signature_hash"].(string))
	idp.SetEncryptAssertionAlgorithm(enc)
	//idp.SetEnableMetadataEndpoint(saml2_map["metadata_endpoint"].(bool))
	idp.SetEnableMetadataEndpoint(true)
	idp.SetMessageTtl(int32(saml2_map["message_ttl"].(int)))
	idp.SetMessageTtlTolerance(int32(saml2_map["message_ttl_tolerance"].(int)))
	//idp.SetActiveBindings(convertInterfaceToStringSet(""))

	b, err := convertMapArrToActiveBinding(saml2_map["bindings"])
	if err != nil {
		return err
	}
	idp.SetActiveBindings(b)

	return nil
}

func convertSPFederatedConnectionsMapArrToDTOs(idp IdPRole, d *schema.ResourceData, sp interface{}) ([]api.FederatedConnectionDTO, error) {
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

			enc, err := mapTFEncryptionToSaml2(GetAsString(d, fmt.Sprintf("sp.%d.saml2.0.encrypt_algorithm", spIdx), idp.GetEncryptAssertionAlgorithm()))
			if err != nil {
				return result, err
			}
			spChannel.SetSignatureHash(GetAsString(d, fmt.Sprintf("sp.%d.saml2.0.signature_hash", spIdx), idp.GetSignatureHash()))
			spChannel.SetEncryptAssertionAlgorithm(enc)
			spChannel.SetMessageTtl(GetAsInt32(d, fmt.Sprintf("sp.%d.saml2.0.message_ttl", spIdx), idp.GetMessageTtl()))
			spChannel.SetMessageTtlTolerance(GetAsInt32(d, fmt.Sprintf("sp.%d.saml2.0.message_ttl_tolerance", spIdx), idp.GetMessageTtlTolerance()))

			// TODO : support attribute profile specific for SPChannel

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

func convertSPFederatedConnectionDTOsToMapArr(fcs []api.FederatedConnectionDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	for _, fc := range fcs {

		spChannel, err := fc.GetSPChannel()
		if err != nil {
			return result, err
		}
		sp_map := map[string]interface{}{
			"name": fc.GetName(),
		}

		enc, err := mapSaml2EncryptionToTF(spChannel.GetEncryptAssertionAlgorithm())
		if err != nil {
			return result, err
		}

		if spChannel.GetOverrideProviderSetup() {

			saml2_map := map[string]interface{}{
				"want_authn_req_signed": spChannel.GetWantAuthnRequestsSigned(),
				// NOT SUPPORETD BY SERVER "want_req_signed":     spChannel.GetWantSignedRequests(),
				// NOT SUPPORETD BY SERVER "sign_reqs":           spChannel.GetSignRequests(),
				"sign_reqs":             true,
				"signature_hash":        spChannel.GetSignatureHash(),
				"encrypt_algorithm":     enc,
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

func convertActiveBindingToMapArr(ac []string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	if ac == nil {
		return result, nil
	}

	resultMap := map[string]interface{}{
		"http_post":     contains(ac, "SAMLR2_HTTP_POST"),
		"http_redirect": contains(ac, "SAMLR2_HTTP_REDIRECT"),
		"soap":          contains(ac, "SAMLR2_SOAP"),
		"artifact":      contains(ac, "SAMLR2_ARTIFACT"),
		"local":         contains(ac, "SAMLR2_LOCAL"),
	}

	result = append(result, resultMap)

	return result, nil
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func convertMapArrToActiveBinding(ac_arr interface{}) ([]string, error) {
	var ac []string
	tfmapLs, err := asTFMapSingle(ac_arr)
	if err != nil {
		return ac, err
	}
	if tfmapLs == nil || len(tfmapLs) == 0 {
		return ac, nil
	}

	if tfmapLs["http_post"].(bool) {
		ac = append(ac, "SAMLR2_HTTP_POST")
	}
	if tfmapLs["http_redirect"].(bool) {
		ac = append(ac, "SAMLR2_HTTP_REDIRECT")
	}
	if tfmapLs["soap"].(bool) {
		ac = append(ac, "SAMLR2_SOAP")
	}
	if tfmapLs["artifact"].(bool) {
		ac = append(ac, "SAMLR2_ARTIFACT")
	}
	if tfmapLs["local"].(bool) {
		ac = append(ac, "SAMLR2_LOCAL")
	}
	return ac, nil
}
