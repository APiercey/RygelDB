package commands

import (
	"example.com/rygel/core"
)

type InsertCommand struct {
  collectionName string
  data map[string]interface{}
}

func (c InsertCommand) Execute(s *core.Store) (string, bool) {
  item, err := core.BuildItem(c.data)

  if err != nil {
    return err.Error(), false
  }

  result := s.InsertItem(c.collectionName, item)

  if result {
    return "OK", true
  } else {
    return "ERR Could not store document", false
  }
}

func (c InsertCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  if len(c.data) == 0 {
    return false
  }

  return true
}

