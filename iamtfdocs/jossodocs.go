package main

import (
	"fmt"
	"os"

	"github.com/atricore/terraform-provider-iamtf/docs"
)

func main() {

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Printf("Usage jossodocs: out src")
	}

	out := args[0]
	src := args[1]

	err := docs.GenerateDocs(out, src)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
