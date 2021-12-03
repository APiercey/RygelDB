package commands

import (
  cs "rygel/core/store" 
)

type DefineCollectionCommand struct {
  Store *cs.Store
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

