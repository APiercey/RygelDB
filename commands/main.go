package commands

import (
  "example.com/kv_store/store" 
)

type Command interface {
  Execute(store *store.Store) (result string, store_was_updated bool)
}
