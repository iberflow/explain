package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ignasbernotas/explain/config"
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
	"github.com/rivo/tview"
	"strings"
)

type CommandOptions struct {
	view      *tview.Flex
	clickFunc func(region string)
}

func NewCommandOptions() *CommandOptions {
	o := &CommandOptions{}
	o.view = o.build()

	return o
}

func (c *CommandOptions) Layout() *tview.Flex {
	return c.view
}

func (c *CommandOptions) SetClickFunc(opts *man.List, callback OptionClickFunc) *CommandOptions {
	c.clickFunc = ClickFunc(opts, callback)

	return c
}

func (c *CommandOptions) Clear() *CommandOptions {
	c.view.Clear()

	return c
}

func (c *CommandOptions) SetOptions(list *man.List) *CommandOptions {
	c.Clear()

	list = list.Unique(false)
	list.Sort()

	for i, u := range list.Options() {
		c.addOption(i, u)
	}

	return c
}

func (c *CommandOptions) addOption(index int, option *man.Option) *CommandOptions {
	optionBox := tview.NewTextView()
	optionBox.SetBorderPadding(0, 0, 2, 2)
	optionBox.SetText(text.FormatDescription(strings.TrimSpace(option.Description)))
	optionBox.SetRegionClickFunc(c.clickFunc)
	optionBox.SetWordWrap(true)
	optionBox.SetBorder(false)
	optionBox.SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true)

	titleText := text.ColorOption(index, option)
	title := Title("◉ "+titleText, 1, false)
	title.SetBorderPadding(1, 0, 2, 2)
	title.SetRegionClickFunc(c.clickFunc)
	title.SetTextColor(tcell.GetColor(config.FlagColor))
	title.SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true)

	title.SetBorder(false)

	c.view.AddItem(title, 2, 1, false)
	c.view.AddItem(optionBox, 0, 1, true)

	return c
}

func (c *CommandOptions) build() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorderPadding(1, 4, 0, 0)

	return flex
}
