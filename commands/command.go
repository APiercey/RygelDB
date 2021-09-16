package commands

import (
  "example.com/rygel/core" 
)

type Command interface {
  Execute(store *core.Store) (result string, store_was_updated bool)
  Valid() bool
}
