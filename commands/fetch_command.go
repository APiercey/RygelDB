package commands

import (
	"encoding/json"
	comp "rygel/comparisons"
	"rygel/core"
)

type FetchCommand struct {
  Store *core.Store
  Limit int
  CollectionName string
  Predicates comp.PredicateCollection
}

func (c FetchCommand) candidateItems(s *core.Store) []core.Item {
  collection := s.Collections[c.CollectionName]

  return collection.Items
}

func (c FetchCommand) Execute() (string, bool) {
  matchingDataOfItems := []map[string]interface{}{}

  for _, item := range c.candidateItems(c.Store) {
    if len(matchingDataOfItems) == c.Limit {
      break
    }

    if c.Predicates.SatisfiedBy(item) {
      matchingDataOfItems = append(matchingDataOfItems, item.Data)
    }
  }

  out, err := json.Marshal(matchingDataOfItems)

  if err != nil { panic (err) }

  return string(out), false
}

func (c FetchCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  return true
}

