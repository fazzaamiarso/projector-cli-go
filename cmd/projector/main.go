package main

import (
	"fmt"
	"log"

	"github.com/fazzaamiarso/projector/pkg/projector"
)


func main () {
	opts, err := projector.GetOptions()
	if err != nil {
		log.Fatalf("unable to get options %+v", err)
	}
	
	fmt.Printf("%+v", opts )
}