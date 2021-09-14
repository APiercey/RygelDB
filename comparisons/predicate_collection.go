package comparisons

import (
	"example.com/rygel/store"
)

type PredicateCollection struct {
  predicates []Predicate
}

func (pc PredicateCollection) SatisfiedBy(item store.Item) bool {
	if len(pc.predicates) == 0 {
		return true
	}

	for _, wp := range pc.predicates {
		if !wp.Filter(item) {
			return false
		}
	}

	return true
}

func (pc *PredicateCollection) AddPredicate(predicate Predicate) {
	pc.predicates = append(pc.predicates, predicate)
}

func BuildPredicateCollection() PredicateCollection {
  return PredicateCollection{predicates: []Predicate{}}
}
