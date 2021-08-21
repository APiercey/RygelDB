package commands

import (
  "example.com/rygel/store" 
)

type DefineIndexCommand struct {
  collectionName string
  path string
}

func (c DefineIndexCommand) Execute(s *store.Store) (string, bool) {
  index := store.BuildIndex(c.path)

  if s.AddIndex(c.collectionName, index) {
    return "OK", true
  } else {
    return "ERR Could not define index", false
  }
}
