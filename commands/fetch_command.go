package commands

import (
	"encoding/json"

	"example.com/rygel/store"
	comp "example.com/rygel/comparisons"
)

type FetchCommand struct {
  limit int
  collectionName string
  predicates comp.PredicateCollection
}

func (c FetchCommand) Execute(s *store.Store) (string, bool) {
  items := []map[string]interface{}{}

  numFoundItems := 0

  for _, item := range s.Collections[c.collectionName].Items {
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

func (c FetchCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

