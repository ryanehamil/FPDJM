package fyre

import (
	"time"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/ryanehamil/FPDJM/src/attractor"
)

type fyreApp struct {
	a   fyne.App
	w   fyne.Window
	t  *canvas.Text
	r   *canvas.Raster
	att *attractor.System
}

func (fa *fyreApp) AddRaster(raster *canvas.Raster) {
	fa.w.SetContent(raster)
}

func (fa *fyreApp) Run() {

	fa.w.Resize(fyne.NewSize(800, 600))

	fa.att.Start()

	fa.w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeySpace {
			if (fa.att.Running){ fa.att.Stop() }
			start := time.Now()
			fa.t.Text = "regenerating"
			fa.t.Refresh()
			fa.att.Initialize()
			fa.att.Start()
			fa.r.Refresh()
			fa.t.Text = "regenerated in " + time.Since(start).String() + " settings: " + fa.att.Parameters.String()
			fa.t.Refresh()
		}
	})

	// Show the window
	fa.w.ShowAndRun()

}

func New(att *attractor.System) (fa fyreApp) {

	fa.a = app.New()
	fa.w = fa.a.NewWindow("Main Window")
	fa.att = att

	fa.r = canvas.NewRaster(fa.att.Plot)
	fa.r.SetMinSize(fyne.NewSize(800, 800))
	
	// fa.AddRaster(fa.r)

	fa.t = canvas.NewText("Text", color.White)

	content := container.New(layout.NewVBoxLayout(), fa.t, fa.r)

	fa.w.SetContent( content)

	return fa
}