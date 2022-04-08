package pdjm

import "sort"

type Histogram struct {
	OriginalPoints []Coordinate3d
	ScaledPoints   []Coordinate3dInt
	CurrentScale   Coordinate2dInt
}

type byY []Coordinate3d

func (s byY) Len() int {
	return len(s)
}
func (s byY) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byY) Less(i, j int) bool {
	return s[i].X < s[j].X
}

type byX []Coordinate3d

func (s byX) Len() int {
	return len(s)
}
func (s byX) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byX) Less(i, j int) bool {
	return s[i].X < s[j].X
}

// Function to add a coordinate to the histogram but check if the x,y coordinates exists and increment the z
func (h *Histogram) Add(c Coordinate2d) {
	// for i := 0; i < len((*h).OriginalPoints); i++ {
	// 	if (*h).OriginalPoints[i].X == c.X && (*h).OriginalPoints[i].Y == c.Y {
	// 		(*h).OriginalPoints[i].Z++
	// 		return
	// 	}
	// }
	h.OriginalPoints = append(h.OriginalPoints, Coordinate3dInitfrom2d(c))
}

func (h *Histogram) Sort() {
	sort.Sort(byY(h.OriginalPoints))
	sort.Sort(byX(h.OriginalPoints))
}

func (h *Histogram) Scale(width, height int) {
	if (h.CurrentScale == Coordinate2dInt{width, height}) {
		return
	} else {
		// Clear ScaledPoints and append the scaled points
		h.ScaledPoints = nil
		for i := 0; i < len(h.OriginalPoints); i++ {
			x := int((h.OriginalPoints[i].X + 2.0) / 4.0 * float64(width))
			y := int((h.OriginalPoints[i].Y + 2.0) / 4.0 * float64(height))
			z := h.OriginalPoints[i].Z
			h.ScaledPoints = append(h.ScaledPoints, Coordinate3dInt{x, y, z})
		}
		h.CurrentScale = Coordinate2dInt{width, height}
	}
}
