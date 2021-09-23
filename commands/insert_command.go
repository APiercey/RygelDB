package commands

import (
	"example.com/rygel/core"
)

type insertCommand struct {
  collectionName string
  data map[string]interface{}
}

func (c insertCommand) Execute(s *core.Store) (string, bool) {
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

func (c insertCommand) Valid() bool {
  if c.collectionName == "" {
    return false
  }

  if len(c.data) == 0 {
    return false
  }

  return true
}

