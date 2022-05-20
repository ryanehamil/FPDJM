package attractor

import (
	"fmt"
	"image"
	"image/color"

	"github.com/ryanehamil/FPDJM/logger"
	"github.com/ryanehamil/FPDJM/utils"
)

type System struct {
	Runs int
	Parameters
	NextCoordinate Coordinate2d
	Histogram
	// img *image.NRGBA
}

func New(attractorName string, defaultRuns int) (att *System) {
	switch expression := attractorName; expression {
	case "Peter de Jong":
		att = &System{}
		att.Runs = defaultRuns
		att.Initialize()
		return att
	}
	return nil

}

func (d *System) Initialize() {
	if d.Runs == 0 {
		d.Runs = 100
	}
	d.Parameters.Random()
	d.Histogram = Histogram{}
	d.NextCoordinate = Coordinate2d{1, 1}
}

func (dj *System) InitializeTemplate() {
	a, b, c, d := 0.970, -1.899, 1.381, -1.506
	dj.Parameters = Parameters{a, b, c, d}
	dj.Histogram = Histogram{}
	dj.NextCoordinate = Coordinate2d{1, 1}
}

func (d *System) Run() {
	for i := 0; i < d.Runs; i++ {
		d.NextCoordinate.X, d.NextCoordinate.Y = deJong(d.NextCoordinate.X, d.NextCoordinate.Y, d.Parameters.a, d.Parameters.b, d.Parameters.c, d.Parameters.d)
		d.Histogram.Add(d.NextCoordinate)
	}
}

func (d *System) Plot(w, h int) image.Image {

	d.Histogram.Scale(w, h)

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	for i := 0; i < len(d.Histogram.ScaledPoints); i++ {
		point := d.Histogram.ScaledPoints[i]
		colR32, colG32, colB32, _ := img.At(point.X, point.Y).RGBA()
		colR, colG, colB := uint8(colR32), uint8(colG32), uint8(colB32)
		img.Set(point.X, point.Y, color.RGBA{colR + 1, colG + 1, colB + 1, 255})
	}

	// Grade on a curve
	var maxBrightness uint8
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			colR32, _, _, _ := img.At(x, y).RGBA()
			if uint8(colR32) > maxBrightness {
				maxBrightness = uint8(colR32)
			}
		}
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			multiplier := (255 / maxBrightness) * 2
			colR32, colG32, colB32, _ := img.At(x, y).RGBA()
			colR, colG, colB := uint8(colR32), uint8(colG32), uint8(colB32)
			colR, colG, colB = utils.Min8(colR*multiplier, 255), utils.Min8(colG*multiplier, 255), utils.Min8(colB*multiplier, 255)
			img.Set(x, y, color.RGBA{colR, colG, colB, 255})
		}
	}

	return img
}

func (d *System) String() {
	logger.Log(fmt.Sprintln("DeJong Histogram OriginalPoints:"))
	logger.Log(fmt.Sprintf("%v", d.Histogram.OriginalPoints))
	logger.Log(fmt.Sprintln("DeJong Histogram ScaledPoints:"))
	logger.Log(fmt.Sprintf("%v", d.Histogram.ScaledPoints))
}
