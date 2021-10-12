package commands

import (
	"strconv"

	"example.com/rygel/core"
	comp "example.com/rygel/comparisons" 
)


type removeItemCommand struct {
  collectionName string
  limit int
  predicates comp.PredicateCollection
}

func (c removeItemCommand) Execute(s *core.Store) (string, bool) {
  numFoundItems := 0

  keptItems := []core.Item{}
  for _, item := range s.Collections[c.collectionName].Items {
    keep := true

    if numFoundItems == c.limit {
      keep = false
    }

    if !c.predicates.SatisfiedBy(item) {
      keep = false
    }

    if keep {
      keptItems = append(keptItems, item)
      numFoundItems += 1
    }
  }

  collection := s.Collections[c.collectionName]
  collection.ReplaceItems(keptItems)
  s.Collections[c.collectionName] = collection

  return "Removed " + strconv.Itoa(numFoundItems) + " items", false
}

func (c removeItemCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

