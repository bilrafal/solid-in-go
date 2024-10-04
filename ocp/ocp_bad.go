package ocp

import "fmt"

type FilterBad struct {
	//
}

func (f FilterBad) GetByColor(
	product []Product, color Color) []Product {
	result := make([]Product, 0)

	for _, p := range product {
		if p.color == color {
			result = append(result, p)
		}
	}

	return result
}

func (f FilterBad) GetBySize(
	product []Product, size Size) []Product {
	result := make([]Product, 0)

	for _, p := range product {
		if p.size == size {
			result = append(result, p)
		}
	}

	return result
}

func (f FilterBad) GetByColorAndSize(
	product []Product, size Size, color Color) []Product {
	result := make([]Product, 0)

	for _, p := range product {
		if p.size == size && p.color == color {
			result = append(result, p)
		}
	}

	return result
}

func UseFilterBad() {

	apple := Product{name: "Apple", color: Green, size: Small}
	tree := Product{name: "Tree", color: Green, size: Big}
	house := Product{name: "House", color: Blue, size: Big}

	products := []Product{apple, tree, house}

	badFilter := FilterBad{}
	filteredProductsByColor := badFilter.GetByColor(products, Green)
	fmt.Println("filtered products by color bad: ", filteredProductsByColor)

	filteredProductsBySize := badFilter.GetBySize(products, Big)
	fmt.Println("filtered products by size bad: ", filteredProductsBySize)

	filteredProductsBySizeAndColor := badFilter.GetByColorAndSize(products, Big, Green)
	fmt.Println("filtered big green products bad: ", filteredProductsBySizeAndColor)
}
