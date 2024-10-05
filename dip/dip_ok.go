package dip

import (
	"design-patterns/solid-in-go/dip/models"
	"design-patterns/solid-in-go/dip/rels_store"
	"fmt"
)

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*models.Person
}

type Research struct {
	browser RelationshipBrowser
}

func (r Research) Investigate() {
	for _, person := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child 	called", person.Name)
	}
}

func UseInvestigate() {
	parent := models.Person{Name: "John"}
	child1 := models.Person{Name: "Rafa"}
	child2 := models.Person{Name: "Marta"}

	relationships := rels_store.Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
