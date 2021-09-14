package comparisons

import (

	"example.com/rygel/store"
)

type Predicate struct {
	Path []string `json:"path"`
	Operator string `json:"operator"`
	Value interface{} `json:"value"`
}

func (wp Predicate) Filter(item store.Item) bool {
  value, presence := wp.pluckValue(item)

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

func (wp Predicate) pluckValue(item store.Item) (interface{}, bool) {
  steps := wp.Path[:len(wp.Path) - 1]
  key := wp.Path[len(wp.Path) - 1]

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
