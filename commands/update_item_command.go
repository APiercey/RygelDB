package commands

import (
	"strconv"

	comp "rygel/comparisons"
	cs "rygel/core/store"
	"rygel/core"
	"rygel/common"
)

type UpdateItemCommand struct {
  Store *cs.Store
  CollectionName string
  Limit int
  Predicates comp.PredicateCollection
  Data common.Data
}

func (c UpdateItemCommand) Execute() (string, bool) {
  collection := c.Store.Collections[c.CollectionName]
  numFoundItems := 0

  f := func(item *core.Item) bool {
    if numFoundItems == c.Limit {
      return false
    }

    if c.Predicates.SatisfiedBy(item) {
      item.SetData(c.Data)
      numFoundItems += 1
    }  

    return true
  }

  collection.Enumerate(f, false)

  return "Updated " + strconv.Itoa(numFoundItems) + " items", true
}

func (c UpdateItemCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  return true
}

