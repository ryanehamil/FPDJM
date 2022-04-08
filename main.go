package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"github.com/ryanehamil/FPDJM/pdjm"
)

// main function
func main() {
	// Create an new app
	a := app.New()

	// Create a new window
	w := a.NewWindow("Main Window")

	dejong := pdjm.DeJong{}
	dejong.Initialize()
	dejong.Run(1000000)

	raster := canvas.NewRaster(dejong.Plot)

	w.SetContent(raster)

	w.Resize(fyne.NewSize(800, 600))

	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeySpace {
			dejong.Initialize()
			dejong.Run(1000000)
			raster.Refresh()
		}
	})

	// Show the window
	w.ShowAndRun()

	// using fyne handle a click

}
