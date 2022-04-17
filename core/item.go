package core

import (
	"rygel/common"
)

type Item struct {
  Data common.Data
  wasUpdated bool
  flaggedToRemove bool
}

func (i Item) WasUpdated() bool {
  return i.wasUpdated
}

func (i Item) ShouldBeRemoved() bool {
  return i.flaggedToRemove
}

func (i *Item) FlagToRemove() {
  i.flaggedToRemove = true
}

func (i *Item) SetData(newData common.Data) {
  i.Data = newData
  i.wasUpdated = true
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

func BuildItem(data common.Data) (Item, error) {
  return Item{Data: data, wasUpdated: false, flaggedToRemove: false}, nil
}
