package commands

import (
	"encoding/json"
	comp "rygel/comparisons"
	"rygel/core"
)

type fetchCommand struct {
  limit int
  collectionName string
  predicates comp.PredicateCollection
}

func (c fetchCommand) RawStatement() string {
  return "";
}

func (c fetchCommand) candidateItems(s *core.Store) []core.Item {
  collection := s.Collections[c.collectionName]
  candidateItems := c.predicates.IndexedItems(collection)

  if len(candidateItems) > 0 {
    return candidateItems
  } 

  return collection.Items
}

func (c fetchCommand) Execute(s *core.Store) (string, bool) {
  matchingDataOfItems := []map[string]interface{}{}

  for _, item := range c.candidateItems(s) {
    if len(matchingDataOfItems) == c.limit {
      break
    }

    if c.predicates.SatisfiedBy(item) {
      matchingDataOfItems = append(matchingDataOfItems, item.Data)
    }
  }

  out, err := json.Marshal(matchingDataOfItems)

  if err != nil { panic (err) }

  return string(out), false
}

func (c fetchCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

