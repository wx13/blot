package main

import (
	"blot"
	"fmt"
)

func main() {

	plot := blot.NewBlot()
	line := blot.Line{
		X: []float64{0, 100, 200, 300},
		Y: []float64{0, 100, 0, 100},
	}
	plot.AddLine(line)
	canvas := plot.Plot("example", 500, 300)
	fmt.Println(canvas)

}

