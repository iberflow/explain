package ui

import "github.com/rivo/tview"

const PageDashboard = "dashboard"
const PageCommandEdit = "commandEdit"
const PageSearch = "search"

type Pages struct {
	pages *tview.Pages
}

func NewPages() *Pages {
	return &Pages{
		pages: tview.NewPages(),
	}
}

func (p *Pages) IsPage(name string) bool {
	pageName, _ := p.pages.GetFrontPage()

	return pageName == name
}

func (p *Pages) Add(name string, primitive tview.Primitive) *Pages {
	p.pages.AddPage(name, primitive, true, false)

	return p
}

func (p *Pages) Show(name string) *Pages {
	p.pages.SwitchToPage(name)

	return p
}

func (p *Pages) Layout() *tview.Pages {
	return p.pages
}
