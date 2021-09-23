package commands

import (
  "example.com/rygel/core" 
)

type defineIndexCommand struct {
  collectionName string
  path string
}

func (c defineIndexCommand) Execute(s *core.Store) (string, bool) {
  // index := store.BuildIndex(c.path)

  // if s.AddIndex(c.collectionName, index) {
  //   return "OK", true
  // } else {
  //   return "ERR Could not define index", false
  // }
  return "TODO Implement", false
}

func (c defineIndexCommand) Valid() bool {
  // if c.collectionName == "" {
  //   return false
  // }

  // return true

  // TODO: Implement
  return false
}

