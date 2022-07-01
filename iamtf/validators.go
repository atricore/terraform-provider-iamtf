package iamtf

import (
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func stringInSlice(valid []string) schema.SchemaValidateDiagFunc {
	return func(i interface{}, k cty.Path) diag.Diagnostics {
		v, ok := i.(string)
		if !ok {
			return diag.Errorf("expected type of %v to be string", k)
		}
		for _, str := range valid {
			if v == str {
				return nil
			}
		}
		return diag.Errorf("expected %v to be one of %v, got %s", k, strings.Join(valid, ","), v)
	}
}

func intInSlice(valid []int) schema.SchemaValidateDiagFunc {
	return func(i interface{}, k cty.Path) diag.Diagnostics {
		v, ok := i.(int)
		if !ok {
			return diag.Errorf("expected type of %v to be int", k)
		}
		for _, str := range valid {
			if v == str {
				return nil
			}
		}
		return diag.Errorf("expected %v to be one of %v, got %d", k, valid, v)
	}
}
