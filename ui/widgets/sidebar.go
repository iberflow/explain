package widgets

import (
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/rivo/tview"
)

type Sidebar struct {
	view    *tview.List
	options *man.List
}

func NewSidebar() *Sidebar {
	s := &Sidebar{}
	s.view = s.build()

	return s
}

func (s *Sidebar) Select(index int) {
	s.view.SetCurrentItem(index)
}

func (s *Sidebar) SetSelectionFunc(selected func(index int)) {
	s.view.SetSelectedFunc(func(i int, str string, s2 string, r rune) {
		selected(i)
	})
}

func (s *Sidebar) SetOptions(options *man.List) *tview.List {
	s.view.Clear()

	for _, opt := range options.Options() {
		s.view.AddItem(opt.String(), opt.Description, 0, nil)
	}

	s.view.ScrollToSelected()

	return s.view
}

func (s *Sidebar) Layout() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(Title("Options", 1, true), 3, 1, false).
		AddItem(s.view, 0, 1, true)
}

func (s *Sidebar) build() *tview.List {
	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(false)
	list.KeepSelectedItemInView(false)
	list.SetSelectOnNavigation(true)
	list.SetSelectedFocusOnly(false)
	list.SetHighlightFullLine(true)
	list.SetBorderPadding(0, 1, 3, 3)

	return list
}
