package attractor

import (
	"math/rand"
	"time"
	"fmt"
)

func f64toString(f float64) string {
	return fmt.Sprintf("%f", f)
}


// Parameters is a map of parameter names to values
type Parameters struct {
	a float64
	b float64
	c float64
	d float64
}

// Initialize sets the parameters to their initial random values
// Uses NormFloat64 *2 to get a random value between -2 and 2
func (p *Parameters) Random() {
	rand.Seed(time.Now().UnixNano())
	p.a = rand.NormFloat64() * 2
	p.b = rand.NormFloat64() * 2
	p.c = rand.NormFloat64() * 2
	p.d = rand.NormFloat64() * 2
}

func (p *Parameters) String() string {
	return "a: " + f64toString(p.a) + " b: " + f64toString(p.b) + " c: " + f64toString(p.c) + " d: " + f64toString(p.d)
}