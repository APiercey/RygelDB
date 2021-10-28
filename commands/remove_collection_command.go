package commands

import (
  "rygel/core" 
)

type removeCollectionCommand struct {
  collectionName string
}

func (c removeCollectionCommand) Execute(s *core.Store) (string, bool) {
  if s.UndefineCollection(c.collectionName) {
    return "OK", true
  } else {
    return "ERR Could not undefine collection", false
  }
}
func (c removeCollectionCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  return true
}

