package main

import (
	"github.com/ryanehamil/FPDJM/attractor"
	"github.com/ryanehamil/FPDJM/fyre"
)

// main function
func main() {

	fpdjm := fyre.New(attractor.New("Peter de Jong", 10000000))

	fpdjm.Run()

}
