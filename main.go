package main

import (
	"context"
	"flag"
	"log"

	"github.com/atricore/terraform-provider-iamtf/iamtf"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {

	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: iamtf.Provider}

	if debugMode {
		// TODO: update this string with the full name of your provider as used in your configs
		err := plugin.Debug(context.Background(), "github.com/atricore/josso-terraform-provider", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: iamtf.Provider,
	})
}
