package comparisons

import (
	"reflect"

	"example.com/rygel/store"
)

type WherePredicate struct {
	Path []string `json:"path"`
	Operator string `json:"operator"`
	Value interface{} `json:"value"`
}

func (wp WherePredicate) Filter(item store.Item) bool {
  value, presence := wp.pluckValue(item)

  if !presence { return false }

  return wp.compare(value)
}

func (wp WherePredicate) compare(value interface{}) bool {
  switch wp.Operator {
    case "=":
      return value == wp.Value
    case "!=":
      return value != wp.Value
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

func greaterThan(leftValue interface{}, rightValue interface{}) bool {
  leftSide := reflect.TypeOf(leftValue).Kind()
  rightSide := reflect.TypeOf(rightValue).Kind()

  if leftSide.String() != rightSide.String() { return false }

  if leftSide == reflect.Float64 { return leftValue.(float64) > rightValue.(float64) }
  if leftSide == reflect.String { return leftValue.(string) > rightValue.(string) }

  return false
}

func greaterThanOrEqual(leftValue interface{}, rightValue interface{}) bool {
  leftSide := reflect.TypeOf(leftValue).Kind()
  rightSide := reflect.TypeOf(rightValue).Kind()

  if leftSide.String() != rightSide.String() { return false }

  if leftSide == reflect.Float64 { return leftValue.(float64) >= rightValue.(float64) }
  if leftSide == reflect.String { return leftValue.(string) >= rightValue.(string) }

  return false
}

func lessThan(leftValue interface{}, rightValue interface{}) bool {
  leftSide := reflect.TypeOf(leftValue).Kind()
  rightSide := reflect.TypeOf(rightValue).Kind()

  if leftSide.String() != rightSide.String() { return false }

  if leftSide == reflect.Float64 { return leftValue.(float64) < rightValue.(float64) }
  if leftSide == reflect.String { return leftValue.(string) < rightValue.(string) }

  return false
}

func lessThanOrEqual(leftValue interface{}, rightValue interface{}) bool {
  leftSide := reflect.TypeOf(leftValue).Kind()
  rightSide := reflect.TypeOf(rightValue).Kind()

  if leftSide.String() != rightSide.String() { return false }

  if leftSide == reflect.Float64 { return leftValue.(float64) <= rightValue.(float64) }
  if leftSide == reflect.String { return leftValue.(string) <= rightValue.(string) }

  return false
}

func (wp WherePredicate) pluckValue(item store.Item) (interface{}, bool) {
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
