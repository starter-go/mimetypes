package main

import (
	"testing"

	"github.com/starter-go/application"
	"github.com/starter-go/mimetypes/modules/mimetypes"
	"github.com/starter-go/mimetypes/src/test/golang/unit/unitnames"
	"github.com/starter-go/units"
)

func getModuleT() application.Module {
	return mimetypes.ModuleForTest()
}

func TestFindType(t *testing.T) {
	args := []string{}
	mod := getModuleT()
	cfg := &units.Config{
		Args:       args,
		Cases:      unitnames.FindType,
		Module:     mod,
		Properties: nil,
		T:          t,
		UsePanic:   false,
	}
	units.Run(cfg)
}

func TestListTypes(t *testing.T) {
	args := []string{}
	mod := getModuleT()
	cfg := &units.Config{
		Args:       args,
		Cases:      unitnames.ListTypes,
		Module:     mod,
		Properties: nil,
		T:          t,
		UsePanic:   false,
	}
	units.Run(cfg)
}
