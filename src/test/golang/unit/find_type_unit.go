package unit

import (
	"encoding/json"

	"github.com/starter-go/i18n"
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/mimetypes/src/test/golang/unit/unitnames"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// FindTypeUnit ... 单元测试示例
type FindTypeUnit struct {

	//starter:component

	_as func(units.Units) //starter:as(".")

	TM mimetypes.Manager //starter:inject("#")

}

func (inst *FindTypeUnit) _impl() units.Units { return inst }

// Units ...
func (inst *FindTypeUnit) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     unitnames.FindType,
		Enabled:  true,
		Priority: 0,
		Test:     inst.test1,
	})

	return list
}

// Units ...
func (inst *FindTypeUnit) test1() error {

	// by type name
	info, err := inst.TM.Find("text/html", nil)
	if err != nil {
		return err
	}
	logTypeInfo("find type by name: %s", info)

	// by suffix
	langs := []i18n.Language{"de_de", "en_us", "fr_fr", "it_it", "zh_cn"}
	for _, lang := range langs {
		opt := &mimetypes.Options{Language: lang}
		info, err = inst.TM.FindBySuffix(".html", opt)
		if err != nil {
			return err
		}
		logTypeInfo("find type by suffix: %s", info)
	}

	return nil
}

func logTypeInfo(format string, info *mimetypes.Info) {
	bin, err := json.MarshalIndent(info, "", "\t")
	if err == nil {
		str := string(bin)
		vlog.Info(format, str)
	}
}
