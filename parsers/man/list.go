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

func (ol *List) Unique(concatDescription bool) *List {
	groups := make(map[string][]*Option)
	list := NewList(make([]*Option, 0))

	for _, opt := range ol.options {
		if _, value := groups[opt.Name]; !value {
			groups[opt.Name] = make([]*Option, 0)
		}
		groups[opt.Name] = append(groups[opt.Name], opt)
	}

	for _, group := range groups {
		var desc string
		var params []string
		for i, opt := range group {
			if concatDescription {
				desc += opt.Description
			}

			if len(opt.Parameters) > 0 {
				params = append(params, opt.Parameters...)
			}

			if i == len(group)-1 {
				if concatDescription {
					opt.Description = desc
				}

				opt.Parameters = params
				list.Add(opt)
			}
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
