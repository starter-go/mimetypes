package lib

import (
	"strings"

	"github.com/starter-go/mimetypes"
)

type mtHolder struct {
	reg *mimetypes.Registration
}

func (inst *mtHolder) add(src *mimetypes.Registration) {

	// check newer
	if src == nil {
		return
	}

	// check older
	dst := inst.reg
	if dst == nil {
		inst.reg = src
		return
	}

	// for fields
	if src.Info.Icon != "" {
		dst.Info.Icon = src.Info.Icon
	}
	if src.Info.Label != "" {
		dst.Info.Label = src.Info.Label
	}
	if src.Info.Description != "" {
		dst.Info.Description = src.Info.Description
	}
	if src.Info.Language != "" {
		dst.Info.Language = src.Info.Language
	}

	if src.Name != "" {
		dst.Name = src.Name
	}

	dst.Priority = src.Priority
	dst.Suffixes = inst.mixSuffixes(src.Suffixes, dst.Suffixes)
}

func (inst *mtHolder) normalizeSuffix(suffix string) string {
	suffix = strings.TrimSpace(suffix)
	suffix = strings.ToLower(suffix)
	const dot = "."
	if strings.HasPrefix(suffix, dot) {
		return suffix
	}
	return dot + suffix
}

func (inst *mtHolder) mixSuffixes(suffixes1, suffixes2 []string) []string {

	set := make(map[string]bool)
	for _, item := range suffixes1 {
		suffix := inst.normalizeSuffix(item)
		set[suffix] = true
	}
	for _, item := range suffixes2 {
		suffix := inst.normalizeSuffix(item)
		set[suffix] = true
	}

	all := make([]string, 0)
	for suffix := range set {
		if suffix != "" {
			all = append(all, suffix)
		}
	}
	return all
}
