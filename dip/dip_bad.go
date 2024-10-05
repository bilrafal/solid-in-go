package dip

import (
	"design-patterns/solid-in-go/dip/models"
	"fmt"
)

// ResearchBad is a HLM
type ResearchBad struct {
	//	break DIP
	relationships Relationships
}

func (r ResearchBad) Investigate() {
	// Investigate method uses the internals of the LLM (relations)
	relations := r.relationships.relations

	for _, rel := range relations {
		if rel.From.Name == "John" {
			fmt.Printf("John has a child called: %s", rel.To.Name)
		}

	}
}

func UseBadInvestigate() {
	parent := models.Person{Name: "John"}
	child1 := models.Person{Name: "Rafa"}
	child2 := models.Person{Name: "Marta"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := ResearchBad{relationships}
	r.Investigate()
}
