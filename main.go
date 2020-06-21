package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/zserge/lorca"
)

func main() {
	address := runWebapp()

	ui, err := lorca.New(address, "", 490, 525)
	if err != nil {
		panic(fmt.Sprintf("%+v", errors.WithStack(err)))
	}
	defer ui.Close()

	<-ui.Done()
}
