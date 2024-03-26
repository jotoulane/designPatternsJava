package LSP

type Square struct {
	sideLength int
}

func (s *Square) SetSideLength(sideLength int) {
	s.sideLength = sideLength
}

func (s *Square) Width() int {
	return s.sideLength
}

func (s *Square) Length() int {
	return s.sideLength
}

func NewSquare(sideLength int) *Square {
	return &Square{sideLength: sideLength}
}
