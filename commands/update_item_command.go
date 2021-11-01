package commands

import (
	"strconv"

	comp "rygel/comparisons"
	"rygel/core"
	"rygel/common"
)

type updateItemCommand struct {
  collectionName string
  limit int
  predicates comp.PredicateCollection
  data common.Data
}

func (c updateItemCommand) RawStatement() string {
  return "";
}

func (c updateItemCommand) Execute(s *core.Store) (string, bool) {
  numFoundItems := 0

  for i := 0; i < len(s.Collections[c.collectionName].Items); i++ {
    if numFoundItems == c.limit {
      break
    }

    if c.predicates.SatisfiedBy(s.Collections[c.collectionName].Items[i]) {
      s.Collections[c.collectionName].Items[i].SetData(c.data)

      numFoundItems += 1
    }  
  }

  return "Updated " + strconv.Itoa(numFoundItems) + " items", false
}

func (c updateItemCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

