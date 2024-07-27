package main

import (
	"os"

	"github.com/starter-go/mimetypes/modules/mimetypes"
	"github.com/starter-go/starter"
)

func main() {
	m := mimetypes.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
