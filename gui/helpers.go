package gui

import "github.com/rivo/tview"

func titleWidget(title string, pY int, alignMiddle bool) *tview.TextView {
	widget := tview.NewTextView().SetText(title)
	widget.SetBorder(false)
	widget.SetBorderPadding(pY, pY, 1, 1)

	align := 0
	if alignMiddle {
		align = 1
	}

	widget.SetTextAlign(align)

	return widget
}