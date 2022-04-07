package pdjm

type Histogram struct {
	OriginalPoints []Coordinate3d
	ScaledPoints   []Coordinate3dInt
	CurrentScale   Coordinate2dInt
}

// Function to add a coordinate to the histogram but check if the x,y coordinates exists and increment the z
func (h *Histogram) Add(c Coordinate2d) {
	for i := 0; i < len((*h).OriginalPoints); i++ {
		if (*h).OriginalPoints[i].X == c.X && (*h).OriginalPoints[i].Y == c.Y {
			(*h).OriginalPoints[i].Z++
			return
		}
	}
	(*h).OriginalPoints = append((*h).OriginalPoints, Coordinate3dInitfrom2d(c))
}

func (h *Histogram) Scale(x, y int) {
	// Clear ScaledPoints and append the scaled points
	(*h).ScaledPoints = nil
	for i := 0; i < len((*h).OriginalPoints); i++ {
		x := int((*h).OriginalPoints[i].X * float64(x))
		y := int((*h).OriginalPoints[i].Y * float64(y))
		z := (*h).OriginalPoints[i].Z
		(*h).ScaledPoints = append((*h).ScaledPoints, Coordinate3dInt{x, y, z})
	}
	(*h).CurrentScale = Coordinate2dInt{x, y}
}
