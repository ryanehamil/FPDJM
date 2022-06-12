package attractor

type Histogram struct {
	Points int
	MaxZ int
	CurrentScale	Coordinate2dInt
	OriginalMap		map[Coordinate2d]int
	ScaledMap		map[Coordinate2dInt]int
}

// Function to add a coordinate to the histogram but check if the x,y coordinates exists and increment the z
func (h *Histogram) Add(c Coordinate2d) {
	if h.OriginalMap == nil {
		h.OriginalMap = make(map[Coordinate2d]int)
	}
	if _, ok := h.OriginalMap[c]; ok {
		h.OriginalMap[c]++
	} else {
		h.OriginalMap[c] = 1
	}
}

func (h *Histogram) Scale(width, height int) {
	if (h.CurrentScale == Coordinate2dInt{width, height}) {
		return
	} else {
		h.CurrentScale = Coordinate2dInt{width, height}
		shortSide := 0.0
		if ( width < height ) {
			shortSide = float64(width)
		} else {
			shortSide = float64(height)
		}		
		offsetX := float64((float64(width) - shortSide) / 2)
		offsetY := float64((float64(height) - shortSide) / 2)
		// Clear ScaledPoints and append the scaled points
		h.ScaledMap = make(map[Coordinate2dInt]int)

		// Loop through the original points and scale them
		for key, element := range h.OriginalMap {
			scaledX := int(float64(key.X+2.0) / 4 * shortSide + offsetX)
			scaledY := int(float64(key.Y+2.0) / 4 * shortSide + offsetY) * -1 + height
			scaledPoint := Coordinate2dInt{scaledX, scaledY}
			h.ScaledMap[scaledPoint] += element
			if element > h.MaxZ {
				h.MaxZ = element
			}
		}
	}
}
