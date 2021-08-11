package ui

import (
	"github.com/ignasbernotas/explain/ui/widgets"
	"github.com/rivo/tview"
)

type Widgets struct {
	sidebar *widgets.Sidebar

	commandLine    *widgets.CommandLine
	commandOptions *widgets.CommandOptions
	commandForm    *tview.Modal

	selectedArgument *widgets.SelectedArgument

	pages *Pages

	help *widgets.Help
}

func NewWidgets() *Widgets {
	return &Widgets{}
}
