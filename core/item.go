package core

import (
	"example.com/rygel/common"
)

type Item struct {
  Data map[string]interface{}
  IsStale bool
}

func (i *Item) MarkAsStale() {
  i.IsStale = true
}

func (i Item) PluckValueOnPath(dp common.DataPath) (interface{}, bool) {
  steps := dp.Steps()
  key := dp.Key()

  structure := i.Data

  for _, step := range steps {
    traversedStructure, presence := structure[step]

    if presence {
      structure = traversedStructure.(map[string]interface{})
    } else {
      // It would be possible to traverse into arrays
      // but I wont implement this yet
      return nil, false
    }
  }

  value, presence := structure[key]
  return value, presence
}

func BuildItem(data map[string]interface{}) (Item, error) {
  return Item{Data: data, IsStale: false}, nil
}
