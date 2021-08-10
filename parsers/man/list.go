package man

import (
	"sort"
	"strings"
)

type List struct {
	options []*Option
}

func NewList(options []*Option) *List {
	return &List{options: options}
}

func (ol *List) Options() []*Option {
	return ol.options
}

func (ol *List) First() *Option {
	if len(ol.options) == 0 {
		return nil
	}

	return ol.options[0]
}

func (ol *List) Add(option *Option) *List {
	ol.options = append(ol.options, option)

	return ol
}

func (ol *List) Unique() *List {
	keys := make(map[string]bool)
	list := NewList(make([]*Option, 0))
	for _, entry := range ol.options {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list.Add(entry)
		}
	}

	return list
}

func (ol *List) Sort() {
	sort.Slice(ol.options, func(i, j int) bool {
		return ol.options[i].Name < ol.options[j].Name
	})
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
