package args

import "strings"

type List struct {
	items []*Arg
}

func (al *List) Items() []*Arg {
	return al.items
}

func (al *List) Add(arg *Arg) {
	al.items = append(al.items, arg)
}

func (al *List) String() string {
	var str []string

	for _, item := range al.items {
		str = append(str, item.Name)
	}

	return strings.Join(str, " ")
}
