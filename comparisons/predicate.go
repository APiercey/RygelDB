package comparisons

import (
	"rygel/core"
	"rygel/common"
)

type Predicate struct {
	Path common.DataPath 
	Operator string 
	Value interface{} 
}

func (p Predicate) IndexedItems(index core.Index) []core.Item {
  if index.ContainsValue(p.Value) {
    return index.CopiedItems(p.Value)
  } else {
    return []core.Item{}
  }
}

func (wp Predicate) SatisfiedBy(item core.Item) bool {
  value, presence := item.PluckValueOnPath(wp.Path)

  if !presence { return false }

  return wp.compare(value)
}

func (wp Predicate) compare(value interface{}) bool {
  switch wp.Operator {
    case "=":
      return equals(value, wp.Value)

    case "!=":
      return notEquals(value, wp.Value)

    case ">":
      return greaterThan(value, wp.Value)

    case ">=":
      return greaterThanOrEqual(value, wp.Value)

    case "<":
      return lessThan(value, wp.Value)

    case "<=":
      return lessThanOrEqual(value, wp.Value)

    default:
      return false
  }
}

func (wp Predicate) pluckValue(item core.Item) (interface{}, bool) {
  steps := wp.Path.Steps()
  key := wp.Path.Key()

  structure := item.Data

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

