package demotypes

import "github.com/starter-go/mimetypes"

// DemoTypes ...
type DemoTypes struct {

	//starter:component

	_as func(mimetypes.Registry) //starter:as(".")

}

func (inst *DemoTypes) _impl() mimetypes.Registry { return inst }

// ListRegistrations ...
func (inst *DemoTypes) ListRegistrations() []*mimetypes.Registration {

	list := make([]*mimetypes.Registration, 0)

	list = append(list, &mimetypes.Registration{
		Info: mimetypes.Info{
			Type: "text/html",

			Label:       "{{strings.mimetypes.html.label}}",
			Description: "{{strings.mimetypes.html.description}}",
			Language:    "zh_cn",
			Icon:        "/a/text/html.png",
		},
		Suffixes: []string{".html", ".htm"},
	})

	list = append(list, &mimetypes.Registration{
		Info: mimetypes.Info{
			Type: "text/plain",

			Label:       "Plain Text",
			Description: "Plain Text(*.txt)",
			Language:    "zh_cn",
			Icon:        "/a/text/plain.png",
		},
		Suffixes: []string{".txt", ".text"},
	})

	list = append(list, &mimetypes.Registration{
		Info: mimetypes.Info{
			Type:        "image/png",
			Label:       "png format image",
			Description: "Portable Network Graphics",
			Language:    "zh_cn",
			Icon:        "/a/b/c.png",
		},
		Suffixes: []string{".png"},
	})

	return list
}
