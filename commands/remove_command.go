package commands

import (
  "example.com/kv_store/store" 
)

type RemoveCommand struct {
  collectionName string
  key string
}

func (c RemoveCommand) Execute(s *store.Store) (string, bool) {
  if s.RemoveItem(c.collectionName, c.key) {
    return "OK", true
  } else {
    return "ERR", false
  }
}
