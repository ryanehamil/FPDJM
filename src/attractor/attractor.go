package attractor

import (
	"image"
	"time"
	"sync"
	"image/color"

)



type System struct {
	Running bool
	Runs int
	Parameters
	NextCoordinate Coordinate2d
	Histogram
	Width int
	Height int
	img *image.RGBA
	imgRuns int
}

func New(attractorName string) (att *System) {
	switch expression := attractorName; expression {
	case "Peter de Jong":
		att = &System{}
		att.InitializeTemplate()
		return att
	}
	return nil

}

func (d *System) Initialize() {
	d.Parameters.Random()
	d.Histogram = Histogram{}
	d.NextCoordinate = Coordinate2d{1, 1}
}

func (dj *System) InitializeTemplate() {
	a, b, c, d := 2.38767, -1.22713, -0.39595, -4.67104
	dj.Parameters = Parameters{a, b, c, d}
	dj.Histogram = Histogram{}
	dj.NextCoordinate = Coordinate2d{1, 1}
}

func (d *System) Start() {
	d.Running = true

	// Start the goroutine with timer
	go func() {
		for d.Running {
			d.Run(100000)
			

			// Sleep for a second
			time.Sleep(time.Second)
		} // end for
	}() // end goroutine
	
}
func (d *System) Stop()	{
	d.Running = false
}

func (d *System) Run(count int) {
	for i := 0; i < count; i++ {
		d.NextCoordinate.X, d.NextCoordinate.Y = deJong(d.NextCoordinate.X, d.NextCoordinate.Y, d.Parameters.a, d.Parameters.b, d.Parameters.c, d.Parameters.d)
		d.Histogram.Add(d.NextCoordinate)
		d.Runs++
	}
	
}



func (d *System) Plot(w, h int) image.Image {
	if d.Width == w && d.Height == h && d.imgRuns == d.Runs {
		return d.img
	} else if d.Width == w && d.Height == h && d.imgRuns != d.Runs {
		img := d.img

		var wg sync.WaitGroup
		for key, element := range d.Histogram.ScaledMap {
			wg.Add(1)
			normalized := (float64(element) / float64(d.Histogram.MaxZ)) * 255.0
			go func(key Coordinate2dInt, element int, normalized float64) {
				defer wg.Done()
				img.Set(key.X, key.Y, color.RGBA{uint8(normalized), uint8(normalized), uint8(normalized), 255})
			}(key, element, normalized)
		}
		wg.Wait()

		d.Histogram = Histogram{}

		d.img = img

		return img

	} else {
		
		d.Histogram.Scale(w, h)
		img := image.NewRGBA(image.Rect(0, 0, w, h))

		// Reset the image to white
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				img.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}

		// loop through the scaled map and set the pixels to the color of the scaled map
		// Send this to multiple goroutines to speed up the process
		var wg sync.WaitGroup
		for key, element := range d.Histogram.ScaledMap {
			wg.Add(1)
			normalized := (float64(element) / float64(d.Histogram.MaxZ)) * 255.0
			go func(key Coordinate2dInt, element int, normalized float64) {
				defer wg.Done()
				img.Set(key.X, key.Y, color.RGBA{uint8(normalized), uint8(normalized), uint8(normalized), 255})
			}(key, element, normalized)
		}
		wg.Wait()

		d.Histogram = Histogram{}

		
		d.img = img

		return img
	}
}