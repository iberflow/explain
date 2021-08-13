package widgets

import (
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
	"github.com/rivo/tview"
	"strings"
)

type SelectedArgument struct {
	title       *tview.TextView
	description *tview.TextView
	arguments   *tview.TextView
	argumentsFormat   *tview.TextView
}

func NewSelectedArgument() *SelectedArgument {
	s := &SelectedArgument{}
	s.title = s.buildTitle()
	s.description = s.buildDescription()
	s.arguments = s.buildArguments()
	s.argumentsFormat = s.buildArgumentFormats()

	return s
}

func (c *SelectedArgument) Select(option *man.Option) {
	c.title.SetText(text.ColorOption(1, option))
	c.description.SetText(text.FormatDescription(option.Description)).ScrollToBeginning()

	title := `[::d]Argument formats:`
	if len(option.Parameters) < 2 {
		title = ""
	}
	c.argumentsFormat.SetText(title)
	c.arguments.SetText(drawArgumentList(option.Parameters)).ScrollToBeginning()
}

func (c *SelectedArgument) SetClickFunc(opts *man.List, callback func(index int)) *SelectedArgument {
	c.description.SetRegionClickFunc(ClickFunc(opts, callback))

	return c
}

func (c *SelectedArgument) Layout() *tview.Flex {
	args := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(c.argumentsFormat, 3, 1, false).
		AddItem(c.arguments, 0, 5, true)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(c.title, 1, 1, false).
		AddItem(c.description, 0, 1, true)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(content, 0, 6, false).
		AddItem(args, 0, 3, true)

	return layout
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

func (c *SelectedArgument) buildArgumentFormats() *tview.TextView {
	formatTitle := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[::d]Argument formats`)
	formatTitle.SetBorderPadding(0, 0, 2, 2)

	return formatTitle
}

func (c *SelectedArgument) buildArguments() *tview.TextView {
	args := tview.NewTextView()
	args.SetText("").
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetWordWrap(true).
		SetRegions(true)
	args.SetBorderPadding(0, 0, 2, 2)
	args.SetBorder(false)

	return args
}

func drawArgumentList(args []string) string {
	if len(args) < 2 {
		return ""
	}
	var str []string
	for i := 0; i < len(args); i++ {
		str = append(str, "[::d]<"+args[i]+`>`)
	}

	return strings.Join(str, "\n")
}
