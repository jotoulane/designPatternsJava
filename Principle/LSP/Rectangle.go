package LSP

type Rectangle struct {
	length int
	width  int
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetLength(length int) {
	r.length = length
}

func (r *Rectangle) Width() int {
	return r.width
}

func (r *Rectangle) Length() int {
	return r.length
}

func NewRectangle(length, width int) *Rectangle {
	return &Rectangle{length: length, width: width}
}
