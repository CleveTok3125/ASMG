package ui

import "github.com/rivo/tview"

func (u *UI) build() {
	u.Left = u.buildLeft()
	u.Right = u.buildRight()
	u.Top = u.buildTop()
	u.Middle = u.buildMiddle()
	u.Bottom = u.buildBottom()

	center := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(u.Top, 0, 1, false).
		AddItem(u.Middle, 0, 3, false).
		AddItem(u.Bottom, 0, 1, false)

	u.Root = tview.NewFlex().
		AddItem(u.Left, 0, 1, false).
		AddItem(center, 0, 2, false).
		AddItem(u.Right, 0, 1, false)
}
