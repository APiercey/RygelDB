package commands

import (
	"rygel/core"
	cs "rygel/core/store"
	"rygel/common"
)

type InsertCommand struct {
  Store *cs.Store
  CollectionName string
  Data common.Data
}

func (c InsertCommand) Execute() (string, bool) {
  item, err := core.BuildItem(c.Data)

  if err != nil {
    return err.Error(), false
  }

  result := c.Store.InsertItem(c.CollectionName, item)

  if result {
    return "OK", true
  } else {
    return "ERR Could not store document", false
  }
}

func (c InsertCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  if len(c.Data) == 0 {
    return false
  }

  return true
}

