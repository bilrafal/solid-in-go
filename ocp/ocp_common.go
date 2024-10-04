package ocp

type Color string

const (
	Blue  Color = "blue"
	Green Color = "green"
)

type Size int

const (
	Small Size = 1
	Big   Size = 20
)

type Product struct {
	name  string
	color Color
	size  Size
}
