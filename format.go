package gowild

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)


// TODO
var Formats = map[string]func() string {
	"date-time":             genDateTime,
	"date":                  genDate,
	"time":                  genTime,
	// TODO
	// "duration":              genDuration,
	// "period":                genPeriod,
	// "hostname":              genHostname,
	// "email":                 genEmail,
	// "ip-address":            genIPV4,
	// "ipv4":                  genIPV4,
	// "ipv6":                  genIPV6,
	// "uri":                   genURI,
	// "iri":                   genURI,
	// "uri-reference":         genURIReference,
	// "uriref":                genURIReference,
	// "iri-reference":         genURIReference,
	// "uri-template":          genURITemplate,
	// "regex":                 genRegex,
	// "json-pointer":          genJSONPointer,
	// "relative-json-pointer": genRelativeJSONPointer,
	// "uuid":                  genUUID,
}


// TODO
// return current date-time
func genDateTime() string {
	return time.Now().Format("2006-01-02T15:04:05.000Z")
}


// TODO
func genDate() string {
	return gofakeit.Date().String()
}

func genTime() string {
	return gofakeit.Date().String()
}

