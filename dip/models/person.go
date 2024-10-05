package models

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	Name string
}

type Info struct {
	From         *Person
	Relationship Relationship
	To           *Person
}
