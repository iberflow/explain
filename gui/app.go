package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ignasbernotas/explain/config"
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"github.com/ignasbernotas/explain/text"
	"github.com/rivo/tview"
	"strings"
)

type Widgets struct {
	sidebar *tview.List

	commandLine    *tview.TextView
	commandOptions *tview.Flex
	commandForm    *tview.Modal

	optionTitle       *tview.TextView
	optionDescription *tview.TextView

	pages *tview.Pages
}

func NewWidgets() *Widgets {
	return &Widgets{}
}

type App struct {
	gui     *tview.Application
	widgets *Widgets

	documentationOptions *man.List

	command        *args.Command
	commandOptions *man.List

	activeOption *man.Option
}

const showBorders = false

func NewApp(documentationOptions *man.List, command *args.Command, commandOptions *man.List) *App {
	return &App{
		documentationOptions: documentationOptions,

		command:        command,
		commandOptions: commandOptions,

		gui:     tview.NewApplication(),
		widgets: NewWidgets(),
	}
}

func (a *App) Draw() {
	a.widgets.sidebar = a.sidebar()
	a.widgets.commandLine = a.commandLine()
	a.widgets.commandOptions = a.optionList()
	a.widgets.optionDescription = a.currentOption()
	a.widgets.commandForm = a.commandForm()
	a.widgets.optionTitle = a.optionTitle()
	a.widgets.pages = a.buildPages()

	a.setupKeyBindings()

	if err := a.gui.
		SetRoot(a.widgets.pages, true).
		EnableMouse(true).
		Run(); err != nil {
		panic(err)
	}
}

func (a *App) optionTitle() *tview.TextView {
	t := titleWidget("Welcome!", 0, false)
	t.SetBorderPadding(0, 0, 2, 2)

	return t
}

func (a *App) buildPages() *tview.Pages {
	sidebar := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleWidget("Options", 1, true), 3, 1, false).
		AddItem(a.widgets.sidebar, 0, 1, true)

	content := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(titleWidget("Command", 1, true), 3, 0, false).
		AddItem(a.widgets.commandLine, 3, 0, false).
		AddItem(a.widgets.optionTitle, 1, 0, false).
		AddItem(a.widgets.optionDescription, 0, 4, false).
		AddItem(a.widgets.commandOptions, 0, 6, false)

	changeCommand := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.widgets.commandForm, 0, 1, true)

	container := tview.NewFlex()
	container.AddItem(sidebar, 25, 1, true)
	container.AddItem(content, 0, 5, false)

	pages := tview.NewPages()
	pages.AddPage("dashboard", container, true, true)
	pages.AddPage("changeCommand", changeCommand, true, false)

	return pages
}

func (a *App) setupKeyBindings() {
	a.gui.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			pageName, _ := a.widgets.pages.GetFrontPage()
			if pageName == "dashboard" {
				a.gui.Stop()
			} else {
				a.widgets.pages.SwitchToPage("dashboard")
			}

			return nil
		}
		if event.Key() == tcell.KeyCtrlQ {
			return nil
		}
		if event.Rune() == '?' {
			a.widgets.pages.SwitchToPage("changeCommand")

			return nil
		}
		return event
	})
}

func (a *App) commandForm() *tview.Modal {
	var cmd *args.Command
	changed := func(text string) {
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			return
		}
		commands := args.Parse(text)
		cmd = commands[0]
	}

	modal := tview.NewModal().
		AddInputText([]string{"Command: "}, a.command.String(), changed).
		SetText("Edit command").
		AddButtons([]string{"Save", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Save" {
				if cmd != nil {
					a.command = cmd
				}

				a.widgets.commandLine.Clear()
				a.widgets.commandLine.SetText(text.RenderCommand(a.command, a.documentationOptions))
				a.widgets.pages.SwitchToPage("dashboard")
			}
		})

	return modal
}

func (a *App) commandLine() *tview.TextView {
	cmd := tview.NewTextView()
	cmd.SetText(text.RenderCommand(a.command, a.documentationOptions)).
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true).
		SetTextAlign(1)

	cmd.SetRegionClickFunc(a.regionClickFunc())

	cmd.SetBorder(showBorders).
		SetBorderPadding(0, 0, 2, 2)

	return cmd
}

func (a *App) currentOption() *tview.TextView {
	activeOption := tview.NewTextView()
	activeOption.SetText("").
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetWordWrap(true).
		SetRegions(true)
	activeOption.SetBorderPadding(0, 0, 2, 2)
	activeOption.SetBorder(showBorders)
	activeOption.SetRegionClickFunc(a.regionClickFunc())

	return activeOption
}

func (a *App) optionList() *tview.Flex {
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow)

	flex.SetBorderPadding(1, 4, 0, 0)

	for i, opt := range a.commandOptions.Options() {
		optionBox := tview.NewTextView()
		optionBox.SetBorderPadding(0, 0, 2, 2)
		optionBox.SetRegionClickFunc(a.regionClickFunc())
		optionBox.SetText(text.FormatDescription(strings.TrimSpace(opt.Description)))
		optionBox.SetBorder(false)
		optionBox.SetToggleHighlights(true).
			SetDynamicColors(true).
			SetRegions(true)

		titleText := text.Underline(text.MarkRegion(i, opt.String(), true))
		title := titleWidget("◉ "+titleText, 1, false)
		title.SetBorderPadding(1, 0, 2, 2)
		title.SetTextColor(tcell.GetColor(config.FlagColor))
		title.SetRegionClickFunc(a.regionClickFunc())
		title.SetToggleHighlights(true).
			SetDynamicColors(true).
			SetRegions(true)

		title.SetBorder(false)

		flex.AddItem(title, 2, 1, false)
		flex.AddItem(optionBox, 0, 1, true)
	}

	return flex
}

func (a *App) sidebar() *tview.List {
	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(showBorders)
	list.KeepSelectedItemInView(false)
	list.SetSelectOnNavigation(true)
	list.SetSelectedFocusOnly(false)
	list.SetHighlightFullLine(true)
	list.SetBorderPadding(0, 1, 3, 3)
	list.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		opts := a.documentationOptions.Options()
		a.widgets.optionDescription.
			SetText(text.FormatDescription(opts[i].Description)).
			ScrollToBeginning()

		a.widgets.optionTitle.
			SetTextColor(tcell.GetColor(config.FlagColor)).
			SetText(opts[i].String())
	})

	for _, opt := range a.documentationOptions.Options() {
		list.AddItem(opt.String(), opt.Description, 0, nil)
	}

	return list
}

func (a *App) regionClickFunc() func(region string) {
	return func(region string) {
		region = text.StripColor(region)
		found := a.documentationOptions.Search(region)
		if found > 0 {
			a.widgets.sidebar.SetCurrentItem(found)
			a.widgets.sidebar.SetOffset(found, 0)
		}
	}
}
