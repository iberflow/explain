package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
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

	manPageOptions []*man.Option
	activeOption   *man.Option
}

const showBorders = false

func NewApp(sidebarItems []*man.Option, command *args.Command) *App {
	return &App{
		manPageOptions: sidebarItems,
		command:        command,
		gui:            tview.NewApplication(),
		widgets:        NewWidgets(),
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
	a.widgets.optionDescription = a.getActiveOptionBox()
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

func (a *App) getActiveOptionBox() *tview.TextView {
	activeOption := tview.NewTextView()
	activeOption.SetText("Select an argument").
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true)

	activeOption.SetBorder(showBorders).
		SetBorderPadding(0, 0, 2, 2)

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

func (a *App) getCommandOptionsBox() *tview.List {
	options := tview.NewList().ShowSecondaryText(false)
	options.SetBorder(showBorders)
	options.SetSelectedFocusOnly(false)
	options.SetHighlightFullLine(true)
	options.SetBorderPadding(1, 1, 3, 3)
	options.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		a.widgets.optionDescription.SetText(a.manPageOptions[i].Description)
		a.widgets.optionTitle.
			SetTextColor(tcell.GetColor(args.FlagColor)).
			SetText(a.manPageOptions[i].String())
	})

	for _, opt := range a.manPageOptions {
		if len(opt.Name) == 0 {

		}

		options.AddItem(opt.String(), opt.Description, 0, nil)
	}

	return options
}
