package comparisons

import (
	"example.com/rygel/store"
)

type PredicateCollection struct {
  wherePredicates []WherePredicate
}

func (pc PredicateCollection) SatisfiedBy(item store.Item) bool {
	if len(pc.wherePredicates) == 0 {
		return true
	}

	for _, wp := range pc.wherePredicates {
		if wp.Filter(item) {
			return true
		}
	}

	return false
}

func (pc *PredicateCollection) AddPredicate(predicate WherePredicate) {
	pc.wherePredicates = append(pc.wherePredicates, predicate)
}

func BuildPredicateCollection() PredicateCollection {
  return PredicateCollection{wherePredicates: []WherePredicate{}}
}
