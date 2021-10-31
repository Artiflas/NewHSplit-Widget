package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create new app
	a := app.New()

	// New title and window
	w := a.NewWindow("NewHSplit Widget in Fyne")

	// Resize window
	w.Resize(fyne.NewSize(400, 400))

	// First widget
	label1 := canvas.NewText("Text1", color.Black)
	label2 := canvas.NewText("Text2", color.Black)
	w1 := widget.NewIcon(theme.CancelIcon())
	btn1 := widget.NewButton("Play ", func() {

	})

	// Setup content
	w.SetContent(
		container.NewHSplit(
			container.NewVBox(
				label1,
				w1,
			),
			// Second Section
			container.NewVBox(
				label2,
				btn1,
			),
		),
	)

	w.ShowAndRun()
}
