package dip

import "design-patterns/solid-in-go/dip/models"

// Relationships is a LLM, i.e.: storage
type Relationships struct {
	relations []models.Info
}

func (r *Relationships) AddParentAndChild(
	parent, child *models.Person,
) {
	r.relations = append(r.relations, models.Info{parent, models.Parent, child})
	r.relations = append(r.relations, models.Info{child, models.Child, parent})
}
