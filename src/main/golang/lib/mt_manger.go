package lib

import (
	"fmt"
	"strings"

	"github.com/starter-go/i18n"
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/vlog"
)

// DefaultTypeManager ...
type DefaultTypeManager struct {

	//starter:component

	_as func(mimetypes.Manager) //starter:as("#")

	Sources []mimetypes.Registry //starter:inject(".")
	I18n    i18n.Service         //starter:inject("#")

	defaultLanguages []i18n.Language
	cache            *mtCache
}

func (inst *DefaultTypeManager) _impl() mimetypes.Manager { return inst }

func (inst *DefaultTypeManager) loadCache() (*mtCache, error) {
	c := &mtCache{}
	c.init()
	src := inst.Sources
	for _, r1 := range src {
		if r1 == nil {
			continue
		}
		r2s := r1.ListRegistrations()
		for _, r2 := range r2s {
			c.add(r2)
		}
	}
	c.complete()
	return c, nil
}

func (inst *DefaultTypeManager) tryGetCache() *mtCache {
	c := inst.cache
	if c == nil {
		cache, err := inst.loadCache()
		if err != nil {
			panic(err)
		}
		c = cache
		inst.cache = c
	}
	return c
}

// Find ...
func (inst *DefaultTypeManager) Find(t mimetypes.Type, opt *mimetypes.Options) (*mimetypes.Info, error) {

	c := inst.tryGetCache()
	c.lock()
	defer c.unlock()

	key := keyForType(t)
	reg := c.table[key]
	if reg == nil {
		return nil, fmt.Errorf("no type info, key=%s", key)
	}

	info := inst.makeResultItem(reg, opt)
	return info, nil
}

// FindBySuffix ...
func (inst *DefaultTypeManager) FindBySuffix(suffix string, opt *mimetypes.Options) (*mimetypes.Info, error) {

	c := inst.tryGetCache()
	c.lock()
	defer c.unlock()

	key := keyForSuffix(suffix)
	reg := c.table[key]
	if reg == nil {
		return nil, fmt.Errorf("no type info, key=%s", key)
	}

	info := inst.makeResultItem(reg, opt)
	return info, nil
}

// List ...
func (inst *DefaultTypeManager) List(f mimetypes.Filter, opt *mimetypes.Options) []*mimetypes.Info {
	return inst.innerListItems(f, opt)
}

// ListAll ...
func (inst *DefaultTypeManager) ListAll(opt *mimetypes.Options) []*mimetypes.Info {
	return inst.innerListItems(nil, opt)
}

func (inst *DefaultTypeManager) myDefaultFilterFn(_ *mimetypes.Info) bool {
	return true
}

func (inst *DefaultTypeManager) innerListItems(fx mimetypes.Filter, opt *mimetypes.Options) []*mimetypes.Info {

	c := inst.tryGetCache()
	c.lock()
	defer c.unlock()

	if fx == nil {
		fx = inst.myDefaultFilterFn
	}

	src := c.list
	dst := make([]*mimetypes.Info, 0)
	for _, item1 := range src {
		if fx(&item1.Info) {
			item2 := inst.makeResultItem(item1, opt)
			dst = append(dst, item2)
		}
	}
	return dst
}

// func (inst *DefaultTypeManager) getDefaultLanguage2() (i18n.Language, i18n.Language) {
// 	all := inst.defaultLanguages
// 	if all == nil {
// 		all = inst.I18n.Available()
// 		inst.defaultLanguages = all
// 	}
// 	size := len(all)
// 	if size > 0 {
// 		return all[0], all[size-1]
// 	}
// 	return "", ""
// }

func (inst *DefaultTypeManager) makeResultItem(src *mimetypes.Registration, opt *mimetypes.Options) *mimetypes.Info {

	if opt == nil {
		opt = new(mimetypes.Options)
	}

	// lang1, lang2 := inst.getDefaultLanguage2()
	langlist1 := inst.I18n.Available()
	langlist2 := []i18n.Language{opt.Language}
	langlist2 = append(langlist2, langlist1...)

	res := inst.I18n.GetResources(langlist2...)
	s2 := &src.Info
	dst := &mimetypes.Info{}
	*dst = *s2

	// load i18n
	inst.findI18nString(dst.Label, res, func(text2 string) {
		dst.Label = text2
	})
	inst.findI18nString(dst.Description, res, func(text2 string) {
		dst.Description = text2
	})

	dst.Language = opt.Language
	return dst
}

func (inst *DefaultTypeManager) findI18nString(text string, res i18n.Resources, callback func(text2 string)) {
	const (
		prefix = "{{"
		suffix = "}}"
	)
	want := strings.TrimSpace(text)
	if strings.HasPrefix(want, prefix) && strings.HasSuffix(want, suffix) {
		i1 := len(prefix)
		i2 := len(want) - len(suffix)
		want = strings.TrimSpace(want[i1:i2])
	} else {
		return
	}
	text2, err := res.GetString(want)
	if err != nil {
		vlog.Warn(err.Error())
		return
	}
	callback(text2)
}

// Reload ...
func (inst *DefaultTypeManager) Reload() error {
	inst.cache = nil
	return nil
}
