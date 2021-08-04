package commands

import (
  "example.com/kv_store/store" 
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
