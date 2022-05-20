package attractor

import "math"

// Peter de Jong attractor function
// https://en.wikipedia.org/wiki/De_Jong_attractor
func deJong(xi, yi, a, b, c, d float64) (x, y float64) {
	x = math.Sin(a*yi) + math.Cos(b*xi)
	y = math.Sin(c*xi) + math.Cos(d*yi)
	return x, y
}
