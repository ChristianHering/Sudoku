package main

import (
	"log"

	"github.com/zserge/lorca"
)

func main() {
	address := runWebapp()

	ui, err := lorca.New(address, "", 490, 525)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	<-ui.Done()
}
