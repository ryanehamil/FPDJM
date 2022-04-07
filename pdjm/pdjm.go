package pdjm

import (
	"fmt"

	"github.com/ryanehamil/FPDJM/logger"
)

type DeJong struct {
	Parameters
	NextCoordinate Coordinate2d
	Histogram
}

func (d *DeJong) Initialize() {
	d.Parameters.Random()
	d.NextCoordinate = Coordinate2d{1, 1}
}

func (d *DeJong) Run(count int) {
	for i := 0; i < count; i++ {
		ThisCoordinate := d.NextCoordinate
		d.NextCoordinate.X = deJongX(d.Parameters, ThisCoordinate)
		d.NextCoordinate.Y = deJongY(d.Parameters, ThisCoordinate)
		d.Histogram.Add(d.NextCoordinate)
	}
}

func (d *DeJong) String() {
	logger.Log(fmt.Sprintf("%v", d.Histogram))
}
