package blot

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

func (b *Blot) Plot() string {
	return ""
}
