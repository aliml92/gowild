package gowild

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
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



var drafts = map[string]Draft{
	"4": 		jsonschema.Draft4,
	"6": 		jsonschema.Draft6,
	"7": 		jsonschema.Draft7,
	"2019": 	jsonschema.Draft2019,
	"2020": 	jsonschema.Draft2020,
	"latest":	jsonschema.Draft2020,
}

type Schema *jsonschema.Schema
type Draft *jsonschema.Draft

type Generator struct {
	schema Schema
	draft  Draft
}

func NewGenerator(url, draft string) (*Generator, error) {
	g := &Generator{}
	if v, ok := drafts[draft]; ok {
		g.draft = v
	} else {
		g.draft = drafts["latest"]
	}
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft4
	schema, err := compiler.Compile(url)
	if err != nil {
		return nil, err
	}
	g.schema = schema
	return g, nil
}  




func (g *Generator) GenerateOne() string {
	var b strings.Builder
	generate("", true, true, &b, g.schema)
	res := b.String()
	return res
}  



func generate(key string, start bool, end bool, b *strings.Builder, schema Schema) {
	if start { b.WriteString("{") }
	if key != "" {
		b.WriteString(strconv.Quote(key))
		b.WriteString(":")
	}
	jsonType := schema.Types[0]
	switch jsonType {
		case jsonString:
			genStringValue(key, b, schema)
		case jsonNumber:
			genNumberValue(key, b, schema)
		case jsonInteger:
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
			fmt.Println("array is on") 	
			genArrayValue(key, b, schema)
		case jsonBoolean:
			genBooleanValue(key, b, schema)
		case jsonNull:
			b.WriteString("")			
	}
	if end { b.WriteString("}") } else { b.WriteString(",") }
}




func genStringValue(key string, b *strings.Builder, schema Schema) {
	var s string
	if schema.Format != "" {
		q := strconv.Quote(Formats[schema.Format](schema.MinLength, schema.MaxLength))
		b.WriteString(q)
		return  
	}
	// TODOS
	// 1. check pattern
	// 2. check content encoding
	// 3. check content media type 
	for _, enum := range schema.Enum {
		switch v := enum.(type) {
		case string:
			s = strconv.Quote(v)
			b.WriteString(s)
		case int:
			s = strconv.Quote(string(rune(v)))
			b.WriteString(s)
		}
		return
	}  
	if schema.MaxLength != -1 {
		s = gofakeit.Word()
		if len(s) > schema.MaxLength {
			s = s[:schema.MaxLength]
		}
	} else if schema.MinLength != -1 {
		s = gofakeit.Word()
		for len(s) < schema.MinLength {
			s += s
		}
	} else {
		s = gofakeit.Word()
	}
	s = strconv.Quote(s)
	b.WriteString(s)
}


func genIntegerValue(key string, b *strings.Builder, schema Schema) {
	i := gofakeit.Uint16()
	b.WriteString(strconv.FormatUint(uint64(i), 10))
	// create integer value
}

func genNumberValue(key string, b *strings.Builder, schema Schema) {
	i := float32(gofakeit.Int8())
	if schema.MultipleOf != nil {
		f, _ := schema.MultipleOf.Float32()
		i = f * i
	}
	b.WriteString(fmt.Sprintf("%f", i)) 
	// create number value
}

// TODO
func genArrayValue(key string, b *strings.Builder, schema Schema) {
	switch v := schema.Items.(type) {
	case *jsonschema.Schema:
		b.WriteString("[")
		fmt.Printf("items type1: %v\n", v) 
		generate("", false, false, b, v)
		b.WriteString("]")
	case []*jsonschema.Schema:
		fmt.Printf("items type2: %v\n", v) 
		b.WriteString("[")
		generate("", false, false, b, v[0])
		b.WriteString("]")
	default:
		fmt.Printf("items type3: %v\n", v) 	
	} 
}


func genBooleanValue(key string, b *strings.Builder, schema Schema) {
	if gofakeit.Bool() {
		b.WriteString("true")
	} else {
		b.WriteString("false")
	}
}

