package main

import (
	"github.com/ryanehamil/FPDJM/src/attractor"
	"github.com/ryanehamil/FPDJM/src/fyre"
)

// main function
func main() {

	fpdjm := fyre.New(attractor.New("Peter de Jong"))

	fpdjm.Run()

}
