package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"github.com/ryanehamil/FPDJM/logger"
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

	logger.Log("Starting DeJong attractor")
	dejong.Run(200000)
	logger.Log("DeJong attractor complete")

	// // Create a new label
	// l := widget.NewLabel("Nothing here")

	// // Add the label to the window
	// w.SetContent(l)

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {

			if (dejong.Histogram.CurrentScale != pdjm.Coordinate2dInt{X: w, Y: h} && w > 10 && h > 10) {
				dejong.Histogram.CurrentScale = pdjm.Coordinate2dInt{X: w, Y: h}
				dejong.Histogram.Scale(w, h)
			}

			// check if the x,y is in the histogram and return the z
			for i := 0; i < len(dejong.Histogram.ScaledPoints); i++ {
				if dejong.Histogram.ScaledPoints[i].X == x && dejong.Histogram.ScaledPoints[i].Y == y {
					return color.RGBA{uint8(0), uint8(0), uint8(math.Min(float64(dejong.Histogram.ScaledPoints[i].Z*200), 255.0)), 0xff}
				}
			}

			return color.RGBA{
				uint8(0),
				uint8(0),
				uint8(0),
				0xff}
		})

	w.SetContent(raster)

	w.Resize(fyne.NewSize(800, 600))

	// Show the window
	w.ShowAndRun()

}
