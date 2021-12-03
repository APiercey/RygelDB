package store

import (
	"testing"
	"rygel/common"
  "rygel/core"
)

func TestCreateCollection(t *testing.T) {
  store := BuildStore("test")

  if !store.CreateCollection("flowers") {
    t.Log("Could not create collection")
    t.Fail()
  }

  if len(store.Collections) != 1 {
    t.Log("Collection not present in store")
    t.Fail()
  }
}

func TestStoreInsertItem(t *testing.T) {
  store := BuildStore("test")
  store.CreateCollection("flowers")

  item, _ := core.BuildItem(common.Data{
    "Birds of Paradise": 1,
    "Dasies": 2,
  })

  if !store.InsertItem("flowers", item) {
    t.Log("Could not insert item", item)
    t.Fail()
  }

  if len(store.Collections["flowers"].Items) != 1 {
    t.Log("Item not present in collection", item)
    t.Fail()
  }

}
