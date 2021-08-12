package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ignasbernotas/explain/ui/history"
	"github.com/ignasbernotas/explain/ui/widgets"
	"github.com/rivo/tview"
)

type App struct {
	gui       *tview.Application
	widgets   *Widgets
	processor *Processor
	history   *history.History
}

func NewApp(processor *Processor) *App {
	return &App{
		gui:       tview.NewApplication(),
		widgets:   NewWidgets(),
		processor: processor,
		history:   history.NewHistory(),
	}
}

func (a *App) Draw() {
	a.widgets.help = widgets.NewHelp(false)
	a.widgets.sidebar = a.sidebar()
	a.widgets.commandLine = a.commandLine()
	a.widgets.selectedArgument = a.selectedArgument()
	a.widgets.commandOptions = a.commandOptions()
	a.widgets.commandForm = a.commandForm()
	a.widgets.pages = a.buildPages()
	a.setupKeyBindings()
	a.setupHistory()

	if len(a.processor.command.String()) == 0 {
		a.widgets.pages.Show(PageCommandEdit)
	}

	if err := a.gui.
		SetRoot(a.widgets.pages.Layout(), true).
		EnableMouse(true).
		Run(); err != nil {
		panic(err)
	}
}

func (a *App) buildPages() *Pages {
	pages := NewPages()

	content := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.widgets.commandLine.Layout(), 0, 1, false).
		AddItem(a.widgets.selectedArgument.Layout(), 0, 2, false).
		AddItem(a.widgets.commandOptions.Layout(), 0, 4, false)

	dashboard := tview.NewFlex()
	dashboard.AddItem(a.widgets.sidebar.Layout(), 25, 1, true)
	dashboard.AddItem(content, 0, 5, false)
	dashboard.AddItem(a.widgets.help.Layout(), 0, 1, false)

	pages.Add(PageDashboard, dashboard)

	changeCommand := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.widgets.commandForm, 0, 1, true)

	pages.Add(PageCommandEdit, changeCommand)
	pages.Show(PageDashboard)

	return pages
}

func (a *App) sidebar() *widgets.Sidebar {
	sidebar := widgets.NewSidebar()
	sidebar.SetSelectionFunc(func(index int) {
		opts := a.processor.DocumentationOptions().Options()

		a.widgets.selectedArgument.Select(opts[index])
		a.history.Add(history.OptionSelect, index)
	})
	sidebar.SetOptions(a.processor.DocumentationOptions())
	a.widgets.sidebar = sidebar

	return sidebar
}

func (a *App) selectedArgument() *widgets.SelectedArgument {
	opt := a.processor.CommandOptions().First()

	arg := widgets.NewSelectedArgument()
	if opt != nil {
		arg.Select(opt)
	}
	arg.SetClickFunc(a.processor.DocumentationOptions(), func(index int) {
		a.widgets.sidebar.Select(index)
	})

	return arg
}

func (a *App) commandLine() *widgets.CommandLine {
	line := widgets.NewCommandLine()
	line.SetCommand(a.processor.Command(), a.processor.DocumentationOptions())
	line.SetClickFunc(a.processor.DocumentationOptions(), func(index int) {
		a.widgets.sidebar.Select(index)
	})

	return line
}

func (a *App) setupKeyBindings() {
	a.gui.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			if a.widgets.pages.IsPage(PageDashboard) {
				a.gui.Stop()
			} else {
				a.widgets.pages.Show(PageDashboard)
			}

			return nil
		}
		if event.Key() == tcell.KeyCtrlQ {
			return nil
		}
		if event.Rune() == '[' || event.Key() == tcell.RuneLArrow {
			a.history.GoBack()
		}
		if event.Rune() == ']' || event.Key() == tcell.RuneRArrow {
			a.history.GoForward()
		}
		if event.Rune() == '/' {
			a.widgets.pages.Show(PageCommandEdit)

			return nil
		}
		return event
	})
}

func (a *App) setupHistory() {
	a.history.SetListenFunc(func(i *history.Item) {
		// cursor move events
		a.widgets.help.SetDebug(fmt.Sprintf("current index: %d, items: %d", a.history.Index(), a.history.Count()))

		if i.Type == history.CursorMove {
			if itm := a.history.GetItem(); itm != nil {
				if itm.Type == history.OptionSelect {
					a.widgets.help.SetDebug(fmt.Sprintf("Option selected: %d", itm.Value.(int)))
					a.history.DisableNextEvent()
					a.widgets.sidebar.Select(itm.Value.(int))
				}
				if itm.Type == history.CommandChange {
					a.updateCommand(itm.Value.(string))
				}
			}
		}
	})
	a.history.Add(history.OptionSelect, 0)
}

func (a *App) commandForm() *tview.Modal {
	var cmd string
	changed := func(text string) {
		cmd = text
	}

	modal := tview.NewModal().
		AddInputText([]string{"Command: "}, a.processor.Command().String(), changed).
		SetText("Edit command").
		AddButtons([]string{"Save", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Save" {
				a.history.Add(history.CommandChange, a.processor.Command().String())
				a.updateCommand(cmd)
				a.history.Add(history.CommandChange, a.processor.Command().String())
			}
		})

	return modal
}

func (a *App) commandOptions() *widgets.CommandOptions {
	opts := widgets.NewCommandOptions()
	opts.SetClickFunc(a.processor.DocumentationOptions(), func(index int) {
		a.widgets.sidebar.Select(index)
	})
	opts.SetOptions(a.processor.CommandOptions())

	return opts
}

func (a *App) updateCommand(cmd string) {
	if len(cmd) > 0 {
		a.processor.LoadCommand(cmd)
	}
	a.widgets.sidebar.SetOptions(a.processor.DocumentationOptions())
	a.widgets.sidebar.SetOptions(a.processor.DocumentationOptions())
	a.widgets.commandLine.SetCommand(a.processor.Command(), a.processor.DocumentationOptions())
	a.widgets.commandOptions.SetOptions(a.processor.CommandOptions())
	a.widgets.pages.Show(PageDashboard)
}
