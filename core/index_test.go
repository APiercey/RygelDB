package core

import (
	"testing"
	"rygel/common"
)

func TestIndexPathSerialized(t *testing.T) {
  expected := "foo,bar,comma,fun,[weird],val\"u\"es"
  serializedPath := common.DataPath{RealPath: []string{"foo", "bar", "comma,fun", "[weird]", "val\"u\"es"}}
  result := serializedPath.SerializedPath()

  if result != expected {
    t.Log("Does not match", result)
    t.Fail()
  }
}
