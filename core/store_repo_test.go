package core

import (
	"testing"
)

func setupRepo() StoreRepo {
  store1 := BuildStore("test1")
  store2 := BuildStore("test2")

  storeRepo := StoreRepo{
    Stores: []Store{store1, store2},
  }

  return storeRepo
}

func TestStoreRepoFindByName(t *testing.T) {
  storeRepo := setupRepo() 
  store := storeRepo.FindByName("test1")

  if store.Name != "test1" {
    t.Log("Could not find correct store")
    t.Fail()
  }
}

