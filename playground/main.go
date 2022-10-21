package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sanity-io/litter"
	"github.com/santhosh-tekuri/jsonschema/v5"
)





func main(){
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft6
	schema, err := compiler.Compile("./201/AuthorizeResponse.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%#v\n", err)
		os.Exit(1)
	}
	litter.Dump(schema)

	var jsonStr strings.Builder
	var dfs func(root *jsonschema.Schema)
	dfs = func(root *jsonschema.Schema) {
		if root == nil {
			return
		}
		
	} 
}