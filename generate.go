package gowild

import (
	"strconv"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)


const (
	jsonString 	= "string"
	jsonNumber 	= "number"
	jsonInteger	= "integer"
	jsonObject 	= "object"
	jsonArray   = "array"
	jsonBoolean = "boolean"
	jsonNull	= "null"
)


func Generate(url string) (string, error) {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft4
	schema, err := compiler.Compile(url)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	generate("", &b, schema)
	return b.String(), nil
}



func generate(key string, b *strings.Builder, schema *jsonschema.Schema) {
	b.WriteString("{")
	if key != "" {
		b.WriteString(strconv.Quote(key))
		b.WriteString(":")
	}
	for k, s := range schema.Properties {
		generate(k, b, s)
	}
	jsonType := schema.Types[0]
	switch jsonType {
		case jsonString:
			genStringValue(key, b, schema)
		case jsonInteger:
			genIntegerValue(key, b, schema)
		case jsonNumber:
			genNumberValue(key, b, schema)
		case jsonArray:
			genArrayValue(key, b, schema)
		case jsonBoolean:
			genBooleanValue(key, b, schema)
		case jsonNull:
			b.WriteString("")		
	}
	b.WriteString("}")
}




func genStringValue(key string, b *strings.Builder, schema *jsonschema.Schema) {
	if schema.Format != "" {
		q := strconv.Quote(Formats[schema.Format]())
		b.WriteString(q)
		return  
	}
	q := strconv.Quote("fakestring")
	b.WriteString(q)
	// 1. check pattern 
    // 2. check content encoding
	// 3. check content media type 
	// 4. special function to create string, check if gofakeit has an option for key
	// b.WriteString(func(key, s.MinLength, s.MaxLength)
}


func genIntegerValue(key string, b *strings.Builder, schema *jsonschema.Schema) {
	b.WriteString("1")
	// create integer value
}

func genNumberValue(key string, b *strings.Builder, schema *jsonschema.Schema) {
	b.WriteString("1.54")
	// create number value
}

func genArrayValue(key string, b *strings.Builder, schema *jsonschema.Schema) {
	b.WriteString("[\"sampletext1\", \"sampletext2\"]")
	// create arrat value
}


func genBooleanValue(key string, b *strings.Builder, schema *jsonschema.Schema) {
	b.WriteString("true")
	// create boolean value
}

