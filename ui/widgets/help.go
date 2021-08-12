package widgets

import (
	"github.com/rivo/tview"
)

type Help struct {
	view       *tview.Flex
	debug      *tview.TextView
	printDebug bool
}

func NewHelp(printDebug bool) *Help {
	h := &Help{printDebug: printDebug}
	h.debug = tview.NewTextView()
	h.view = h.draw()

	return h
}

func (h *Help) SetDebug(debug string) {
	if h.printDebug {
		h.debug.SetText(h.debug.GetText(false) + "\n" + debug)
	}
}

func (h *Help) Layout() *tview.Flex {
	return h.view
}

func (h *Help) draw() *tview.Flex {
	help := tview.NewFlex().SetDirection(tview.FlexRow)
	help.AddItem(Title("Shortcuts", 1, true), 3, 1, false)
	help.AddItem(shortcut("/", "Change command"), 1, 1, false)
	help.AddItem(shortcut("[", "History back"), 1, 1, false)
	help.AddItem(shortcut("[", "History forward"), 1, 1, false)
	help.AddItem(tview.NewBox(), 1, 1, false)
	help.AddItem(shortcut("esc", "Back / Quit"), 1, 1, false)
	help.AddItem(tview.NewBox(), 3, 1, false)
	help.AddItem(h.debug, 0, 10, false)

	return help
}

func shortcut(shortcut, title string) *tview.TextView {
	shortcut = "[#4CF90F:#0C2D04:b] " + shortcut + " [#ffffff:#000000:d] - " + title
	text := tview.NewTextView().SetText(shortcut)
	text.SetDynamicColors(true)
	text.SetRegions(true)

	return text
}
