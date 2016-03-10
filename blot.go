package blot

import "fmt"

type Style struct {
}

type Line struct {
	X, Y  []float64
	Style Style
	Label string
}

type Blot struct {
	Lines []Line
}

func NewBlot() *Blot {
	return &Blot{}
}

func (b *Blot) AddLine(line Line) {
	b.Lines = append(b.Lines, line)
}

func (b *Blot) Plot(id string, width, height int) string {

	elem := fmt.Sprintf(`<canvas id="%s" width="%d" height="%d"></canvas>`, id, width, height)
	script := "<script>"
	script += fmt.Sprintf(`canvas = document.getElementById("%s");`, id)
	script += `context = canvas.getContext("2d");`

	for _, line := range b.Lines {
		script += fmt.Sprintf("context.moveTo(%f,%f);", line.X[0], line.Y[0])
		for k := range line.X {
			script += fmt.Sprintf("context.lineTo(%f,%f);", line.X[k], line.Y[k])
		}
		script += "context.stroke();"
	}

	script += "</script>"

	return elem + script

}

