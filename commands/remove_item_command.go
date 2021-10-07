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

func (c removeItemCommand) candidateItems(s *core.Store) []core.Item {
  collection := s.Collections[c.collectionName]
  candidateItems := c.predicates.IndexedItems(collection)

  if len(candidateItems) > 0 {
    return candidateItems
  } 

  return collection.Items
}

func (c removeItemCommand) Execute(s *core.Store) (string, bool) {
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

func (c removeItemCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

