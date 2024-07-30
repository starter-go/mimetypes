package mimetypes

import (
	"mime"

	"github.com/starter-go/i18n"
)

////////////////////////////////////////////////////////////////////////////////

// Type 是表示类型的字符串
type Type string

func (t Type) String() string {
	return string(t)
}

// Pure 取纯粹的类型 string （不含 params）
func (t Type) Pure() string {
	str := t.String()
	t2, _, err := mime.ParseMediaType(str)
	if err == nil {
		return t2
	}
	return ""
}

// Full 取完整的类型 string
func (t Type) Full() string {
	return string(t)
}

// Parse 解析这个类型
func (t Type) Parse() (string, map[string]string, error) {
	str := t.String()
	return mime.ParseMediaType(str)
}

// New 根据传入的参数新建类型
func New(t string, p map[string]string) Type {
	str := mime.FormatMediaType(t, p)
	return Type(str)
}

////////////////////////////////////////////////////////////////////////////////

// Info 包含与某个类型相关的信息
type Info struct {
	Type        Type          `json:"type"`
	Label       string        `json:"label"`
	Description string        `json:"description"`
	Icon        string        `json:"icon"`
	Language    i18n.Language `json:"lang"`
}

////////////////////////////////////////////////////////////////////////////////
