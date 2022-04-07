package pdjm

import "math"

// internal de Jong X function
func deJongX(p Parameters, i Coordinate2d) (x float64) {
	x = math.Sin(p.a*i.Y) + math.Sin(p.b*i.X)
	return x
}

// internal de Jong Y function
func deJongY(p Parameters, i Coordinate2d) (y float64) {
	y = math.Sin(p.c*i.X) + math.Sin(p.d*i.Y)
	return y
}
