package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ignasbernotas/explain/config"
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
	"github.com/rivo/tview"
)

type Widgets struct {
	options           *tview.List
	command           *tview.TextView
	optionDescription *tview.TextView
	optionTitle       *tview.TextView
}

func NewWidgets() *Widgets {
	return &Widgets{}
}

type App struct {
	gui     *tview.Application
	widgets *Widgets

	command *args.Command

	documentationOptions *man.OptionList
	commandOptions       *man.OptionList

	activeOption *man.Option
}

const showBorders = false

func NewApp(documentationOptions *man.OptionList, command *args.Command, commandOptions *man.OptionList) *App {
	return &App{
		documentationOptions: documentationOptions,
		command:              command,

		commandOptions: commandOptions,

		gui:     tview.NewApplication(),
		widgets: NewWidgets(),
	}
}

func titleWidget(title string, pY int) *tview.TextView {
	widget := tview.NewTextView().SetText(title)
	widget.SetBorder(showBorders)
	widget.SetBorderPadding(pY, pY, 1, 1)
	widget.SetTextAlign(1)

	return widget
}

func (a *App) Draw() {
	a.widgets.options = a.getCommandOptionsBox()
	a.widgets.command = a.getCommandBox()
	a.widgets.optionDescription = a.getOptionDescriptionBox()
	a.widgets.optionTitle = titleWidget("Current command", 0)

	flex := tview.NewFlex().
		AddItem(
			tview.NewFlex().
				SetDirection(tview.FlexRow).
				AddItem(titleWidget("Options", 1), 4, 1, false).
				AddItem(a.widgets.options, 0, 1, true),
			0, 1, true,
		).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(titleWidget("Command", 1), 4, 1, false).
				AddItem(a.widgets.command, 5, 1, false).
				AddItem(a.widgets.optionTitle, 3, 1, false).
				AddItem(a.widgets.optionDescription, 0, 1, false),
			0, 3, false,
		)

	if err := a.gui.
		SetRoot(flex, true).
		EnableMouse(true).
		Run(); err != nil {
		panic(err)
	}
}

func (a *App) getOptionDescriptionBox() *tview.TextView {
	activeOption := tview.NewTextView()
	activeOption.SetText("Select an argument").
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetWordWrap(true).
		SetRegions(true)

	activeOption.SetBorder(showBorders).
		SetBorderPadding(0, 0, 2, 2)

	activeOption.SetRegionClickFunc(func(region string) {
		region = text.StripColor(region)
		found := a.documentationOptions.Search(region)
		if found > 0 {
			activeOption.Clear()
			a.widgets.options.SetCurrentItem(found)
			a.widgets.options.SetOffset(found, 0)
		}
	})

	activeOption.Highlight("")

	return activeOption
}

func (a *App) getCommandBox() *tview.TextView {
	command := tview.NewTextView()
	command.SetText(a.command.StringRegions()).
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true)

	command.SetBorder(showBorders).
		SetBorderPadding(1, 1, 2, 2)

	return command
}

func (a *App) getActiveCommandOptionsBox() *tview.Flex {
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleWidget("Options", 1), 4, 1, false).
		AddItem(a.widgets.options, 0, 1, true)

	return flex
}

func (a *App) getCommandOptionsBox() *tview.List {
	options := tview.NewList().ShowSecondaryText(false)
	options.SetBorder(showBorders)
	options.KeepSelectedItemInView(false)
	options.SetSelectOnNavigation(true)
	options.SetSelectedFocusOnly(false)
	options.SetHighlightFullLine(true)
	options.SetBorderPadding(1, 1, 3, 3)
	options.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		a.widgets.optionDescription.
			SetText(text.FormatDescription(a.documentationOptions.Options()[i].Description)).
			ScrollToBeginning()

		a.widgets.optionTitle.
			SetTextColor(tcell.GetColor(config.FlagColor)).
			SetText(a.documentationOptions.Options()[i].String())
	})

	for _, opt := range a.documentationOptions.Options() {
		options.AddItem(opt.String(), opt.Description, 0, nil)
	}

	return options
}
