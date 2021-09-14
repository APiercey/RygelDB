package comparisons

import (
	"example.com/rygel/store"
)

type PredicateCollection struct {
  wherePredicates []Predicate
}

func (pc PredicateCollection) SatisfiedBy(item store.Item) bool {
	if len(pc.wherePredicates) == 0 {
		return true
	}

	for _, wp := range pc.wherePredicates {
		if !wp.Filter(item) {
			return false
		}
	}

	return true
}

func (pc *PredicateCollection) AddPredicate(predicate Predicate) {
	pc.wherePredicates = append(pc.wherePredicates, predicate)
}

func BuildPredicateCollection() PredicateCollection {
  return PredicateCollection{wherePredicates: []Predicate{}}
}
