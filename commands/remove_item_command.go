package commands

import (
  "example.com/rygel/store" 
)

type RemoveItemCommand struct {
  collectionName string
  key string
}

func (c RemoveItemCommand) Execute(s *store.Store) (string, bool) {
  if s.RemoveItem(c.collectionName, c.key) {
    return "OK", true
  } else {
    return "ERR", false
  }
}
