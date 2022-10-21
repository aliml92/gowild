package gowild

import (


	"github.com/brianvoe/gofakeit/v6"
)


const (
	DefaultDateTimeFormat = "2006-01-02T15:04:05.000Z" 
)



// TODO
var Formats = map[string]func(int, int) string {
	"date-time":             genDateTime,
	"date":                  genDate,
	"time":                  genTime,
	// TODO
	// "duration":              genDuration,
	// "period":                genPeriod,
	"hostname":              genHostname,
	"email":                 genEmail,
	"ip-address":            genIPV4,
	"ipv4":                  genIPV4,
	"ipv6":                  genIPV6,
	"uri":                   genURI,
	"iri":                   genURI,
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
func genDateTime(min, max int) string {
	return gofakeit.Date().Format(DefaultDateTimeFormat)
}


// TODO
func genDate(min, max int) string {
	return gofakeit.Date().Format(DefaultDateTimeFormat)
}

func genTime(min, max int) string {
	return gofakeit.Date().String()
}


func genHostname(min, max int) string {
	return gofakeit.DomainName()
}

func genEmail(min, max int) string {
	return gofakeit.Email()
}

func genIPV4(min, max int) string {
	return gofakeit.IPv4Address()
} 

func genIPV6(min, max int) string {
	return gofakeit.IPv6Address()
}

func genURI(min, max int) string {
	return gofakeit.URL()
}