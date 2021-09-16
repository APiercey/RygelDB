package comparisons

import (
	"reflect"
)

func equals(leftValue interface{}, rightValue interface{}) bool {
  return leftValue == rightValue
}

func notEquals(leftValue interface{}, rightValue interface{}) bool {
  return leftValue != rightValue
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
