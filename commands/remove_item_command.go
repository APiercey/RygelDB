package commands

import (
	"strconv"

	"rygel/core"
	cs "rygel/core/store"
	comp "rygel/comparisons" 
)

type RemoveItemCommand struct {
  Store *cs.Store
  CollectionName string
  Limit int
  Predicates comp.PredicateCollection
}

func (c RemoveItemCommand) Execute() (string, bool) {
  numFoundItems := 0

  keptItems := []core.Item{}
  for _, item := range c.Store.Collections[c.CollectionName].Items {
    keep := true

    if numFoundItems == c.Limit {
      keep = false
    }

    if !c.Predicates.SatisfiedBy(item) {
      keep = false
    }

    if keep {
      keptItems = append(keptItems, item)
      numFoundItems += 1
    }
  }

  collection := c.Store.Collections[c.CollectionName]
  collection.ReplaceItems(keptItems)
  c.Store.Collections[c.CollectionName] = collection

  return "Removed " + strconv.Itoa(numFoundItems) + " items", true
}

func (c RemoveItemCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  return true
}

