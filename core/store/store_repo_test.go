package store

import (
	"testing"

	"os"
)

func TestStoreRepoCreateStore(t *testing.T) {
  defer func() { os.RemoveAll("./" + t.Name()) }()

  storeRepo := InitializeFromDir("./" + t.Name())
  store, err := storeRepo.Create("new_store")

  if store.Name != "new_store" || err != nil {
    t.Log("Could not find correct store")
    t.Fail()
  }
}

func TestStoreRepoFindByName(t *testing.T) {
  defer func() { os.RemoveAll("./" + t.Name()) }()

  storeRepo := InitializeFromDir("./" + t.Name())
  storeRepo.Create("test1")
  storeRepo.Create("test2")

  store, err := storeRepo.FindByName("test2")

  if store.Name != "test2" || err != nil {
    t.Logf("Could not find correct store. Store found: %s", store.Name)
    t.Fail()
  }
}

func TestStoreRepoPersistence(t *testing.T) {
  defer func() { os.RemoveAll("./" + t.Name()) }()

  firstStoreRepo := InitializeFromDir("./" + t.Name())
  firstStoreRepo.Create("test")

  secondStoreRepo := InitializeFromDir("./" + t.Name())
  store, err := secondStoreRepo.FindByName("test")

  if err != nil || store.Name != "test" {
    t.Logf("Store persistence does not work. Err: %s", err)
    t.Fail()
  }

}
