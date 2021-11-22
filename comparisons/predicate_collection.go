package comparisons

import (
	"rygel/core"
	"rygel/common"
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

func (pc PredicateCollection) IsOverlapping(dp common.DataPath) bool {
  for _, pred := range pc.predicates {
		if pred.Path.Equals(dp) {
			return true
		}
  }

	return false
}

func BuildPredicateCollection() PredicateCollection {
  return PredicateCollection{predicates: []Predicate{}}
}
