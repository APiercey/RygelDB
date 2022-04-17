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
  collection := c.Store.Collections[c.CollectionName]
  numFoundItems := 0

  f := func(item *core.Item) bool {
    if numFoundItems == c.Limit {
      return false
    }


    if c.Predicates.SatisfiedBy(item) {
      item.FlagToRemove()
      numFoundItems += 1
    }

    return true
  }

  collection.Enumerate(f, false)

  return "Removed " + strconv.Itoa(numFoundItems) + " items", true
}

func (c RemoveItemCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  return true
}

