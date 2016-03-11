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
		Style: blot.Style{
			Color: "#ff0000",
			Dashed: false,
		},
	}
	plot.AddLine(line)
	line = blot.Line{
		X: []float64{2, 3, 4, 5},
		Y: []float64{-1, 1, -1, 1},
		Style: blot.Style{
			Color: "#0000ff",
			Dashed: true,
		},
	}
	plot.AddLine(line)
	canvas := plot.Plot("example", 500, 300)
	fmt.Println(canvas)

}

