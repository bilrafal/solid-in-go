package lsp

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type SquareBad struct {
	Rectangle
}

func NewSquareBad(size int) *SquareBad {
	sq := SquareBad{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *SquareBad) SetWidth(width int) {
	s.width = width

	//this is NOT OK
	s.height = width
}

func (s *SquareBad) SetHeight(height int) {
	s.height = height

	//this is NOT OK
	s.width = height

}

func UseSizedBad() {
	var sized Sized

	rc := &Rectangle{2, 3}
	sized = rc

	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("UseSizedBad: Expected an area of rectangle:", expectedArea,
		", got area: ", actualArea, "\n")

	sq := NewSquareBad(5)
	sized = sq

	width = sized.GetWidth()
	sized.SetHeight(10)
	actualArea = sized.GetWidth() * sized.GetHeight()
	fmt.Print("UseSizedBad: Expected an area of squareBad:", expectedArea,
		", got area: ", actualArea, "\n")
}
