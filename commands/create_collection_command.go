package commands

import (
  "example.com/kv_store/store" 
)

type CreateCollectionCommand struct {
  collectionName string
}

func (c CreateCollectionCommand) Execute(s *store.Store) (string, bool) {
  if s.CreateCollection(c.collectionName) {
    return "OK", true
  } else {
    return "ERR Could not define collection", false
  }
}
