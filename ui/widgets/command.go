package widgets

import (
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
	"github.com/rivo/tview"
)

type CommandLine struct {
	view    *tview.TextView
	options *man.List
}

func NewCommandLine() *CommandLine {
	c := &CommandLine{}
	c.view = c.build()

	return c
}

func (c *CommandLine) SetCommand(command *args.Command, opts *man.List) *CommandLine {
	c.view.SetText(text.RenderCommand(command, opts))

	return c
}

func (c *CommandLine) SetClickFunc(opts *man.List, callback OptionClickFunc) *CommandLine {
	fn := ClickFunc(opts, callback)
	c.view.SetRegionClickFunc(fn)

	return c
}

func (c *CommandLine) Layout() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(Title("Command", 1, true), 3, 1, false).
		AddItem(c.view, 0, 1, true)
}

func (c *CommandLine) build() *tview.TextView {
	cmd := tview.NewTextView()
	cmd.SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true).
		SetTextAlign(1)

	cmd.SetBorder(false).
		SetBorderPadding(0, 0, 2, 2)

	return cmd
}
