package lib

import (
	"sort"
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

func (inst *mtCache) innerPut(item *mimetypes.Registration) {

	// put to table with name
	t1 := item.Info.Type
	key := keyForType(t1)
	inst.table[key] = item

	// put to table with suffix
	for _, suffix := range item.Suffixes {
		key = keyForSuffix(suffix)
		inst.table[key] = item
	}

}

func (inst *mtCache) add(item *mimetypes.Registration) {
	if item == nil {
		return
	}
	t1 := item.Info.Type
	// compute name
	name, _, err := t1.Parse()
	if err == nil {
		item.Name = name
	}
	inst.list = append(inst.list, item)
}

func (inst *mtCache) complete() {

	// sort 根据项目的优先级排序
	sort.Sort(inst)

	// put to table
	all := inst.list
	for _, item := range all {
		inst.innerPut(item)
	}
}

func (inst *mtCache) Len() int {
	return len(inst.list)
}
func (inst *mtCache) Less(i1, i2 int) bool {
	p1 := inst.list[i1].Priority
	p2 := inst.list[i2].Priority
	return p1 < p2
}
func (inst *mtCache) Swap(i1, i2 int) {
	l := inst.list
	l[i1], l[i2] = l[i2], l[i1]
}
