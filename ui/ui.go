package ui

import "github.com/rivo/tview"

type UI struct {
	App  *tview.Application
	Root tview.Primitive

	Left   tview.Primitive
	Right  tview.Primitive
	Top    tview.Primitive
	Middle tview.Primitive
	Bottom tview.Primitive
}

func New(app *tview.Application) *UI {
	u := &UI{
		App: app,
	}

	u.build()

	return u
}
