package comparisons

import (
	"testing"

	"example.com/rygel/store"
)

func TestFilteringByEquals(t *testing.T) {
  item, _ := store.BuildItem(map[string]interface{}{
    "test": 1,
  })

  predicate := Predicate{Path: []string{"test"}, Operator: "=", Value: 1}

  if !predicate.SatisfiedBy(item) {
    t.Log("= operator is broken")
    t.Fail()
  }
}

func TestFilteringByNotEquals(t *testing.T) {
  item, _ := store.BuildItem(map[string]interface{}{
    "test": 1,
  })

  predicate := Predicate{Path: []string{"test"}, Operator: "1=", Value: 2}

  if predicate.SatisfiedBy(item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringByGreaterThan(t *testing.T) {
  item, _ := store.BuildItem(map[string]interface{}{
    "test": 1,
  })

  predicate := Predicate{Path: []string{"test"}, Operator: ">", Value: 0}

  if predicate.SatisfiedBy(item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringByGreaterThanOrEquals(t *testing.T) {
  item, _ := store.BuildItem(map[string]interface{}{
    "test": 1,
  })

  predicate := Predicate{Path: []string{"test"}, Operator: ">=", Value: 1}

  if predicate.SatisfiedBy(item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringStaleItems(t *testing.T) {
  item, _ := store.BuildItem(map[string]interface{}{
    "Birds of Paradise": 1,
  })

  item.MarkAsStale()

  predicate := Predicate{Path: []string{"Birds of Paradise"}, Operator: "=", Value: 1}

  if predicate.SatisfiedBy(item) {
    t.Log("Does not filter ignore stale items")
    t.Fail()
  }
}

func TestFilteringByLessThan(t *testing.T) {
  item, _ := store.BuildItem(map[string]interface{}{
    "test": 1,
  })

  predicate := Predicate{Path: []string{"test"}, Operator: "<", Value: 2}

  if predicate.SatisfiedBy(item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringByLessThanOrEquals(t *testing.T) {
  item, _ := store.BuildItem(map[string]interface{}{
    "test": 1,
  })

  predicate := Predicate{Path: []string{"test"}, Operator: "<=", Value: 1}

  if predicate.SatisfiedBy(item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}