package main

import (
	"log"

	"github.com/aliml92/gowild"
)





func main(){
	jsonStr, err := gowild.Generate("Authorize.json")
	if err != nil {
		log.Printf("err: %v", err)
	}
	log.Printf("json string: %v", jsonStr)
}