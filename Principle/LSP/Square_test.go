package LSP

import (
	"testing"
)

func TestNew(t *testing.T) {
	rectangle := NewRectangle(5, 10)
	square := NewSquare(7)

	PrintArea(rectangle)
	PrintArea(square)
}

func PrintArea(q Quadrilateral) {
	area := q.Width() * q.Length()
	println(area)
}
