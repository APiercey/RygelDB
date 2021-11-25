package commands

import (
  "rygel/core" 
)

type DefineCollectionCommand struct {
  Store *core.Store
  CollectionName string
}

func (c DefineCollectionCommand) Execute() (string, bool) {
  if c.Store.CreateCollection(c.CollectionName) {
    return "OK", true
  } else {
    return "ERR Could not define collection", false
  }
}

func (c DefineCollectionCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  return true
}

