package docs

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/atricore/terraform-provider-iamtf/iamtf"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

type GenContext struct {
	out string
	src string
}

var CTX GenContext

var IGNORE_PROPS = [...]string{"element_id"}

// Generates documentation markdonw files in the destination path
func GenerateDocs(out string, src string) error {

	CTX = GenContext{
		out: out,
		src: src,
	}

	var errWrap error

	fmt.Print("Generating docs ... \n")
	fmt.Printf("SRC: %s\n", CTX.src)
	fmt.Printf("OUT: %s\n", CTX.out)

	err := initOut(CTX.out)
	if err != nil {
		errWrap = errors.Wrap(err, "init out")
	}

	// 1. Provider documentation
	err = genProvider("josso", iamtf.Provider())
	if err != nil {
		errWrap = errors.Wrap(err, "josso")
	}

	// 2. Identity Appliance
	err = genResource("iamtf_identity_appliance", iamtf.ResourceIdentityAppliance())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_identity_appliance out")
	}

	// 3. Identity Provider
	err = genResource("iamtf_idp", iamtf.ResourceIdP())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_idp")
	}

	// 4. Identity Sources/Vaults
	err = genResource("iamtf_idsource_db", iamtf.ResourcedbidSource())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_idsource_db")
	}
	err = genResource("iamtf_idsource_ldap", iamtf.ResourceIdSourceLdap())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_idsource_ldap")
	}
	err = genResource("iamtf_idvault", iamtf.ResourceIdVault())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_idvault")
	}

	// 5. Applications
	err = genResource("iamtf_app_agent", iamtf.ResourceJosso1Re())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_app_agent")
	}
	err = genResource("iamtf_app_oidc", iamtf.ResourceOidcRp())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_app_oidc")
	}
	err = genResource("iamtf_app_saml2", iamtf.ResourceExtSaml2Sp())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_app_saml2")
	}

	// 6. Execution Environments
	err = genResource("iamtf_execenv_tomcat", iamtf.ResourceTomcatExecenv())
	if err != nil {
		errWrap = errors.Wrap(err, "iamtf_execenv_tomcat")
	}

	return errWrap
}

func genProvider(name string, p *schema.Provider) error {

	return genDocs(name, "", p.Schema)

}

func genResource(name string, r *schema.Resource) error {
	return genDocs(name, r.Description, r.Schema)
}

func genDocs(name string, description string, m map[string]*schema.Schema) error {

	var errWrap error

	// Open writer
	w, err := buildWriter(CTX.out, name)
	if err != nil {
		return err
	}

	// Write header
	err = printHeader("", name, w, 1)
	if err != nil {
		errWrap = errors.Wrap(err, "header")
	}
	fmt.Fprintf(w, "\n%s\n\n", description)

	// Write properties
	f := func(p string, n string, s *schema.Schema, depth int) error {
		err = printProperty(p, n, s, w, depth)
		if err != nil {
			errWrap = errors.Wrap(err, "property")
		}
		return nil
	}

	// Process children
	err = walkSchemaMap("", name, m, f, 2)
	if err != nil {
		errWrap = errors.Wrap(err, "walker")
	}

	// Print footer
	err = printFooter("", name, w, 1)
	if err != nil {
		errWrap = errors.Wrap(err, "footer")
	}
	w.Flush()

	return errWrap

}

func printHeader(path string, name string, w *bufio.Writer, depth int) error {

	for i := 0; i < depth; i++ {
		fmt.Fprint(w, "#")
	}
	fmt.Fprintf(w, " %s\n", name)

	// Add manual documentation
	copyContent(fmt.Sprintf("%s/%s_%s_h.md", CTX.src, path, name), w)

	return nil
}

func printFooter(path string, name string, w *bufio.Writer, depth int) error {

	// Add manual documentation
	fmt.Fprint(w, "\n")
	copyContent(fmt.Sprintf("%s/%s_%s_f.md", CTX.src, path, name), w)
	fmt.Fprint(w, "\n")

	return nil
}

func printBody(path string, name string, s *schema.Schema, w *bufio.Writer, depth int) error {
	fmt.Fprintf(w, "\n%s\n\n", s.Description)

	copyContent(fmt.Sprintf("%s/%s_%s_b.md", CTX.src, path, name), w)

	fmt.Fprintf(w, "* type: %s\n", strings.TrimPrefix(s.Type.String(), "Type"))

	if s.Required {
		fmt.Fprintf(w, "* required")
	}
	if s.Optional {
		fmt.Fprintf(w, "* optional: %t\n", s.Optional)
	}

	if s.Computed {
		fmt.Fprintf(w, "* computed: %t\n", s.Computed)
	}

	fmt.Fprintf(w, "\n")

	return nil
}

// Prints the schema property definition.
func printProperty(p string, n string, s *schema.Schema, w *bufio.Writer, depth int) error {

	var errWrap error

	// TODO : errors!
	err := printHeader(p, n, w, depth)
	if err != nil {
		errWrap = errors.Wrap(err, "header")
	}

	err = printBody(p, n, s, w, depth)
	if err != nil {
		errWrap = errors.Wrap(err, "body")
	}

	err = printFooter(p, n, w, depth)
	if err != nil {
		errWrap = errors.Wrap(err, "footer")
	}

	return errWrap
}

//
// This walks an element by exuting the provided function on all its children
//
// ns : namespace of the element to process
// name : of the element to process
func walkSchemaMap(
	ns string,
	name string,
	props map[string]*schema.Schema,
	fn func(ns string, name string, sc *schema.Schema, depth int) error, depth int) error {

	keys := keys(props)
	sort.Strings(keys)
	// Get keys from map, each key is a property
	for _, prop_name := range keys {

		if isIgnored(prop_name) {
			continue
		}

		prop_value := props[prop_name]
		prop_ns := fmt.Sprintf("%s_%s", ns, name)

		// Cal the function that handles the prop
		err := fn(prop_ns, prop_name, prop_value, depth)
		if err != nil {
			return err
		}

		// If the prop is a list, walk its children
		switch prop_value.Type {
		case schema.TypeList:
			o := prop_value.Elem
			if r, ok := o.(*schema.Resource); ok {
				err = walkSchemaMap(prop_ns, prop_name, r.Schema, fn, depth+1)
				if err != nil {
					return err
				}
			}
		case schema.TypeMap:
			// TODO !
		case schema.TypeSet:
			// TODO !
		}

	}
	return nil

}

func isIgnored(p string) bool {
	for _, i := range IGNORE_PROPS {
		if i == p {
			return true
		}
	}
	return false
}

func keys(m map[string]*schema.Schema) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}

func buildWriter(path string, name string) (*bufio.Writer, error) {
	var f *os.File
	var err error
	var writer *bufio.Writer
	fname := fmt.Sprintf("%s/%s.md", path, name)

	// change to either true|false
	outputToScreen := false
	if outputToScreen {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(fname, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return writer, fmt.Errorf("Error opening file %s: %v", fname, err)
		}
	}

	return bufio.NewWriter(f), nil

}

func initOut(dname string) error {
	if _, err := os.Stat(dname); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dname, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// If file does not exists, ignore it
func copyContent(fname string, w *bufio.Writer) error {
	// open input file

	src, err := ioutil.ReadFile(fname)
	if err != nil {
		// If file not found , ignore and return
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	_, err = w.Write(src)

	return err
}
