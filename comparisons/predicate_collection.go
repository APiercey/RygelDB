package comparisons

import (
	"example.com/rygel/core"
)

type PredicateCollection struct {
  predicates []Predicate
}

func (pc PredicateCollection) SatisfiedBy(item core.Item) bool {
	if len(pc.predicates) == 0 {
		return true
	}

	for _, wp := range pc.predicates {
		if !wp.SatisfiedBy(item) {
			return false
		}
	}

	return true
}

func (pc *PredicateCollection) AddPredicate(predicate Predicate) {
	pc.predicates = append(pc.predicates, predicate)
}

// func (pc PredicateCollection) OverLappingPaths(paths []string) {
// 	pc.predicates = append(pc.predicates, predicate)
// }

func BuildPredicateCollection() PredicateCollection {
  return PredicateCollection{predicates: []Predicate{}}
}
