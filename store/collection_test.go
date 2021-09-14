package store

import (
	"testing"
)

func TestCollectionInsertItem(t *testing.T) {
  collection := BuildCollection("flowers")

  item, _ := BuildItem(map[string]interface{}{
    "Birds of Paradise": 1,
    "Dasies": 2,
  })

  if !collection.InsertItem(item) {
    t.Log("Could not insert item", item)
    t.Fail()
  }

  if len(collection.Items) != 1 {
    t.Log("Item not present in collection", item)
    t.Fail()
  }
}
