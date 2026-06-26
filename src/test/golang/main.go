package main

import (
	"os"

	"github.com/starter-go/mimetypes/modules/mimetypes"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := mimetypes.ModuleForTest()
	c := new(units.Context)

	c.Arguments = a
	c.Module = m
	c.UsePanic = true

	units.Run(c)
}
