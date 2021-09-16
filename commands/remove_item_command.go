package commands

import (
	"strconv"

	"example.com/rygel/core"
	comp "example.com/rygel/comparisons" 
)

type RemoveItemCommand struct {
  collectionName string
  limit int
  predicates comp.PredicateCollection
}

func (c RemoveItemCommand) Execute(s *core.Store) (string, bool) {
  numFoundItems := 0

  collection := s.Collections[c.collectionName]
  items := collection.Items

  for _, item := range items {
    if numFoundItems == c.limit {
      break
    }

    if c.predicates.SatisfiedBy(item) {
      item.MarkAsStale()
      numFoundItems += 1
    }
  }

  collection.Items = items
  s.Collections[c.collectionName] = collection

  return "Removed " + strconv.Itoa(numFoundItems) + " items", false
}

func (c RemoveItemCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

