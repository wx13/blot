package main

import (
	"blot"
	"fmt"
)

func main() {

	plot := blot.NewBlot()
	line := blot.Line{
		X: []float64{0, 1, 2, 3},
		Y: []float64{0, 1, 0, 1},
	}
	plot.AddLine(line)
	canvas := plot.Plot("example", 500, 300)
	fmt.Println(canvas)

}

