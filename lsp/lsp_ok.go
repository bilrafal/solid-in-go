package lsp

type Square struct {
	size int
}

func NewSquare(size int) *Square {
	return &Square{size: size}
}

// Rectangle creates a method to return the instance of Rectangle
func (s *Square) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}
