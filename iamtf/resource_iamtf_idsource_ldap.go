package iamtf

import (
	"context"
	"fmt"
	"strings"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// TODO : credentialQueryString;
// TODO : updateableCredentialAttribute;
// TODO : updatePasswordEnabled;

func ResourceIdSourceLdap() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdSourceLdapCreate,
		ReadContext:   resourceIdSourceLdapRead,
		UpdateContext: resourceIdSourceLdapUpdate,
		DeleteContext: resourceIdSourceLdapDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "element ID",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "id source name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "id source description",
			},
			"initial_ctx_factory": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Java JNDI initial context factory",
				Default:     "com.sun.jndi.ldap.LdapCtxFactory",
			},
			"provider_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "LDAP server connection url: ldaps://localhost:636",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "credential to connect to the LDAP server",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "credential to connect to the LDAP server",
			},
			"authentication": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringInSlice([]string{"none", "strong", "simple"}),
				Default:          "simple",
				Description:      "credential to connect to the LDAP server",
			},
			"search_scope": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringInSlice([]string{"base", "one", "subtree", "chidlren"}),
				Default:          "subtree",
				Description:      "LDAP search scope",
			},
			"users_ctx_dn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DN to search for users",
			},
			"userid_attr": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "uid",
				Description: "LDAP attribute containing a user identifier",
			},
			"groups_ctx_dn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DN to search for groups",
			},
			"groupid_attr": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "cn",
				Description: "LDAP attribute containing a group identifier",
			},
			"groupmember_attr": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "uniquemember",
				Description: "LDAP attribute containing a user identifier in a group",
			},
			"group_match_mode": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "DN",
				ValidateDiagFunc: stringInSlice([]string{"DN", "UID", "PRINCIPAL"}),
				Description:      "Specifies the type of value stored as a groupmember of a group",
			},
			"user_attributes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of LDAP attributes and the name to be used as claim for a user",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "LDAP attribute",
						},
						"claim": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "claim name reported to applications",
						},
					},
				},
			},
			"referrals": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringInSlice([]string{"follow", "ignore"}),
				Default:          "follow",
				Description:      "how to process referrals in a directory node",
			},
			"operational_attrs": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Require LDAP operational attributes (useful for LDAP password policy management)",
			},
			"extension": customClassSchema(),
		},
	}
}

func resourceIdSourceLdapCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIdSourceLdapCreate", "ida", d.Get("ida").(string))

	idSourceLdap, err := buildIdSourceLdapDTO(d)
	if err != nil {
		return diag.Errorf("failed to build idSourceLdap: %v", err)
	}
	l.Trace("resourceIdSourceLdapCreate", "ida", d.Get("ida").(string), "name", *idSourceLdap.Name)

	a, err := getJossoClient(m).CreateIdSourceLdap(d.Get("ida").(string), idSourceLdap)
	if err != nil {
		l.Debug("resourceIdSourceLdapCreate %v", err)
		return diag.Errorf("failed to create idSourceLdap: %v", err)
	}

	if err = buildIdSourceLdapResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdSourceLdapCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIdSourceLdapCreate OK", "ida", d.Get("ida").(string), "name", *idSourceLdap.Name)

	return nil
}
func resourceIdSourceLdapRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	// if idaName is empty, it means that the resource is being imported
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceIdSourceLdapRead", "ida", idaName, "name", d.Id())
	idSourceLdap, err := getJossoClient(m).GetIdSourceLdap(idaName, d.Id())
	if err != nil {
		l.Debug("resourceIdSourceLdapRead %v", err)
		return diag.Errorf("resourceIdSourceLdapRead: %v", err)
	}
	if idSourceLdap.Name == nil || *idSourceLdap.Name == "" {
		l.Debug("resourceIdSourceLdapRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIdSourceLdapResource(idaName, d, idSourceLdap); err != nil {
		l.Debug("resourceIdSourceLdapRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIdSourceLdapRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdSourceLdapUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdSourceLdapUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	idSourceLdap, err := buildIdSourceLdapDTO(d)
	if err != nil {
		l.Debug("resourceIdSourceLdapUpdate %v", err)
		return diag.Errorf("failed to build idSourceLdap: %v", err)
	}

	a, err := getJossoClient(m).UpdateIdSourceLdap(d.Get("ida").(string), idSourceLdap)
	if err != nil {
		l.Debug("resourceIdSourceLdapUpdate %v", err)
		return diag.Errorf("failed to update idSourceLdap: %v", err)
	}

	if err = buildIdSourceLdapResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdSourceLdapUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIdSourceLdapUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdSourceLdapDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIdSourceLdapDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdSourceLdap(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdSourceLdapDelete %v", err)
		return diag.Errorf("failed to delete idSourceLdap: %v", err)
	}

	l.Debug("resourceIdSourceLdapDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildIdSourceLdapDTO(d *schema.ResourceData) (api.LdapIdentitySourceDTO, error) {
	var err, errWrap error
	dto := api.NewLdapIdentitySourceDTO()
	dto.ElementId = PtrSchemaStr(d, "element_id")
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	dto.InitialContextFactory = PtrSchemaStr(d, "initial_ctx_factory")
	dto.ProviderUrl = PtrSchemaStr(d, "provider_url")
	dto.SecurityPrincipal = PtrSchemaStr(d, "username")
	dto.SecurityCredential = PtrSchemaStr(d, "password")
	dto.SecurityAuthentication = PtrSchemaStr(d, "authentication")
	dto.LdapSearchScope = PtrSchemaStr(d, "search_scope")
	dto.UsersCtxDN = PtrSchemaStr(d, "users_ctx_dn")
	dto.PrincipalUidAttributeID = PtrSchemaStr(d, "userid_attr")
	dto.RolesCtxDN = PtrSchemaStr(d, "groups_ctx_dn")
	dto.RoleAttributeID = PtrSchemaStr(d, "groupid_attr")
	dto.UidAttributeID = PtrSchemaStr(d, "groupmember_attr")
	dto.RoleMatchingMode = PtrSchemaStr(d, "group_match_mode")

	if v, ok := d.Get("user_attributes").([]interface{}); ok {
		var s string
		s, err = flattenUserAttrs(v)
		if err != nil {
			return *dto, err
		}
		dto.UserPropertiesQueryString = api.PtrString(s)
	} else {
		return *dto, fmt.Errorf("invalid user_attributes value %#v", d.Get("user_attributes"))
	}

	if err != nil {
		errWrap = errors.Wrap(err, "user_attributes")
	}
	cc_dto, err := convertCustomClassMapArrToDTO(d.Get("extension"))
	if err != nil {
		errWrap = errors.Wrap(err, "extension")
	}
	dto.CustomClass = cc_dto
	dto.Referrals = PtrSchemaStr(d, "referrals")
	dto.IncludeOperationalAttributes = PtrSchemaBool(d, "operational_attrs")

	return *dto, errWrap
}

/*func buildIdSourceLdapResource(d *schema.ResourceData, dto api.LdapIdentitySourceDTO) error {
	return buildIdSourceLdapResourceForIda("", d, dto)
}*/

func buildIdSourceLdapResource(idaName string, d *schema.ResourceData, dto api.LdapIdentitySourceDTO) error {
	// if idaName is not empty store it in d struct
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("ida", idaName)
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	d.Set("initial_ctx_factory", cli.StrDeref(dto.InitialContextFactory))
	d.Set("provider_url", cli.StrDeref(dto.ProviderUrl))

	d.Set("username", cli.StrDeref(dto.SecurityPrincipal))
	d.Set("password", cli.StrDeref(dto.SecurityCredential))
	d.Set("authentication", cli.StrDeref(dto.SecurityAuthentication))
	d.Set("search_scope", cli.StrDeref(dto.LdapSearchScope))
	d.Set("users_ctx_dn", cli.StrDeref(dto.UsersCtxDN))
	d.Set("userid_attr", cli.StrDeref(dto.PrincipalUidAttributeID))
	d.Set("groups_ctx_dn", cli.StrDeref(dto.RolesCtxDN))
	d.Set("groupid_attr", cli.StrDeref(dto.RoleAttributeID))
	d.Set("groupmember_attr", cli.StrDeref(dto.UidAttributeID))
	d.Set("group_match_mode", cli.StrDeref(dto.RoleMatchingMode))

	atrs, err := unflattenUserAttrs(cli.StrDeref(dto.UserPropertiesQueryString))
	if err != nil {
		return err
	}
	d.Set("user_attributes", atrs)

	// User attributes!
	// dto.UserPropertiesQueryString : key=value,key1=value1

	d.Set("referrals", cli.StrDeref(dto.Referrals))
	d.Set("operational_attrs", cli.BoolDeref(dto.IncludeOperationalAttributes))

	customClass, err := convertCustomClassDTOToMapArr(dto.CustomClass)
	if err != nil {
		return err
	}
	_ = d.Set("extension", customClass)

	return nil
}

func flattenUserAttrs(attrs []interface{}) (string, error) {
	var reg string
	for _, e := range attrs {

		var m map[string]interface{}
		var ok bool
		var claim, attribute string

		if m, ok = e.(map[string]interface{}); !ok {
			return reg, fmt.Errorf("invalid attribute map %#v", e)
		}

		if claim, ok = m["attribute"].(string); ok {
			reg += claim + "="
		} else {
			return reg, fmt.Errorf("invalid attribute map %#v", e)
		}

		if attribute, ok = m["claim"].(string); ok {
			reg += attribute + ","
		}

	}

	if last := len(reg) - 1; last >= 0 && reg[last] == ',' {
		reg = reg[:last]
	}

	return reg, nil
}

func unflattenUserAttrs(attrs string) ([]interface{}, error) {

	var result []interface{}

	mappings := strings.Split(attrs, ",")
	for _, mapping := range mappings {
		if r, err := unmarshalAttr(mapping); err != nil {
			return nil, err
		} else {
			result = append(result, r)
		}
	}

	return result, nil

}

// Receives a string in the format <key>=<value> and returnsa map
func unmarshalAttr(mapping string) (map[string]interface{}, error) {

	reg := strings.SplitN(mapping, "=", 2)
	if len(reg) != 2 {
		return nil, fmt.Errorf("invalid string for attribute %s", mapping)
	}
	return map[string]interface{}{
		"claim":     reg[1],
		"attribute": reg[0],
	}, nil
}
