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
	generate("", true, true, &b, schema)
	res := b.String()
	return res[1:b.Len()-1], nil
}



func generate(key string, start bool, end bool, b *strings.Builder, schema *jsonschema.Schema) {
	if start { b.WriteString("{") }
	if key != "" {
		b.WriteString(strconv.Quote(key))
		b.WriteString(":")
	}
	jsonType := schema.Types[0]
	switch jsonType {
		case jsonString:
			genStringValue(key, b, schema)
		case jsonInteger:
			genNumberValue(key, b, schema)
		case jsonNumber:
			genIntegerValue(key, b, schema)
		case jsonObject:
			l := len(schema.Properties)
			if l > 0 {
				var i int 
				for k, s := range schema.Properties {
					var start, end bool
					if i == 0{
						start = true
					}
					if i == l-1 {
						end = true
					}
					generate(k, start, end,  b, s)
					i++
				}
			}	
		case jsonArray:	
			genArrayValue(key, b, schema)
		case jsonBoolean:
			genBooleanValue(key, b, schema)
		case jsonNull:
			b.WriteString("")			
	}
	if end { b.WriteString("}") } else { b.WriteString(",") }
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

