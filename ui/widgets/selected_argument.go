package widgets

import (
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
	"github.com/rivo/tview"
	"strings"
)

type SelectedArgument struct {
	layout          *tview.Flex
	optionArguments *tview.Flex

	name            *tview.TextView
	description     *tview.TextView
	arguments       *tview.TextView
	argumentsFormat *tview.TextView
}

func NewSelectedArgument() *SelectedArgument {
	s := &SelectedArgument{}
	s.name = s.buildTitle()
	s.description = s.buildDescription()
	s.arguments = s.buildArguments()
	s.argumentsFormat = s.buildArgumentFormats()
	s.optionArguments = s.buildOptionArgs()
	s.layout = s.buildLayout()

	return s
}

func (s *SelectedArgument) Select(option *man.Option) {
	s.name.SetText(text.ColorOption(1, option))
	s.description.SetText(text.FormatDescription(option.Description)).ScrollToBeginning()

	title := `[::d]Argument formats:`
	if len(option.Parameters) < 2 {
		title = ""
		s.layout.ResizeItem(s.optionArguments, 0,0)
	} else {
		s.layout.ResizeItem(s.optionArguments, 0,3)
	}
	s.argumentsFormat.SetText(title)
	s.arguments.SetText(drawArgumentList(option.Parameters)).ScrollToBeginning()
}

func (s *SelectedArgument) SetClickFunc(opts *man.List, callback func(index int)) *SelectedArgument {
	s.description.SetRegionClickFunc(ClickFunc(opts, callback))

	return s
}

func (s *SelectedArgument) buildLayout() *tview.Flex {
	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(s.name, 1, 1, false).
		AddItem(s.description, 0, 1, true)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(content, 0, 6, false).
		AddItem(s.optionArguments, 0, 3, true)

	return layout
}

func (s *SelectedArgument) Layout() *tview.Flex {
	return s.layout
}

func (s *SelectedArgument) buildOptionArgs() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(s.argumentsFormat, 3, 1, false).
		AddItem(s.arguments, 0, 5, true)
}

func (s *SelectedArgument) buildTitle() *tview.TextView {
	t := Title("Welcome!", 0, false)
	t.SetBorderPadding(0, 0, 2, 2)
	t.SetDynamicColors(true)
	t.SetRegions(true)
	t.SetRegionClickFunc(func(region string) {
		// do nothing
	})

	return t
}

func (s *SelectedArgument) buildDescription() *tview.TextView {
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

func (s *SelectedArgument) buildArgumentFormats() *tview.TextView {
	formatTitle := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[::d]Argument formats`)
	formatTitle.SetBorderPadding(0, 0, 2, 2)

	return formatTitle
}

func (s *SelectedArgument) buildArguments() *tview.TextView {
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
