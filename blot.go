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
	Lines            []Line
	Width, Height    float64
	ScaleX, ScaleY   float64
	OffsetX, OffsetY float64
	Margin           float64
}

func NewBlot() *Blot {
	return &Blot{}
}

func (b *Blot) AddLine(line Line) {
	b.Lines = append(b.Lines, line)
}

func (b *Blot) Scale(xIn, yIn float64) (float64, float64) {
	xOut := (xIn - b.OffsetX) * b.ScaleX
	yOut := b.Height - (yIn-b.OffsetY)*b.ScaleY
	return xOut, yOut
}

func (b *Blot) SetSize(width, height int) {
	b.Width = float64(width)
	b.Height = float64(height)
	minX, maxX, minY, maxY := b.GetMinMax()
	b.OffsetX = minX
	b.OffsetY = minY
	b.ScaleX = b.Width / (maxX - minX)
	b.ScaleY = b.Height / (maxY - minY)
	b.Margin = b.Width / 20.0
}

func (b *Blot) GetMinMax() (minX, maxX, minY, maxY float64) {
	minX = b.Lines[0].X[0]
	maxX = b.Lines[0].X[0]
	minY = b.Lines[0].Y[0]
	maxY = b.Lines[0].Y[0]
	for _, line := range b.Lines {
		for k := range line.X {
			if line.X[k] < minX {
				minX = line.X[k]
			}
			if line.X[k] > maxX {
				maxX = line.X[k]
			}
			if line.Y[k] < minY {
				minY = line.Y[k]
			}
			if line.Y[k] > maxY {
				maxY = line.Y[k]
			}
		}
	}
	return
}

func (b *Blot) MakeAxes() string {

	script := ""
	script += fmt.Sprintf("context.moveTo(%f, %f);", b.Margin, b.Margin)
	script += fmt.Sprintf("context.lineTo(%f, %f);", b.Margin, b.Height - b.Margin)
	script += fmt.Sprintf("context.lineTo(%f, %f);", b.Width - b.Margin, b.Height - b.Margin)
	script += fmt.Sprintf("context.lineTo(%f, %f);", b.Width - b.Margin, b.Margin)
	script += fmt.Sprintf("context.lineTo(%f, %f);", b.Margin, b.Margin)

	minX, maxX, minY, maxY := b.GetMinMax()
	fontSize := int(0.5*b.Margin)
	script += fmt.Sprintf(`context.font = "%dpx Arial";`, fontSize)
	script += `context.textAlign = "left";`
	script += `context.textBaseline = "middle";`
	script += fmt.Sprintf(`context.fillText("%g", 0, %f);`, maxY, b.Margin)
	script += `context.textAlign = "left";`
	script += `context.textBaseline = "middle";`
	script += fmt.Sprintf(`context.fillText("%g", 0, %f);`, minY, b.Height-b.Margin)
	fmt.Println(minX, maxX, minY, maxY)

	return script
}

func (b *Blot) Plot(id string, width, height int) string {

	b.SetSize(width, height)

	elem := fmt.Sprintf(`<canvas id="%s" width="%d" height="%d"></canvas>`, id, width, height)
	script := "<script>"
	script += fmt.Sprintf(`canvas = document.getElementById("%s");`, id)
	script += `context = canvas.getContext("2d");`

	script += b.MakeAxes()

	for _, line := range b.Lines {
		script += b.PlotLine(line)
	}

	script += "</script>"

	return elem + script

}

func (b *Blot) PlotLine(line Line) string {
	x, y := b.Scale(line.X[0], line.Y[0])
	script := fmt.Sprintf("context.moveTo(%f,%f);", x, y)
	for k := range line.X {
		x, y := b.Scale(line.X[k], line.Y[k])
		script += fmt.Sprintf("context.lineTo(%f,%f);", x, y)
	}
	script += "context.stroke();"
	return script
}
