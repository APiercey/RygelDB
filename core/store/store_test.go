package store

import (
	"testing"
	"rygel/common"
  "rygel/core"
  "os"
)

func useTempStore(name string, f func(Store)) {
  tmpDir := "./tmp-test-store"

  if err := os.MkdirAll(tmpDir, 0755); err != nil {
    panic(err)
  }

  defer func() {
    if err := os.RemoveAll(tmpDir); err != nil {
      panic(err)
    }
  }()

  store := BuildStore(name, tmpDir)

  f(store)
}

func TestCreateCollection(t *testing.T) {
  useTempStore("test", func(store Store) {
    if !store.CreateCollection("flowers") {
      t.Log("Could not create collection")
      t.Fail()
    }

    if len(store.Collections) != 1 {
      t.Log("Collection not present in store")
      t.Fail()
    }
  })
}

func TestStoreInsertItem(t *testing.T) {
  useTempStore("test", func(store Store) {
    store.CreateCollection("flowers")

    item, _ := core.BuildItem(common.Data{
      "Birds of Paradise": 1,
      "Dasies": 2,
    })

    if !store.InsertItem("flowers", item) {
      t.Log("Could not insert item", item)
      t.Fail()
    }

    amount := 0
    store.Collections["flowers"].Enumerate(func(item *core.Item) bool {
      amount += 1
      return true
    }, false)

    if amount != 1 {
      t.Log("Item not present in collection. Found:", amount)
      t.Fail()
    }
  })
}
