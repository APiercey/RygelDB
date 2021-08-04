package commands

import (
  "example.com/rygel/store" 
)

type DefineCollectionCommand struct {
  collectionName string
}

func (c DefineCollectionCommand) Execute(s *store.Store) (string, bool) {
  if s.CreateCollection(c.collectionName) {
    return "OK", true
  } else {
    return "ERR Could not define collection", false
  }
}
