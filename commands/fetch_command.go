package commands

import (
	"encoding/json"

	"example.com/rygel/core"
	comp "example.com/rygel/comparisons"
)

type fetchCommand struct {
  limit int
  collectionName string
  predicates comp.PredicateCollection
}

func (c fetchCommand) Execute(s *core.Store) (string, bool) {
  items := []map[string]interface{}{}

  numFoundItems := 0

  collection := s.Collections[c.collectionName]

  for _, item := range collection.Items {
    if numFoundItems == c.limit {
      break
    }

    if c.predicates.SatisfiedBy(item) {
      items = append(items, item.Data)
      numFoundItems += 1
    }
  }

  out, err := json.Marshal(items)

  if err != nil { panic (err) }

  return string(out), false
}

func (c fetchCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

