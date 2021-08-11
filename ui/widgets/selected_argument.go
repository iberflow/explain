package widgets

import (
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
	"github.com/rivo/tview"
)

type SelectedArgument struct {
	title       *tview.TextView
	description *tview.TextView
}

func NewSelectedArgument() *SelectedArgument {
	s := &SelectedArgument{}
	s.title = s.buildTitle()
	s.description = s.buildDescription()

	return s
}

func (c *SelectedArgument) Select(option *man.Option) {
	c.title.SetText(text.ColorOption(1, option))
	c.description.SetText(text.FormatDescription(option.Description)).ScrollToBeginning()
}

func (c *SelectedArgument) SetClickFunc(opts *man.List, callback func(index int)) *SelectedArgument {
	c.description.SetRegionClickFunc(ClickFunc(opts, callback))

	return c
}

func (c *SelectedArgument) Layout() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(c.title, 3, 1, false).
		AddItem(c.description, 0, 1, true)
}

func (c *SelectedArgument) buildTitle() *tview.TextView {
	t := Title("Welcome!", 0, false)
	t.SetBorderPadding(0, 0, 2, 2)
	t.SetDynamicColors(true)
	t.SetRegions(true)
	t.SetRegionClickFunc(func(region string) {
		// do nothing
	})

	return t
}

func (c *SelectedArgument) buildDescription() *tview.TextView {
	activeOption := tview.NewTextView()
	activeOption.SetText("").
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetWordWrap(true).
		SetRegions(true)
	activeOption.SetBorderPadding(0, 0, 2, 2)
	activeOption.SetBorder(false)

	return activeOption
}
