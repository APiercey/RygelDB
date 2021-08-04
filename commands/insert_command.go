package commands

import (
  "example.com/kv_store/store" 
)

type InsertCommand struct {
  collectionName string
  key string
  data map[string]interface{}
}

func (c InsertCommand) Execute(s *store.Store) (string, bool) {
  item := store.BuildItem(c.key, c.data)

  if s.InsertItem(c.collectionName, item) {
    return "OK", true
  } else {
    return "ERR Could not store document", false
  }
}
