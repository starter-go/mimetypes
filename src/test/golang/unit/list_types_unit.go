package unit

import (
	"fmt"

	"github.com/starter-go/mimetypes"
	"github.com/starter-go/mimetypes/src/test/golang/unit/unitnames"
	"github.com/starter-go/units"
)

// ListTypesUnit ... 单元测试示例
type ListTypesUnit struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

	TM mimetypes.Manager //starter:inject("#")

}

func (inst *ListTypesUnit) _impl() units.Units { return inst }

// Units ...
func (inst *ListTypesUnit) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     unitnames.ListTypes,
		Enabled:  true,
		Priority: 0,
		Test:     inst.test1,
	})

	return list
}

// Units ...
func (inst *ListTypesUnit) test1() error {
	opt := &mimetypes.Options{
		Lang: "en_us",
	}
	all := inst.TM.ListAll(opt)
	for i, info := range all {
		text1 := fmt.Sprintf("list.items[%d].", i)
		logTypeInfo(text1+"info: %s", info)
	}
	return nil
}
