package widgets

import (
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
)

type OptionClickFunc func(index int)

func ClickFunc(opts *man.List, callback OptionClickFunc) func(region string) {
	return func(region string) {
		region = text.StripColor(region)
		index := opts.Search(region)
		if index > 0 {
			callback(index)
		}
	}
}
