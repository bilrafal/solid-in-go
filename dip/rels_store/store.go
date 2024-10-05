package rels_store

import "design-patterns/solid-in-go/dip/models"

// Relationships is a LLM, i.e.: storage
type Relationships struct {
	relations []models.Info
}

func (r Relationships) FindAllChildrenOf(name string) []*models.Person {
	result := make([]*models.Person, 0)
	for i, rel := range r.relations {
		if rel.Relationship == models.Parent && rel.From.Name == name {
			result = append(result, r.relations[i].To)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(
	parent, child *models.Person,
) {
	r.relations = append(r.relations, models.Info{parent, models.Parent, child})
	r.relations = append(r.relations, models.Info{child, models.Child, parent})
}
