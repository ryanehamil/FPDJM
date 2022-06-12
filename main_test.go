package main

import (
	"testing"
	"github.com/ryanehamil/FPDJM/src/attractor"
	"github.com/ryanehamil/FPDJM/src/fyre"
)

func TestAttractor(t *testing.T) {
	att := attractor.New("Peter de Jong", 10000000)
	att.Run()
	att.Plot(800, 800)
}