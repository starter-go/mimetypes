package mimetypes

import (
	"github.com/starter-go/application"
	"github.com/starter-go/i18n/modules/i18n"
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/mimetypes/gen/main4mimetypes"
	"github.com/starter-go/mimetypes/gen/test4mimetypes"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
)

// Module  ...
func Module() application.Module {
	mb := mimetypes.NewMainModule()
	mb.Components(main4mimetypes.ExportComponents)

	mb.Depend(i18n.Module())
	mb.Depend(starter.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := mimetypes.NewTestModule()
	mb.Components(test4mimetypes.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}
