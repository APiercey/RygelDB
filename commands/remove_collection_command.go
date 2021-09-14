package commands

import (
  "example.com/rygel/store" 
)

type RemoveCollectionCommand struct {
  collectionName string
}

func (c RemoveCollectionCommand) Execute(s *store.Store) (string, bool) {
  if s.UndefineCollection(c.collectionName) {
    return "OK", true
  } else {
    return "ERR Could not undefine collection", false
  }
}
func (c RemoveCollectionCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

