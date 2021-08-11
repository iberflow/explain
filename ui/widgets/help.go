package widgets

import (
	"github.com/rivo/tview"
)

type Help struct {
	view *tview.Flex
}

func NewHelp() *Help {
	h := &Help{}
	h.view = h.draw()

	return h
}

func (h *Help) Layout() *tview.Flex {
	return h.view
}

func (h *Help) draw() *tview.Flex {
	help := tview.NewFlex().SetDirection(tview.FlexRow)
	help.AddItem(Title("Shortcuts", 1, true), 3, 1, false)
	help.AddItem(shortcut("?", "Change command"), 1, 1, false)
	help.AddItem(shortcut("[", "Go back"), 1, 1, false)
	help.AddItem(shortcut("[", "Go forward"), 1, 1, false)

	return help
}

func shortcut(shortcut, title string) *tview.TextView {
	shortcut = "[#4CF90F:#0C2D04:b] " + shortcut + " [#ffffff:#000000:d] - " + title
	text := tview.NewTextView().SetText(shortcut)
	text.SetDynamicColors(true)
	text.SetRegions(true)

	return text
}
