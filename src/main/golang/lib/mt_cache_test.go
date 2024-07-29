package lib

import (
	"testing"

	"github.com/starter-go/mimetypes"
)

func TestCacheSort(t *testing.T) {

	builder := func(cache *mtCache, priority int, name string) {
		cache.add(&mimetypes.Registration{
			Priority: priority,
			Name:     name,
		})
	}

	c := &mtCache{}
	c.init()

	builder(c, 66, "666")
	builder(c, 11, "110")
	builder(c, 0, "zero")
	builder(c, 99, "jiu2")
	builder(c, 55, "wu2")
	builder(c, 77, "qi2")
	builder(c, 88, "ba2")
	builder(c, 33, "san3")

	c.complete()

	list := c.list
	for i, item := range list {
		t.Logf("[mimetypes.Registration index:%d priority:%d name:%s]", i, item.Priority, item.Name)
	}
	// 注意：这里的 list.items 顺序必须是从小到大，这样才能让高优先级的项 在 put 操作时覆盖掉同key的低优先级项.
}
