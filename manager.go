package mimetypes

import "github.com/starter-go/i18n"

// Options 表示可选的查询条件
type Options struct {
	Language i18n.Language
}

// Filter 定义类型管理器函数
type Filter func(info *Info) bool

// Manager 提供管理类型的接口
type Manager interface {
	Find(t Type, opt *Options) (*Info, error)

	FindBySuffix(suffix string, opt *Options) (*Info, error)

	List(f Filter, opt *Options) []*Info

	ListAll(opt *Options) []*Info

	Reload() error
}
