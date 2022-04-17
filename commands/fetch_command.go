package commands

import (
	"encoding/json"
	comp "rygel/comparisons"
	cs "rygel/core/store"
	"rygel/core"
)

type FetchCommand struct {
  Store *cs.Store
  Limit int
  CollectionName string
  Predicates comp.PredicateCollection
}

func (c FetchCommand) Execute() (string, bool) {
  matchingDataOfItems := []map[string]interface{}{}

  // TODO: Function to access collection
  collection := c.Store.Collections[c.CollectionName]

  f := func(item *core.Item) bool {
    if len(matchingDataOfItems) == c.Limit { return false }

    if c.Predicates.SatisfiedBy(item) {
      matchingDataOfItems = append(matchingDataOfItems, item.Data)
    }

    return true
  }

  collection.Enumerate(f, false)

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

