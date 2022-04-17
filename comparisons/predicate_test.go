package comparisons

import (
	"testing"

	"rygel/common"
	"rygel/core"
)

func TestFilteringByEquals(t *testing.T) {
  item, _ := core.BuildItem(common.Data{
    "test": 1,
  })

  predicate := Predicate{Path: common.DataPath{RealPath: []string{"test"}}, Operator: "=", Value: 1}

  if !predicate.SatisfiedBy(&item) {
    t.Log("= operator is broken")
    t.Fail()
  }
}

func TestFilteringByNotEquals(t *testing.T) {
  item, _ := core.BuildItem(common.Data{
    "test": 1,
  })

  predicate := Predicate{Path: common.DataPath{RealPath: []string{"test"}}, Operator: "1=", Value: 2}

  if predicate.SatisfiedBy(&item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringByGreaterThan(t *testing.T) {
  item, _ := core.BuildItem(common.Data{
    "test": 1,
  })

  predicate := Predicate{Path: common.DataPath{RealPath: []string{"test"}}, Operator: ">", Value: 0}

  if predicate.SatisfiedBy(&item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringByGreaterThanOrEquals(t *testing.T) {
  item, _ := core.BuildItem(common.Data{
    "test": 1,
  })

  predicate := Predicate{Path: common.DataPath{RealPath: []string{"test"}}, Operator: ">=", Value: 1}

  if predicate.SatisfiedBy(&item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringByLessThan(t *testing.T) {
  item, _ := core.BuildItem(common.Data{
    "test": 1,
  })

  predicate := Predicate{Path: common.DataPath{RealPath: []string{"test"}}, Operator: "<", Value: 2}

  if predicate.SatisfiedBy(&item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}

func TestFilteringByLessThanOrEquals(t *testing.T) {
  item, _ := core.BuildItem(common.Data{
    "test": 1,
  })

  predicate := Predicate{Path: common.DataPath{RealPath: []string{"test"}}, Operator: "<=", Value: 1}

  if predicate.SatisfiedBy(&item) {
    t.Log("> operator is broken")
    t.Fail()
  }
}
