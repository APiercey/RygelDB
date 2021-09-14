package commands

import (
  "example.com/rygel/store" 
)

type Command interface {
  Execute(store *store.Store) (result string, store_was_updated bool)
  Valid() bool
}
