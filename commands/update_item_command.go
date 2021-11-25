package commands

import (
	"strconv"

	comp "rygel/comparisons"
	"rygel/core"
	"rygel/common"
)

type UpdateItemCommand struct {
  Store *core.Store
  CollectionName string
  Limit int
  Predicates comp.PredicateCollection
  Data common.Data
}

func (c UpdateItemCommand) Execute() (string, bool) {
  numFoundItems := 0

  for i := 0; i < len(c.Store.Collections[c.CollectionName].Items); i++ {
    if numFoundItems == c.Limit {
      break
    }

    if c.Predicates.SatisfiedBy(c.Store.Collections[c.CollectionName].Items[i]) {
      c.Store.Collections[c.CollectionName].Items[i].SetData(c.Data)

      numFoundItems += 1
    }  
  }

  return "Updated " + strconv.Itoa(numFoundItems) + " items", true
}

func (c UpdateItemCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  return true
}

