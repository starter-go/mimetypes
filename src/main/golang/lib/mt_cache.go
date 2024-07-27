package lib

import (
	"sync"

	"github.com/starter-go/mimetypes"
)

type mtCache struct {
	mutex sync.Mutex
	list  []*mimetypes.Registration
	table map[Key]*mimetypes.Registration
}

func (inst *mtCache) init() {
	inst.list = make([]*mimetypes.Registration, 0)
	inst.table = make(map[Key]*mimetypes.Registration)
}

func (inst *mtCache) lock() {
	inst.mutex.Lock()
}

func (inst *mtCache) unlock() {
	inst.mutex.Unlock()
}

func (inst *mtCache) put(item *mimetypes.Registration) {

	if item == nil {
		return
	}

	t1 := item.Info.Type

	// compute name
	name, _, err := t1.Parse()
	if err == nil {
		item.Name = name
	}

	// put to table with name
	key := keyForType(t1)
	inst.table[key] = item

	// put to table with suffix
	for _, suffix := range item.Suffixes {
		key = keyForSuffix(suffix)
		inst.table[key] = item
	}

	// append to list
	inst.list = append(inst.list, item)
}
