package commands

import (
  "example.com/rygel/core" 
)

type DefineCollectionCommand struct {
  collectionName string
}

func (c DefineCollectionCommand) Execute(s *core.Store) (string, bool) {
  if s.CreateCollection(c.collectionName) {
    return "OK", true
  } else {
    return "ERR Could not define collection", false
  }
}

func (c DefineCollectionCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

