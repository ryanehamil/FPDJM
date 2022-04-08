package pdjm

import (
	"fmt"
	"image"
	"image/color"

	"github.com/ryanehamil/FPDJM/logger"
)

type DeJong struct {
	Parameters
	NextCoordinate Coordinate2d
	Histogram
	// img *image.NRGBA
}

func (d *DeJong) Initialize() {
	d.Parameters.Random()
	d.Histogram = Histogram{}
	d.NextCoordinate = Coordinate2d{1, 1}
}

func (d *DeJong) Run(count int) {
	logger.Log(fmt.Sprintln("DeJong Running..."))
	for i := 0; i < count; i++ {
		ThisCoordinate := d.NextCoordinate
		d.NextCoordinate.X = deJongX(d.Parameters, ThisCoordinate)
		d.NextCoordinate.Y = deJongY(d.Parameters, ThisCoordinate)
		d.Histogram.Add(d.NextCoordinate)

		// if math.Mod(float64(i), 100000) == 0 {
		// 	logger.Log(fmt.Sprintf("Runs Complete: %d/%d", i, count))
		// }
	}

	// d.Histogram.Sort()

	logger.Log(fmt.Sprintln("DeJong Complete"))
}

func (d *DeJong) Plot(w, h int) image.Image {

	d.Histogram.Scale(w, h)

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	for i := 0; i < len(d.Histogram.ScaledPoints); i++ {
		col, _, _, _ := img.At(d.Histogram.ScaledPoints[i].X, d.Histogram.ScaledPoints[i].Y).RGBA()
		col8 := uint8(int(col) + d.Histogram.ScaledPoints[i].Z*10)
		img.Set(d.Histogram.ScaledPoints[i].X, d.Histogram.ScaledPoints[i].Y, color.RGBA{col8, col8, col8, 255})
	}

	return img
}

func (d *DeJong) String() {
	logger.Log(fmt.Sprintln("DeJong Histogram OriginalPoints:"))
	logger.Log(fmt.Sprintf("%v", d.Histogram.OriginalPoints))
	logger.Log(fmt.Sprintln("DeJong Histogram ScaledPoints:"))
	logger.Log(fmt.Sprintf("%v", d.Histogram.ScaledPoints))
}
