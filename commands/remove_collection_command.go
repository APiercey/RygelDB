package commands

import (
  "rygel/core" 
)

type RemoveCollectionCommand struct {
  Store *core.Store
  CollectionName string
}

func (c RemoveCollectionCommand) Execute() (string, bool) {
  if c.Store.UndefineCollection(c.CollectionName) {
    return "OK", true
  } else {
    return "ERR Could not undefine collection", false
  }
}
func (c RemoveCollectionCommand) Valid() bool {
  if c.CollectionName == "" {
    return false
  }

  return true
}

