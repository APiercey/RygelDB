package core

import (
  "os"
  "rygel/common"
)

type StoreRepo struct {
  Dir string
  Stores []Store
}

func (sr StoreRepo) FindByName(name string) *Store {
  for _, store := range sr.Stores {
    if store.Name == name {
      return &store
    }
  }

  panic("Could not find store")
}

func (sr *StoreRepo) Create(name string) {
  for _, store := range sr.Stores {
    if store.Name == name {
      panic("Store already exists")
    }
  }

  f, err := os.Create(sr.Dir + name + ".store")

  defer f.Close()

  common.HandleErr(err)

  _, err2 := f.WriteString(name + "\n")

  common.HandleErr(err2)

  store := BuildStore(name)

  sr.Stores = append(sr.Stores, store)
}

func InitializeFromDir(dir string) StoreRepo {

  return StoreRepo{Dir: dir}
}
