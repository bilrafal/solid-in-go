package ocp

import "fmt"

// Specification interface type is open for extension, meaning you can implement this
// interface, but it's closed for modification, which means that you are unlikely to ever modify the
// Specification interface and in a similar fashion you are unlikely to ever modify Filter type
type Specification interface {
	IsSatisfied(p Product) bool
}

type ColorSpecification struct {
	color Color
}

func (cs *ColorSpecification) IsSatisfied(p Product) bool {
	return p.color == cs.color
}

type SizeSpecification struct {
	size Size
}

func (ss *SizeSpecification) IsSatisfied(p Product) bool {
	return p.size == ss.size
}

// composite specification
type AndSpecification struct {
	specs []Specification
}

func (as *AndSpecification) IsSatisfied(p Product) bool {
	for _, spec := range as.specs {
		if !spec.IsSatisfied(p) {
			return false
		}
	}
	return true
}

type Filter struct {
}

func (f Filter) Get(
	products []Product, spec Specification) []Product {

	result := make([]Product, 0)
	for _, product := range products {
		if spec.IsSatisfied(product) {
			result = append(result, product)
		}
	}

	return result
}

func UseFilterOK() {

	apple := Product{name: "Apple", color: Green, size: Small}
	tree := Product{name: "Tree", color: Green, size: Big}
	house := Product{name: "House", color: Blue, size: Big}

	products := []Product{apple, tree, house}

	filter := Filter{}
	colorSpec := &ColorSpecification{color: Green}
	filteredProductsByColor := filter.Get(products, colorSpec)
	fmt.Println("filtered green products ok: ", filteredProductsByColor)

	sizeSpec := &SizeSpecification{size: Big}
	filteredProductsBySize := filter.Get(products, sizeSpec)
	fmt.Println("filtered big products ok: ", filteredProductsBySize)

	sizeAndColorSpec := &AndSpecification{specs: []Specification{colorSpec, sizeSpec}}
	filteredProductsBySizeAndColor := filter.Get(products, sizeAndColorSpec)
	fmt.Println("filtered big green products ok: ", filteredProductsBySizeAndColor)
}
