package lib

import (
	"strings"

	"github.com/starter-go/mimetypes"
)

// Key 表示类型的主键
type Key string

func keyForSuffix(suffix string) Key {
	suffix = strings.TrimSpace(suffix)
	suffix = strings.ToLower(suffix)
	if strings.HasPrefix(suffix, ".") {
		return Key(suffix)
	}
	return Key("." + suffix)
}

func keyForType(t mimetypes.Type) Key {
	name, _, err := t.Parse()
	if err == nil {
		name = strings.TrimSpace(name)
		name = strings.ToLower(name)
		return Key(name)
	}
	return "type/unknown"
}
