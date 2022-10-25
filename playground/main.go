package main

import (
	"log"

	"github.com/aliml92/gowild"
)





func main(){
	gen, err := gowild.NewGenerator("StopTransaction.json", "4")
	if err != nil {
		log.Printf("err: %v", err)
	}
	log.Printf("json string: %v", gen.GenerateOne())

}