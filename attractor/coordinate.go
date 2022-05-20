package attractor

import (
	"fmt"
	"math/rand"
	"time"
)

// Simple x y coordinate system
type Coordinate2d struct {
	X float64
	Y float64
}

type Coordinate2dInt struct {
	X int
	Y int
}

// Creates coords of 1,1
func (i *Coordinate2d) Random() {
	rand.Seed(time.Now().UnixNano())
	i.X = rand.NormFloat64()
	i.Y = rand.NormFloat64()
}

// to string for debugging
func (i Coordinate2d) String() string {
	// first convert floast64 to string
	x := fmt.Sprintf("%f", i.X)
	y := fmt.Sprintf("%f", i.Y)
	return "(" + x + "," + y + ")"
}

type Coordinate3d struct {
	X float64
	Y float64
	Z int
}

type Coordinate3dInt struct {
	X int
	Y int
	Z int
}

// to string for debugging
func (i Coordinate3d) String() string {
	// first convert floast64 to string
	x := fmt.Sprintf("%f", i.X)
	y := fmt.Sprintf("%f", i.Y)
	z := fmt.Sprintf("%d", i.Z)
	return "(" + x + "," + y + "," + z + ")"
}

func Coordinate3dInitfrom2d(c Coordinate2d) Coordinate3d {
	return Coordinate3d{c.X, c.Y, 1}
}
