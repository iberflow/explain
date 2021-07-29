package man

import "strings"

type List struct {
	options []*Option
}

func NewList(options []*Option) *List {
	return &List{options: options}
}

func (ol *List) Options() []*Option {
	return ol.options
}

func (ol *List) Add(option *Option) *List {
	ol.options = append(ol.options, option)

	return ol
}

func (ol *List) Search(match string) int {
	for i, o := range ol.options {
		if o.String() == match {
			return i
		}

		alias := strings.TrimLeft(match, "-")
		if len(alias) == 1 && o.Alias == alias {
			return i
		}
	}

	return 0
}
