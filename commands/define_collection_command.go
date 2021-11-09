package commands

import (
  "rygel/core" 
)

type defineCollectionCommand struct {
  collectionName string
}

func (c defineCollectionCommand) Execute(s *core.Store) (string, bool) {
  if s.CreateCollection(c.collectionName) {
    return "OK", true
  } else {
    return "ERR Could not define collection", false
  }
}

func (c defineCollectionCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

