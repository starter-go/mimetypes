package mimetypes

// Registration 表示类型的注册信息
type Registration struct {
	Name     string
	Info     Info
	Suffixes []string
}

// Registry 是用于注册类型的接口
type Registry interface {
	ListRegistrations() []*Registration
}
