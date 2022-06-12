package utils

import "math"

func Min8(a, b uint8) uint8 {
	return uint8(math.Min(float64(a), float64(b)))
}
